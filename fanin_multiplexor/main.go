package main

import (
	"fmt"
	"sync"
	"time"
)

type payload struct {
	name  string
	value int
}

// комбинирует несколько входов в 1 выходной канал
// иными словами мультиплексируем каналы
// порядок вывода не гарантируется
// реальные приложения - хотим лимитировать параллельно выполянемые запросы к внешнему сервису. своего рода rate limiter где в каждый момент времени выполянется 1 запрос
// простой пример - отправка почты несколько горутин разбирают разные типы очередей формируют тело письма тело сообщения получателей и прочие параметры
// и 1 горутина получает на обработку все эти сообщения для отправки через внешний сервис в 1 поток

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(2)
	producers := make([]<-chan payload, 0, 3)
	producers = append(producers, producer("Alice", done, &wg))
	producers = append(producers, producer("Jack", done, &wg))
	producers = append(producers, producer("Max", done, &wg))

	fanIn := make(chan payload, 0)

	wg.Add(2)
	consumer(producers, done, &wg, fanIn)

	go func() {
		for {
			select {
			case <-done:
				return
			// читаем чтобы показать результаты
			case v := <-fanIn:
				fmt.Printf("fanin got %v\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
	close(fanIn)
}

// возвращает канал, в который мы будем писать полезную нагрузку
// сигнальный канал done
// sleep для имитации задержек
// waitgroup для выполнения graceful shutdown
func producer(name string, done <-chan struct{}, wg *sync.WaitGroup) <-chan payload {
	ch := make(chan payload)
	var i = 1
	go func() {
		defer wg.Done()
		for {
			select {
			// оба канала небуферизированные:
			// выполнение блокируется пока done не будет снаружи закрыт
			// пока ch не будет снаружи прочитан
			// при этом в большинстве случаев будет происходить блокировка на запись в канал ch так как горутина consumer читающая канал producer
			// еще могла не запуститься к этому моменту

			case <-done:
				close(ch)
				fmt.Println(name, "completed")
				return
			case ch <- payload{
				name:  name,
				value: i,
			}:
				fmt.Println(name, "produced", i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	return ch
}

func consumer(channels []<-chan payload, done <-chan struct{}, wg *sync.WaitGroup, fanIn chan<- payload) {
	// для каждого канала создаем горутину для независимого чтения каждого из каналов
	for i, ch := range channels {
		i := i + 1
		ch := ch
		go func() {
			defer wg.Done()
			fmt.Println("started consume of producer", i)
			for {
				select {
				case <-done:
					fmt.Println("consume of producer", i, "completed")
					return
				case v := <-ch:
					fmt.Println("consumer of producer", i, "got value", v.value, "from", v.name)
					// все результаты пишем в 1 канал
					fanIn <- v
				}
			}
		}()
	}
}
