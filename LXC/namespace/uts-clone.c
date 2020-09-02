#define _GNU_SOURCE
#include <sched.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

#define NOT_OK_EXIT(code, msg); {if(code == -1){perror(msg); exit(-1);}}

static int child_func(void *hostname) {
    // 设置主机名
    sethostname(hostname, strlen(hostname));
    // 用一个新的bash来替换掉当前的子进程
    // 执行完execlp后，子进程没有退出，也没有创建新的进程
    // 只是当前子进程不再运行自己的代码，而是去执行bash的代码
    // bash退出后子进程执行完毕
    execlp("bash", "bash", (char *) NULL);

    // 从这里开始代码将不会被执行到，因为当前进程已经被上面的bash替换掉了

    return 0;
}

static char child_stack[1024 * 1024];

int main(int argc, char * argv[]) {
    pid_t child_pid;
    
    if (argc < 2) {
        printf("Usage: %s <child-hostname>\n", argv[0]);
        return -1;
    }

    // 创建并启动子进程，调用该函数后，父进程将继续往后执行，也就是执行后面的waitpid
    // 子进程将执行child_func这个函数
    // 栈是从高位向低位增长，所以这里要指向高位地址
    // CLONE_NEWUTS 表示创建新的 UTS namespace
    // SIGCHLD 是子进程退出后返回给父进程的信号，跟namespace无关
    // 传给child_func的参数
    child_pid = clone(child_func, child_stack + sizeof(child_stack), CLONE_NEWUTS | SIGCHLD, argv[1]);

    NOT_OK_EXIT(child_pid, "clone");

    waitpid(child_pid, NULL, 0); // 等待子进程结束

    return 0;
}