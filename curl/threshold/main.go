package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

func main(){

	//msg := "d8c31f91a178a475e55d8e775f58ae4ae552f8d1abc286066ce5e38df089a07f"
	r_str := "18166797309966579161800646903866802141440892938426307975805173284433762670563"
	s_str := "64813725468785706509270427409786231621391982271097046604578969891811131620247"
	v_str := "0"

	R, _ := new(big.Int).SetString(r_str, 10)
	S, _ := new(big.Int).SetString(s_str, 10)
	V, _ := new(big.Int).SetString(v_str, 10)


	var sign []byte
	sign = append(sign, common.LeftPadBytes(R.Bytes(), 32)...)
	sign = append(sign, common.LeftPadBytes(S.Bytes(), 32)...)
	if V.Sign() == 0 {
		sign = append(sign, byte(0))
	} else {
		sign = append(sign, byte(1))
	}
	// 存[]byte的16进制字符串
	TxSign := hex.EncodeToString(sign)

	fmt.Printf("tx_sign=%v\n",TxSign)


	// tx_sign = 282a0afb344647554da4f59e11bb0abab760605a8f28200b4d7fa247235b57e38f4b44c393ee940b5725914665d1352b04b94111d69d83b548dc7efc70577b9700

	// hex_string := "0269a7ef20a5e745a5cc3b8fe12e44351752fe4ed58be5ae49e8667ee18befa6f3aecce7f895a090693629c71162cfb1438191cf0426cd5f5883c2645a4be15000"

	// tx_sign,err := hex.DecodeString(hex_string)

	// if err != nil{
	// 	fmt.Printf("err=%v\n",err)
	// }

	// fmt.Printf("tx_sign=%v\n",tx_sign)


}
