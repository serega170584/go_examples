package main

import (
	"fmt"
	"time"
)

type CounterInterface interface {
	inc(cnt int)
	alert()
}

type Counter struct {
	partsCnt     uint
	windowSize   uint
	timeBaseUnit time.Duration
	startTime    time.Time
	cnts         map[time.Time]int
	threshold    uint
}

func New(partsCnt uint, windowSize uint, timeBaseUnit time.Duration, startTime time.Time, threshold uint) *Counter {
	interval := int(windowSize) / int(partsCnt)
	cnts := make(map[time.Time]int)
	for i := 0; i < int(partsCnt); i++ {
		t := startTime.Add(-time.Duration(i*interval) * timeBaseUnit)
		cnts[t] = 0
	}
	return &Counter{
		partsCnt:     partsCnt,
		windowSize:   windowSize,
		timeBaseUnit: timeBaseUnit,
		startTime:    startTime,
		cnts:         cnts,
		threshold:    threshold,
	}
}

func (c *Counter) inc(cnt int) {
	interval := int(c.windowSize) / int(c.partsCnt)
	partNum := int(time.Now().Sub(c.startTime)) / int(time.Duration(interval)*c.timeBaseUnit)
	t := c.startTime.Add(time.Duration(partNum*interval) * c.timeBaseUnit)
	c.cnts[t] += cnt
}

func (c *Counter) alert() {
	interval := int(c.windowSize) / int(c.partsCnt)
	partNum := int(time.Now().Sub(c.startTime)) / int(time.Duration(interval)*c.timeBaseUnit)
	t := c.startTime.Add(time.Duration(partNum*interval) * c.timeBaseUnit)
	cnt := 0
	for i := 0; i < int(c.partsCnt); i++ {
		cnt += c.cnts[t.Add(-time.Duration(i)*c.timeBaseUnit)]
		if cnt > int(c.threshold) {

		}
	}
}

// | | | | | | |
func main() {
	var a uint
	var b uint
	a = 2
	b = 3
	fmt.Println(a + b)
	t := time.Now()
	t1 := t.Add(-3 * time.Second)
	//fmt.Println(uint(t.Sub(t1)))
	fmt.Println(uint(t.Sub(t1)) / uint(2*time.Second))
}
