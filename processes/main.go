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
	stage Stage
	value int
}

func FirstProcessing(in chan job) chan job {
	out := make(chan job)

	go func() {
		for val := range in {
			fmt.Println(val)
			val.stage, val.value = FirstStage, int(float64(rand.Intn(100)*val.value)/math.Pi)
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
			val.stage, val.value = SecondStage, int(float64(rand.Intn(100)*val.value)/math.E)
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
			val.stage, val.value = ThirdStage, int(rand.Intn(100)*val.value/123)
			out <- val
		}
		close(out)
	}()

	return out
}

func main() {
	cnt := 20
	in := make(chan job, cnt)

	for i := 0; i < cnt; i++ {
		var j job
		j.num, j.value = i, rand.Intn(100)
		in <- j
	}
	close(in)

	out := ThirdProcessing(SecondProcessing(FirstProcessing(in)))
	for val := range out {
		fmt.Println(val)
	}
}
