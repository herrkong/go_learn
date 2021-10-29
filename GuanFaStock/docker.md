### 了解虚拟化技术 docker容器使用


#### 
linux docker将程序和其依赖的环境打包在一起 


#### 1 docker的cmd与 entrypoint区别
ENTRYPOINT指向你的Python脚本本身. 当然你也可以用CMD命令指向Python脚本. 但是通常用ENTRYPOINT可以表明你的docker镜像只是用来执行这个python脚本,也不希望最终用户用这个docker镜像做其他操作.

#### 创建自己的镜像文件image


#### 创建Dockerfile文件


#### docker api 
Docker 提供了一个与 Docker 守护进程交互的 API (称为Docker Engine API)

pull ubuntu镜像
docker pull ubuntu

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



