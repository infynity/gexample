package main

import (
	"flag"
	"os"
)

func main(){

	cch := make(chan int)
	select {
	case  cch<-1:
	default:
		//fmt.Println(123)
	}

	return


	flagset:=flag.FlagSet{}
	if flagset.Lookup("version").Value.(flag.Getter).Get().(bool) {
		//fmt.Println(version.String("nsqd"))
		os.Exit(0)
	}
}
