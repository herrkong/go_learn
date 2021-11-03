
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

/**
 * pthread_join 阻塞回收子线程 相当有 wait
 */

void* tfun4(void* args){
	return (void*)1100;
}
int main0000018() {
	pthread_t tid;
	//  成功返回0
    int ret= pthread_create(&tid, NULL, tfun4, NULL);  // 返回线程id
    if(ret !=0){
    	perror("pthread_create erro");
    	exit(-1);
    }
    int *retval;
    // 参数类型void**
    // 参数1 线程id  参数2 回调函数返回值 tfun4
    ret = pthread_join(tid,(void**)&retval);
    if(ret!=0){
     	perror("pthread_create erro");
        exit(-1);
    }
    // 返回值类型 -- return (void*)1100;
     printf("resutl---%d",retval);
	return 0;
}