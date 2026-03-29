package main

import (
	"context"
	"fmt"
	"sync"
	"time"
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

	fmt.Println(a == b) // ?

	fmt.Println(handler(&horizontalLayout)) // ?
}

func handler(layout *string) string {
	if layout == nil {
		return ""
	}

	return *layout
}

-----------

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

--------

package main

import "fmt"

func main() {
	sl := make([]int, 1, 3)
	fmt.Println(sl) // [0] l=1 c=3

	appendSlice(sl, 1)
	fmt.Println(sl)     // // [0] l=1 c=3 // [0, 1]
	fmt.Println(sl[:2]) //  // [0, 1] l=2 c=3

	copySlice(sl, []int{2, 3})
	fmt.Println(sl) // [2] l=1  c= 3

	mutateSlice(sl, 1, 4)
	fmt.Println(sl) // out of range
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
}

func mutateSlice(sl []int, ind, val int) {
	sl[ind] = val
}

----------

package main

import (
"fmt"
"time"
)

func main() {
	sl := make([]int, 0, 10)

	for i := 0; i < 100_000_0; i++ {
		go func() {
			sl = append(sl, i)
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(len(sl)) // 80к<90к
}

-------------

package main

import "fmt"

// Что выведен функция main?
func main() {
	m := make(map[int]int, 2)

	val := m[55]
	fmt.Println(val) // 0

	m[1] = 10
	m[2] = 20

	fmt.Println(m)
}

------

package main

import (
"fmt"
"time"
)

type User struct {
	Id   int
	Name string
}

func main() {

	go func() {

		a:=1
		defer func (){func1(a)}

		a = 2

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()
			var u *User

			name := u.Name
			fmt.Printf("name: %s\n", name)
		}()

		time.Sleep(time.Second)
	}()

	time.Sleep(time.Second)
}

---------



// testSearcher - функция проверяет время отклика поисковика.
// Если поисковик не ответил, или ответил с ошибкой, то возвращает ошибку
func testSearcher(ctx context.Context, url string) (time.Duration, error) {
	// просто заглушка
	return 1 * time.Second, nil
}
type Res struct {
	url string
	respTime time.Duration
}
// getFastestSearcher - возвращает самый быстрый поисковик из списка и его время ответа
func getFastestSearcher(ctx context.Context, urls []string) (url string, respTime time.Duration, err error) {

	resChan := make(chan Res, 1)
	wg := sync.WaitGroup

	// создаем канал с каким то размером
	// запускаем го рутины количество = буферу( range по каналу, делаем запрос)
	// итерируемся по urls и кладем в канал, закрываем канал

	for _, u :=range urls {
		wg.Add(1)

		go func (){
			defer wg.Done()
			d, err := testSearcher(ctx, u)
			if err != nil {
				return
			}

			<-p
			resChan <- Res{url: u, respTime: d}
		}()
	}

	go func (){
		wg.Wait()
		close(resChan)
	}()

	for _, res := range resChan {
		if respTime == time.Duration(0) {
			url = res.url
			respTime = res.respTime
		}
		if res.respTime < respTime {
			url = res.url
			respTime  = res.respTime
		}
	}
	if url == "" {
		return "", 0, errors.New("all request failed")
	}


	return url, respTime
}

На разминку - ошибся в двух местах, исправил
Строки - ОК. Не был решителен по вопросу о безопасности строк, как менять символ.
Слайсы (append,copy) - Тут понравилось что решил, мало кто решает.
Слайсы, конкурентность пакет синк - ОК, мыслит.
Мамы занет, все решил, но не знает про swiss table , это расстроило.
Панику решил, ок. Вопрос со * - что возвращает recover ответил


--------------





