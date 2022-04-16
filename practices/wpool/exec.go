package wpool

import (
	"context"
	"fmt"
	"sync"
)

type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	Done         chan struct{}
}

func New(wCount int) WorkerPool {
	return WorkerPool{
		workersCount: wCount,
		jobs:         make(chan Job, wCount),
		results:      make(chan Result, wCount),
		Done:         make(chan struct{}),
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		// fan out worker goroutines
		// reading from jobs channel and
		// pushing cacls into results channel
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp WorkerPool) GenerateFrom(jobsBulk []Job) {
	for i := range jobsBulk {
		wp.jobs <- jobsBulk[i]
	}
	close(wp.jobs)
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// fan-in job execution multiplexing results into the results channel
			results <- job.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}
