
https://blog.csdn.net/dreams_deng/article/details/104201715



##### 进程是分配资源的最小单位,  线程是最小的执行单位

线程在进程内部，共享进程地址空间


 pthread_create()   fork()
  *    pthread_self()     getpid()
  *    pthread_exit()     exit()
  *    pthread_join()     wait/waitpid
  *    pthread_cancle()   kill
  *    pthread_detach()


