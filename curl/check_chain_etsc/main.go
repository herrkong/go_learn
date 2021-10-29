package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

type EtscData struct {
	Hash    string `json:"hash"`
	Value   float64 `json:"value"`
	ToAddress string `json:"to"`
	FromAddress string `json:"from"`
}

type EtscChainData struct {
	Error    int        `json:"hash"`
	Data     []EtscData    `json:"data"`
}


func main() {
	hashPush := "0xbc3b3fc58270f2d60e0fce0f5633acb73f3420caf8434f89c2380b26d7795429"
	pushVal := float64(1)
	pushAddress := "0xD6570023e894E773062860681da36c3f9e0215bA"
	checkUrl := "https://api.etscscan.com/client/explore/transaction?hash=0xbc3b3fc58270f2d60e0fce0f5633acb73f3420caf8434f89c2380b26d7795429&type=detail"
	//https://api.etscscan.com/client/explore/transaction?hash=0xbc3b3fc58270f2d60e0fce0f5633acb73f3420caf8434f89c2380b26d7795429&type=detail
	EtscChainData, err := EtscChainGet(checkUrl)
	if err != nil {
		//return status.STATUS_CHECK_CHAIN_FAILED, fmt.Sprintf("链上检查接口请求失败：%s，错误信息：%s", checkUrl, err.Error())
		fmt.Printf("err=%v\n",err)
	}

	error_code := EtscChainData.Error
	 if error_code!=0{
		//return status.STATUS_CHECK_CHAIN_FAILED, fmt.Sprintf("链上检查接口请求失败：%s，error_code=%d != 0", checkUrl, error_code)
		fmt.Printf("error_code=%v\n",error_code)
	}

	// if len(EtscChainData.Data) == 0 {
	// 	return status.STATUS_CHECK_CHAIN_FAILED, fmt.Sprintf("链上检查接口请求失败：%s，错误信息：%s", checkUrl, "no data")
	// }
	etsc_data := EtscChainData.Data[0]

	result_hash := etsc_data.Hash

	result_value := etsc_data.Value

	result_address := etsc_data.ToAddress

	fmt.Printf("result_hash=%s,result_value=%v,result_address=%s\n",result_hash,result_value,result_address)

	// resultInfo, _ := json.Marshal(EtscChainData)
	// resultJson := gjson.Parse(string(resultInfo))
	// //logs.Info(fmt.Sprintf("Etsc resultJson= %v\n", resultJson))
	// error_code := resultJson.Get("error_code").Int()
	// if error_code!=0{
	// 	return status.STATUS_CHECK_CHAIN_FAILED, fmt.Sprintf("链上检查接口请求失败：%s，error_code=%d != 0", checkUrl, error_code)
	// }
	// result_data:= resultJson.Get("data").Array()[0]
	// result_hash := result_data.Get("hash").String()
	// result_value := result_data.Get("value").Float()
	// result_address := result_data.Get("to").String()
	
	if !strings.EqualFold(result_hash,hashPush){
		//return status.STATUS_CHECK_CHAIN_FAILED, fmt.Sprintf("链上检查请求%s，数据 hash：%s 与 链上 hash %s 不相等", checkUrl, hashPush, result_hash)
		fmt.Printf("hash no the same")
	}

	if result_address == pushAddress && math.Abs(pushVal-result_value) < float64(0.000001) {
		//return status.STATUS_CHECK_CHAIN_PASS, "SUCCESS"
		fmt.Printf("amount error\n")
	}
	// return status.STATUS_CHECK_CHAIN_FAILED, fmt.Sprintf("链上检查接口请求失败：%s，充币地址：%s，充币金额%f;链上地址:%s,链上金额%f;不相等。",
	// 	checkUrl, pushAddress, pushVal, result_address, result_value)
	fmt.Printf("check success\n")
}

func EtscChainGet(url string) (Result *EtscChainData , err error) {
	if response, err := http.Get(url); err != nil {
		logs.Error("error occur when get the Etsc chain result!", err)
		return nil, err
	} else {
		if Resultdata, err := ioutil.ReadAll(response.Body); err != nil {
			logs.Error("error occur when read the chain result body!", err)
			return nil, err
		} else {
			Result_ptr := &EtscChainData{}
			err = json.Unmarshal(Resultdata, Result_ptr)
			return Result_ptr, err
		}
	}
}
