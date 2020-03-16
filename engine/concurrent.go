package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) // 100个协程
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %d : %v", itemCount, item)
			fmt.Println("获得数据。。。。。。")

			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
