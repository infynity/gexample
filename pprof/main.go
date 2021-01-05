package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			log.Println(Add("https://github.com/EDDYCJY"))
			time.Sleep(1e9)
		}
	}()

	err :=http.ListenAndServe("0.0.0.0:6060", nil)
	fmt.Println(err)
}
var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}