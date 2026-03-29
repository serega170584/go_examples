ectpackage main

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

fmt.Println(a == b) // ?

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
fmt.Println(len(strChine)) //
}

-----------


package main

import "fmt"

func main() {
sl := make([]int, 1, 3)
fmt.Println(sl) // [0] l=1, c=3

appendSlice(sl, 1)
fmt.Println(sl)     // ?[0] 1
fmt.Println(sl[:2]) // ?[0, 1]

copySlice(sl, []int{2, 3})
fmt.Println(sl) // ? [2]

mutateSlice(sl, 1, 4) [2] l=1
fmt.Println(sl) // [2]
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


-------

package main

import "fmt"

// Что выведен функция main?
func main() {
var m map[int]int

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

func main() {
sl := make([]int, 0, 10)

for i := 0; i < 100_000; i++ {
go func() {
sl = append(sl, i)
}()
}

time.Sleep(time.Second)

fmt.Println(len(sl)) // 90_000 ??
}

---------



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
panic("aaaaa")
_ = a.Foo
}()

time.Sleep(time.Second)
}()

time.Sleep(time.Second)
}
---------------



// testSearcher - функция проверяет время отклика поисковика.
// Если поисковик не ответил, или ответил с ошибкой, то возвращает ошибку
func testSearcher(ctx context.Context, name string) (time.Duration, error) {
// просто заглушка
return 1 * time.Second, nil
}

type result struct {
name string
respTime time.Duration
}

// getFastestSearcher - возвращает самый быстрый поисковик из списка и его время ответа
func getFastestSearcher(ctx context.Context, searchers []string) (name string, respTime time.Duration, err error) {
if len(searchers) == 0 {
return "", 0, errors.New("no searchers")
}

var r result{}

resCh := make(chan result, len(searchers))

for _, s := range searchers {

go func() {
d, err := testSearcher(ctx, s)
if err != nil {
return
}
resCh <- result(name: s, respTime: d)
}()
}

fastTime := int.MaxInt32()
name := ""

for i := 0, i< len(searchers), i++ {
select {
case: <-ctx.Done()
return "", 0, ctx.Err()
case: r := <-resCh
if r.name := "" && r.respTime < fastTime {
fastTime = r.respTime
name = r.name
}
}


}
if  r.name == "" {
return 0, "", errors.New("errr")
}

return name, fastTime, nil
}


}

func producer() <- chan int {

}



