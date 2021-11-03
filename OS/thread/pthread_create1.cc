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
 
/**
 * 循环创建多个线程
 */
int main() {
	pthread_t tid;
	int ret=0;
	for (int i = 0; i < 10; i++) {
		// 错误
		/**
		 * son thread.... pid = 3521 ...  140422606616320-args--2
son thread.... pid = 3521 ...  140422589830912-args--5
son thread.... pid = 3521 ...  140422598223616-args--3
son thread.... pid = 3521 ...  140422615009024-args--1
son thread.... pid = 3521 ...  140422581438208-args--7
son thread.... pid = 3521 ...  140422564652800-args--7
son thread.... pid = 3521 ...  140422573045504-args--7
son thread.... pid = 3521 ...  140422556260096-args--8
son thread.... pid = 3521 ...  140422479345408-args--9
son thread.... pid = 3521 ...  140422470952704-args--10
 出现重复数据，
 为什么： for循环不断执行当，但i=5的时候在栈中，子线程回调才第一次执行，才切换成内核态取出i=5,实际的需要的是0
 解决方法： 直接传递值
		 */
		ret= pthread_create(&tid, NULL, tfun, (void*) &i);  // 返回线程id
		if (ret !=0 ) {
			perror("pthread_create fail");
			exit(-1);
		}
	}
	sleep(1);
	return 0;
}