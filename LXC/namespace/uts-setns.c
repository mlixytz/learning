#define _GNU_SOURCE
#include <fcntl.h>
#include <sched.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

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
    ret = setns(fd, 0);
    NOT_OK_EXIT(ret, "open");

    execlp("bash", "bash", (char *) NULL);

    return 0;
}