package natgeo

import (
	"errors"
	"log"
	"os"
	consul "github.com/hashicorp/consul/api"
)

type ConfigDiscoveryServer interface {
	GetValue(name string) (string, error)
}

var ConfigServer ConfigDiscoveryServer

type EnvVarConfigServer struct{}

func (srv EnvVarConfigServer) GetValue(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	res := os.Getenv(name)
	return res, nil

}

const consulAddress = "amidemo2:8500"

type ConsulConfigServer struct{
	client *consul.Client
}

func NewConsulServer() (*ConsulConfigServer, error) {

	cfgSrv := &ConsulConfigServer{}
	config := consul.DefaultConfig()
	config.Address = consulAddress
	client, err := consul.NewClient(config)
	if err != nil {
		log.Fatal("cannot create consul client")
		return nil, err
	}

	cfgSrv.client = client

	return cfgSrv, nil
}

func (srv ConsulConfigServer) GetValue(name string) (string, error) {

	if srv.client == nil {
		return "", errors.New("Cannot connect to consul client")
	}

	kv := srv.client.KV()
	val, _, err := kv.Get(name, nil)

	log.Println(val)

	if name == "" {
		return "", errors.New("Name cannot be empty")
	}
	if err != nil {
		return "", errors.New("Value for a given key is not found")
	}
	if val == nil {
		return "", errors.New("Value for a given key is not found")
	}

	return string(val.Value), nil

}

func init() {
	//ConfigServer = &EnvVarConfigServer{}
	configServer, err := NewConsulServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	ConfigServer = configServer
}
