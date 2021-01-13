package lib

import (
	"fmt"
	"strings"
)

func init() {
	E.AddFunction("methodMatch", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments)==2{
			 k1,k2:=arguments[0].(string),arguments[1].(string)
			return MethodMatch(k1,k2),nil
		}
		return nil,fmt.Errorf("methodMatch error")
	})
}
func MethodMatch(key1 string, key2 string) bool{
	ks:=strings.Split(key2," ")
	for _,s:=range ks{
		if s==key1{
			return true
		}
	}
	return false

}
