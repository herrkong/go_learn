package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	//"github.com/skycoin/skycoin/src/transaction"
)


type Eth struct {
	Url      string
	From     string
	To       string
	Value    decimal.Decimal
	GasPrice uint64 // 默认的gasPrice
	GasLimit uint64 // 默认的gasLimit
	ExtTrxTx string // 通用存储data
	tx       *types.Transaction
}



func main(){
	extinfo := "7b2274797065223a22307830222c226e6f6e6365223a22307830222c226761735072696365223a22307831373438373665383030222c226d61785072696f72697479466565506572476173223a6e756c6c2c226d6178466565506572476173223a6e756c6c2c22676173223a22307835323038222c2276616c7565223a22307833373832646163653964393030303030222c22696e707574223a223078222c2276223a2230783236222c2272223a22307864303734326439623635623338356638663866666237396136623366386562656261313833666230623563303365393133386261333037383138613234313035222c2273223a22307835653335653761366166623035663133306464393163346138623530636331396433383430353133363535376138356363383434366432353835613031626365222c22746f223a22307864363537303032336538393465373733303632383630363831646133366333663965303231356261222c2268617368223a22307862373733346633663533326365653633353561326134383335386238303564363436343038326334656330373563373964326561663134333461333533656466227d"
	transaction,err:= fromString(extinfo)
	if err !=nil{
		fmt.Printf("get transaciton error")
	}
	fmt.Printf("transaction=%v\n",transaction)


}



func fromString(extInfo string) (*types.Transaction, error) {
	data, err := hex.DecodeString(extInfo)
	if err != nil {
		return nil, err
	}
	var tx *types.Transaction
	if err := json.Unmarshal(data, &tx); err != nil {
		return nil, err
	}
	return tx, nil
}