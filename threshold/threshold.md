

##### transfer demo
eth_transfer demo
// 构建转账脚本 丢给tss server签出3个big_int r s v 


./tss_tools_linux eth_transfer --config=eth_transfer/input.yaml 
./signer_linux.sh 


// 解构eth





#### DISTRIBUTED KEY GENERATION

分布式密钥生成DKG(Distributed Key Generation)相对于传统KGC(Key Generation Center)的特点是在于前者密钥对的生成不依赖于任何可信的第三方


#### tss_server 


#####  三个定时任务 (beego task)
task0: CheckSignError,检查签名正确性, 熔断机制，lark报警
task1: DistributeCmdPeerId,
task2: SetCmdFail,30s还没签好,直接设为失败


##### tss_server， tss_peer对象创建



##### tss_peer 

pullsigncmd
pushsignresult

sync.map  加解锁 peer端口

若干个peer上报tx_sign，均能得到r,s 

key_load 

key_dump

key_check




#### 本地部署一套 tss门限 验证身份
私钥用dkg打散



