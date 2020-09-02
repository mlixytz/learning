'''
    计数器限流法
    
    采用滑动窗口的方式实现，如果不采用滑动窗口的话，会出现临界问题
    举例：
        假设每秒允许访问100次，则设置一个1秒钟的滑动窗口，窗口中有10个格子，
    每个格子100ms，窗口每100ms移动一次，格子里存储计数器的值。每次移动比较
    窗口最后一个格子的和第一个格子，如果差值大于100，则限流。（格子越多越平滑）。
'''

import threading
import time


counter = 0

# 列表中元素存储的值为：{time，count}
windows = []

def accept():
    if grant():
        print("请求完成！")
    else:
        print("被限制了！")

def grant():
    now = int(time.time() / 1000))
    for window in windows:
        if 
