
// 查找进程
ps -ef | grep withdraw_linux 

// 列出所有正在使用的端口和关联的进程 
netstat -nap 

// wc -l 统计个数
// grep 
netstat -nap | grep withdraw | grep 25298 | wc -l

//查看cpu占用率负载等
top 

// 测试连接 远程登录
telnet 