package main

import (
	"fmt"
	"sync"
	"time"
)


var kv sync.Map

func set(key string,value interface{},expire time.Duration){
	kv.Store(key,value)
	time.AfterFunc(expire, func() {
		kv.Delete(key)
	})
}
func main(){

	set("id",123,time.Second*5)
	set("name",2333,time.Second*8)

	for  {
		fmt.Println(kv.Load("id"))
		fmt.Println(kv.Load("name"))
		//time.Sleep(1e9)
	}

}