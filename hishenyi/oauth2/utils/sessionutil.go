package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

var sessionStore = sessions.NewCookieStore([]byte("123456"))

func init() {
	sessionStore.Options.Domain="oauth.jtthink.com"
	sessionStore.Options.Path="/"
	sessionStore.Options.MaxAge=0  //关掉浏览器就清掉session
}
func SaveUserSession(c *gin.Context,userID string ){
	s,err:=sessionStore.Get(c.Request,"LoginUser")
	if err!=nil{
		panic(err.Error())
	}
	s.Values["userID"]=userID
	err=s.Save(c.Request,c.Writer)//save 保存
	if err!=nil{
		panic(err.Error())
	}
}
func GetUserSession(r *http.Request)   string{
	if s,err:=sessionStore.Get(r,"LoginUser");err==nil{
		if s.Values["userID"]!=nil{
			return s.Values["userID"].(string)
		}
	}
	return ""
}

