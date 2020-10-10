package Injector

import (
	"github.com/shenyisyn/goft-expr/src/expr"
	"reflect"
)

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory=NewBeanFactory()
}
type BeanFactoryImpl struct {
	beanMapper BeanMapper
	ExprMap map[string]interface{}

}
func(this *BeanFactoryImpl) Set(vlist ...interface{}){
	if vlist==nil || len(vlist)==0{
		return
	}
	for _,v:=range vlist{
		this.beanMapper.add(v)
	}
}
func(this *BeanFactoryImpl) Get(v interface{}) interface{}{
	if v==nil {
		return nil
	}
	get_v:=this.beanMapper.get(v)
	if get_v.IsValid(){
		return get_v.Interface()
	}
	return nil
}
//处理依赖注入
func(this *BeanFactoryImpl) Apply(bean interface{}) {
	if bean==nil{
		return
	}
	v:=reflect.ValueOf(bean)

	if v.Kind()==reflect.Ptr{
		v=v.Elem()
	}

	if v.Kind()!=reflect.Struct {
		return
	}

	for i:=0;i<v.NumField();i++{
		field:=v.Type().Field(i)
		if v.Field(i).CanSet() &&  field.Tag.Get("inject")!=""{
			if get_v:=this.Get(field.Type);get_v!=nil{
				v.Field(i).Set(reflect.ValueOf(get_v))
				continue
			}
			if field.Tag.Get("inject")!="-"{
				//表达式方式支持
				ret:=expr.BeanExpr(field.Tag.Get("inject"),this.ExprMap)
				if ret!=nil && !ret.IsEmpty(){
					retValue:=ret[0]
					if retValue!=nil{
						this.Set(retValue)
						v.Field(i).Set(reflect.ValueOf(retValue))
					}

				}
			}


		}
	}
}
func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper:make(BeanMapper),ExprMap:make(map[string]interface{})}
}