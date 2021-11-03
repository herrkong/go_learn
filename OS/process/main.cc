#include <iostream>
#include <stdlib.h>
#include <vector>
#include <map>
#include <set>
#include <algorithm>
#include <unistd.h>

using namespace std;

// fork

//RETURN VALUES
    //  Upon successful completion, fork() returns a value of 0 to the child process and returns the process ID of the child process to the parent process.  Otherwise, a value of -1 is
    //  returned to the parent process, no child process is created, and the global variable errno is set to indicate the error.


int main() {
    //printf("子进程pid=%d，父进程ppid=%d\n",getpid(),getppid());
    pid_t pid = fork();
    if (pid == 0) {
        printf("child process =%d\n",getpid());
    }else{
        printf("pid = %d\n",pid); // 父进程中返回子进程的pid 
        printf("father process =%d\n",getpid());
    }

    return 0;
}


