package main

import (
	"fmt"
)

type User interface {
	GetRole() string
}
type Member struct {}
func(this *Member) GetRole() string {
		return   "会员用户"
}



type Admin struct {}
func(this *Admin) GetRole() string {
	return  "后台管理用户"
}


const (
	Mem=iota
	Adm
)
func CreateUserraw(t int) User{
	switch t {
	case Mem:
		return new(Member)
	case Adm:
		return new(Admin)
	default:
		return new(Member)
	}
}





type AbstractFactory interface {
	CreateUser() User
}


type MemberFactory struct {}
func(this *MemberFactory) CreateUser() User{
	return &Member{}
}


type AdminFactory struct {}
func(this *AdminFactory) CreateUser() User{
	return &Admin{}
}


func testptr(p  *string){

	fmt.Println(&p)
	//都是值传递
	p=nil
	fmt.Println(p)

}


func testslice(sb []int){
	fmt.Printf("%p\n",sb)
	fmt.Printf("%p\n",&sb)
	sb=nil
	//sb[1]=123
fmt.Println(sb)
	fmt.Printf("%p\n",sb)
	fmt.Printf("%p\n",&sb)

}
func main()  {

	sb := new(string)
	fmt.Println(&sb,sb)
	testptr(sb)




	fmt.Println(&sb,sb)

	slc:=[]int{1,2,3}
	fmt.Println(slc)
	fmt.Printf("%p\n",slc)
	fmt.Printf("%p\n",&slc)

	fmt.Println()
	testslice(slc)
	fmt.Println(slc)

	//sb=nil
	//
	//
	//fmt.Println(sb,666)
	 var fact AbstractFactory=new(AdminFactory)
	 fmt.Println(fact.CreateUser().GetRole())
}