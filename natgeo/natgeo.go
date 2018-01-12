package natgeo

import (
	"errors"
	"os"
)

type ConfigDiscoveryServer interface {
	GetValue(name string) (string, error)
}

var ConfigServer ConfigDiscoveryServer

type EnvVarConfigServer struct{}

func (srv *EnvVarConfigServer) GetValue(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	res := os.Getenv(name)
	return res, nil

}

func init() {
	ConfigServer = &EnvVarConfigServer{}
}
