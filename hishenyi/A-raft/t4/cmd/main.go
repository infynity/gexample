package main

import (
	"flag"
	"goraft/src/lib"
	"log"
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
	 //启动gin
	 lib.CacheServer().Run(":"+lib.SysConfig.Port)








}
