package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"goauth/utils"
	"log"
	"net/http"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	err:=clientStore.Set("clienta", &models.Client{
		ID:     "clienta",
		Secret: "123",
		Domain: "http://localhost:8080",
	})
	if err!=nil{
		log.Fatal(err)
	}
	manager.MapClientStorage(clientStore)


	srv := server.NewDefaultServer(manager)
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)
	r:=gin.New()
	 r.Use(utils.ErrorHandler())
	//响应授权码
	r.GET("/auth", func(context *gin.Context) {
		err:=srv.HandleAuthorizeRequest(context.Writer,context.Request)
		if err!=nil{
			log.Println(err)
		}
	})
	r.Any("/login", func(c *gin.Context) {
		data:=map[string]string{
			"error":"",
		}
		if c.Request.Method=="POST"{
			uname,upass:=c.PostForm("userName"),c.PostForm("userPass")
			if uname+upass=="shenyi123"{
				utils.SaveUserSession(c,uname)
				c.Redirect(302,"/auth?"+c.Request.URL.RawQuery)
				return
			}else{
				data["error"]="用户名密码错误"
			}
		}
		c.HTML(200, "login.html", data)
	})
	r.LoadHTMLGlob("public/*.html")
	r.Run(":80")

}
func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		  if userID=utils.GetUserSession(r);userID==""{
			  w.Header().Set("Location", "/login?"+r.URL.RawQuery)
			  w.WriteHeader(302)
		  }
		return
}
