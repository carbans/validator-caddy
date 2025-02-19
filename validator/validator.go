package validator

import (
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// ValidatorMiddleware representa el middleware de validación
type ValidatorMiddleware struct {
	Config *Config `json:"config,omitempty"`
}

// CaddyModule registra el módulo en Caddy
func (ValidatorMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.validator",
		New: func() caddy.Module { return new(ValidatorMiddleware) },
	}
}

// ServeHTTP procesa las solicitudes HTTP
func (m *ValidatorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Verifica que la URL del validador esté configurada
	if m.Config == nil || m.Config.ValidatorURL == "" {
		http.Error(w, "Validator service not configured", http.StatusInternalServerError)
		return nil
	}

	// Aquí podrías agregar lógica para validar la petición antes de enviarla al backend
	fmt.Printf("Procesando petición con Validator en: %s\n", m.Config.ValidatorURL)

	// Llama al siguiente middleware en la cadena
	return next.ServeHTTP(w, r)
}

func init() {
	caddy.RegisterModule(ValidatorMiddleware{})
}
