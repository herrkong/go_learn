//Pull外部Server的transfer指令
package server

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	ori_redis "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"runtime/debug"
	"strconv"
	"strings"
	"threshold/src/external_server/redis"
	"threshold/src/public/apis"
	"threshold/src/public/config"
	"threshold/src/public/data/external_data"
	"threshold/src/public/enum"
	"threshold/src/public/models"
	pubUtils "threshold/src/public/utils"
	"time"
)


func (self *ExternalService) handleCmd(id int) (item *external_data.ExternalServerCmdItem, ret error) {
	// var txObj orm.TxOrmer
	// dbObj := orm.NewOrm()
	// if txObj, ret = dbObj.Begin(); ret != nil {
	// 	logs.Error("[handleCmd] dbObj.Begin() err", ret)
	// 	return nil, ret
	// }
	// defer func() {
	// 	if errs := recover(); errs != nil {
	// 		ret = txObj.Rollback()
	// 		logs.Error("%+v, %s", errs, string(debug.Stack()))
	// 	} else {
	// 		ret = txObj.Commit()
	// 	}
	// }()
	// cmd := new(models.IwalaTssTransferCmd)
	// if err := txObj.QueryTable("iwala_tss_transfer_cmd").Filter("id", id).ForUpdate().One(cmd); err != nil {
	// 	panic(err)
	// }
	// if cmd.Sign == "" {
	// 	logs.Error("[handleCmd] transfer_cmd no sign, cmd id:%d", cmd.Id)
	// 	return nil, nil
	// }
	// cmd.State = uint8(enum.CMD_STATE_TEMP)
	// if _, err := txObj.Update(cmd, "state"); err != nil {
	// 	panic(err)
	// }
	// updateError := func(err error) {
	// 	logs.Error("[updateError] err: %+v, stack:%s", err, string(debug.Stack()))
	// 	cmd.State = uint8(enum.CMD_STATE_ERROR)
	// 	cmd.Err = err.Error()
	// 	cmd.Sign = ""
	// 	if _, err := txObj.Update(cmd, "state", "err", "sign"); err != nil {
	// 		panic(err)
	// 	}
	// }
	// coin := strings.ToLower(cmd.CurrencyMark)
	// var err error
	// var gasPrice uint64

	

	// item = &external_data.ExternalServerCmdItem{
	// 	Id:           int64(cmd.Id),
	// 	Coin:         coin,
	// 	SssChain:     int32(chain),
	// 	TssChain:     int32(chain),
	// 	From:         nextSignAddress,
	// 	To:           cmd.ToAddress,
	// 	Amount:       cmd.Amount,
	// 	Summary:      summary,
	// 	Sign:         cmd.Sign,
	// 	TransferType: cmd.TransferType,
	// 	Curve:        int32(curve),
	// 	PublicKey:    publicKey,
	// }
	// //更改这条记录为处理中
	// cmd.State = uint8(enum.CMD_STATE_PROCESSING)
	// cmd.TxId = summary
	// cmd.ExtEthGasprice = gasPrice
	// cmd.ExtTrxTx = extInfo
	// if _, err := txObj.Update(cmd, "state", "tx_id", "ext_eth_gasprice", "ext_trx_tx"); err != nil {
	// 	panic(err)
	// }
	// return item, nil
}

// 向external_client回复cmd
func (self *ExternalService) PullCmdFromExternalServerReply(ctx context.Context, in *external_data.PullCmdFromExternalServerRequest) (reply *external_data.PullCmdFromExternalServerReply, ret error) {
	logs.Debug(fmt.Sprintf("receive cmd pull request, params: %+v", in))
	dbObj := orm.NewOrm()
	transferType := in.TransferType
	if transferType == "" {
		transferType = "tss"
	}
	//查询待处理的记录
	var cmds []*models.IwalaTssTransferCmd
	if !dbObj.QueryTable("iwala_tss_transfer_cmd").Filter("state", enum.CMD_STATE_PROCESSING).Filter("transfer_type", transferType).Filter("update_time__gt", time.Now().Add(-10*time.Minute)).Exist() {
		if _, err := dbObj.QueryTable("iwala_tss_transfer_cmd").Filter("state", enum.CMD_STATE_UNPROCESS).Filter("transfer_type", transferType).OrderBy("id").All(&cmds); err != nil {
			return nil, err
		}
	}
	var items []*external_data.ExternalServerCmdItem
	for _, cmd := range cmds {
		item, err := self.handleCmd(cmd.Id)
		if err == nil && item != nil {
			items = append(items, item)
		}
		//一次只处理一条数据
		break
	}
	return &external_data.PullCmdFromExternalServerReply{
		Code: &external_data.ExternalCommonCode{
			Code: config.SUCCESS,
		},
		Items: items,
	}, nil
}

var (
	ToAddressIdMap map[string]int
)

func init() {
	ToAddressIdMap = make(map[string]int)
}

func UpdateAndCommit(dbObj orm.TxOrmer, md interface{}, cols ...string) {
	if _, err := dbObj.Update(md, cols...); err != nil {
		panic(err)
	}
	if err := dbObj.Commit(); err != nil {
		panic(err)
	}
}

