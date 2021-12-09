package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	value int64
}

type Result struct {
	job *Job
	sum int64
}

var jobC = make(chan *Job, 100)
var resultC = make(chan *Result, 100)
var wg sync.WaitGroup

func worker(ch1 chan<- *Job) {
	//循环生成int64类型的随机数，发送到jobC

	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Second)

	}
}

func leader(ch1 <-chan *Job, reCh1 chan<- *Result) {
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			job: job,
			sum: sum,
		}
		reCh1 <- newResult
	}
}

func main() {
	wg.Add(1)
	go worker(jobC)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go leader(jobC, resultC)
	}
	for result := range resultC {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
