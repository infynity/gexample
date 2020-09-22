package main

import (
	"fmt"
)


type Ack struct {

	da int
	sd string
	faker map[string]int
}

func main(){




	//sii :=int8(2)
	//fmt.Println(&sii)
	//sii2 :=int64(12)
	//fmt.Println(&sii2)
	//sii32 :=int8(12)
	//fmt.Println(&sii32)
	//return

	//slc:=[]int16{1,2,3}
	//slc:=[]int32{1,2,3}   //4个字节
	slc:=[]int64{1,2,3}

	//slc := []byte{12,1,222}
	//slc := []string{"abc","abcde","abcdefghjikaaaaaaaaaaaaaaaaaaaaaaaa"}


	//slc := []Ack{{sd: "as"},{da: 123},{da: 33333333}}
	fmt.Printf("%p\n",&slc)
	fmt.Printf("%p,%v\n",&slc[0],slc[0])
	fmt.Printf("%p\n",&slc[1])
	fmt.Printf("%p\n",&slc[2])

	ptr := &slc[0]
	fmt.Println(*ptr)

	return

	i:=12
	fmt.Println(&i)

	si :=[]byte{1}
	fmt.Println(&si)


}
