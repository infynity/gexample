package main

import (
	"flag"
	"fmt"
	"goraft/src/lib"
	"log"
	"time"
)
func main()  {
	cfile:=""
	flag.StringVar(&cfile, "c", "", "your config file ")
	flag.Parse()
	if cfile==""{
		log.Fatal("config file error")
	}
	 err:=lib.BootStrap(cfile)
	 if err!=nil{
	 	log.Fatal(err)
	 }
	 for {
		fmt.Println("当前主节点是:",lib.RaftNode.Leader())
	 	time.Sleep(time.Second*1)
	 }



}
