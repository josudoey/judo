package cmd

import (
	"github.com/caarlos0/env/v11"
)

func MustEnvParseAs[T any]() T {
	cfg, err := env.ParseAs[T]()
	if err != nil {
		panic(err)
	}

	return cfg
}
