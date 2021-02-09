package conf

import (
	"os"
	"fmt"
	"strings"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

func InProductionEnv() bool {
    if strings.ToUpper(os.Getenv("DEPLOY_ENV")) == "PROD" {
        return true
    } else {
        return false
    }
}

type DBConfig struct {
	Name string `yaml:"db_name"`
	Endpoint string `yaml:"endpoint"`
	Port string `yaml:"port"`
	Auth struct {
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"auth"`
}

func (c DBConfig) DBConnectionURL() string {
	return fmt.Sprintf("%s:%s", c.Endpoint, c.Port)
}

type AppConfig struct {
	Server struct {
		Address string `yaml:"address"`
		HttpPort string `yaml:"httpPort"`
	} `yaml:"server"`
	
	Api struct {
		Version string `yaml:"version"`
		RoutePrefix string `yaml:"routePrefix"`	
	} `yaml:"api"`
	
	DB DBConfig `yaml:"db"`
}

func (c *AppConfig) ReadFromFile(filename string) error {
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    err = yaml.Unmarshal(buf, c)
    if err != nil {
        return fmt.Errorf("Erro ao ler o arquivo %q: %v", filename, err)
    }

    return nil
}

func (c *AppConfig) ApiRoutePrefix() string {
	return fmt.Sprintf("%s/%s", c.Api.RoutePrefix, c.Api.Version)
}

func (c *AppConfig) HttpConnectionURL() string {
	return fmt.Sprintf("%s:%s", c.Server.Address, c.Server.HttpPort)
}