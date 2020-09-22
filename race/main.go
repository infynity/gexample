package main

import (
	"fmt"
	"sync"
	"time"
	//"time"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里我们将 res 放入到 myMap
	myMap[n] = res //concurrent map writes?
}

func main() {

	for i := 0; i < 100; i++ {
		go test(i)
	}

	time.Sleep(2e9)
	for i, v := range myMap {
		fmt.Println(i, v)
	}

	return

	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	//var ch = make(chan bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			fmt.Println(1)
			//time.Sleep(time.Millisecond)
			//count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
