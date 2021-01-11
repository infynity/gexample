package main

import (
	"ff/utils"
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
	//fmt.Println(231348 + 41559 + 90641 + 10485 + 14681)
	//return
	r := gin.Default()
	 //r.LoadHTMLGlob("public/*")

	codeUrl,_:=url.ParseRequestURI("http://localhost:8080/getcode")

	loginUrl:="http://oauth.jtthink.com/auth?" +
		"response_type=code&client_id=clienta&redirect_uri="+
		codeUrl.String()
	loginUrl2:=oauth2Config.AuthCodeURL("myclient")

	fmt.Println(loginUrl2,123123)
//http://oauth.jtthink.com/auth?client_id=clienta&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fgetcode&response_type=code&scope=all&state=myclient 123123

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


	r.GET("/info", func(context *gin.Context) {
		token:=context.Query("token")
		ret:=utils.GetUserInfo(authServerURL+"/info",token,true)
		context.Writer.Header().Add("Content-type","application/json")
		context.String(200,ret)
	})


	r.Run(":8080")
}