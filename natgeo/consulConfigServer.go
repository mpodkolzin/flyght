package natgeo

import (
	"errors"
	"log"

	consul "github.com/hashicorp/consul/api"
)

const consulAddress = "amidemo2:8500"

type ConsulConfigServer struct {
	client *consul.Client
}

func NewConsulServer(address string) (*ConsulConfigServer, error) {

	cfgSrv := &ConsulConfigServer{}
	config := consul.DefaultConfig()
	config.Address = address
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
