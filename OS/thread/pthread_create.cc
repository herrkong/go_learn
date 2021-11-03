#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <signal.h>
#include <sys/time.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <pthread.h>
#include <errno.h>
#include <string.h>
 
 
void* tfun(void* args){
	int* j=(int *)args;
	printf("son thread.... pid = %d ...  %lu-args--%d \n",getpid(),pthread_self(),*j);
   // printf("---%d---",*j);
	return NULL;
}
int main(){
    /**
     *  线程共享资源：
     *  1. 文件描述符表
     *  2. 当前工作目录
     *  3. 用户ID 和组 ID
     *  4. 内存地址空间(.text/.data/.bss/堆/共享库)[共享全局变量]
     *  非共享：
     *1. 线程id
     *2. 栈
     *  优点： 效率高
     */
 
	// lwp 标识线程身份，cpu调度区分   线程id:  进程内部不同线程区分
	pthread_t tid;
	// 线程id,线程属性，回调函数，回调函数参数
	// 打印主线程
	printf("main thread..pid = %d ... pthread %lu \n",getpid(),pthread_self());
    int j=1000;
	int ret = pthread_create(&tid,NULL,tfun,&j);  // 返回线程id
	sleep(1);  //子线程还没有输出，父进程已经死了，避免
	if(ret == -1){
		perror("pthread_create fail");
		exit(-1);
	}
	// gcc single.c -o single.o -lpthread
     // 第三方库要加 -lpthread
	return 0;
}