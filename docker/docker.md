

#### docker k8s 

linux docker将程序和其依赖的环境打包在一起 

k8s 基于容器的集群管理平台 kubernetes

##### docker 和虚拟机的不同

##### docker 命令

pull ubuntu镜像
docker pull ubuntu

docker 运行容器 
docker run -it ubuntu 

列出所有正在运行的container
docker container ls

列出本机所有容器 包括正在运行的
docker container ls -all 

docker  build -t ubuntu

#### 编写docker文件

.dockerignore文件 排除若干文件

From ubuntu
COPY . /root/darwin   //从上下文目录中复制文件或者目录到容器里指定路径。
WORKDIR /root/darwin  






#### 1 docker的cmd与 entrypoint区别
ENTRYPOINT指向你的Python脚本本身. 当然你也可以用CMD命令指向Python脚本. 但是通常用ENTRYPOINT可以表明你的docker镜像只是用来执行这个python脚本,也不希望最终用户用这个docker镜像做其他操作.

#### 创建自己的镜像文件image


#### 创建Dockerfile文件


#### docker api 
Docker 提供了一个与 Docker 守护进程交互的 API (称为Docker Engine API)


命令行启动进入容器
docker run -it ubuntu /bin/bash

后台运行docker
docker run -itd --name ubuntu-test ubuntu /bin/bash


运行程序
docker run ubuntu /bin/echo "Hello world"

目录 


#### docker-compose
创建compose文件 

docker-compose up   
attach 
