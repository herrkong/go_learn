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
 


void* tfun1(void* args){
	int j=(int)args;
	printf("son thread...tfun1.... pid = %d ...  %lu-args--%d \n",getpid(),pthread_self(),j);
   // printf("---%d---",*j);
	return NULL;
}
int main0000016() {
	pthread_t tid;
	int ret=0;
	for (int i = 0; i < 10; i++) {
		ret= pthread_create(&tid, NULL, tfun1, (void*) i);  // 返回线程id
		if (ret !=0 ) {
			perror("pthread_create fail");
			exit(-1);
		}
	}
	sleep(1);
	return 0;
}