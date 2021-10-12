package runner_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/efsn/go-workshop-101/internal/runner"
	"github.com/efsn/go-workshop-101/internal/work"
)

const timeout = 3 * time.Second

func TestRunner_Add(t *testing.T) {
	log.Println("start work")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminate due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("terminate due to interrupt")
			os.Exit(2)
		}
	}
	runWork()
	log.Println("process ended")
}

func runWork() {
	var w work.Worker
	w = runner.New(1)
	w.Task()
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("process - task #%d", i)
		time.Sleep(time.Duration(i))
	}
}
