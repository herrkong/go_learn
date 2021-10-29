// 00000000000000000000000a688906bd8b00000

package main

import(
	"fmt"
	"math"
	"math/big"
	//"strings"
	"github.com/shopspring/decimal"
)



func main(){

	//0xce550ddfdec6d0000
	//ce550ddfdec6d0000
	//00000000000000000000000ce550ddfdec6d0000
	tempString := "00000000000000000000000ce550ddfdec6d0000"
	val := new(big.Int)
	val.SetString(tempString, 16)
	
	tmpval := decimal.NewFromBigInt(val, 0)
	digit := decimal.NewFromFloat(math.Pow10(18))
	dec := decimal.NewFromFloat(math.Pow10(8))
	tmpNum := tmpval.Div(digit)
	tmpNum = tmpNum.Mul(dec).Floor()
	value := tmpNum.Div(dec).String()

	fmt.Println("value=",value)


}