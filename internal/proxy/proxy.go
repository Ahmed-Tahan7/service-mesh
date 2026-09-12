package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	target       *url.URL
	reverseProxy *httputil.ReverseProxy
}

func New(target string) (*Proxy, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy target %q: %w", target, err)
	}

	if targetURL.Scheme == "" || targetURL.Host == "" {
		return nil, fmt.Errorf("proxy target must include scheme and host: %q", target)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	return &Proxy{
		target:       targetURL,
		reverseProxy: reverseProxy,
	}, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.reverseProxy.ServeHTTP(w, r)
}