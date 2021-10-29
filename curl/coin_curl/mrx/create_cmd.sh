#!/bin/bash

#账号
rpcuser=iouyegoW409Iuyt
#密码
rpcpassword=aT3ddxbtLrboiTu

coin=mrx

host=https://testmrxrpc.digifinex.org/wallet/
#RPC_PORT=8332
#PORT=8333




####如果导地址，用这个
while read LINE;do

echo "curl -m 5 -s --user ${rpcuser}:${rpcpassword} -H 'content-type: text/plain;' --data-binary '{\"jsonrpc\":\"1.0\",\"id\":\"curltest\",\"method\":\"importaddress\",\"params\": [\"$LINE\", \"\", false] }' $host |  perl -lpe 's/(0x[0-9a-f]{1,})/hex($1)/e'  | jq ."  >>load_address.sh

done  <mrx.load_address.txt

####如果导公钥，用这个

#while read LINE;do
##echo "curl -m 5 -s --user ${rpcuser}:${rpcpassword} -H 'content-type: text/plain;' --data-binary '{\"jsonrpc\":\"1.0\",\"id\":\"curltest\",\"method\":\"importpubkey\",\"params\": [\"$LINE\", \"\", false] }' http://127.0.0.1:\"$RPC_PORT\"/ |  perl -lpe 's/(0x[0-9a-f]{1,})/hex($1)/e'  | jq ."  >>load_pubkeys.sh
#done  <bsv.load.pubkeys.txt
