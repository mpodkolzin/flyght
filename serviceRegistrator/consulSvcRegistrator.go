package main

import (
	"log"

	"github.com/hashicorp/consul/api"
)

type ConsulSvcRegistrator struct {
	client *api.Client
}

func NewRegistrator() (*ConsulSvcRegistrator, error) {
	reg := &ConsulSvcRegistrator{}
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal("Could not initialize service registrator")
		return nil, err
	}

	reg.client = client
	return reg, nil
}

func (sr ConsulSvcRegistrator) Register() error {

	log.Println("Consule register call")
	return nil
}

func (sr ConsulSvcRegistrator) Unregister() error {
	log.Println("Consule unregister call")
	return nil
}
