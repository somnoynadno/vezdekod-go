package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FILE = "../input.txt"

func main() {
	data, err := os.ReadFile(FILE)
	if err != nil{
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		t, err := time.ParseDuration(line)
		if err != nil {
			fmt.Printf("parse error occurred, skipping task %d: %v\n", i, err)
		}

		fmt.Printf("task %d started...\n", i)
		time.Sleep(t)
		fmt.Printf("task %d finished!\n", i)
	}
}
