

#### grpc和restful api架构风格的对比

grpc 
对接口有着严格的定义 封装了底层的传输方式(tcp udp)和序列化方式(xml json 二进制)和通信细节
让远程调用服务更加简单 

restful 
http + json架构  post get delete 
所有浏览器都支持restful 但不一定支持grpc 

请求响应模式 每次只能处理一个请求 拖慢服务器的效率 
grpc支持双向通信 流式通信等处理


grpc默认使用protobuffer来序列化数据 轻便 简单 高效压缩数据消息大小 二进制传输 对结构化数据进行序列化和反序列化 便于通信和传输

restful的json  读起来方便 简单

