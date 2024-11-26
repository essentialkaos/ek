package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

// Encode converts query struct to URL-encoded string
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
				buf.WriteRune('=')
				queryFormatStringSlice(&buf, u)
			}

		case []fmt.Stringer:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatStringerSlice(&buf, u)
			}

		case []int:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []int8:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []int16:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []int32:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []int64:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []uint:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []uint8:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []uint16:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []uint32:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []uint64:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatNumSlice(&buf, u)
			}

		case []float32:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatFloatSlice(&buf, u)
			}

		case []float64:
			if len(u) > 0 {
				buf.WriteRune('=')
				queryFormatFloatSlice(&buf, u)
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

func queryFormatString(v string) string {
	return url.QueryEscape(v)
}

func queryFormatNumber(v any) string {
	return fmt.Sprintf("%d", v)
}

func queryFormatFloat(v any) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", v), "0"), ".")
}

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

func queryFormatNumSlice[T mathutil.Integer](buf *bytes.Buffer, v []T) {
	for _, vv := range v {
		buf.WriteString(queryFormatNumber(vv))
		buf.WriteRune(',')
	}

	buf.Truncate(buf.Len() - 1)
}

func queryFormatFloatSlice[T mathutil.Float](buf *bytes.Buffer, v []T) {
	for _, vv := range v {
		buf.WriteString(queryFormatFloat(vv))
		buf.WriteRune(',')
	}

	buf.Truncate(buf.Len() - 1)
}
