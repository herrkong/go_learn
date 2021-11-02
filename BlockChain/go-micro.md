#### go-micro

push中使用的架构  基于golang的微服务框架 还有go-kit等


push 

center

RegisterCenterHandler

chain 



go-micro 1.0


                        service

        client                              server

broker codec  register                  selector  transport

broker : 消息发布和订阅的接口 支持http rabbitmq redis
codec : 通讯格式 protobuf json
Registry : 用于实现服务的注册和发现 Go-Micro etcd、kubernetes


go-micro 3.0

            services

client                   server 



