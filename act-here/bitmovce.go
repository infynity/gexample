package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func main2(){
	fmt.Println(9>>1)
	// 1001     100
	fmt.Println(9<<1)
	// 1001   10010
}
//-------------------------------------------------------------------------
func main3() {
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

func maina(){
	arr := [5]int{1,2,3,4,5}


	fmt.Printf("%p\n",&arr)
	slc := arr[:]
	fmt.Println(&slc[0])
	fmt.Printf("%p\n",&slc)

	slc[2]=100

	fmt.Println(slc,arr)
}


//-------------------------------------------------------------------------

func mainqw() {

	var s1 []int
	s2 := make([]int,0)
	s4 := make([]int,0)

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)),*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}


//-------------------------------------------------------------------------

func mainls() {
	s := []int{1,2,3,4,5}
	for _, v:=range s {
		s =append(s, v)
		fmt.Printf("len(s)=%v\n",len(s))
	}
}



//-------------------------------------------------------------------------

func mainss() {
	fmt.Println(1 ^ 2 ^ 1) //2
	fmt.Println(1 ^ 1)     //0
	fmt.Println(0 ^ 3)     //3

	tmpHash := make(map[string]int, 0)
	tmpHash["asd"] = 1
	tmpHash["dsa"] = 1

	for index,_ := range tmpHash {

		fmt.Println(index)
	}

	 ok := tmpHash["dsaa"]
	fmt.Println(ok)
}


//-------------------------------------------------------------------------


func Bar(vl int, width int) string {
	return fmt.Sprintf("%s%*c", strings.Repeat("█", vl/10), vl/10-width+1,
		([]rune(" ▏▎▍▌▋▋▊▉█"))[vl%10])
}

func mainz() {

	fmt.Println("2020-10-22">="2020-10-21")
	return
	fmt.Println(tnslc())

	return
	for i := 0; i <= 100; i++ {
		fmt.Printf("\f%s%d%%", Bar(i, 20), i)
		time.Sleep(200 * time.Millisecond)
	}
}


func tnslc()(res []int){
	//res = make([]int,0)
	return res
}


//-------------------------------------------------------------------------
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func main是(){
	dates := GetBetweenDates("2020-09-02", "2020-10-31")
	fmt.Println(dates)

}


//-------------------------------------------------------------------------


func maindd(){
	p:=5
	ptr := &p
	change(&p)
	fmt.Println(&p,ptr,&ptr)
	var ps *int
	//ps =new(int)
	*ps = 1 //Potential nil pointer dereference   todo panic



}

func change(a *int){
	fmt.Println(a,&a)
	*a = 0
}


//-------------------------------------------------------------------------


func mainaa(){
	var a = 7.98
	var p = &a
	var pp = &p

	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)

	fmt.Println("p = ", p)
	fmt.Println("address of p = ", &p)

	fmt.Println("pp = ", pp)

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
}

//-------------------------------------------------------------------------
