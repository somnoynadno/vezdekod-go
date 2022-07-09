package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const FILE = "../input.txt"

func worker(i int, t time.Duration) {
	fmt.Printf("task %d started...\n", i)
	time.Sleep(t)
	fmt.Printf("task %d finished!\n", i)
}

func main() {
	data, err := os.ReadFile(FILE)
	if err != nil{
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	var wg sync.WaitGroup
	wg.Add(len(lines))

	for i, line := range lines {
		t, err := time.ParseDuration(line)
		if err != nil {
			fmt.Printf("parse error occurred, skipping task %d: %v\n", i, err)
		}

		i := i
		go func() {
			defer wg.Done()
			worker(i, t)
		}()
	}

	wg.Wait()
}
