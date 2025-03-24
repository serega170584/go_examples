package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type TimeInterval struct {
	t   time.Time
	cnt int64
}

type RateCounter struct {
	intervalCnt         int64
	firstTs             time.Time
	intervals           chan *TimeInterval
	currentInterval     *TimeInterval
	intervalDuration    time.Duration
	calculationInterval time.Duration
	cnt                 int64
	mu                  sync.Mutex
}

func NewRateCounter(
	intervalCnt int64,
	firstTs time.Time,
	intervalDuration time.Duration,
	calcualtionInterval time.Duration,
) *RateCounter {
	intervals := make(chan *TimeInterval, intervalCnt)
	return &RateCounter{
		intervalCnt:      intervalCnt,
		firstTs:          firstTs,
		intervals:        intervals,
		intervalDuration: intervalDuration,
		currentInterval: &TimeInterval{
			t: firstTs,
		},
		calculationInterval: calcualtionInterval,
		mu:                  sync.Mutex{},
	}
}

func (rc *RateCounter) add(cnt int64) {
	ct := time.Now()
	intervalStart := time.Duration(1) * rc.intervalDuration
	go func() {
		i := 1
		for ct.After(rc.currentInterval.t.Add(intervalStart)) {
			i++
			intervalStart = time.Duration(i) * rc.intervalDuration
		}
		if i > 1 {
			rc.mu.Lock()
			rc.currentInterval = &TimeInterval{
				t:   rc.currentInterval.t.Add(time.Duration(i-1) * rc.intervalDuration),
				cnt: cnt,
			}
			rc.intervals <- rc.currentInterval
			rc.mu.Unlock()
		} else {
			atomic.AddInt64(&rc.currentInterval.cnt, cnt)
		}
	}()
}

func (rc *RateCounter) calculateIntervalCntSum() {
	ticker := time.NewTicker(rc.calculationInterval)
	ct := time.Now()
	go func() {
		for {
			select {
			case <-ticker.C:
				rc.mu.Lock()
				var cntSum int64 = 0
				firstIntervalTime := ct.Add(-time.Duration(rc.intervalCnt) * rc.intervalDuration)
				out := make(chan *TimeInterval, rc.intervalCnt)
				for ti := range rc.intervals {
					if ti.t.After(firstIntervalTime) {
						out <- ti
						cntSum += ti.cnt
					}
				}
				atomic.AddInt64(&rc.cnt, cntSum)
				rc.intervals = out
				rc.mu.Unlock()
			}
		}
	}()
}

func main() {

}
