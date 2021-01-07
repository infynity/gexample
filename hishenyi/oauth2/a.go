package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

func main()  {
	r := gin.Default()
	 //r.LoadHTMLGlob("public/*")

	codeUrl,_:=url.ParseRequestURI("http://localhost:8080/getcode")

	loginUrl:="http://oauth.jtthink.com/auth?" +
		"response_type=code&client_id=clienta&redirect_uri="+
		codeUrl.String()

	fmt.Println(codeUrl)


	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "a-index.html", map[string]string{
			"loginUrl":loginUrl,
		})
	})
	r.GET("/getcode", func(c *gin.Context) {
		code,_:=c.GetQuery("code")
		c.JSON(200,gin.H{"code":code})
	})
	r.LoadHTMLGlob("public/*")

	r.Run(":8080")
}