func (self *ExternalService) SaveResult(item *external_data.ExternalServerCmdResult) (ret error) {
	var txObj orm.TxOrmer
	dbObj := orm.NewOrm()
	if txObj, ret = dbObj.Begin(); ret != nil {
		logs.Error("save result, begin transaction error.", ret)
		return ret
	}

	defer func() {
		if errs := recover(); errs != nil {
			if err := txObj.Rollback(); err != nil {
				ret = fmt.Errorf("%+v, rollback error: %s, %s", errs, err.Error(), string(debug.Stack()))
			} else {
				ret = fmt.Errorf("%+v, %s", errs, string(debug.Stack()))
			}
		}
	}()

	// 拉原cmd
	transfer := new(models.IwalaTssTransferCmd)
	if err := txObj.QueryTable("iwala_tss_transfer_cmd").Filter("id", item.Id).ForUpdate().One(transfer); err != nil {
		panic(err)
	}
	if transfer.State == uint8(enum.CMD_STATE_SUCCESS) && transfer.TxId != "" {
		panic(fmt.Errorf("the transaction has been successful before, so you can't retransfer it again, %+v", transfer))
	}

	// 签名失败
	if item.Code.Code != config.SUCCESS {
		transfer.State = uint8(enum.CMD_STATE_ERROR)
		transfer.Err = item.Code.Err
		UpdateAndCommit(txObj, transfer, "state", "err")
		return
	}

	// 额外安全检查
	if id, ok := ToAddressIdMap[transfer.ToAddress]; ok && transfer.Id <= id {
		transfer.State = uint8(enum.CMD_STATE_ERROR)
		transfer.Err = fmt.Sprintf("critical error memory, last is %d", id)
		UpdateAndCommit(txObj, transfer, "state", "err")
		return
	} else if id, err := ori_redis.Int64(redis.Do("7", "hget", "tss_id", transfer.ToAddress)); err == nil && int64(transfer.Id) <= id {
		transfer.State = uint8(enum.CMD_STATE_ERROR)
		transfer.Err = fmt.Sprintf("critical error redis, last is %d", id)
		UpdateAndCommit(txObj, transfer, "state", "err")
		return
	}

	// 签名成功

	// 还原thresholdCoin
	param := apis.CoinParam{
		Coin:           transfer.CurrencyMark,
		Chain:          transfer.Chain,
		From:           transfer.FromAddress,
		To:             transfer.ToAddress,
		Value:          pubUtils.MustDecimal(transfer.Amount),
		ExtTrxTx:       transfer.ExtTrxTx,
		ExtEthGasPrice: transfer.ExtEthGasprice,
		Memo:           transfer.Memo,
	}
	thresholdCoin, _, _, err := apis.NewThresholdCoin(param)
	if err != nil {
		transfer.State = uint8(enum.CMD_STATE_ERROR)
		transfer.Err = fmt.Sprintf("new threshold obj error, %s", err.Error())
		UpdateAndCommit(txObj, transfer, "state", "err")
		return
	}

	// 附加签名
	transfer.TxSign = item.Txsign // 下面的都需要update tx_sign字段
	if extInfo, err := thresholdCoin.WithSignature(item.Txsign); err != nil {
		transfer.State = uint8(enum.CMD_STATE_ERROR)
		transfer.Err = fmt.Sprintf("WithSignature error, %s", err.Error())
		UpdateAndCommit(txObj, transfer, "tx_sign", "state", "err")
		return
	} else {
		transfer.ExtTrxTx = extInfo // 下面的都需要update ext_trx_tx字段
	}

	// 是否需要继续签名
	if thresholdCoin.NeedSign() {
		transfer.State = uint8(enum.CMD_STATE_UNPROCESS)
		UpdateAndCommit(txObj, transfer, "tx_sign", "ext_trx_tx", "state")
		return
	}

	// 不需要继续签名的则开始广播
	txId, err := thresholdCoin.Broadcast()
	if err != nil {
		transfer.State = uint8(enum.CMD_STATE_ERROR)
		transfer.Err = fmt.Sprintf("broadcast error, %s", err.Error())
		UpdateAndCommit(txObj, transfer, "tx_sign", "ext_trx_tx", "state", "err")
		return
	}

	transfer.TxId = txId
	transfer.State = uint8(enum.CMD_STATE_SUCCESS)

	// 广播成功后清数据
	//transfer.ExtTrxTx = ""
	//transfer.Sign = ""

	ToAddressIdMap[transfer.ToAddress] = transfer.Id
	if _, err := redis.Do("7", "hset", "tss_id", transfer.ToAddress, transfer.Id); err != nil {
		logs.Error("set tss_id error", err)
	}

	UpdateAndCommit(txObj, transfer, "tx_sign", "ext_trx_tx", "tx_id", "state", "sign")
	return
}

func (self *ExternalService) PushResultToExternalServerReply(ctx context.Context, in *external_data.PushResultToExternalServerRequest) (*external_data.PushResultToExternalServerReply, error) {
	logs.Info(fmt.Sprintf("receive cmd execution results, params: %+v", in))
	for _, item := range in.Items {
		if err := self.SaveResult(item); err != nil {
			logs.Error(err)
		}
	}
	return &external_data.PushResultToExternalServerReply{
		Code: &external_data.ExternalCommonCode{
			Code: config.SUCCESS,
		},
	}, nil
}
