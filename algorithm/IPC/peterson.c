// Peterson互斥算法

#define FALSE 0
#define TRUE 1
#define N 2 // 进程数量

int turn; // 现在轮到谁
int interested[N]; // 所有值初始化为0

void enter_region(int process) {
    int other;
    other = 1 - process;
    interested[process] = TRUE;
    turn = process;
    while (turn == process && interested[other] == TRUE);
}

void level_region(int process) {
    interested[process] = FALSE;
}
