package main

import (
	_ "expvar"
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const FILE = "../input.txt"

// Job holds the attributes needed to perform unit of work.
type Job struct {
	Name  string
	Delay time.Duration
}

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

func (w Worker) start(wg *sync.WaitGroup) {
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to my jobQueue.
				fmt.Printf("worker%d: started %s, blocking for %f seconds\n", w.id, job.Name, job.Delay.Seconds())
				time.Sleep(job.Delay)
				fmt.Printf("worker%d: completed %s!\n", w.id, job.Name)
				wg.Done()
			case <-w.quitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.id)
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
	wg         *sync.WaitGroup
}

func (d *Dispatcher) run(tasksNum int) {
	var wg sync.WaitGroup
	d.wg = &wg
	d.wg.Add(tasksNum)

	fmt.Println("dispatcher is started!")
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool)
		worker.start(d.wg)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				fmt.Printf("fetching workerJobQueue for: %s\n", job.Name)
				workerJobQueue := <-d.workerPool
				fmt.Printf("adding %s to workerJobQueue\n", job.Name)
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) await() {
	d.wg.Wait()
}

func main() {
	var (
		maxWorkers   = flag.Int("max_workers", 2, "The number of workers to start")
		maxQueueSize = flag.Int("max_queue_size", 10, "The size of job queue")
	)
	flag.Parse()

	data, err := os.ReadFile(FILE)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	// Create the job queue.
	jobQueue := make(chan Job, *maxQueueSize)

	// Start the dispatcher.
	dispatcher := NewDispatcher(jobQueue, *maxWorkers)
	dispatcher.run(len(lines))

	for i, line := range lines {
		t, err := time.ParseDuration(line)
		if err != nil {
			fmt.Printf("parse error occurred, skipping task %d: %v\n", i, err)
		}

		job := Job{Name: "task" + strconv.Itoa(i), Delay: t}
		jobQueue <- job
	}

	dispatcher.await()
	fmt.Println("all jobs done, dispatcher exits")
}
