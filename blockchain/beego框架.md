### beego 

#### 模块设计

session 模块
是用来存储客户端用户，session 模块目前只支持 cookie 方式的请求，如果客户端不支持 cookie，那么就无法使用该模块。
钱包组的项目中没有涉及和用户保持连接 没有用这个模块

logs模块

config模块

cache模块
四种缓存模型
memory 内存
file 
redis
memcache

httplib模块  context模块
发送http请求 类似curl调用 封装request请求传参 和response解析 有可能要配置代理 发送http请求 
为了安全性 有些内网ip的机器只能通过公网ip的机器去发送请求 比如sss_server所在机器都得配置代理

#### model设计


