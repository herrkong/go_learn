
#### 内核态 用户态

cpu指令集
汇编语言 对应着cpu指令  指令的集合 就是指令集 
CISC复杂指令集  RISC精简指令集 

这些能直接操作硬件的指令非常危险 操作系统屏蔽非系统调用 

用户态下通过系统调用来申请内存 控制硬件 ring3 
内核态   操作系统内核中运行 执行所有cpu指令集  ring0(cpu指令集操作权限) 

https://segmentfault.com/a/1190000039774784



#### 内核态用户态切换开销大

用户态 调用系统调用(文件操作 socket 等) 调用完成再回到用户态  
当前用户状态的保存和恢复  内核态做的检查工作

https://sites.google.com/site/linux31family/home/home-1/--7/linux-c--1

