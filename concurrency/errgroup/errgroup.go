package errgroup

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type Group struct {
	wg      sync.WaitGroup
	err     error
	cancel  context.CancelFunc
	errOnce sync.Once
	ctx     context.Context
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel, ctx: ctx}, ctx
}

func (g *Group) process(f func(context.Context) error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done()
	}()
	g.wg.Add(1)
	err = f(g.ctx)
}

func (g *Group) Go(f func(context.Context) error) {
	if g.ctx != nil {
		select {
		case <-g.ctx.Done():
			return
		default:
			g.process(f)
		}
	} else {
		g.process(f)
	}
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
