package main

import (
	_ "threshold/src/public/log"
	_ "threshold/src/external_client/db"
	pub_config "threshold/src/public/config"
	"threshold/src/external_client/config"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/client/orm"
	"github.com/farmerx/gorsa"
	"os"
	"threshold/src/public/client"
	"threshold/src/public/data/external_data"
	"threshold/src/public/enum"
	"threshold/src/public/lark"
	"threshold/src/public/models"
	"time"
	"flag"
	mysql "threshold/src/external_client/db"
	"threshold/src/public/utils"
	"os/exec"
	"strings"
)

//签名校验
func CheckExternalCmdSign(item *external_data.ExternalServerCmdItem) error {
	if err := gorsa.RSA.SetPublicKey(config.Rsa.ExternalPubkey); err != nil {
		return err
	}
	sign, err := hex.DecodeString(item.Sign)
	if err != nil {
		return err
	}
	raw, err := gorsa.RSA.PubKeyDECRYPT(sign)
	if err != nil {
		return err
	}
	preSign := fmt.Sprintf("%d|%s|%s", item.Id, item.To, utils.MustDecimal(item.Amount).String())//数量以decimal转string为准
	logs.Debug("preSign:%s", preSign)
	if string(raw) != preSign {
		logs.Error(fmt.Sprintf("new sign not match, expect:%s, got:%s", preSign, string(raw)))
		// 兼容老的sign，todo:完全上高精度之后，这个逻辑需要去掉
		oldPreSign := fmt.Sprintf("%d|%s|%.8f", item.Id, item.To, utils.MustFloat(item.Amount))
		logs.Debug("oldPreSign:%s", oldPreSign)
		if string(raw) != oldPreSign {
			logs.Error(fmt.Sprintf("old sign not match, expect:%s, got:%s", oldPreSign, string(raw)))
			return errors.New("sign not match")
		}else{
			logs.Info("old sign check pass, preSign:%s", oldPreSign)
		}
	}else{
		logs.Info("new sign check pass, preSign:%s", preSign)
	}

	return nil
}

//风控检查
func CheckRisk(item *external_data.ExternalServerCmdItem) (bool, string) {
	//笔数、额度、mitm、防证书篡改
	return false, "success"
}

//计算签名
func CalcManagerCmdSign(id int64, summary string) (string, error) {
	if err := gorsa.RSA.SetPrivateKey(config.Rsa.ManagerPrikey); err != nil {
		return "", err
	}
	raw := fmt.Sprintf("%d|%s", id, summary)
	logs.Debug("CalcManagerCmdSign, preSign:%s", raw)
	if bytes, err := gorsa.RSA.PriKeyENCTYPT([]byte(raw)); err != nil {
		logs.Error("CalcManagerCmdSign, err:%+v", err)
		return "", err
	} else {
		return hex.EncodeToString(bytes), nil
	}
}

//外部指令存储
func SaveCmd(countryType pub_config.CountryType, item *external_data.ExternalServerCmdItem) (int64, error) {
	cmdInfo := map[string]interface{}{
		"coin":   item.Coin,
		"from":   item.From,
		"to":     item.To,
		"amount": item.Amount,
	}
	jsonData, err := json.Marshal(cmdInfo)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	dbObj := orm.NewOrm()
	switch enum.TransferType(item.TransferType) {
	case enum.TSS:
		cmd := &models.IwalaTssManagerCmd{
			Country:       string(countryType),
			ExternalCmdId: int(item.Id),
			Type:          int(enum.CMD_TYPE_SIGN),
			Finish:        int(enum.CMD_STATE_UNFINISHED),
			Summary:       item.Summary,
			Chain:         int(item.TssChain),
			From:          item.From,
			Curve:         int(item.Curve),
			PublicKey:     item.PublicKey,
			T:             config.ManagerEnv.PeerT,
			N:             config.ManagerEnv.PeerN,
			JsonData:      string(jsonData),
		}
		id, err := dbObj.Insert(cmd)
		if err != nil {
			logs.Error(err)
			return 0, err
		}
		sign, err := CalcManagerCmdSign(id, cmd.Summary)
		if err != nil {
			logs.Error(err)
			return 0, err
		}
		cmd.Sign = sign
		return dbObj.Update(cmd, "sign")
	case enum.SSS:
		cmd := &models.IwalaSssManagerCmd{
			Country:       string(countryType),
			ExternalCmdId: int(item.Id),
			Type:          int(enum.CMD_TYPE_SIGN),
			Finish:        int(enum.CMD_STATE_UNFINISHED),
			Summary:       item.Summary,
			Chain:         int(item.SssChain),
			Curve:         int(item.Curve),
			From:          item.From,
			PublicKey:     item.PublicKey,
			T:             config.ManagerEnv.PeerT,
			N:             config.ManagerEnv.PeerN,
			JsonData:      string(jsonData),
		}
		id, err := dbObj.Insert(cmd)
		if err != nil {
			logs.Error(err)
			return 0, err
		}
		sign, err := CalcManagerCmdSign(id, cmd.Summary)
		if err != nil {
			logs.Error(err)
			return 0, err
		}
		cmd.Sign = sign
		return dbObj.Update(cmd, "sign")
	default:
		return 0, fmt.Errorf("unknow type, %s", item.TransferType)
	}

}

