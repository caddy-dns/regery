package regery

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	dnsregery "github.com/olivierdeckers/dns-regery"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *dnsregery.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.regery",
		New: func() caddy.Module { return &Provider{new(dnsregery.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIToken = caddy.NewReplacer().ReplaceAll(p.Provider.APIToken, "")
	p.Provider.Secret = caddy.NewReplacer().ReplaceAll(p.Provider.Secret, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	regery {
//	    api_token <api_token>
//	    secret <secret>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				if d.NextArg() {
					p.Provider.APIToken = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "secret":
				if p.Provider.Secret != "" {
					return d.Err("Secret already set")
				}
				if d.NextArg() {
					p.Provider.Secret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	if p.Provider.Secret == "" {
		return d.Err("missing secret")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
