package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/essentialkaos/ek/v13/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Query is a map[string]any used for query
type Query map[string]any

// ////////////////////////////////////////////////////////////////////////////////// //

// QueryPayload is an interface for query payload with custom encoder
type QueryPayload interface {
	// ToQuery encodes payload for using in query string
	ToQuery(name string) string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Set sets query parameter
func (q Query) Set(name string, value any) bool {
	if q == nil || name == "" || value == nil {
		return false
	}

	q[name] = value

	return true
}

// SetIf sets query parameter if given condition is true
func (q Query) SetIf(cond bool, name string, value any) bool {
	if q == nil || name == "" || !cond {
		return false
	}

	q[name] = value

	return true
}

// Get returns parameter with given name
func (q Query) Get(name string) any {
	if q == nil || name == "" {
		return nil
	}

	return q[name]
}

// Delete deletes parameter with given name
func (q Query) Delete(name string) bool {
	if q == nil || name == "" {
		return false
	}

	delete(q, name)

	return true
}

// DeleteIf deletes parameter with given name if condition is true
func (q Query) DeleteIf(cond bool, name string) bool {
	if q == nil || name == "" || !cond {
		return false
	}

	if q[name] != nil {
		delete(q, name)
	}

	return true
}

// Encode encodes query parameters into a URL-encoded string
func (q Query) Encode() string {
	var buf bytes.Buffer

	for k, v := range q {
		if k == "" {
			continue
		}

		buf.WriteString(url.QueryEscape(k))

		switch u := v.(type) {
		case string:
			if u != "" {
				buf.WriteRune('=')
				buf.WriteString(queryFormatString(u))
			}

		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			buf.WriteRune('=')
			buf.WriteString(queryFormatNumber(v))

		case float32, float64:
			buf.WriteRune('=')
			buf.WriteString(queryFormatFloat(v))

		case nil:
			// noop

		case fmt.Stringer:
			vv := url.QueryEscape(u.String())

			if vv != "" {
				buf.WriteRune('=')
				buf.WriteString(vv)
			}

		case QueryPayload:
			vv := url.QueryEscape(u.ToQuery(k))

			if vv != "" {
				buf.WriteRune('=')
				buf.WriteString(vv)
			}

		case []string:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatStringSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatStringSlice(&buf, u)
				}
			}

		case []fmt.Stringer:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatStringerSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatStringerSlice(&buf, u)
				}
			}

		case []int:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []int8:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []int16:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []int32:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []int64:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []uint:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []uint8:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []uint16:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []uint32:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []uint64:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatNumSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatNumSlice(&buf, u)
				}
			}

		case []float32:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatFloatSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatFloatSlice(&buf, u)
				}
			}

		case []float64:
			if len(u) > 0 {
				if strings.HasSuffix(k, "[]") {
					buf.WriteRune('=')
					queryFormatFloatSliceSplit(&buf, k, u)
				} else {
					buf.WriteRune('=')
					queryFormatFloatSlice(&buf, u)
				}
			}

		default:
			buf.WriteRune('=')
			buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", v)))
		}

		buf.WriteRune('&')
	}

	if buf.Len() == 0 {
		return ""
	}

	buf.Truncate(buf.Len() - 1)

	return buf.String()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// queryFormatString formats string for query
func queryFormatString(v string) string {
	return url.QueryEscape(v)
}

// queryFormatNumber formats number for query
func queryFormatNumber(v any) string {
	return fmt.Sprintf("%d", v)
}

// queryFormatFloat formats float for query
func queryFormatFloat(v any) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", v), "0"), ".")
}

// queryFormatStringSlice formats slice of strings for query
func queryFormatStringSlice(buf *bytes.Buffer, v []string) {
	l := buf.Len()

	for _, vv := range v {
		if vv != "" {
			buf.WriteString(queryFormatString(vv))
			buf.WriteRune(',')
		}
	}

	if l != buf.Len() {
		buf.Truncate(buf.Len() - 1)
	}
}

// queryFormatStringSliceSplit formats slice of strings for query with key
func queryFormatStringSliceSplit(buf *bytes.Buffer, k string, v []string) {
	buf.WriteString(v[0])

	for _, vv := range v[1:] {
		buf.WriteRune('&')
		buf.WriteString(queryFormatString(k))
		buf.WriteRune('=')
		buf.WriteString(queryFormatString(vv))
	}
}

// queryFormatStringerSlice formats slice of fmt.Stringer compatible objects for query
func queryFormatStringerSlice(buf *bytes.Buffer, v []fmt.Stringer) {
	l := buf.Len()

	for _, vv := range v {
		vvv := vv.String()

		if vvv != "" {
			buf.WriteString(queryFormatString(vvv))
			buf.WriteRune(',')
		}
	}

	if l != buf.Len() {
		buf.Truncate(buf.Len() - 1)
	}
}

// queryFormatStringerSliceSplit formats slice of fmt.Stringer compatible objects for query
// with key
func queryFormatStringerSliceSplit(buf *bytes.Buffer, k string, v []fmt.Stringer) {
	buf.WriteString(v[0].String())

	for _, vv := range v[1:] {
		buf.WriteRune('&')
		buf.WriteString(queryFormatString(k))
		buf.WriteRune('=')
		buf.WriteString(queryFormatString(vv.String()))
	}
}

// queryFormatNumSlice formats slice of numbers for query
func queryFormatNumSlice[T mathutil.Integer](buf *bytes.Buffer, v []T) {
	for _, vv := range v {
		buf.WriteString(queryFormatNumber(vv))
		buf.WriteRune(',')
	}

	buf.Truncate(buf.Len() - 1)
}

// queryFormatNumSliceSplit formats slice of numbers for query with key
func queryFormatNumSliceSplit[T mathutil.Integer](buf *bytes.Buffer, k string, v []T) {
	buf.WriteString(queryFormatNumber(v[0]))

	for _, vv := range v[1:] {
		buf.WriteRune('&')
		buf.WriteString(queryFormatString(k))
		buf.WriteRune('=')
		buf.WriteString(queryFormatNumber(vv))
	}
}

// queryFormatFloatSlice formats slice of floats for query
func queryFormatFloatSlice[T mathutil.Float](buf *bytes.Buffer, v []T) {
	for _, vv := range v {
		buf.WriteString(queryFormatFloat(vv))
		buf.WriteRune(',')
	}

	buf.Truncate(buf.Len() - 1)
}

// queryFormatFloatSliceSplit formats slice of floats for query with key
func queryFormatFloatSliceSplit[T mathutil.Float](buf *bytes.Buffer, k string, v []T) {
	buf.WriteString(queryFormatFloat(v[0]))

	for _, vv := range v[1:] {
		buf.WriteRune('&')
		buf.WriteString(queryFormatString(k))
		buf.WriteRune('=')
		buf.WriteString(queryFormatFloat(vv))
	}
}
