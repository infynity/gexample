package lib

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/hashicorp/raft"
	"io/ioutil"
	"log"
)

type Config struct {
	ServerName string  `yaml:"server-name"`
	ServerID string `yaml:"server-id"`
	LogStore string
	StableStore string
	Transport string
	Servers []raft.Server
	Port string
}

func NewConfig() *Config {
	return &Config{}
}
func loadConfigFile(path string )  []byte  {
	b,err:=ioutil.ReadFile(path)
	if err!=nil{
		log.Println(err)
		return nil
	}
	return b
}
func LoadConfig(path string) (*Config,error){
	config:=NewConfig()
	if b:=loadConfigFile(path);b!=nil{
		err:=yaml.Unmarshal(b,config)
		if err!=nil{
			return nil,err
		}
		return config,err
	}else {
		return nil,fmt.Errorf("加载配置失败")
	}

}