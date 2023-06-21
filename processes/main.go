package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	InitialState State = iota
	FirstState
	SecondState
	ThirdState
)

type State int

type job struct {
	num   int
	value int
	state State
}

func FirstProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for j := range in {
			fmt.Println(j)
			j.state = FirstState
			j.value = int((float64(j.value)*8 + math.Pi) / float64(rand.Intn(9)+1))
			out <- j
		}
		close(out)
	}()

	return out
}

func SecondProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for j := range in {
			fmt.Println(j)
			j.value = int((float64(j.value)*8 + math.E) / float64(rand.Intn(9)+1))
			j.state = SecondState
			out <- j
		}
		close(out)
	}()

	return out
}

func ThirdProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for j := range in {
			fmt.Println(j)
			j.value = int(j.value * 16 / (rand.Intn(9) + 1))
			j.state = ThirdState
			out <- j
		}
		close(out)
	}()

	return out
}

func main() {
	cnt := 10
	in := make(chan job, 10)

	for i := 0; i < cnt; i++ {
		var j job
		j.num, j.value = i, i
		in <- j
	}
	close(in)

	out := ThirdProcessing(SecondProcessing(FirstProcessing(in)))
	for j := range out {
		fmt.Println(j)
	}
}
