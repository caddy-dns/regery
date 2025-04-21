Regery module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Regery.

## Caddy module name

```
dns.providers.regery
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "regery",
        "api_token": "YOUR_API_TOKEN",
        "secret": "YOUR_API_SECRET"
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally
{
  acme_dns regery {
    api_token YOUR_API_TOKEN
    secret YOUR_API_SECRET
  }
}
```

```
# one site
tls {
  dns regery {
    api_token YOUR_API TOKEN
    secret YOUR_API_SECRET
  }
}
```
