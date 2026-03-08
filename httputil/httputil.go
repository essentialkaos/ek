// Package httputil provides methods for working with HTTP request/responses
package httputil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"net/http"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetRequestAddr returns host and port info from request
func GetRequestAddr(r *http.Request) (string, string) {
	if r == nil || r.Host == "" || r.URL == nil {
		return "", ""
	}

	if !strings.ContainsRune(r.Host, ':') {
		return r.Host, GetPortByScheme(r.URL.Scheme)
	}

	host, port, _ := net.SplitHostPort(r.Host)

	return host, port
}

// GetRequestHost returns host from request struct
func GetRequestHost(r *http.Request) string {
	host, _ := GetRequestAddr(r)
	return host
}

// GetRequestPort returns port from request struct
func GetRequestPort(r *http.Request) string {
	_, port := GetRequestAddr(r)
	return port
}

// GetRemoteAddr returns network address that sent the request
func GetRemoteAddr(r *http.Request) (string, string) {
	if r == nil {
		return "", ""
	}

	addr := r.RemoteAddr

	if addr == "" || !strings.ContainsRune(addr, ':') {
		return "", ""
	}

	host, port, _ := net.SplitHostPort(addr)

	return host, port
}

// GetRemoteHost returns host that sent the request
func GetRemoteHost(r *http.Request) string {
	host, _ := GetRemoteAddr(r)
	return host
}

// GetRemotePort returns port of the host that sent the request
func GetRemotePort(r *http.Request) string {
	_, port := GetRemoteAddr(r)
	return port
}

// GetPortByScheme returns port for supported scheme
func GetPortByScheme(s string) string {
	switch strings.ToLower(s) {
	case "http", "ws":
		return "80"
	case "https", "wss":
		return "443"
	case "ftp":
		return "21"
	}

	return ""
}

// IsURL checks if given value is valid URL or not
func IsURL(s string) bool {
	return IsHTTPS(s) || IsHTTP(s) || IsFTP(s)
}

// IsHTTP returns true if given URL contains http scheme
func IsHTTP(s string) bool {
	return strings.HasPrefix(s, "http://")
}

// IsHTTPS returns true if given URL contains https scheme
func IsHTTPS(s string) bool {
	return strings.HasPrefix(s, "https://")
}

// IsFTP returns true if given URL contains ftp scheme
func IsFTP(s string) bool {
	return strings.HasPrefix(s, "ftp://")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetDescByCode returns response code description
//
// Deprecated: Use [http.StatusText] instead
func GetDescByCode(code int) string {
	return http.StatusText(code)
}
