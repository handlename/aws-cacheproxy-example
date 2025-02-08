package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"
)

type Server struct {
	Config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		Config: config,
	}
}

// Handler handle http request via ServeMux.
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	logger.Debug("request received",
		zap.String("path", r.URL.Path),
		zap.Any("query", r.URL.Query()))

	u, err := url.Parse(r.URL.Query().Get("url"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid url")
		return
	}

	if !s.AllowedURL(u) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "proxy for requested URL is not allowed")
		return
	}

	content, status, err := s.Fetch(u)
	if err != nil {
		logger.Warn("failed to fetch content",
			zap.String("url", u.String()),
			zap.Error(err))

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(content)

	logger.Info("proxied",
		zap.String("url", u.String()),
		zap.Int("status", status),
		zap.Int("content_length", len(content)))
}

type SpecialChar rune

const (
	SpecialCharWildcard SpecialChar = '*'
)

// AllowedURL checks if the given URL is allowed to be accessed.
func (s *Server) AllowedURL(u *url.URL) bool {
	for _, host := range s.Config.AllowedHosts {
		if host.Name != u.Host {
			continue
		}

		for _, path := range host.Paths {
			if strings.HasSuffix(path, string(SpecialCharWildcard)) {
				prefix := strings.TrimSuffix(path, string(SpecialCharWildcard))
				if strings.HasPrefix(u.Path, prefix) {
					logger.Debug("matched with wildcard pattern", zap.String("host", host.Name), zap.String("path", path))
					return true
				}
			}

			if path == u.Path {
				logger.Debug("matched with exact pattern", zap.String("host", host.Name), zap.String("path", path))
				return true
			}
		}
	}

	return false
}

// Fetch fetches the content of the given URL.
func (s *Server) Fetch(u *url.URL) (body []byte, status int, err error) {
	r, err := http.Get(u.String())
	if err != nil {
		logger.Warn("failed to fetch content",
			zap.String("url", u.String()),
			zap.Error(err))

		return nil, 0, err
	}
	defer r.Body.Close()

	body, err = io.ReadAll(r.Body)
	if err != nil {
		logger.Warn("failed to read fetched body",
			zap.Error(err))
	}

	return body, r.StatusCode, nil
}
