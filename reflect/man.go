package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string  `json:"asd"`
}


func main(){
	user := User{Id: 1,Name: "asd"}
	v:=reflect.TypeOf(user)
	value:=reflect.ValueOf(user)
	fmt.Println(v,v.Kind(),value,v.NumField(),v.Field(0),v.Field(1))

	for i:=0;i<v.NumField();i++{
		fmt.Println(
			value.Kind(),
			value.Type().Field(i).Tag.Get("json"),
			v.Field(i).Tag,
			)
	}
}