package lib

import "sync"

var Cache *sync.Map
var once sync.Once
func getCache() *sync.Map  {
	once.Do(func() {
		Cache=&sync.Map{}
	})
	return Cache
}
func Set(key string ,value string ){
	Cache.Store(key,value)
}
func Get(key string ) interface{}  {
	if v,ok:= Cache.Load(key);ok{
		return v
	}
	return nil
}