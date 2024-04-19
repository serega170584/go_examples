package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	file *os.File
}

func NewLogger() *Logger {
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	return &Logger{file: file}
}

func (l *Logger) Log(s string) error {
	_, err := l.file.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}

type CustomLogger struct {
	logger *Logger
	ch     chan string
	wg     *sync.WaitGroup
}

func NewCustomLogger(l *Logger) *CustomLogger {
	ch := make(chan string, 100)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			_ = l.Log(val)
		}
	}()
	return &CustomLogger{logger: l, ch: ch, wg: wg}
}

func (cl *CustomLogger) Log(s string) error {
	defer handlaPanic()
	cl.ch <- s
	return nil
}

func handlaPanic() {
	e := recover()
	err := e.(error)
	fmt.Println(err.Error())
}

func (sl *CustomLogger) Close() error {
	close(sl.ch)
	sl.wg.Wait()
	return sl.logger.Close()
}

func main() {
	l := NewLogger()
	cl := NewCustomLogger(l)
	for i := 0; i < 200; i++ {
		go cl.Log("asdasdasdasd")
	}
	_ = cl.Close()
}
