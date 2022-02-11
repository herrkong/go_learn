

#### slice 和 array 的区别 slice的实现

数据结构和内存

函数传参

#### go中实现函数的timeout有哪几种方式

1 select + time.After
2 context提供了withTimeout的方法 
3 time.NewTimer


#### Context的使用场景

1 rpc调用
在主goroutine上有4个RPC，RPC2/3/4是并行请求的，我们这里希望在RPC2请求失败之后，直接返回错误，并且让RPC3/4停止继续计算。这个时候，就使用的到Context。

2 PipeLine
pipeline模式就是流水线模型，流水线上的几个工人，有n个产品，一个一个产品进行组装。其实pipeline模型的实现和Context并无关系，没有context我们也能用chan实现pipeline模型。但是对于整条流水线的控制，则是需要使用上Context的。

3 超时请求
我们发送RPC请求的时候，往往希望对这个请求进行一个超时的限制。当一个RPC请求超过10s的请求，自动断开。当然我们使用CancelContext，也能实现这个功能（开启一个新的goroutine，这个goroutine拿着cancel函数，当时间到了，就调用cancel函数）

4 HTTP服务器的request互相传递数据


#### goroutine和线程的关系 什么情况下会创建新的线程



#### 什么是代码的可测性 什么样的代码会有好的可测性

代码的可测试性就是根据代码编写单元测试的容易程度。对测试性差的代码往往难以编写单元测试

依赖注入是实现代码可测试性的最有效的手段



#### 怎么优化go语言服务的内存占用










