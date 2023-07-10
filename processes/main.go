package main

import (
	"fmt"
	"math"
	"math/rand"
)

type State int

type job struct {
	num   int
	state State
	value int
}

const (
	InitialStage State = iota
	FirstStage
	SecondStage
	ThirdStage
)

func main() {
	cnt := 20
	in := make(chan job, cnt)
	go func() {
		for i := 0; i < cnt; i++ {
			in <- job{num: i, value: i}
		}
		close(in)
	}()

	res := ThirdProcessing(SecondProcessing(FirstProcessing(in)))
	for job := range res {
		fmt.Println(job)
	}
}

func FirstProcessing(in chan job) chan job {
	out := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = FirstStage
			job.value = int(float64(job.value+1000) * math.Pi)
			out <- job
		}
		close(out)
	}()
	return out
}

func SecondProcessing(in chan job) chan job {
	out := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = SecondStage
			job.value = int(float64(job.value) * math.E)
			out <- job
		}
		close(out)
	}()
	return out
}

func ThirdProcessing(in chan job) chan job {
	out := make(chan job)
	go func() {
		for job := range in {
			fmt.Println(job)
			job.state = ThirdStage
			job.value = job.value / (rand.Intn(5) + 1)
			out <- job
		}
		close(out)
	}()
	return out
}
