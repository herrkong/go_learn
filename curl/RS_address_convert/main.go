package main

// Reed Solomon Encoding and Decoding for Apl and other rs-address-coin

import (
	"fmt"
	"crypto/sha256"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"math/big"
	//"github.com/astaxie/beego/httplib"
)

var (
	CoinType string
	Base_32_length int
	Base_10_length int
	Alphabet string
	Initial_codeword []int
	Gexp  []int
	Glog  []int
	Codeword_map []int
)

type Account struct{
	

}

type GetEcBlockIdResponse struct {
	EcBlockId     		string `json:"ecBlockId"`
	EcBlockHeight       string `json:"ecBlockHeight"`
}


type TransactionObject struct{
	EncryptedMessageIsPrunable  bool       `json:"encryptedMessageIsPrunable"`
	MessageIsPrunable  bool                `json:"messageIsPrunable"`
	Phased  bool                           `json:"phased"`
	EcBlockId   string                     `json:"ecBlockId"`
	Broadcast   bool                       `json:"broadcast"`
	DeadlineValue uint32                   `json:"deadlineValue"`
	ReferencedTransactionFullHash  string  `json:"referencedTransactionFullHash"`
	PublicKeyValue  string                 `json:"publicKeyValue"`
	AmountATM string                       `json:"amountATM"`
	FeeATM string                          `json:"feeATM"`
	SenderAccount string                   `json:"senderAccount"`
	RecipientId   string                   `json:"recipientId"`
	RecipientPublicKey string              `json:"recipientPublicKey"`
	Attachment  string                     `json:"attachment"`
	Validate    bool                       `json:"validate"`
}


// secretphrase : 0bc92158a6bfdf7f5f9f4f2c8f92f64a
// RS account : APL-R86T-Q6VH-R4VZ-45SGW
// account id : 3331260409037166745
// publicKey  : 14f7f751c31e0335058aa0c8ca80a5b282e0959d16de4524ec34ea529d23eec2



// secretphrase : aaf60a5aa5a11cf5882938131355b5a2
// RS account : APL-JVXK-XLZY-T3UP-G9U7K
// account id : 16759943620440125361
// publicKey  : 7473bf8635be3a48382f009058252a37f6c8fdb2ddb64f3fb22f122c3018c7d1


func main(){

	// generate apl key pair
	//privkey,pubkey,err := NewKeyPair()
	seed := "0bc92158a6bfdf7f5f9f4f2c8f92f64a"
	//seed := "aaf60a5aa5a11cf5882938131355b5a2"
	privkey,pubkey,err := NewKeyPairFromSeed(seed)

	if err != nil{
		fmt.Printf("Generate Key Pair failed!\n")
	}
	fmt.Printf("privateKey=%v\n,publicKey=%v\n",privkey,pubkey)

	// pubkey2account_id
	convert_account_id := Pubkey2Account(pubkey)

	fmt.Printf("convert_account_id=%v\n",convert_account_id)

	// PrivateKey2SecretPhrase
	// SecretPhrase := PrivateKey2SecretPhrase(privkey)

	// fmt.Printf("SecretPhrase=%v\n",SecretPhrase)
	// NewKeyPairFromSeed



	// set accountid
	//accountid := "18093681334230470778"
	//accountid:= "10245873926842220553"
	//input_rs_account := "APL-W22B-3NRF-24TK-AA7EW"

	accountid := convert_account_id

	// set apl parameter
	CoinType = "APL-"
	Initial_codeword = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	Gexp = []int{1, 2, 4, 8, 16, 5, 10, 20, 13, 26, 17, 7, 14, 28, 29, 31, 27, 19, 3, 6, 12, 24, 21, 15, 30, 25, 23, 11, 22, 9, 18, 1}
	Glog = []int{0, 0, 1, 18, 2, 5, 19, 11, 3, 29, 6, 27, 20, 8, 12, 23, 4, 10, 30, 17, 7, 22, 28, 26, 21, 25, 9, 16, 13, 14, 24, 15};
	Codeword_map = []int{3, 2, 1, 0, 7, 6, 5, 4, 13, 14, 15, 16, 12, 8, 9, 10, 11}
	Base_32_length = 13
    Base_10_length = 20
	Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"

	//fmt.Printf("initial_codeword=%v\n,gexp=%v\n,codeword_map=%v\n,base_32_length=%v\n,base_10_length=%v\n,alphabet=%v\n",Initial_codeword,Gexp,Codeword_map,Base_32_length,Base_10_length,Alphabet)

	//fmt.Printf("coin_type=%s,account_id=%v\n",CoinType,accountid)
	rs_address := AccountId2RsAddress(accountid)
	fmt.Printf("rs_address=%s\n",rs_address)

	// if rs_address == input_rs_account{
	// 	fmt.Printf("check success!\n")
	// }

	senderAccount := "APL-R86T-Q6VH-R4VZ-45SGW"
	recipientId := "16759943620440125361"
	attachment := "OrdinaryPayment"
	amountATM := "100000000"
	broadcast := true
	validate := true
	publicKey := pubkey
	recipientPublicKey := "7473bf8635be3a48382f009058252a37f6c8fdb2ddb64f3fb22f122c3018c7d1"
	
	unsignedTransactionBytes := Createtransaction(senderAccount,recipientId,amountATM,attachment,broadcast,validate,publicKey,recipientPublicKey)

	fmt.Printf("unsignedTransactionBytes=%v\n",unsignedTransactionBytes)



}

