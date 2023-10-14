package main

import (
	"errors"
	"fmt"
)

type Error struct {
	err error
}

func (e Error) Error() string {
	return fmt.Sprintf("%v", e.err)
}

func New(err error) Error {
	return Error{err: fmt.Errorf("Test 1 %w", err)}
}

func (e Error) Unwrap() error {
	return e.err
}

type Error1 struct {
	err error
}

func (e Error1) Error() string {
	return fmt.Sprintf("%v", e.err)
}

func (e Error1) Unwrap() error {
	return e.err
}

func NewNew(err error) Error1 {
	return Error1{err: fmt.Errorf("Test 2 %w", err)}
}

type Error2 struct {
	err error
}

func (e Error2) Error() string {
	return fmt.Sprintf("%v", e.err)
}

func (e Error2) Unwrap() error {
	return e.err
}

func NewNewNew(err error) Error2 {
	return Error2{err: fmt.Errorf("Test 3 %w", err)}
}

func main() {
	fmt.Println(returnError() == nil)
	fmt.Println(returnErrorPtr() == nil)
	fmt.Println(returnCustomError() == nil)
	fmt.Println(returnCustomErrorPtr() == nil)
	fmt.Println(returnMyErr() == nil)

	err := errors.New("Test")
	err1 := New(err)
	err2 := NewNew(err1)
	err3 := NewNew(err2)
	//err = errors.New("Test")
	//err3 := New(err)

	fmt.Println(errors.Is(err, err1))
	fmt.Println(errors.Is(err2, err3))
	fmt.Println(errors.As(err2, &err1))
	fmt.Println(errors.As(err2, &err3))
	//err := Error2{test: "123"}
	//err1 := Error2{test: "1233344"}
	//fmt.Println(errors.Is(err, err1))
	//fmt.Println(errors.As(err, &err1))
	//err := errors.New("sdasdadasd")
	//err1 := errors.New("sdasdadasdaaaaaaaa")
	//fmt.Println(test().(type))
	//err := Error{}
	//var testErr Error
	//err := Error{}
	//err2 := fmt.Errorf("error2: [%w]", err)
	//fmt.Printf("%v", err2)
	//fmt.Println(errors.Unwrap(err1))
	//err2 := fmt.Errorf("error2: [%w]", err1)
	//fmt.Println(err2)
	//fmt.Println(errors.Unwrap(err2))
	//fmt.Println(errors.Is(err, testErr))
	//fmt.Println(errors.As(err, &testErr))
	//var e NegativeError
	//e = 123455
	//fmt.Printf("%v\n", e)
	//testErr := errors.New("Test error")
	//err4 := errors.New("asdasdasdasdasd")
	//err1 := fmt.Errorf("Line: %s\n", "123")
	//err2 := fmt.Errorf("Line: %s %v\n", "456", err1)
	//err3 := fmt.Errorf("Line: %s %v\n", "789", err2)
	//fmt.Printf("%v\n", err3)

	//fmt.Printf("%b\n", err() == nil)
	//fmt.Printf("%b\n", err1() == nil)
	//fmt.Printf("%b\n", err2() == nil)
	//fmt.Printf("%s\n", err2().Error())
	//fmt.Printf("%b\n", err3() == nil)
	//fmt.Printf("%b\n", err4() == nil)
}

func err() error {
	var err error
	return err
}

func err1() *error {
	var err *error
	return err
}

func err2() error {
	var err Error
	return err
}

func err3() error {
	var err *Error
	return err
}

func err4() error {
	return nil
}

//	return "Error11111"
//}
//
//func (e Error) setErr(err error) {
//	e.test = err
//}
//
//func (e Error) Unwrap() error {
//	return fmt.Errorf("dadadasdadasdasdadadadadadsda")
//}
//
//type Error1 struct {
//	err error
//}
//
//func (e Error1) Error() string {
//	return "Error2222222"
//}
//
//func test() error {
//	return Error{}
//}
//
//type Error2 struct {
//	test string
//}
//

//type NegativeError int
//
//func (e NegativeError) Error() string {
//	return fmt.Sprintf("%d", e)
//}

func returnError() error {
	var err error
	return err
}

func returnErrorPtr() *error {
	var err *error
	return err
}

func returnCustomError() error {
	var customErr MyErr
	return customErr
}

func returnCustomErrorPtr() error {
	var customErr *MyErr
	return customErr
}

func returnMyErr() *MyErr {
	return nil
}

type MyErr struct{}

func (me MyErr) Error() string {
	return "my err string"
}
