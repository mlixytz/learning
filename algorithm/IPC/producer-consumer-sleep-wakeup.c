// 生产者消费者问题 sleep-wakeup 伪代码实现

#define TRUE 1
#define N 100 // 缓冲区中的槽数目
int count = 0; // 缓冲区中数据项数目

void producer(void) {
    int item;
    while (TRUE) {
        item = produce_item();
        if (count == N) { sleep(); }
        insert_item(item);
        count = count + 1;
        if (count == 1) { wakeup(consumer);}
    }
}

void consumer(void) {
    int item;

    while (TRUE) {
        if (count == 0) { sleep(); }
        item = remove_item();
        count = count - 1;
        if (count == N - 1) { wakeup(producer); }
        consumer_item(item);
    }
}