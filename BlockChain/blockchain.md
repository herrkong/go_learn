

#### 区块链
一个去中心化的账本 每个节点都可以参与记账 记账能得到激励 挖矿就是打包别人的交易 获得激励的过程 
因为每一笔交易 都会有一部分费用被矿工拿了



#### btc eth 
##### 公私钥对的生成 地址的生成 
privatekey-pubkey  key-pair  
基于哪种曲线 secp256  p256  ed25519

address其实就是pubkey的hash


注意btc的余额体系 是一种utxo模型(unspent outputs) 未花费的交易
通过listunspent 去查询某个address的余额 

你所拥有的是花费这个output的权利,也就是说,你的私钥可以制定这个output成为谁的input.

bob 给alice 转 3个btc
bob     alice     bob
--> 5   -->3      -->1.99998424    (自己指定去向 指定手续费的多少)

--->2                                

eth 就类似于银行的账户余额概念 

#### 交易构建
rpc方法 节点中转账 sendtoaddress

createtransaction 构建转账脚本 
利用私钥签名 签名算法 

将signedtransaction broadcast广播


##### 门限签名 
sss架构
将私钥打散 仅在签名时拼接成完成私钥签名 

tss架构
不直接用私钥签名 而是通过threshold门限给出signedtransactionbytes


#### 加密算法  共识算法 签名算法

#### 非对称加密 对称加密等

##### 加密算法
// 对称加密 
AES

//非对称加密
rsa 
escda (btc )



secpk256
ed25519

##### 共识算法
###### PoW  (proof of work)工作量证明 
完成某项工作的凭证 原理： 寻找一个以若干个0为开头的哈希串

比如矿工们争夺这个新区块的打包权 通过工作量证明pow的机制 
输入是一个给定字符串+nonce 
所有矿工们都在不断的进行hash计算 设定的难度值规定要进行多少次hash才合法   知道找到指定多少个0开头的hash串 

打包一个区块 区块中包含多少交易  
sha256哈希函数 
H(x)
无论x输入多长，都输出64个字符，共32字节（byte），256位（bit） ，输出有2^256可能 产生一样的hash的概率几乎不可能
输出只包含数字0~9和字母A~F，大小写不敏感

btc使用了两种hash算法  先进行sha256 再进行RipeMD160  （常用的has算法还有md5 sha512等等）

###### PoS 权益证明 Proof Of Stake 
提出币龄的概念 ： 币龄 = 持有量 * 时间   币龄越大更有可能获得打包记账权 
为了避免囤币 独占打包权 设置最大记账概率

###### DPoS 委任权益证明 Delegated Proof of Stake 
类似人民代表大会制度 选出一定数量的人大代表 来进行打包出块 如果不能履行职责就回被替代


####  币种
btc eth  dash mrx berry/luk etsc ddr chnd ltnm vlx sol 

#### erc20代币 
智能合约 contract  


####  挖矿

有些节点一直在尝试打包新的交易到区块 并附加到区块链上  这就是矿工在挖矿 这个过程为获得经济奖励 

发起转账交易需要支付一定费用给矿工。(input - output)(这里不会在交易中指定手续费地址，因为还不知道是哪个矿工打包交易呢)

coinbase中指向矿工地址的output的金额 就是该区块交易的中所有手续费 和 挖矿激励之和

矿工们首先争夺这个新区块的打包权 通过工作量证明pow的机制 

所有矿工们都在不断的进行hash计算 直到计算值小于设定的难度值 



##### 可能出现的一个问题

两个矿工在同一时间打包了新的有效区块 两个区块数据不同 （coinbase address 不同 merke hash root 不同）,   有可能都会上链 这就产生了分叉，
看后续谁先接上下一个区块，另外一条链就会废弃，区块链的数据最终都会保持一致（最长分叉的共识算法），一般来说经过6个区块确认的交易几乎不可更改。


##### 区块数据结构

区块 
区块大小  Block Size
区块
Nonce 交易计数器
交易列表



区块头 数据结构
Version 
Merkle Root hash    Merkle数 root hash 
previous block hash 上一个blcok hash
difficulty target  难度目标
timestamp
nonce 


区块数据    
区块第一个交易是coinbase交易 给矿工打包产生区块的激励 现在大概6个btc 每四年减半一次

打包的其他交易数据
transaction



