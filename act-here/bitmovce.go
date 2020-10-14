package main

import "fmt"

func main2(){
	fmt.Println(9>>1)
	// 1001     100
	fmt.Println(9<<1)
	// 1001   10010
}
//-------------------------------------------------------------------------
func main() {
	i := 10 //整形变量 i
	ip := &i //指向整型变量 i 的指针ip,包含了 i 的内存地址
	fmt.Printf("main中i的值为：%v，i 的内存地址为：%v，i的指针的内存地址为：%v\n",i,ip,&ip)
	modifyBypointer(ip)
	fmt.Printf("main中i的值为：%v，i 的内存地址为：%v，i的指针的内存地址为：%v\n",i,ip,&ip)
}

func modifyBypointer(i *int) {
	fmt.Printf("modifyBypointeri的值为：%v, i 的内存地址为：%v，i的指针的内存地址为：%v\n",*i,i,&i)
	*i = 11
}

//-------------------------------------------------------------------------
