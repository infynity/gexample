package lib

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"sync"
)

var Cache *sync.Map
var once sync.Once
func getCache() *sync.Map  {
	once.Do(func() {
		Cache=&sync.Map{}
	})
	return Cache
}
func Set(key string ,value string ){
	getCache().Store(key,value)
}
func Get(key string ) interface{}  {
	if v,ok:= getCache().Load(key);ok{
		return v
	}
	return nil
}
type CacheRequest struct {
	Key string `json:"key" binding:"required,min=1"`
	Value string `json:"value" binding:"omitempty,min=1"`
}

func NewCacheRequest() *CacheRequest {
	return &CacheRequest{}
}
func CacheMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		 defer func() {
		 	if e:=recover();e!=nil{
		 		context.JSON(400,gin.H{"message":e})
			}
		 }()
		 context.Next()
	}
}
func Error(err error)  {
	 if err!=nil{
	 	panic(err)
	 }

}
func CacheServer() *gin.Engine{
	r:=gin.New()
	r.Use(CacheMiddleware())


	r.Handle("POST","/get", func(context *gin.Context) {
		 req:=NewCacheRequest()
		 Error(context.BindJSON(req))
		 if v:=Get(req.Key);v!=nil{
		 	req.Value=v.(string)
		 	context.JSON(200,req)
		 }else{
		 	Error(fmt.Errorf("find no cache"))
		 }
	})


	r.Handle("POST","/set", func(context *gin.Context) {
		req:=NewCacheRequest()
		Error(context.BindJSON(req))
		// Set(req.Key,req.Value) //往我们的sync.Map里插值


		req_bytes,_:=json.Marshal(req)
		future:=RaftNode.Apply(req_bytes,time.Second)


		if e:=future.Error();e!=nil{
			Error(e)
		}else{
			context.JSON(200,gin.H{"message":"OK"})
		}

	})
	return r
}