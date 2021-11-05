
#### channel 底层实现原理

读等待协程队列 recvq，维护了阻塞在读此 channel 的协程列表
写等待协程队列 sendq，维护了阻塞在写此 channel 的协程列表
缓冲数据队列 buf，用环形队列实现，不带缓冲的 channel 此队列 size 则为 0

有缓冲channel
无缓冲channel
