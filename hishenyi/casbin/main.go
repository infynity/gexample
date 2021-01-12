package main

import (
	"cb/lib"
	"github.com/gin-gonic/gin"
)

func main() {
	// /depts
	//sub:= "shenyi" // 想要访问资源的用户。
	//obj:= "/depts" // 将被访问的资源。
	//act:= "POST" // 用户对资源执行的操作。
	//e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")

	//ok:= e.Enforce(sub, obj, act)
	//if ok {
	//	log.Println("运行通过")
	//}

	r:=gin.Default()
	r.Use(lib.Middlewares()...)

	r.GET("/depts", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":"部门列表"})
	})


	r.POST("/depts", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":"批量修改部门列表"})
	})

	r.Run(":8080")


}
