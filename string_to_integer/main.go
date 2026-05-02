package main

import (
	"fmt"
	"time"
)

// v <= 5
//5 3 1
// l=1 r=5 m=3
// 2 3 2  7 > 5
// l=1 r=2 m=1
// 0 1 0 1 < 5
// l=2 r=2 условие выполнено
//

func main() {
	highChan := make(chan int)
	lowChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		highChan <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		lowChan <- 2
	}()

	for {
		isHandled := func() bool {
			select {
			case v := <-highChan:
				fmt.Println("High:", v)
				return true
			default:
				select {
				case v := <-lowChan:
					fmt.Println("Low:", v)
					return true
				default:
					return false
				}
			}
		}()

		if isHandled {
			break
		}
	}
}
