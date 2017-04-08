// Package httputil provides methods for working with HTTP request/responses
package httputil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net/http"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var hasDesc bool
var statusDesc map[int]string

// ////////////////////////////////////////////////////////////////////////////////// //

// GetRequestAddr return host and port info from request
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

// GetRequestHost return host from request struct
func GetRequestHost(r *http.Request) string {
	host, _ := GetRequestAddr(r)
	return host
}

// GetRequestPort return port from request struct
func GetRequestPort(r *http.Request) string {
	_, port := GetRequestAddr(r)
	return port
}

// GetRemoteHost return network address that sent the request
func GetRemoteAddr(r *http.Request) (string, string) {
	addr := r.RemoteAddr

	if addr == "" || !strings.Contains(addr, ":") {
		return "", ""
	}

	addrSlice := strings.Split(addr, ":")

	return addrSlice[0], addrSlice[1]
}

// GetRemoteHost return host that sent the request
func GetRemoteHost(r *http.Request) string {
	host, _ := GetRemoteAddr(r)
	return host
}

// GetRemoteHost return host port that sent the request
func GetRemotePort(r *http.Request) string {
	_, port := GetRemoteAddr(r)
	return port
}

// GetDescByCode return response code description
func GetDescByCode(code int) string {
	if !hasDesc {
		statusDesc = map[int]string{
			100: "Continue",
			101: "Switching Protocols",

			200: "OK",
			201: "Created",
			202: "Accepted",
			203: "Non Authoritative Info",
			204: "No Content",
			205: "Reset Content",
			206: "Partial Content",

			300: "Multiple Choices",
			301: "Moved Permanently ",
			302: "Found",
			303: "See Other",
			304: "Not Modified",
			305: "Use Proxy",
			307: "Temporary Redirect",

			400: "Bad Request",
			401: "Unauthorized",
			402: "Payment Required",
			403: "Forbidden",
			404: "Not Found",
			405: "Method Not Allowed",
			406: "Not Acceptable",
			407: "Proxy Auth Required",
			408: "Request Timeout",
			409: "Conflict",
			410: "Gone",
			411: "Length Required",
			412: "Precondition Failed",
			413: "Request Entity Too Large",
			414: "Request URI TooLong",
			415: "Unsupported Media Type",
			416: "Requested Range Not Satisfiable",
			417: "Expectation Failed",
			418: "Teapot",

			500: "Internal Server Error",
			501: "Not Implemented",
			502: "Bad Gateway",
			503: "Service Unavailable",
			504: "Gateway Timeout",
			505: "HTTP Version Not Supported",
		}

		hasDesc = true
	}

	return statusDesc[code]
}

// IsURL check if given value is url or not
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
