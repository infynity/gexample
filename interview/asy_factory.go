package main

import "fmt"



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
func CreateUser(t int) User{
	switch t {
	case Mem:
		return new(Member)
	case Adm:
		return new(Admin)
	default:
		return new(Member)
	}
}

//抽象工厂


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

func main()  {
	var fact AbstractFactory=new(AdminFactory)
	fmt.Println(fact.CreateUser().GetRole())
}