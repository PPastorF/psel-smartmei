package types

import (
)

type UniqueId int32

type Timestamp string

type Config struct {
	ServerConfig struct {
		HttpPort int32 `yaml:"httpPort"`
	} `yaml:"serverConfig"`
	Api struct {
		ApiVersion string `yaml:"apiVersion"`
	} `yaml:"api"`
}
