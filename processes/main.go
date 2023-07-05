package main

import (
	"fmt"
	"math"
)

type State int

const (
	InitialState State = iota
	FirstStage
	SecondStage
	ThirdStage
)

type job struct {
	num   int
	state State
	value int
}

func main() {
	cnt := 10
	in := make(chan job, cnt)
	go func() {
		for i := 0; i < cnt; i++ {
			in <- job{num: i, value: i}
		}
		close(in)
	}()
	res := ThirdProcessing(SecondProcessing(FirstProcessing(in)))

	for val := range res {
		fmt.Println(val)
	}
}

func FirstProcessing(in chan job) chan job {
	output := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = FirstStage
			job.value = 1000 + int(float64(job.value)*math.Pi)
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
			job.value *= int(float64(job.value) * math.E)
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
			job.value = int(job.value / 1000)
			output <- job
		}
		close(output)
	}()
	return output
}
