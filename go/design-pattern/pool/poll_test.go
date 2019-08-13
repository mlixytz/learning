package pool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxGoroutines   = 25 // 要使用的goroutine的数量
	pooledResources = 2  // 池中资源的数量
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func TestPool(t *testing.T) {
	performQueries := func(query int, p *Pool) {
		conn, err := p.Acquire()
		if err != nil {
			log.Println(err)
			return
		}
		defer p.Release(conn)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
	}
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program.")
	p.Close()
}