//rsa错误信息存储
func SaveRsaError() (int64, error) {
	dbObj := orm.NewOrm()
	rsaError := &models.IwalaRsaError{
		Program: "external",
	}
	//发送lark告警
	if err := lark.SendText(config.LarkEnv.LarkUrl, config.LarkEnv.RobotId, "rsa校验错误", "ras校验错误，查看日志"); err != nil {
		logs.Error(err)
	}
	return dbObj.Insert(rsaError)
}

func QueryRsaError() {
	if pub_config.Program.Mode != pub_config.PROD {
		return
	}
	var rsaErrors []*models.IwalaRsaError
	dbObj := orm.NewOrm()
	if count, err := dbObj.QueryTable("iwala_rsa_error").Filter("program", "external").All(&rsaErrors); err != nil {
		logs.Error("query rsa error.", err)
		os.Exit(-1)
	} else if count >= 2 {
		logs.Error("count >= 2, exit")
		os.Exit(-1)
	}
}

func PushErrorToExternalServer(c client.ManagerClient, cmdItem *external_data.ExternalServerCmdItem, err string) {
	logs.Error(err)
	var resultItems []*external_data.ExternalServerCmdResult

	resultItems = append(resultItems, &external_data.ExternalServerCmdResult{
		Id: cmdItem.Id,
		Code: &external_data.ExternalCommonCode{
			Code: pub_config.ERROR,
			Err:  err,
		},
	})
	if _, err := c.PushResultToExternalServer(resultItems); err != nil {
		logs.Error(err)
	}
}

//拉取外部server的转账指令
func PullCmdFromExternalServer(countryType pub_config.CountryType, serverInfo pub_config.ServerInfo, coinType int64) {
	defer func() {
		if errs := recover(); errs != nil {
			logs.Error(errs)
		}
	}()

	c := client.ManagerClient{
		CountryType: countryType,
		ServerInfo:  serverInfo,
		CoinType:    coinType,
	}
	reply, err := c.PullCmdFromExternalServer(string(config.ManagerEnv.TransferType))
	if err != nil {
		logs.Error(fmt.Sprintf("pull cmd from external server error, %s", err.Error()))
		return
	}
	for _, item := range reply.Items {
		if err = CheckExternalCmdSign(item); err != nil {
			//签名校验错误，返回指令发起人结果
			PushErrorToExternalServer(c, item, fmt.Sprintf("check sign error, %s", err.Error()))
			if _, err := SaveRsaError(); err != nil {
				logs.Error("save rsa error", err)
			}
			QueryRsaError()
			continue
		}
		riskRet, riskStr := CheckRisk(item)
		if riskRet {
			//出发风控，返回指令发起人结果
			logs.Error("trigger risk, %s", riskStr)
			PushErrorToExternalServer(c, item, fmt.Sprintf("trigger risk, %s", riskStr))
			continue
		}
		if _, err := SaveCmd(countryType, item); err != nil {
			logs.Error("save cmd %+v error, %s", item, err.Error())
			PushErrorToExternalServer(c, item, fmt.Sprintf("save cmd %+v error, %s", item, err.Error()))
			continue
		}
	}
}

