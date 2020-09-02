#define _GNU_SOURCE
#include <sched.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

#define NOT_OK_EXIT(code, msg); {if(code == -1){perror(msg); exit(-1);}}

int main(int argc, char * argv[]) {
    int fd, ret;

    if (argc < 2) {
        printf("%s /proc/PID/ns/FILE\n", argv[0]);
        return -1;
    }

    fd = open(argv[1], O_RDONLY);
    NOT_OK_EXIT(fd, "open");

    // 执行完setns后，当前进程将加入指定的namespace
    // 第二个参数为0，表示由系统自己检测fd对应的是那种类型的namespace
    ret = setns(fd, 0)；
    NOT_OK_EXIT(ret, "open")

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