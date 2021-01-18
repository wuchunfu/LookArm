package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type SiteConfig struct{
	Server struct{
		ServerPort string `yaml:"server_port"`
		JwtKey string  `yaml:"jwtkey"`
	}
	Database struct {
		DBhost     string `yaml:"db_host"`
		DBport     string `yaml:"db_port"`
		DBuser     string `yaml:"db_user"`
		DBpassowrd string `yaml:"db_password"`
		DBname     string `yaml:"db_name"`
	}
}

func(c *SiteConfig) LoadCinfig() *SiteConfig {
	yamlFile,err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("err:",err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	
	if err != nil {
		fmt.Println("err:",err)
	}
	return c
}
