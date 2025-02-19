package validator

import (
	"fmt"

	"github.com/caddyserver/caddy/v2"
)

// Config representa la configuración del middleware
type Config struct {
	ValidatorURL string `json:"validator_url,omitempty"`
}

// DefaultConfig define los valores predeterminados
var DefaultConfig = Config{
	ValidatorURL: "http://default-validator.local",
}

// Provision inicializa la configuración del middleware
func (m *ValidatorMiddleware) Provision(ctx caddy.Context) error {
	if m.Config == nil {
		m.Config = &DefaultConfig
	}

	// Verifica si el ValidatorURL está vacío y usa el valor por defecto
	if m.Config.ValidatorURL == "" {
		m.Config.ValidatorURL = DefaultConfig.ValidatorURL
	}

	fmt.Printf("ValidatorMiddleware configurado con URL: %s\n", m.Config.ValidatorURL)
	return nil
}
