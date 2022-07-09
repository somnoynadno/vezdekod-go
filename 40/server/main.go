package main

import (
	"container/list"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var queue *list.List

type Task struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	readerSleepTime = 1 * time.Second
	syncCheckSleepTime = 100 * time.Millisecond
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func addSyncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = time.ParseDuration(t.Duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.Name == "" {
		t.Name = RandStringBytes(10)
	}

	fmt.Printf("pushing task %s (duration %s) to the queue (sync)\n", t.Name, t.Duration)
	queue.PushBack(t)

	for {
		found := false
		for e := queue.Front(); e != nil; e = e.Next() {
			p, ok := e.Value.(Task)
			if ok && t == p{
				found = true
			}
		}

		if !found {
			fmt.Printf("task %s is finished, closing connection\n", t.Name)
			return
		}

		fmt.Printf("[debug] task %s is still in queue...\n", t.Name)
		time.Sleep(syncCheckSleepTime)
	}
}

func addAsyncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = time.ParseDuration(t.Duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.Name == "" {
		t.Name = RandStringBytes(10)
	}

	fmt.Printf("pushing task %s (duration %s) to the queue (async)\n", t.Name, t.Duration)
	queue.PushBack(t)
}

func scheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var res []interface{}
	for e := queue.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var s time.Duration
	for e := queue.Front(); e != nil; e = e.Next() {
		t, ok := e.Value.(Task)
		if ok {
			d, _ := time.ParseDuration(t.Duration)
			s += d
		}
	}

	_, _ = w.Write([]byte(s.String()))
}

func reader(queue *list.List) {
	fmt.Println("starting queue reader...")
	for {
		e := queue.Front()
		if e == nil {
			fmt.Println("task queue is empty, chilling...")
			time.Sleep(readerSleepTime)
		} else {
			t, ok := e.Value.(Task)
			if !ok {
				panic("garbage in queue >_<")
			}

			fmt.Printf("task %s (duration %s) started...\n", t.Name, t.Duration)
			d, _ := time.ParseDuration(t.Duration)
			time.Sleep(d)

			fmt.Printf("task %s (duration %s) finished!\n", t.Name, t.Duration)
			queue.Remove(e)
		}
	}
}

func main() {
	var (
		port = flag.String("port", "8000", "The server port")
	)
	flag.Parse()

	queue = list.New()

	// Start the HTTP handler.
	http.HandleFunc("/add/sync", addSyncHandler)
	http.HandleFunc("/add/async", addAsyncHandler)
	http.HandleFunc("/schedule", scheduleHandler)
	http.HandleFunc("/time", timeHandler)

	go reader(queue)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