#### 闪电网络 
再加一层网络
将很多小额的转账打包成一个交易 交给底层的比特币网络执行，用户不用等待 手续费也低 
比特币服务于小额支付场景

但是闪电网络 还得建立一个中心化的系统去打包小额交易 其实违背了去中心化的设计初衷


#### 隔离见证 Segregated Witness

隔离见证(SegWit)是把交易的签名数据从交易数据中剥离出来，用于解决延展性攻击。

区块中只包含交易数据 记录交易金额去向就好 每个区块可以包含更多的交易 



#### 为什么使用merkel树
Merkle Tree 是一种二叉树，包含了区路中所有交易

完整性验证  : 树中任何一个节点发生变化 merkle tree root 节点就会变化
零知识证明 : 证明树中确实有这个笔交易 但是确不知道这笔交易的具体内容 和 整个树的内容 



#### 5 非对称加密 np问题 rsa算法

p问题 存在多项式时间内能解决的问题 时间复杂度 O(n2)

np问题  能在多项式时间内证明或证伪的问题 

rsa 经典的非对称加密算法
##### 非对称加密
公钥加密 私钥解密


单向函数求解简单 反向函数求解复杂的特性

1 将两个质数相乘很容易 但是将合数分解为质数却很难 
n1 * n2 ---> p     p ---> n1  * n2 

2 (m^e) mod n = c 

m e n ---> c    c---> m e n 


##### rsa算法

1 公私钥的生成
a 选取两个大素数 p q 
b 计算模数 n = p * q 
c 根据欧拉函数 r = (p -1)(q -1)
d 选择一个小于r的整数e  求得e关于模r的模反元素 命名为d （模反元素:d*e≡1(mod r)）
e 销毁p q 
f  (n,e)为公钥 （n,d）为私钥 


p = 3 q = 11 

n = 3 * 11 = 33 

r = (3 -1)(11 -1) = 2 * 10 = 20 

e * d = 20 * a + 1 (a 为正整数 且e与r互质)

则可以选择 e = 3  d = 7 

则公钥 （33，3） 私钥（33，7）


A 发送18给 B 

A 用公钥 加密 message 18 

18^3 mod 33  = 24  

B 用私钥对24 解密 

24 ^ 7 mod 33 =  18 

这里的n在实际使用时是非常大的 


##### 对称加密算法

私钥加密 私钥解密

AES算法 



#### 区块链 矿工挖矿 工作量证明pow 

简单的来说 区块链就是由一个个区块链接而成 
每一个区块中记录了谁转账给谁多少钱等交易信息  这是一个分布式的 去中心化的账本 每一个节点都可以记账 都能保存整个区块链的交易信息

我们传统的概念就是中心化的 比如去银行存钱 银行来作为一个中心化系统记录你存了多少钱 转账多少钱给谁。

在比特币区块链系统中不是银行转账的概念 它建立的是utxo的经济模型  unspent outputs 未花费的输出

比如Alice 有一笔未花费的outputs  这个是上一个人给他转的5个btc 到它的地址 ，alice 现在想要花掉这个3个btc 就把这3个btc转给bob
发现只有1.99994682了（交易单位是satois 10^8）,有一部分作为这笔交易的手续费支付给矿工了。 

矿工们是要竞争最新一个区块的打包权  打包这个区块的矿工除了获得每笔交易的手续费 还会有一笔挖矿激励 这是在每个区块的第一笔coinbase交易 现在大概是6个btc
每四年挖矿激励会减半 矿工总共获得的钱就是挖矿激励 和 手续费的总和  这个手续费的设置 就是发起转账人自己构建交易时决定的 发送多少 找零多少 自己设置  
设置的手续费越高 矿工就肯定优先打包手续费高的交易 优先上链。

矿工们这么竞争打包权的   这个就是工作量算法pow来决定 
大家同时在做hash计算 hash函数就是给定一个字符串 输出都是相同长度的字符串 区块难度系数设定了是4 大家对同一个字符串做hash运算
看谁的hash结果的前4位全都为0 这就是符合条件的输出 这个矿工就获得了打包权
当然不排除这样的情况 两个矿工同时算出结果 同时都在打包出了区块  这就是出线了分叉 
下次看谁先出块  落后的那条分叉链就会被抛弃 不可能两条链每次都同时出块
 

Alice      bob   Alice

 5         3      
---->   --->    
                1.99994682
                -------->


