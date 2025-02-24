package main

import (
	"github.com/tclemos/go-expert-desafio-clean-arch/config"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/webserver"
)

func main() {
	var restConfig config.RESTConfig
	if err := config.LoadConfig(".env", &restConfig); err != nil {
		panic(err)
	}

	webserver.Start(restConfig)
}
