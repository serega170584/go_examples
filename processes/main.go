package main

import (
	"fmt"
	"math"
)

type State int

const (
	InitialStage State = iota
	FirstStage
	SecondStage
	ThirdStage
)

type job struct {
	num   int
	state State
	value int64
}

func main() {
	cnt := 20
	in := make(chan job, cnt)
	for i := 0; i < cnt; i++ {
		job := job{num: i, value: int64(i)}
		in <- job
	}
	close(in)

	output := ThirdProcessing(SecondProcessing(FirstProcessing(in)))
	for job := range output {
		fmt.Println(job)
	}
}

func FirstProcessing(in chan job) chan job {
	output := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = FirstStage
			job.value = job.value + 1000
			output <- job
		}
		close(output)
	}()
	return output
}

func SecondProcessing(in chan job) chan job {
	output := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = SecondStage
			job.value = int64(float64(job.value) * math.Pi)
			output <- job
		}
		close(output)
	}()
	return output
}

func ThirdProcessing(in chan job) chan job {
	output := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = ThirdStage
			job.value = int64(float64(job.value) * math.E)
			output <- job
		}
		close(output)
	}()
	return output
}
