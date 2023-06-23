package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	InitialStage Stage = iota
	FirstStage
	SecondStage
	ThirdStage
)

type Stage int

type job struct {
	num   int
	value int
	state Stage
}

func main() {
	cnt := 20
	in := make(chan job, cnt)

	for i := 0; i < cnt; i++ {
		var j job
		j.num = i
		j.value = rand.Intn(1000)
		in <- j
	}
	close(in)

	out := ThirdProcessing(SecondProcessing(FirstProcessing(in)))

	for range out {
	}
}

func FirstProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for val := range in {
			fmt.Println(val)
			val.value = int(float64(val.value) / math.Pi)
			val.state = FirstStage
			out <- val
		}
		close(out)
	}()

	return out
}

func SecondProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for val := range in {
			fmt.Println(val)
			val.value = int(float64(val.value) * math.E)
			val.state = SecondStage
			out <- val
		}
		close(out)
	}()

	return out
}

func ThirdProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for val := range in {
			fmt.Println(val)
			val.value = val.value - 13
			val.state = ThirdStage
			out <- val
		}
		close(out)
	}()

	return out
}
