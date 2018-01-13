package natgeo

import (
	"log"
)

type ConfigDiscoveryServer interface {
	GetValue(name string) (string, error)
}

var ConfigServer ConfigDiscoveryServer

func init() {
	//ConfigServer = &EnvVarConfigServer{}
	configServer, err := NewConsulServer(consulAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	ConfigServer = configServer
}
