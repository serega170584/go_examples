// Что выведет следующая программа и почему?

// package main

// import "fmt"

// type Person struct {
//     Name string
// }

// func changeName(person *Person) {
//     person.Name = "Alice"
// }

// func main() {
//     person := &Person{
//         Name: "Bob",
//     }
//     fmt.Println(person.Name) // Bob
//     changeName(person)
//     fmt.Println(person.Name) // Bob
// }

// Что выведет следующая программа и почему?

// package main

// import (
//     "fmt"
//     "sync"
// )

// func main() {
//     var max int

//     mu := sync.Mutex{}

//     wg := &sync.WaitGroup{}
//     wg.Add(1000)

//     for i := 1000; i > 0; i-- {
//         i := i
//         go func() {
//             defer wg.Done()
//             defer mu.Unlock()
//             mu.Lock()
//             if i%2 == 0 && i > max {
//                 max = i
//             }
//         }()
//     }

//     wg.Wait()

//     fmt.Printf("Maximum is %d", max) // 1000
// }

// Есть функция unpredictableFunc, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
// Нужно написать обёртку predictableFunc, которая будет работать
// с заданным фиксированным таймаутом (например, 1 секунду).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

func predictableFunc() (int64, bool) {
	ch := make(chan int64)
	var ok bool
	firstCh := make(chan struct{})
	secondCh := make(chan struct{})

	go func() {
		res := unpredictableFunc()
		firstCh <- struct{}{}
		ok = true
		ch <- res
	}()

	go func() {
		time.Sleep(time.Second)
		secondCh <- struct{}{}
		ch <- -1
	}()

	select {
	case <-firstCh:
		res := <-ch
	case <-secondCh:
		res := <-ch
	}

	// res := <-ch
	return res, ok
}
func main() {
	fmt.Println("started")

	res, ok := predictableFunc()
	if !ok {
		fmt.Println("error")
		return
	}

	fmt.Println(res)
}
