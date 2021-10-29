#!/usr/bin/env bash
time=$(date "+%Y%m%d.%H%M%S")
external_server_dir="/data/vhosts/tss.digifinex.org/external_server"

echo "编译external_server"
PROJECT_FOLDER=$(cd "$(dirname "$0")";cd ..;pwd)
cd $PROJECT_FOLDER/bin
GOOS=linux go build -o $external_server_dir/bin/external_server_linux."${time}" ../src/external_server
cd ${external_server_dir}/bin && ln -s -f -n ./external_server_linux.${time} ./external_server_linux && ln -s -f -n ${external_server_dir}/bin/external_server_linux.${time} /tmp/external_server_linux

echo "停止external_server"
killall external_server_linux
sleep 1

echo "启动external_server"
cd ${external_server_dir}/bin ; nohup ./external_server_linux > ${external_server_dir}/bin/nohup.out 2>&1 &
sleep 1

echo "检查运行external_server"
ps -eo pid,lstart,etime,cmd | grep -v grep |grep external_server_linux
