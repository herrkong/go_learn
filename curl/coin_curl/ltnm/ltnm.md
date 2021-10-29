base58Prefixes[PUBKEY_ADDRESS] = std::vector<unsigned char>(1,65); 
base58Prefixes[SCRIPT_ADDRESS] = std::vector<unsigned char>(1,25); 
base58Prefixes[SECRET_KEY] =     std::vector<unsigned char>(1,128); 


base58Prefixes[PUBKEY_ADDRESS] = std::vector<unsigned char>(1,137);
base58Prefixes[SCRIPT_ADDRESS] = std::vector<unsigned char>(1,78);
base58Prefixes[SECRET_KEY] =     std::vector<unsigned char>(1,130);


//测试服热钱包地址
xMEPz3bYeh4fQjhntmQ1hP4Q4GMPx5dpv6

//导入节点
xGupv64Lxcd7JCYmrqP4ih4Yk2JWmiBrPE
xD6w4zJJJM4VDVPa8dG39byo7AoWvk4yHg
xMNNRaYXExYhN88gCtgmB66cZLbLfTdtJM
xDBUXDFJ9xrNrP7kd1AjfmApeuRDBXXsKs
xC7yUj3mZ99VUhAyiCwqFXx3HVDKzcQtCc
xMEPz3bYeh4fQjhntmQ1hP4Q4GMPx5dpv6


// mainnet
TJaQ2GbbqdG2QzWY6dQ7ofTuRhkbZgEAAn

TJaQ2GbbqdG2QzWY6dQ7ofTuRhkbZgEAAn
TEmWBAqZBMhQLHMLNRH6EaP9nrFbiXsEA4
TP2wXm5n7yBcUv6SSghpG4VyF23RYL1vYE
TEr3dPnZ2yVHyB5WroBnkjaBLasJ33M38j
TDnYaub2S9nQbV8jwzxtLWMPyAfQpaaoYd
TNty6E8oXhhaXXfZ8ZR4nMTkjwoUiFVMC2


#### all_amount
curl "http://127.0.0.1:9199/get/json?action=all_amount&coin=ltnm"

curl "http://127.0.0.1:9197/get/json?action=all_amount&coin=ltnm"


####  hash_query_info


curl "http://127.0.0.1:9199/get/json?action=hash_query_info&coin=ltnm&hash=557437221053d681acd804c033e6c58f3bd3830ff36a0020a5c9f0d6dbcece9a&address=xMEPz3bYeh4fQjhntmQ1hP4Q4GMPx5dpv6&amount=200.00000000"


curl "http://127.0.0.1:9197/get/json?action=hash_query_info&coin=ltnm&hash=557437221053d681acd804c033e6c58f3bd3830ff36a0020a5c9f0d6dbcece9a&address=xMEPz3bYeh4fQjhntmQ1hP4Q4GMPx5dpv6&amount=200.00000000"



http://3.35.159.185:38552/   with username and passwd : iqlas@cumuluslogic.net:P@ssword123

curl --user main:main --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' : http://dns.seed2.latinum.staking.zeeve.net:38555/ | jq


http://main:main@dns.seed2.latinum.staking.zeeve.net:38555/wallet/testwallet/