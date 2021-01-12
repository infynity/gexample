package lib

import (
	//"github.com/casbin/casbin"
	"github.com/casbin/casbin/v2"
    "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"

)

func CheckLogin() gin.HandlerFunc{
	return func(context *gin.Context) {
	 	  if context.Request.Header.Get("token")==""{
				context.AbortWithStatusJSON(400,gin.H{"message":"token required"})
		  }else{
		  		 context.Set("user_name",context.Request.Header.Get("token"))
		  		 context.Next()
		  }
	}
}


func RBAC() gin.HandlerFunc  {
	// e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")
	adapter,err:=gormadapter.NewAdapterByDB( Gorm)

	if err!=nil{
		log.Fatal()
	}
	e,err:= casbin.NewEnforcer("resources/model.conf",adapter)
	if err!=nil{
		log.Fatal()
	}
	err=e.LoadPolicy()
	if err!=nil{
		log.Fatal()
	}

	return func(context *gin.Context) {
		user,_:=context.Get("user_name")
		access,err:=e.Enforce(user,context.Request.RequestURI,context.Request.Method)
		if err!=nil || !access{
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