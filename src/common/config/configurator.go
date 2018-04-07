package config

type Configurator interface {
	Load(c interface{})
}
