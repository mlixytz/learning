package singleton

import "sync"

// Singleton 单例结构
type Singleton struct{}

var instance *Singleton
var mu sync.Mutex
var once sync.Once

// GetInstanceLazyLoading 懒汉模式
func GetInstanceLazyLoading() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// GetInstanceWithMutex 带锁的单例模式
func GetInstanceWithMutex() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// GetInstanceWithMutexDoubleCheck 带锁的双重检查
func GetInstanceWithMutexDoubleCheck() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

// GetInstance 最完美的单例
func GetInstance() *Singleton {
	once.Do(func() { instance = &Singleton{} })
	return instance
}
