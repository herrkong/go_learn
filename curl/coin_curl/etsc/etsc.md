
decimal  18 
100 000000000000000000

0xe45D60e41B7bdA46d2306c07Da411b8b7af6A0bA

0xc9fA462De139b7eb98b36422a54548BFD1de8350


每一笔转账需要销毁转账金额的0.8%，销毁地址： 0x0000000000000000000000000000000000000000

19500.00 

156.00 



0x0000000000000000000000000000000000000000


归集检查的问题

地址余额查询失败


transfer  nonce=23 , 销毁 nonce=24 还是失败了,两笔交易间隔时间1min10s左右，销毁可能失败

1min  12 block 

提币 归集 销毁 三个转账 始终有机会接近 会造成广播失败  nonce 太接近 每次构建转账时  nonce 随机增加2~5 




#### all_amount
curl "http://127.0.0.1:9199/get/json?action=all_amount&coin=etsc"

curl "http://127.0.0.1:9197/get/json?action=all_amount&coin=etsc"



####  hash_query_info


curl "http://127.0.0.1:9199/get/json?action=hash_query_info&coin=etsc&hash=0x406e35697eac90943795533e15834f9d04565e2de16bc8cc7592156b4a6da115&address=0x0be17519f8aefb9a707f30f7bf76f94d86d176c9&amount=4960.00000000"   


curl "http://127.0.0.1:9197/get/json?action=hash_query_info&coin=etsc&hash=0x406e35697eac90943795533e15834f9d04565e2de16bc8cc7592156b4a6da115&address=0x0be17519f8aefb9a707f30f7bf76f94d86d176c9&amount=4960.00000000"   
