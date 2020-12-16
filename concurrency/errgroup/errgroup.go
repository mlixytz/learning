package errgroup

import (
	"context"
	"sync"
)

type Group struct {
	wg  sync.WaitGroup
	err error
	cf  context.CancelFunc
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	c, cancel := context.WithCancel(context.Background())
	return &Group{
		cf: cancel,
	}, c
}

func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		err := f()
		if g.err == nil && err != nil {
			g.err = err
			g.cf()
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.err == nil {
		g.cf()
	}
	return g.err
}