func SetCoinType(coin_type string){
	CoinType = coin_type
}


func GetCoinType() string{
	return CoinType
}


func PrivateKey2SecretPhrase(PrivateKey string) (SecretPhrase []string ){





	return SecretPhrase
}

func AccountId2RsAddress(accountId string) string {
	return GetCoinType() + RSEncode(accountId)
}

//todo
func RSAddress2AccountId(rs string) (accountId float64){
	return accountId
}

//todo
func RSEncode(accountId string) (rs_address string){
	plain_string := accountId
	length := len(plain_string)
	plain_string_10 := make([]int,Base_10_length)

	for i:=0;i< length;i++{
		plain_string_10[i] = int(plain_string[i]) - (int)('0')
	}

	codeword_length:=0
	codeword := make([]int,len(Initial_codeword))

	//base 10 to base 32 conversion
	finish_flag := false
	for !finish_flag{
		new_length:=0
		digit_32:=0
		for i:=0;i< length;i++{
			digit_32 = digit_32 * 10 + plain_string_10[i]
			if (digit_32 >= 32) {
				plain_string_10[new_length] = digit_32 >> 5
				digit_32 &= 31
				new_length += 1
			}else if(new_length > 0){
				plain_string_10[new_length] = 0
				new_length += 1
			}
		}
		length = new_length
		codeword[codeword_length] = digit_32
		codeword_length += 1
		if length <= 0{
			finish_flag = true
		}
	}

	p := make([]int,4)
	for i := Base_32_length - 1;i >=0 ;i--{
		fb := codeword[i] ^ p[3]
		p[3] = p[2] ^ Gmult(30, fb)
		p[2] = p[1] ^ Gmult(6, fb)
		p[1] = p[0] ^ Gmult(9, fb)
		p[0] = Gmult(17, fb)
	}

	copy(codeword[13:],p[0:(len(Initial_codeword)-Base_32_length)])

	cypher_builder := make([]string,20)

	for i:=0;i <17 ; i++{
		codework_index:= Codeword_map[i]
		alphabet_index := codeword[codework_index]
		cypher_builder = append(cypher_builder,string(Alphabet[alphabet_index]))

		if ((i & 3) == 3 && i < 13 ){
			cypher_builder = append(cypher_builder, "-")
		}
	}

	rs_address = toString(cypher_builder)

	return rs_address
}


