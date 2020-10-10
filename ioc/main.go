package main

import (
	"fmt"
	"ruok/ioc/Config"
	. "ruok/ioc/Injector"
	"ruok/ioc/services"
)

func main()  {
	serviceConfig:=Config.NewServiceConfig()

	BeanFactory.ExprMap= map[string]interface{}{
		"ServiceConfig":serviceConfig,
	}


   	BeanFactory.Set(serviceConfig)


	{
		//这里 测试 userService
		userService:=services.NewUserService()
		BeanFactory.Apply(userService)
		fmt.Println(userService.Order)
	}


	{
		//这里 测试 adminService
		adminService:=services.NewAdminService()
		BeanFactory.Apply(adminService)
		fmt.Println(adminService.Order)
	}





}
