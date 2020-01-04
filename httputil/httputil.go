// Package httputil provides methods for working with HTTP request/responses
package httputil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net/http"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetRequestAddr returns host and port info from request
func GetRequestAddr(r *http.Request) (string, string) {
	if r.Host == "" {
		return "", ""
	}

	hostSlice := strings.Split(r.Host, ":")

	switch len(hostSlice) {
	case 2:
		return hostSlice[0], hostSlice[1]
	default:
		return hostSlice[0], "80"
	}
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
	addr := r.RemoteAddr

	if addr == "" || !strings.Contains(addr, ":") {
		return "", ""
	}

	addrSlice := strings.Split(addr, ":")

	return addrSlice[0], addrSlice[1]
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

// GetDescByCode returns response code description
func GetDescByCode(code int) string {
	switch code {
	case 100:
		return "Continue"
	case 101:
		return "Switching Protocols"
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 202:
		return "Accepted"
	case 203:
		return "Non Authoritative Info"
	case 204:
		return "No Content"
	case 205:
		return "Reset Content"
	case 206:
		return "Partial Content"
	case 300:
		return "Multiple Choices"
	case 301:
		return "Moved Permanently "
	case 302:
		return "Found"
	case 303:
		return "See Other"
	case 304:
		return "Not Modified"
	case 305:
		return "Use Proxy"
	case 307:
		return "Temporary Redirect"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 402:
		return "Payment Required"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 406:
		return "Not Acceptable"
	case 407:
		return "Proxy Auth Required"
	case 408:
		return "Request Timeout"
	case 409:
		return "Conflict"
	case 410:
		return "Gone"
	case 411:
		return "Length Required"
	case 412:
		return "Precondition Failed"
	case 413:
		return "Request Entity Too Large"
	case 414:
		return "Request URI TooLong"
	case 415:
		return "Unsupported Media Type"
	case 416:
		return "Requested Range Not Satisfiable"
	case 417:
		return "Expectation Failed"
	case 418:
		return "Teapot"
	case 500:
		return "Internal Server Error"
	case 501:
		return "Not Implemented"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	case 504:
		return "Gateway Timeout"
	case 505:
		return "HTTP Version Not Supported"
	default:
		return "Unknown"
	}
}

// IsURL check if given value is valid URL or not
func IsURL(url string) bool {
	switch {
	case len(url) < 10:
		return false
	case url[0:7] == "http://":
		return true
	case url[0:8] == "https://":
		return true
	case url[0:6] == "ftp://":
		return true
	}

	return false
}
