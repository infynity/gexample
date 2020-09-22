package main

import (
	"fmt"
	"strings"
)

func myClosure(suffix string)  func(string ) string{

	//var suf = suf  // no need

	return func( str string) string{
		if strings.HasSuffix(str,suffix){

			return str

		}else{
			return str + suffix
		}

	}
}

func main(){


	mc := myClosure(".jpg")

	bmp := myClosure(".bmp")
	fmt.Println(mc("abc.jpg"))

	fmt.Println(bmp("abcd"))
	fmt.Println(mc("abcd"))
}
