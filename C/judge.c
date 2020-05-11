#include<unistd.h>
#include<fcntl.h>
#include<stdlib.h>
#include<string.h>
#include<stdio.h>
#include<sys/types.h>
#include<sys/time.h>
#include<windows.h>


#define AC 0
#define PE 1
#define TLE 2
#define MLE 3
#define WA 4
#define RE 5
#define OLE 6
#define CE 7
#define SE 8

struct result {
    int status;
    int timeUsed;
    int memoryUsed;
};

/**
 * 设置进程cpu和内存限制
 */
void setProcessLimit(int time_limit,int memory_limit){
  
}

/*
monitor the user process
*/
void monitor(pid_t pid, int timeLimit, int memoryLimit, struct result *rest) {
    int status;
    struct rusage ru;
    if (wait4(pid, &status, 0, &ru) == -1)
        printf("wait4 failure");
    rest->timeUsed = ru.ru_utime.tv_sec * 1000
            + ru.ru_utime.tv_usec / 1000
            + ru.ru_stime.tv_sec * 1000
            + ru.ru_stime.tv_usec / 1000;
    rest->memoryUsed = ru.ru_maxrss;
    if (WIFSIGNALED(status)) {
        switch (WTERMSIG(status)) {
            case SIGSEGV:
                if (rest->memoryUsed > memoryLimit)
                    rest->status = MLE;
                else
                    rest->status = RE;
                break;
            case SIGALRM:
            case SIGXCPU:
                rest->status = TLE;
                break;
            default:
                rest->status = RE;
                break;
        }
    } else {
        if (rest->timeUsed > timeLimit)
            rest->status = TLE;
        else if (rest->memoryUsed > memoryLimit)
            rest->status = MLE;
        else
            rest->status = AC;
    }
}

int run(char *args[],int timeLimit, int memoryLimit, char *in, char *out){
    pid_t pid = vfork();
    if(pid<0)
        printf("error in fork!\n");
    else if(pid == 0) {
        runCmd(args, timeLimit, memoryLimit, in, out);
    } else {
        struct result rest;
        monitor(pid, timeLimit, memoryLimit, &rest);
        printf("{\"status\":%d,\"timeUsed\":%d,\"memoryUsed\":%d}", rest.status, rest.timeUsed, rest.memoryUsed);
    }
}

/**
 * 分割字符串
 */
void split(char **arr,char *str,const char *del){
  char *s = NULL;
  
}


int main(int argc,char *argv[]){
  char *cmd[20];
  split(cmd,argv[1],"@");
  run(cmd,atoi(argv[2]),atoi(argv[3]),argv[4],argv[5]);
}