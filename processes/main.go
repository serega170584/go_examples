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
	stage Stage
}

func main() {
	cnt := 20
	in := make(chan job, cnt)
	for i := 0; i < cnt; i++ {
		var val job
		val.num, val.value = i, i
		in <- val
	}
	close(in)
	out := ThirdProcessing(SecondProcessing(FirstProcessing(in)))

	for val := range out {
		fmt.Println(val)
	}
}

func FirstProcessing(in chan job) chan job {
	out := make(chan job)
	go func() {
		for val := range in {
			fmt.Println(val)
			val.stage = FirstStage
			val.value = rand.Intn(1000) + rand.Intn(1000)*val.value
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
			val.stage = SecondStage
			val.value = rand.Intn(800) + int(float64(rand.Intn(val.value))*math.Pi)
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
			val.stage = ThirdStage
			val.value = rand.Intn(600) + int(float64(rand.Intn(val.value))*math.E)
			out <- val
		}
		close(out)
	}()
	return out
}
