package main

import (
	"fmt"
	"unsafe"
)

type Integer int

type Operation interface {
	Less(b Integer) bool
	Add(b Integer)
}


func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

//
//func main(){
//	s := make([]byte, 200)
//	ptr := unsafe.Pointer(&s[0])
//	fmt.Println(ptr)
//
//	fmt.Printf("%p",s)
//}

type N struct {
	 Name string `json:"name"`
	 Data int `json:"data"`}

type Error struct {
	errCode uint8
}
func (e *Error) Error() string {
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}
 func main() {
	 var e *Error
	 fmt.Println(e)

	 fmt.Printf("%p\n",e)

	 	ptr := unsafe.Pointer(e)
	fmt.Println(ptr,123)

	 //return
	 checkError(e)
}

func checkError(err error) {
	fmt.Println(err)
	if err != nil {
	//if err != (*Error)(nil) {
		fmt.Println(err)
		panic(err)
	}
}
