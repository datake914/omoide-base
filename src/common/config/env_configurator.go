package config

import "github.com/kelseyhightower/envconfig"

type EnvConfigurator struct {
}

func NewEnvConfigurator() *EnvConfigurator {
	return new(EnvConfigurator)
}

func (*EnvConfigurator) Load(c interface{}) {
	envconfig.MustProcess("OMID", c)
}
