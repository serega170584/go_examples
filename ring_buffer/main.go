package main

import "fmt"

func NewRingBuffer(inCh, outCh chan int) *RingBuffer {
	return &RingBuffer{
		inCh:  inCh,
		outCh: outCh,
	}
}

type RingBuffer struct {
	inCh  chan int
	outCh chan int
}

func (r *RingBuffer) Run() {
	defer close(r.outCh)
	for v := range r.inCh {
		select {
		case r.outCh <- v:
		default:
			<-r.outCh
			r.outCh <- v
		}
	}
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int, 4)
	rb := NewRingBuffer(inCh, outCh)
	go rb.Run()

	for i := 1; i <= 10; i++ {
		inCh <- i
	}

	close(inCh)

	for res := range outCh {
		fmt.Println(res)
	}

}
