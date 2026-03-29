package main

import (
	"fmt"
)

const (
	horizontalLayout = "horizontal"
	verticalLayout   = "vertical"
)

func main() {
	var (
		a int   = 3
		b int32 = 3
	)

	fmt.Println(a == b) // true

	fmt.Println(handler(&horizontalLayout)) // ?
}

func handler(layout *string) string {
	if layout == nil {
		return ""
	}

	return *layout
}

--------------
package main

import "fmt"

// Вопрос: что выведет функция len?
func main() {
	strEn := "John"  // 4 буквы англ.алфавита
	strRU := "Маша"  // 4 буквы русского алфавита
	strChine := "傑基" // 2 китайских иероглифа

	fmt.Println(len(strEn)) // 4
	fmt.Println(len(strRU)) // 8
	fmt.Println(len(strChine)) // 6-8
}

------------
package main

import "fmt"

func main() {
	sl := make([]int, 1, 3)
	fmt.Println(sl) // [0] l=1 c=3

	appendSlice(sl, 1)
	fmt.Println(sl)     // ? [0] l=1 c=3
	fmt.Println(sl[:2]) // ? [0] l=1 c=3

	copySlice(sl, []int{2, 3})
	fmt.Println(sl) // ? [2] l=1 c=3

	mutateSlice(sl, 1, 4)
	fmt.Println(sl) // ?
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val) // l=2, c=3, [0, 1], 0
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
}

func mutateSlice(sl []int, ind, val int) {
	sl[ind] = val
}

---------------
package main

import (
"fmt"
"time"
)

func main() {
	sl := make([]int, 0, 10)

	for i := 0; i < 100_000_000; i++ {
		go func() {
			sl = append(sl, i)
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(len(sl)) // 0 < N < 100_000
}

---------------
package main

import "fmt"

// Что выведен функция main?
func main() {
	var m map[int]int

	m := make(map[int]int, 2)

	val := m[55]
	fmt.Println(val)

	m[1] = 10
	m[2] = 20

	fmt.Println(m)
}

--------------
package main

import (
"fmt"
"time"
)

func main() {

	go func() {

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()

			var a *struct{ Foo int }
			fmt.Println("A")
			_ = a.Foo
		}()

		time.Sleep(time.Second)
	}()

	time.Sleep(time.Second)
}

---------------



// testSearcher - функция проверяет время отклика поисковика.
// Если поисковик не ответил, или ответил с ошибкой, то возвращает ошибку
func testSearcher(ctx context.Context, url string) (time.Duration, error) {
	// просто заглушка
	return 1 * time.Second, nil
}

// getFastestSearcher - возвращает самый быстрый поисковик из списка и его время ответа
func getFastestSearcher(ctx context.Context, urls []string) (name string, respTime time.Duration, err error) {
	m := make(map[string]time.Duration, len(urls))

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	sem := make(chan struct{}, 2)

	for url := range urls {
		go func() {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {<-sem}()

			duration, err := testSearcher(ctx, url)
			if err != nil {
				return "", nil, errors.New("")
			}

			mu.Lock()
			defer mu.Unlock()
			m[url] = duration
		}()
	}

	wg.Wait()

	if len(m) == 0 {
		return "", nil, errors.New("All requests failed")
	}

	minUrl := ""
	minDuration := m[0]

	// [g: 1, t: 2, c: 3]

	// 1. g 1 <= 1
	// 2. t 2 <= 1
	// 3. c 3 <= 1
	for url, duration := range m {
		if duration <= minDuration {
			minUrl = url
			minDuration = duration
		}
	}

	return minUrl, minDuration, nil
}


