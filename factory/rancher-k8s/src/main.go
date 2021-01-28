package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){

	engine := gin.New()

	fmt.Println("mm",engine)


	//select{}

	fmt.Println(123)
	engine.Run(":80")
}

