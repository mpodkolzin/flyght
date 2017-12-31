package main

import (
	"log"
	"os"

	"fmt"

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

	config := api.DefaultConfig()
	cc, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Cannot create consul client")
		return err
	}

	registration := new(api.AgentServiceRegistration)

	registration.ID = "test_id"
	registration.Name = "publisher service"

	address, _ := os.Hostname()
	registration.Address = address

	registration.Port = 9090
	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%d/ping", address, registration.Port)
	registration.Check.Interval = "5s"

	err = cc.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not register service, terminating")
	}

	log.Println("Consule register call")
	return nil
}

func (sr ConsulSvcRegistrator) Unregister() error {
	log.Println("Consule unregister call")
	return nil
}
