package caddy_plugins

import (
	"fmt"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("header_check", headerCheckParseCaddyfile)
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func headerCheckParseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := MiddlewareHeaderCheck{
		headerMap: make(map[string]string),
	}

	for h.Next() {
		var segments []caddyfile.Segment
		for nesting := h.Nesting(); h.NextBlock(nesting); {
			segment := h.NextSegment()
			m.headerMap[segment[0].Text] = segment[1].Text
			segments = append(segments, segment)
		}
	}

	fmt.Println(m.headerMap)

	return m, nil
}