package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Auth is interface for authentication method
type Auth interface {
	// Apply sets authentication data for given request
	Apply(r *http.Request, header string)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// AuthBasic is auth using username and password (RFC 7617)
//
// https://datatracker.ietf.org/doc/html/rfc7617
type AuthBasic struct {
	Username string
	Password string
}

// AuthOAuth is auth using Bearer token (RFC 6750)
//
// https://datatracker.ietf.org/doc/html/rfc6750
type AuthBearer struct {
	Token string
}

// AuthOAuth is auth using OAuth token
//
// https://www.rfc-editor.org/rfc/rfc5849.html#section-3.5.1
type AuthOAuth struct {
	Realm           string
	ConsumerKey     string
	Token           string
	SignatureMethod string
	Signature       string
	Nonce           string
	Version         string
	Timestamp       int64
}

// AuthDigest is auth using Digest Auth (RFC 7616)
//
// https://datatracker.ietf.org/doc/html/rfc7616
type AuthDigest struct {
	Username  string
	Realm     string
	URI       string
	Algorithm string
	Nonce     string
	CNonce    string
	NC        uint
	QOP       string
	Response  string
	Opaque    string
	UserHash  bool
}

// AuthVAPID s auth using Voluntary Application Server Identification (RFC 8292)
//
// https://datatracker.ietf.org/doc/html/rfc8292
type AuthVAPID struct {
	Credential    string
	SignedHeaders string
	Signature     string
}

// AuthAWS4 is auth using AWS Signature Version 4
//
// https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-auth-using-authorization-header.html
type AuthAWS4 struct {
	Credential    string
	SignedHeaders string
	Signature     string
}

// AuthAPIKey is auth using X-API-Key/API-Key header
type AuthAPIKey struct {
	Key string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Apply sets authentication data for given request
func (a AuthBasic) Apply(r *http.Request, header string) {
	r.Header.Set(header, "Basic "+base64.StdEncoding.EncodeToString(
		[]byte(a.Username+":"+a.Password),
	))
}

// Apply sets authentication data for given request
func (a AuthBearer) Apply(r *http.Request, header string) {
	r.Header.Set(header, "Bearer "+a.Token)
}

// Apply sets authentication data for given request
func (a AuthOAuth) Apply(r *http.Request, header string) {
	var buf bytes.Buffer

	buf.WriteString("OAuth ")

	if a.Realm != "" {
		fmt.Fprintf(&buf, "realm=%q, ", a.Realm)
	}

	if a.ConsumerKey != "" {
		fmt.Fprintf(&buf, "oauth_consumer_key=%q, ", a.ConsumerKey)
	}

	if a.Token != "" {
		fmt.Fprintf(&buf, "oauth_token=%q, ", a.Token)
	}

	if a.SignatureMethod != "" {
		fmt.Fprintf(&buf, "oauth_signature_method=%q, ", a.SignatureMethod)
	}

	if a.Signature != "" {
		fmt.Fprintf(&buf, "oauth_signature=%q, ", a.Signature)
	}

	fmt.Fprintf(&buf, "oauth_timestamp=\"%d\", ", a.Timestamp)

	if a.Nonce != "" {
		fmt.Fprintf(&buf, "oauth_nonce=%q, ", a.Nonce)
	}

	if a.Version != "" {
		fmt.Fprintf(&buf, "oauth_version=%q, ", a.Version)
	}

	if buf.Len() > 2 {
		buf.Truncate(buf.Len() - 2)
	}

	r.Header.Set(header, buf.String())
}

// Apply sets authentication data for given request
func (a AuthDigest) Apply(r *http.Request, header string) {
	var buf bytes.Buffer

	buf.WriteString("Digest ")

	if a.Username != "" {
		fmt.Fprintf(&buf, "username=%q, ", a.Username)
	}

	if a.Realm != "" {
		fmt.Fprintf(&buf, "realm=%q, ", a.Realm)
	}

	if a.URI != "" {
		fmt.Fprintf(&buf, "uri=%q, ", a.URI)
	}

	if a.Algorithm != "" {
		fmt.Fprintf(&buf, "algorithm=%s, ", a.Algorithm)
	}

	if a.Nonce != "" {
		fmt.Fprintf(&buf, "nonce=%q, ", a.Nonce)
	}

	if a.CNonce != "" {
		fmt.Fprintf(&buf, "cnonce=%q, ", a.CNonce)
	}

	fmt.Fprintf(&buf, "nc=%08d, ", a.NC)

	if a.QOP != "" {
		fmt.Fprintf(&buf, "qop=%s, ", a.QOP)
	}

	if a.Response != "" {
		fmt.Fprintf(&buf, "response=%q, ", a.Response)
	}

	if a.Opaque != "" {
		fmt.Fprintf(&buf, "opaque=%q, ", a.Opaque)
	}

	if a.UserHash {
		fmt.Fprint(&buf, "userhash=true, ")
	}

	if buf.Len() > 2 {
		buf.Truncate(buf.Len() - 2)
	}

	r.Header.Set(header, buf.String())
}

// Apply sets authentication data for given request
func (a AuthAWS4) Apply(r *http.Request, header string) {
	r.Header.Set("Authorization", fmt.Sprintf(
		"AWS4-HMAC-SHA256 Credential=%s,SignedHeaders=%s,Signature=%s",
		a.Credential, a.SignedHeaders, a.Signature,
	))
}

// Apply sets authentication data for given request
func (a AuthAPIKey) Apply(r *http.Request, header string) {
	r.Header.Set("X-API-Key", a.Key)
	r.Header.Set("API-Key", a.Key)
}
