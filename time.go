package main

import (
	"fmt"
	"sort"
	"time"
)
const (
	Sunday  = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type aji int
func main(){
	fmt.Println(123)

	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Saturday,time.Sunday,Sunday,Saturday)
	fmt.Println(new(aji))


	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}


	whatAmI(new(chan int))
	whatAmI("asd")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var threeD [2][2][3]int
	fmt.Println(threeD)

	s := make([]string, 3)
	fmt.Println("emp:", s)

		s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

		fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e","f")
	fmt.Println("apd:", s,len(s),cap(s))



	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)


	l := s[1:2]

	fmt.Println("sl1:", l)
	fmt.Println(s[5])

	fmt.Println(s)
	fmt.Println([]int{0,1,2,3,4,5})
	fmt.Println(s[1:3])

	for k,v:=range s{
		fmt.Println(k,v)
	}


	sum(1,2,34,44)
	nums := []int{1, 2, 3, 4}
sum(nums ...)



	fn:=intSeq()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	fn2:=intSeq()
	fmt.Println(fn2())
	fmt.Println(fn())

	fmt.Println(fact(5))

	//ci:=ai{id: []int{1,2}}
	//bi:=di{id: []int{1,2}}
	//fmt.Println(ci==ai(bi))  struct containing []int cannot be compared)

	ints :=[]int{3,4,5,1}
	sort.Ints(ints)
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)
	//sorted := sort.IntsAreSorted(ints)

}


type ai struct {
	id []int
}

type di struct {
	id []int


}


func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)

	// 2*fact(1) 1*fact(0) 1
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}