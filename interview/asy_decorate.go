package main

import (
	"net/http"
)


func gethandler()http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func decorate(fn http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("this is decorate"))
		fn(writer,request)
	}
}

func index(writer http.ResponseWriter, request *http.Request){

	writer.Write([]byte("aaaaa"))
}
var sj int
func main(){

	//fmt.Println(&sj==nil,sj)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})

	http.HandleFunc("/asd",gethandler())
	http.HandleFunc("/asdee",decorate(index))

	http.ListenAndServe(":8080",nil)





}
