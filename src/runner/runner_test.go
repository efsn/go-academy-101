package runner_test

import (
	"go-workshop-101/src/runner"
	"log"
	"os"
	"testing"
	"time"
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
	log.Println("process ended")
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("process - task #%d", i)
		time.Sleep(time.Duration(i))
	}
}
