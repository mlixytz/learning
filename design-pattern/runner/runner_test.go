package runner

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	createTask := func() func(int) {
		return func(id int) {
			log.Printf("Processor - Task #%d.", id)
			time.Sleep(time.Duration(id) * time.Second)
		}
	}

	const timeout = 3 * time.Second
	log.Println("Starting work...")
	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}
