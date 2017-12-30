package main

//Registrator : registers services with one of the service directories
type Registrator interface {
	Register() error
	Unregister() error
}
