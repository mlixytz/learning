package errgroup

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {
	var g Group

	// 启动第一个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(5 * time.Second)
		t.Log("exec #1")
		return nil
	})

	// 启动第二个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(10 * time.Second)
		t.Log("exec #2")
		return errors.New("failed to exec #2")
	})

	// 启动第三个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(15 * time.Second)
		t.Log("exec #3")
		return nil
	})

	if err := g.Wait(); err == nil {
		t.Error("Successfully exec all")
	} else {
		t.Log("failed:", err)
	}
}

func TestErrGroupWithCancel(t *testing.T) {
	g, _ := WithContext(context.Background())

	time.Sleep(3 * time.Second)

	// 启动第一个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(5 * time.Second)
		t.Log("exec #1")
		return nil
	})

	// 启动第二个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(10 * time.Second)
		t.Log("exec #2")
		return errors.New("failed to exec #2")
	})

	// 启动第三个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(15 * time.Second)
		t.Log("exec #3")
		return nil
	})

	if err := g.Wait(); err == nil {
		t.Error("Successfully exec all")
	} else {
		t.Log("failed:", err)
	}
}

func TestErrGroupWithCancel2(t *testing.T) {
	g, ctx := WithContext(context.Background())

	// 启动第一个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(5 * time.Second)
		g.cancel()
		t.Log("exec #1")
		return nil
	})

	// 启动第二个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(10 * time.Second)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			t.Log("exec #2")
			return errors.New("failed to exec #2")
		}
	})

	if err := g.Wait(); err == nil {
		t.Error("Successfully exec all")
	} else {
		t.Log("failed:", err)
	}
}

func TestErrGroupRecover(t *testing.T) {
	var g Group

	// 启动第一个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(5 * time.Second)
		t.Log("exec #1")
		a := []int{}
		t.Log(a[2])
		return nil
	})

	// 启动第二个任务，执行成功
	g.Go(func(context.Context) error {
		time.Sleep(10 * time.Second)
		t.Log("exec #2")
		return errors.New("failed to exec #2")
	})

	if err := g.Wait(); err == nil {
		t.Error("Successfully exec all")
	} else {
		t.Log("failed:", err)
	}
}
