package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"time"
)
// api
func RunServer() error {
	r:=gin.New()
	r.Handle("GET","/new", func(ctx *gin.Context) {
		orderSpan:=opentracing.StartSpan("创建订单API")
		defer orderSpan.Finish()

		spanCtx:=opentracing.ContextWithSpan(context.Background(),orderSpan)
		time.Sleep(time.Millisecond*500)
		ret:=GenOrder(spanCtx,"生成订单")
		orderNo:="order"+time.Now().Format("20060102150405")
		orderSpan.SetTag("orderno",orderNo)
		if ret{
			ctx.JSON(200,gin.H{"message":"下单成功"})
		}else{
			orderSpan.SetTag("status","error")
			ctx.JSON(400,gin.H{"message":"下单失败"})
		}
	})
	return r.Run(":8081")
}
// A 生成订单
func GenOrder(spanCtx context.Context,name string) bool  {
	span,ctx:=opentracing.StartSpanFromContext(spanCtx,name)
	defer span.Finish()
	time.Sleep(time.Second*1)
	if GetStock(ctx,"检查库存")>0{
		return true
	}else{
		return false
	}
}
// 获取库存
func GetStock(spanCtx context.Context,name string) int   {
	span,_:=opentracing.StartSpanFromContext(spanCtx,name)
	defer span.Finish()

	time.Sleep(time.Millisecond*300)
	if time.Now().Unix()%2==0{
		return 0
	}
	return 10
}

