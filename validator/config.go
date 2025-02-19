package validator

import "github.com/caddyserver/caddy/caddyfile"

// Config estructura de configuración del middleware
type Config struct {
	ValidatorURL string `json:"validator_url,omitempty"`
}

// UnmarshalCaddyfile permite leer configuración desde el Caddyfile
func (m *ValidatorMiddleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.Val() == "validator_url" && d.NextArg() {
			m.ValidatorURL = d.Val()
		}
	}
	return nil
}
