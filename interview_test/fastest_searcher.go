package main

import (
	"context"
	"fmt"
	"time"
)

// testSearcher - функция проверяет время отклика поисковика.
// Если поисковик не ответил, или ответил с ошибкой, то возвращает ошибку
func testSearcher(ctx context.Context, name string) (time.Duration, error) {
	// просто заглушка
	return 1 * time.Second, nil
}

// getFastestSearcher - возвращает самый быстрый поисковик из списка и его время ответа
func getFastestSearcher(ctx context.Context, searchers []string) (name string, respTime time.Duration, err error) {
	type result struct{
		name string
		duration time.Duration
		err error
	}


	ch := make(chan result,len(searchers))
	ctxT,cancel := ctx.WithTimeout(ctx,5*time.Second)
	defer cancel()

	sem := make(chan struct{},100)

	for _, s :+ range searchers{
		s := s
		go func(){
		sem <- struct{}{}
		defer func(){<- sem}()
		d,e := testSearcher(ctxT,s)
		ch <- result{s,d,e}
	}()
	}

	err = fmt.Errorf("no searchers availeble")
	for range searchers{
		r := <- ch
		if r.err == nil{
			if name == "" || r.duration < respTime{
				name = r.name
				respTime = r.duration
				err = nil
			}
		}else{
			return "",0,fmt.Errorf("err: %w from searchers",r.err)
		}

	}


	return name, respTime,err
}
