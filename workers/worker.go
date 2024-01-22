package workers

import (
	"log"
	"sync"
	"time"
)

type Pool interface {
	Jobs(jobs []int) Pool

	Wait(wg *sync.WaitGroup) Pool
	Run(req chan int, res chan int) error
}

type queue struct {
	jobs         []int
	numberWorker int
	wg           *sync.WaitGroup
}

func NewWorker(number int) Pool {
	return &queue{
		numberWorker: number,
	}
}

func (q *queue) Wait(wg *sync.WaitGroup) Pool {
	q.wg = wg
	return q
}
func (q *queue) Jobs(jobs []int) Pool {
	q.jobs = jobs
	return q
}

func (q *queue) Run(req chan int, res chan int) error {
	for i := 1; i <= q.numberWorker; i++ {
		q.wg.Add(1)
		go worker(i, req, res, q.wg)
	}

	return nil
}
func worker(id int, req chan int, result chan int, wg *sync.WaitGroup) {
	for job := range req {
		log.Printf("woker id = %d is handle job = %d", id, job)
		time.Sleep(time.Second)
		result <- id
	}
	wg.Done()
}
