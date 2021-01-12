package lib

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckLogin() gin.HandlerFunc{
	return func(context *gin.Context) {
		fmt.Println(context.Query("token"),"ooooooooo")

	 	  if context.Request.Header.Get("token")==""{
				context.AbortWithStatusJSON(400,gin.H{"message":"token required"})
		  }else{
		  		 context.Set("user_name",context.Request.Header.Get("token"))
		  		 context.Next()
		  }
	}
}

func RBAC() gin.HandlerFunc  {
	e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")

	return func(context *gin.Context) {
		user,_:=context.Get("user_name")


		log.Println(user,context.Request.RequestURI,context.Request.Method,"qqqqqqqqq")
		if !e.Enforce(user,context.Request.RequestURI,context.Request.Method){
			context.AbortWithStatusJSON(403,gin.H{"message":"forbidden"})
		}else{
			context.Next()
		}
	}
}

func Middlewares() (fs []gin.HandlerFunc)  {
	 fs=append(fs,CheckLogin(),RBAC())
	 return
}