package main

import (
	"fmt"
	//"math/big"
	"math"
	"github.com/shopspring/decimal"
)


func main(){

	// balance_str := "0x1e5b308931e55df689be00"
	
	// fmt.Printf("get_str=%s\n",balance_str[2:])

	// val := new(big.Int)
	// val.SetString(balance_str[2:], 16)
	// balance_decimal := decimal.NewFromBigInt(val, 0)
	
	// digit := decimal.NewFromFloat(math.Pow10(18))
	// balance := balance_decimal.Div(digit)

	// fmt.Printf("balance=%v\n",balance)



	balance_str := "0x1e5b308931e55df689be00"

	balance_decimal,err := decimal.NewFromString(balance_str)
	if err != nil{
		fmt.Printf("balance_decimal error")
	}
	digit := decimal.NewFromFloat(math.Pow10(18))
	balance := balance_decimal.Div(digit)
	fmt.Printf("balance=%v\n",balance)


}