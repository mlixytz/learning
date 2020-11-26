package pool

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Task struct {
	Handler func(...interface{})
	Params  []interface{}
}

type Pool struct {
	taskChan       chan *Task
	capacity       uint64
	runningWorkers uint64
	state          int64
	sync.Mutex
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
	RUNNING = 1
	STOPED  = 0
)

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{
		taskChan: make(chan *Task, capacity),
		capacity: capacity,
		state:    RUNNING,
	}, nil
}

func (p *Pool) run() {
	p.incRunning()
	go func() {
		defer func() {
			p.decRunning()
		}()
		for {
			select {
			case task, ok := <-p.taskChan:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			}
		}
	}()

}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) getRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) getCapacity() uint64 {
	return p.capacity
}

func (p *Pool) getState() int64 {
	p.Lock()
	defer p.Unlock()
	return p.state
}

func (p *Pool) setState(state int64) {
	p.Lock()
	defer p.Unlock()

	p.state = state
}

var ErrPoolAlreadyClosed = errors.New("pool already closed")

func (p *Pool) Put(task *Task) error {
	if p.getState() == STOPED {
		return ErrPoolAlreadyClosed
	}
	p.Lock()
	if p.getRunningWorkers() < p.getCapacity() {
		p.run()
	}
	p.Unlock()

	p.Lock()
	if p.state == RUNNING {
		p.taskChan <- task
	}
	p.Unlock()

	return nil
}

func (p *Pool) Close() {
	p.setState(STOPED)
	for len(p.taskChan) > 0 {
	}
	close(p.taskChan)
}
