package validator

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

var (
	_ caddy.Module                = (*ValidatorMiddleware)(nil)
	_ caddy.Provisioner           = (*ValidatorMiddleware)(nil)
	_ caddyfile.Unmarshaler       = (*ValidatorMiddleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*ValidatorMiddleware)(nil)
)

func init() {
	caddy.RegisterModule(ValidatorMiddleware{})
	httpcaddyfile.RegisterHandlerDirective("validator", parseCaddyfileHandler)
}

type ValidatorMiddleware struct {
	ValidatorURL string `json:"validator_url,omitempty"`

	logger *zap.SugaredLogger
	secret []byte
}

// CaddyModule returns the Caddy module information.
func (ValidatorMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.validator",
		New: func() caddy.Module { return new(ValidatorMiddleware) },
	}
}

// Provision implements the caddy.Provisioner interface.
func (m *ValidatorMiddleware) Provision(ctx caddy.Context) error {
	if m.logger == nil {
		m.logger = ctx.Logger(m).Sugar()
	}

	return nil
}

// ServeHTTP implements the caddy.Handler interface.
func (m ValidatorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request,
	next caddyhttp.Handler,
) error {
	m.logger.Debugf("Esta llegando esto al validator")

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler. Syntax:
//
//	validator_url <url>
func (m *ValidatorMiddleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&m.ValidatorURL) {
			return d.ArgErr()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
	}

	return nil
}

func parseCaddyfileHandler(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler,
	error,
) {
	m := new(ValidatorMiddleware)
	if err := m.UnmarshalCaddyfile(h.Dispenser); err != nil {
		return nil, err
	}

	return m, nil
}
