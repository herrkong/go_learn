package main

import(
	"fmt"
	"crypto/sha256"
	"github.com/shengdoushi/base58"
)

func main(){
	ret_byte,err := DecodeCheck("TVj9YmcwmQ4TSTRcpaYosn6dXcpidj8g4g")
	if err != nil{
		fmt.Printf("%s\n",err)
	}
	fmt.Printf("ret_byte=%v\n",ret_byte)

}


func DecodeCheck(input string) ([]byte, error) {
	decodeCheck, err := Decode(input)

	if err != nil {
		return nil, err
	}

	if len(decodeCheck) < 4 {
		return nil, fmt.Errorf("b58 check error")
	}

	decodeData := decodeCheck[:len(decodeCheck)-4]

	h256h0 := sha256.New()
	h256h0.Write(decodeData)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	if h1[0] == decodeCheck[len(decodeData)] &&
		h1[1] == decodeCheck[len(decodeData)+1] &&
		h1[2] == decodeCheck[len(decodeData)+2] &&
		h1[3] == decodeCheck[len(decodeData)+3] {
		return decodeData, nil
	}
	return nil, fmt.Errorf("b58 check error")
}


func Decode(input string) ([]byte, error) {
	return base58.Decode(input, base58.BitcoinAlphabet)
}