//推送结果到外部server
func PushResultToExternalServer(countryType pub_config.CountryType, serverInfo pub_config.ServerInfo, coinType int64) {
	defer func() {
		if errs := recover(); errs != nil {
			logs.Error(errs)
		}
	}()

	c := client.ManagerClient{
		CountryType: countryType,
		ServerInfo:  serverInfo,
		CoinType:    coinType,
	}
	//开始处理tss转账
	var tssCmds []*models.IwalaTssManagerCmd
	dbObj := orm.NewOrm()
	if _, err := dbObj.QueryTable("iwala_tss_manager_cmd").Filter("finish", enum.CMD_STATE_FINISHED).Filter("country", countryType).All(&tssCmds); err != nil {
		logs.Error("query finished tss cmd error.", err)
		return
	}
	for _, cmd := range tssCmds {
		logs.Info(fmt.Sprintf("tss begin push result, %+v", cmd))
		//结果一条一条发送
		var items []*external_data.ExternalServerCmdResult
		code := pub_config.ERROR
		if cmd.TxSign != "" {
			code = pub_config.SUCCESS
		}
		items = append(items, &external_data.ExternalServerCmdResult{
			Id: int64(cmd.ExternalCmdId),
			Code: &external_data.ExternalCommonCode{
				Code: int64(code),
				Err:  cmd.Err,
			},
			Txsign: cmd.TxSign,
		})
		//更新状态为准备推送
		cmd.Finish = int(enum.CMD_STATE_START_PUSH)
		if _, err := dbObj.Update(cmd, "finish"); err != nil {
			logs.Error(fmt.Errorf("tss update finish state %d error, %s", cmd.Finish, err.Error()))
			continue
		}
		//推送
		if _, err := c.PushResultToExternalServer(items); err != nil {
			logs.Error(fmt.Sprintf("tss push result to external server error, %s", err.Error()))
			//更新状态为推送失败
			cmd.Finish = int(enum.CMD_STATE_PUSH_ERROR)
		} else {
			logs.Info(fmt.Sprintf("tss push result to external server success"))
			//更新状态为推送成功
			cmd.Finish = int(enum.CMD_STATE_PUSH_SUCCESS)
		}
		if _, err := dbObj.Update(cmd, "finish"); err != nil {
			logs.Error(fmt.Errorf("tss update cmd finish state %d error, %s", cmd.Finish, err.Error()))
		}
	}
	//开始处理sss转账
	var sssCmds []*models.IwalaSssManagerCmd
	if _, err := dbObj.QueryTable("iwala_sss_manager_cmd").Filter("finish", enum.CMD_STATE_FINISHED).Filter("country", countryType).All(&sssCmds); err != nil {
		logs.Error("query finished sss cmd error.", err)
		return
	}
	for _, cmd := range sssCmds {
		logs.Info(fmt.Sprintf("sss begin push result, %+v", cmd))
		//结果一条一条发送
		var items []*external_data.ExternalServerCmdResult
		code := pub_config.ERROR
		if cmd.TxSign != "" {
			code = pub_config.SUCCESS
		}
		items = append(items, &external_data.ExternalServerCmdResult{
			Id: int64(cmd.ExternalCmdId),
			Code: &external_data.ExternalCommonCode{
				Code: int64(code),
				Err:  cmd.Err,
			},
			Txsign: cmd.TxSign,
		})
		//更新状态为准备推送
		cmd.Finish = int(enum.CMD_STATE_START_PUSH)
		if _, err := dbObj.Update(cmd, "finish"); err != nil {
			logs.Error(fmt.Errorf("sss update finish state %d error, %s", cmd.Finish, err.Error()))
			continue
		}
		//推送
		if _, err := c.PushResultToExternalServer(items); err != nil {
			logs.Error(fmt.Sprintf("sss push result to external server error, %s", err.Error()))
			//更新状态为推送失败
			cmd.Finish = int(enum.CMD_STATE_PUSH_ERROR)
		} else {
			logs.Info(fmt.Sprintf("sss push result to external server success"))
			//更新状态为推送成功
			cmd.Finish = int(enum.CMD_STATE_PUSH_SUCCESS)
		}
		if _, err := dbObj.Update(cmd, "finish"); err != nil {
			logs.Error(fmt.Errorf("sss update cmd finish state %d error, %s", cmd.Finish, err.Error()))
		}
	}
}

func main() {

	var mode string
	flag.StringVar(&mode, "mode", "", "run mode")
	flag.Parse()
	if mode != `tss` && mode !=`sss` {
		logs.Critical(`run mode error, mode:%s`, mode)
		os.Exit(0)
	}
	config.ManagerEnv.TransferType = enum.TransferType(mode)

	mysql.InitMysql()

	QueryRsaError()
	exit := make(chan int)

	// Lark通报启动情况
	exeFile, _ := exec.LookPath(os.Args[0])
	fileInfo, _ := os.Stat(exeFile)
	_ = lark.SendText(config.LarkEnv.LarkUrl, config.LarkEnv.RobotId, fileInfo.Name()+"启动", strings.Join(os.Args, ` `))

	go func() {
		//定时触发
		pullTimer := time.NewTicker(config.ManagerEnv.PullCommandInterval)
		defer pullTimer.Stop()
		pushTimer := time.NewTicker(config.ManagerEnv.PushResultInterval)
		defer pushTimer.Stop()
		for {
			select {
			case <-pullTimer.C:
				//定时拉取外部服务器指令
				for countryType, serverInfo := range pub_config.ExternalServerEnv.Server {
					PullCmdFromExternalServer(countryType, serverInfo, 1)
				}
			case <-pushTimer.C:
				//定时检查结果，推送到外部server
				for countryType, serverInfo := range pub_config.ExternalServerEnv.Server {
					PushResultToExternalServer(countryType, serverInfo, 1)
				}
			}
		}
	}()
	<-exit
}

// trigger ci
