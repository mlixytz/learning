package work

import "sync"

// Worker 使用Pool的接口类型
type Worker interface {
	Task()
}

// Pool goroutines池
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建goroutines池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// ShutDown 等待所有的goroutine停止工作
func (p *Pool) ShutDown() {
	close(p.work)
	p.wg.Wait()
}
