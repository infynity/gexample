package main

import "ruok/jaeger/lib"

func main()  {

	lib.InitTraceConfig() //初始化配置
    lib.RunServer()

}
