package validator

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// ValidatorMiddleware estructura del módulo
type ValidatorMiddleware struct{}

// CaddyModule registra el módulo en Caddy
func (ValidatorMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.validator",
		New: func() caddy.Module { return new(ValidatorMiddleware) },
	}
}

// ServeHTTP procesa la petición y la envía al servicio de validación
func (m ValidatorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	validationURL := "http://validator-service:9000/check"

	req, err := http.NewRequest("POST", validationURL, r.Body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Si la validación falla, bloquea la solicitud
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Blocked by validation service", http.StatusForbidden)
		return nil
	}

	// Si la validación es correcta, sigue con la cadena de middleware
	return next.ServeHTTP(w, r)
}

func init() {
	caddy.RegisterModule(ValidatorMiddleware{})
}
