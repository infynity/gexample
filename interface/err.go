package main

import "fmt"

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

func Call() *Error {
	return nil//this cause panic
	return &Error{msg: "asd"}
}

func Err() error {
	return Call()
}

func main() {
	e := Err()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("nil")
}