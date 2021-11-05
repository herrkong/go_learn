
#### 1 mysql和redis 数据一致性问题

读redis :  先查redis 查不到 再查mysql 并把查询结果写redis
写redis  
可能产生不一致的情况 : 
先更新mysql 在删除redis之前 mysql挂了没有成功写入 
先删除redis 在更新mysql之前 去读 没有 从mysql加载到redis 此时redis中是脏数据

延时双删策略来避免 mysql和redis 数据不一致

写mysql前后 都删除redis  给redis 设置过期时间
下一次再来读的时候 加载的mysql中的最新值 

#### 缓存穿透 缓存击穿 缓存雪崩 及其解决办法

缓存穿透: 数据在缓存和数据库中都不存在 
设置缺省值 查询为空 也存缓存 过期时间设置较短一点

缓存击穿: 访问非常频繁的缓存过期了 导致查mysql
访问频繁的redis key 不设置过期时间

缓存雪崩: 大量缓存失效 或者redis挂了
数据预热 提前查询比较频繁的key
设置不同的缓存过期时间 分布均匀


#### redis 集群模式

主从模式: master and slave . master可读写 slave 只读
哨兵模式:  master挂了 可将其中一个slave 切换为master 
cluster集群模式 : 分布式存储 节点间通信 超过半数节点挂了 集群才不可用
