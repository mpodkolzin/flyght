package registrator

type Registrator interface {
	Register() error
	Unregister() error
}
