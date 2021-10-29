package main


import (
	"fmt"
	"github.com/shopspring/decimal"
)


func main(){
	balance := decimal.NewFromFloat(float64(588256821.55561516))
	fee:= decimal.NewFromFloat(8.932)
	
	outAll := decimal.NewFromFloat(0)
	changeDec := balance.Sub(fee).Sub(outAll)

	change, _ := changeDec.Float64()
	changeDecRecover := decimal.NewFromFloat(change)

	if !changeDec.Equal(changeDecRecover){
		fmt.Printf("changeDec(%v)!=changeDecRecover(%v)\n",changeDec,changeDecRecover)
	}

	fmt.Printf("balance=%v,fee=%v,change=%v,changeDec=%v,changeDecRecover=%v\n",balance.String(),fee,change,changeDec,changeDecRecover)
}