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
 *  exit: 退出进程
 *  pthred_exit 退出线程
 */
void* tfun3(void* args){
	int j=(int )args;
	if(j==2){
	//	return NULL;  // 返回到函数调用者那里去 ， 退出以后不会往下走了
		pthread_exit(NULL);
	}
	printf("son thread...tfun1.... pid = %d ...  %lu-args--%d \n",getpid(),pthread_self(),j);
   // printf("---%d---",*j);
	return NULL;
}
int main0000017() {
	pthread_t tid;
	int ret=0;
	for (int i = 0; i < 10; i++) {
		ret= pthread_create(&tid, NULL, tfun3, (void*) i);  // 返回线程id
		if (ret == -1) {
			perror("pthread_create fail");
			exit(-1);
		}
	}
	sleep(1);
 
	return 0;
}