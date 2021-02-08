package misc

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		HttpPort string `yaml:"httpPort"`
	} `yaml:"server"`
	
	Api struct {
		ApiVersion string `yaml:"apiVersion"`
	} `yaml:"api"`
	
	Db struct {
		DbName string `yaml:"db_name"`
		Endpoint string `yaml:"endpoint"`
		Port string `yaml:"port"`
		Auth struct {
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
		} `yaml:"auth"`
	} `yaml:"db"`
}

func (c *Config) ReadFromFile(filename string) error {
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