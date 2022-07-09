package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const URL = "http://localhost:8000"

type Task struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

func postTaskAsync(task Task) error {
	fmt.Printf("sending task %s (duration %s)\n", task.Name, task.Duration)

	j, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", URL + "/add/async", bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("http response status: " + resp.Status)
	}

	return nil
}

func postTaskSync(task Task) error {
	fmt.Printf("sending task %s (duration %s)\n", task.Name, task.Duration)

	j, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", URL + "/add/sync", bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("http response status: " + resp.Status)
	}

	return nil
}

func getTime() (time.Duration, error) {
	resp, err := http.Get(URL + "/time")
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return time.ParseDuration(string(body))
}

func getSchedule() ([]Task, error) {
	resp, err := http.Get(URL + "/schedule")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res []Task
	err = json.Unmarshal(body, &res)

	return res, err
}

// scenario1 checks that server time evaluation is correct.
func scenario1() {
	fmt.Println("starting test scenario 1...")

	durations := []time.Duration{3 * time.Second, 2 * time.Second, 1 * time.Second}
	var sum time.Duration

	for i, d := range durations {
		err := postTaskAsync(Task{Duration: d.String(), Name: strconv.Itoa(i)})
		if err != nil {
			panic(err)
		}
		sum += d
	}

	t, err := getTime()
	if err != nil {
		panic(err)
	}

	if t != sum {
		panic(fmt.Sprintf("expect duration %d, got %d", sum, t))
	}

	fmt.Println("expected duration is equal to real one")
	fmt.Println("scenario 1 passed!")
	time.Sleep(sum)
}

// scenario2 checks that schedule is correct and can be parsed.
func scenario2() {
	fmt.Println("starting test scenario 2...")

	durations := []time.Duration{1 * time.Second, 2 * time.Second, 1 * time.Second}
	var sum time.Duration

	for i, d := range durations {
		err := postTaskAsync(Task{Duration: d.String(), Name: strconv.Itoa(i)})
		if err != nil {
			panic(err)
		}
		sum += d
	}

	schedule, err := getSchedule()
	if err != nil {
		panic(err)
	}

	for i, task := range schedule {
		d, err := time.ParseDuration(task.Duration)
		if err != nil {
			panic("malformed task duration " + task.Duration)
		}

		if d != durations[i] {
			panic(fmt.Sprintf("expected duration %s, got %s", durations[i].String(), task.Duration))
		}
	}

	fmt.Println("schedule is equal to expected one")
	fmt.Println("scenario 2 passed!")
	time.Sleep(sum)
}

// scenario3 checks that sync task creation works as expected.
func scenario3() {
	fmt.Println("starting test scenario 3...")

	d := 3 * time.Second
	variance := 1 * time.Second

	start := time.Now()
	err := postTaskSync(Task{Duration: d.String(), Name: "sync"})
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	if elapsed - variance < d && d < elapsed + variance {
		fmt.Println("elapsed time is in acceptable interval")
	} else {
		panic(fmt.Sprintf("expected elapsed time %s +/- %s, got %s",
			d.String(), variance.String(), elapsed.String()))
	}

	schedule, err := getSchedule()
	if err != nil {
		panic(err)
	}
	if schedule != nil {
		panic("schedule is not empty")
	}

	fmt.Println("sync task creation works as expected")
	fmt.Println("scenario 3 passed!")
}

func main() {
	scenario1()
	time.Sleep(1 * time.Second)
	fmt.Println("------------------")

	scenario2()
	time.Sleep(1 * time.Second)
	fmt.Println("------------------")

	scenario3()
	time.Sleep(1 * time.Second)
	fmt.Println("------------------")
}