func RSDecode(accountId float64) (rs_address string){

	return rs_address
}


func Gmult(a int,b int) int{
	if (a == 0 || b == 0){
		return 0
	}
	idx := (Glog[a] + Glog[b]) % 31
	return Gexp[idx]
}


func toString( strs []string) (rs_address string) {
	for i:=0 ; i < len(strs); i++{
		rs_address += strs[i]
	}
	return rs_address
}


func Pubkey2Account(publicKey string )  (account_id string){
	pubkey := []byte(publicKey)
	publicKeyHash := sha256.Sum256(pubkey)
	fmt.Printf("publicKeyHash=%v\n",publicKeyHash)
	if (len(publicKeyHash) < 8) {
		fmt.Printf("invalid public sha256 hash!\n")
	}

	BigInteger := new(big.Int)
	BigInteger.SetBytes([]byte{publicKeyHash[7],publicKeyHash[6],publicKeyHash[5],publicKeyHash[4],publicKeyHash[3],publicKeyHash[2],publicKeyHash[1],publicKeyHash[0]})

	account_id = BigInteger.String()

	return account_id
}



func NewKeyPair() (privateKey string,publicKey string, err error) {
	PubKey, PrivKey, _ := ed25519.GenerateKey(nil)
	privateKey = hex.EncodeToString(PrivKey[:])
	publicKey = hex.EncodeToString(PubKey[:])
	return privateKey,publicKey ,nil
}


// 这里seed就用sss门限库里的私钥 其中32位
//aaf60a5aa5a11cf5882938131355b5a2c7ac341d45cad7acc3a338188c4c2588f410bc92158a6bfdf7f5f9f4f2c8f92f64abb7a117490aaf31ac1a209443cf08
func NewKeyPairFromSeed(seed string) (privateKey string,publicKey string, err error){
	seed_byte := []byte(seed)
	fmt.Printf("len(seed)=%v\n",len(seed))
	PrivKey := ed25519.NewKeyFromSeed(seed_byte)
	privateKey = hex.EncodeToString(PrivKey[:])
	publicKey = hex.EncodeToString(PrivKey[32:])
	return privateKey,publicKey ,nil
}


func Createtransaction(senderAccount string, recipientId string,amountATM string,attachment string,broadcast bool,validate bool,publicKey string,recipientPublicKey string) (  unsignedTransaction string){
	
	unsignedTransaction = TransactionConvert(senderAccount,recipientId,amountATM,attachment,broadcast,validate,publicKey,recipientPublicKey)

	return unsignedTransaction
}



func TransactionConvert(senderAccount string, recipientId string,amountATM string,attachment string,broadcast bool,validate bool,publicKey string,recipientPublicKey string)  ( unsignedTransaction string){
	var transactionObject TransactionObject
	transactionObject.EncryptedMessageIsPrunable = true
	transactionObject.MessageIsPrunable = false
	transactionObject.Phased = false
	//transactionObject.EcBlockId,_= getECBlock()
	transactionObject.EcBlockId = "9906837839936786887"
	transactionObject.Broadcast = false
	transactionObject.DeadlineValue = 600
	transactionObject.PublicKeyValue = publicKey
	transactionObject.AmountATM = amountATM
	transactionObject.FeeATM = "100000000"
	transactionObject.SenderAccount = senderAccount
	transactionObject.RecipientId = recipientId
	transactionObject.RecipientPublicKey = recipientPublicKey
	transactionObject.Attachment = attachment
	transactionObject.Validate = validate

	unsignedTransactionBytes,err := json.Marshal(transactionObject)
	if err != nil{
		fmt.Printf("marshal unsignedTransaction failed,err=%v\n",err)
	}
	unsignedTransaction = hex.EncodeToString(unsignedTransactionBytes)

	//fmt.Printf("unsignedTransaction=%v\n",unsignedTransaction)

	return unsignedTransaction
}







