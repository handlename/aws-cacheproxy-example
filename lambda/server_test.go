package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Server_AllowedURL(t *testing.T) {
	parse := func(s string) *url.URL {
		u, err := url.Parse(s)
		require.NoError(t, err)
		return u
	}

	specs := []struct {
		name   string
		host   ConfigAllowedHost
		url    *url.URL
		expect bool
	}{
		{
			name: "allowed",
			host: ConfigAllowedHost{
				Name:  "allowed.example",
				Paths: []string{"/foo"},
			},
			url:    parse("http://allowed.example/foo"),
			expect: true,
		},
		{
			name: "allowed: multiple paths",
			host: ConfigAllowedHost{
				Name:  "multiple-allowed-paths.example",
				Paths: []string{"/foo", "/bar"},
			},
			url:    parse("http://multiple-allowed-paths.example/bar"),
			expect: true,
		},
		{
			name: "allowed: wildcard",
			host: ConfigAllowedHost{
				Name:  "allowed-by-wildcard.example",
				Paths: []string{"/wild*"},
			},
			url:    parse("http://allowed-by-wildcard.example/wildcat"),
			expect: true,
		},
		{
			name: "not allowed: host name is different",
			host: ConfigAllowedHost{
				Name:  "allowed.example",
				Paths: []string{"/foo"},
			},
			url:    parse("http://not-allowed.example/foo"),
			expect: false,
		},
		{
			name: "not allowed: no paths matched",
			host: ConfigAllowedHost{
				Name:  "allowed.example",
				Paths: []string{"/foo"},
			},
			url:    parse("http://allowed.example/bar"),
			expect: false,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(t *testing.T) {
			conf := &Config{
				AllowedHosts: []ConfigAllowedHost{spec.host},
			}
			server := NewServer(conf)
			t.Logf("url: %+v", spec.url)
			assert.Equal(t, spec.expect, server.AllowedURL(spec.url))
		})
	}
}
