package main

import "fmt"

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}
func (e *Error) Error2() string {
	return e.msg
}

type ff interface {
	Error() string
	Error2() string
}

func Call() *Error {
	return nil//this cause panic
	return &Error{msg: "asd"}
}

func Err() error {
	return Call()
}

func main() {

	fmt.Println(32/10)
	return
	e := Err()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("nil")
}