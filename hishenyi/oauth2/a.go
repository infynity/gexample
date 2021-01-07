package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/url"
)
const (
	authServerURL = "http://oauth.jtthink.com"
)
var (
	oauth2Config = oauth2.Config{
		ClientID:     "clienta",
		ClientSecret: "123",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:8080/getcode",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/auth",//获取授权码 地址
			TokenURL: authServerURL + "/token", //获取token地址
		},
	}

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
		//	 c.JSON(200,gin.H{"code":code})
		token,err:=oauth2Config.Exchange(c,code)
		if err != nil {
			c.JSON(400,gin.H{"message":err.Error()})
		}else{
			c.JSON(200,token)
		}
	})
	r.LoadHTMLGlob("public/*")

	r.Run(":8080")
}