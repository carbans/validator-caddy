package main

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/carbans/validator-caddy/validator"
)

func main() {
	caddy.RegisterModule(validator.ValidatorMiddleware{})
}
