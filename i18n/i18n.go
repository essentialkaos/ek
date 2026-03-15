// Package i18n provides methods and structs for internationalization
package i18n

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/template"

	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/pluralize"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// UNKNOWN_VALUE is the string returned when a data property is missing or unrecognised
const UNKNOWN_VALUE = "???"

// ////////////////////////////////////////////////////////////////////////////////// //

// Text is a localisation string that supports fmt-style formatting and Go templates
type Text string

// Data holds key-value pairs passed to a Text template during rendering
type Data map[string]any

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilBundle is returned when one or more bundles passed to [Fallback] is nil
	ErrNilBundle = errors.New("one or more provided bundles is nil")

	// ErrNilWriter is returned when a nil writer is passed to [Text.Write]
	ErrNilWriter = errors.New("target writer is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Fallback merges a chain of bundles right-to-left, copying missing Text fields
// from lower-priority bundles into the highest-priority one
func Fallback[T any](bundles ...*T) (*T, error) {
	switch {
	case len(bundles) == 0,
		len(bundles) == 1 && bundles[0] == nil:
		return nil, ErrNilBundle
	case len(bundles) == 1:
		err := ValidateBundle(bundles[0])

		if err != nil {
			return nil, err
		}

		return bundles[0], nil
	}

	var bundleType string
	var bundleVars []reflect.Value

	for i := len(bundles) - 1; i >= 0; i-- {
		b := bundles[i]

		if b == nil {
			return nil, ErrNilBundle
		}

		bv := reflect.Indirect(reflect.ValueOf(b))
		bundleVars = append(bundleVars, bv)

		if bundleType == "" {
			bundleType = bv.Type().Name()
		}
	}

	copyFieldsData(bundleVars)

	return bundles[len(bundles)-1], nil
}

// ValidateBundle returns an error if any pointer sub-bundle inside bundle is nil
func ValidateBundle(bundle any) error {
	return checkBundle(reflect.Indirect(reflect.ValueOf(bundle)), "")
}

// IsComplete reports whether all Text fields in bundle are non-empty and returns
// a slice of dot-separated field paths that are still unset
func IsComplete(bundle any) (bool, []string) {
	if bundle == nil {
		return false, nil
	}

	return isCompleteBundle(reflect.Indirect(reflect.ValueOf(bundle)), "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// With renders [Text] as a Go template against data and returns the result as a string.
// Optional extra[0] and extra[1] are prepended and appended to the output respectively.
func (t Text) With(data any, extra ...string) string {
	var buf bytes.Buffer

	t.Write(&buf, data, extra...)

	return buf.String()
}

// Write renders [Text] as a Go template against data and writes the result to wr.
// Optional extra[0] and extra[1] are written before and after the rendered output.
func (t Text) Write(wr io.Writer, data any, extra ...string) error {
	if wr == nil {
		return ErrNilWriter
	}

	if len(extra) > 0 {
		fmt.Fprint(wr, extra[0])
	}

	if t == "" || data == nil ||
		!strings.Contains(string(t), `{{`) ||
		!strings.Contains(string(t), `}}`) {
		writeWithExtra(wr, t, extra...)
		return nil
	}

	tt, err := template.New("").Parse(string(t))

	if err != nil {
		writeWithExtra(wr, t, extra...)
		return fmt.Errorf("can't parse template: %w", err)
	}

	err = tt.Execute(wr, data)

	if err != nil {
		writeWithExtra(wr, t, extra...)
		return fmt.Errorf("can't apply template: %w", err)
	}

	if len(extra) > 1 {
		fmt.Fprint(wr, extra[1])
	}

	return nil
}

// Format interpolates Text using fmt.Sprintf-style arguments
func (t Text) Format(a ...any) string {
	return fmt.Sprintf(string(t), a...)
}

// Error interpolates Text using fmt.Sprintf-style arguments and returns it
// as an error
func (t Text) Error(a ...any) error {
	return fmt.Errorf(string(t), a...)
}

// Add returns the text with prefix prepended and suffix appended
func (t Text) Add(prefix, suffix string) string {
	return prefix + string(t) + suffix
}

// Start returns the text with s prepended
func (t Text) Start(s string) string {
	return s + string(t)
}

// End returns the text with s appended
func (t Text) End(s string) string {
	return string(t) + s
}

// String returns the underlying string value of [Text]
func (t Text) String() string {
	return string(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// S is shortcut for [Text.String]
func (t Text) S() string {
	return string(t)
}

// F is shortcut for [Text.Format]
func (t Text) F(a ...any) string {
	return t.Format(a...)
}

// E is shortcut for [Text.Error]
func (t Text) E(a ...any) error {
	return t.Error(a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if property with given name is present in data map
func (d Data) Has(prop string) bool {
	if len(d) == 0 {
		return false
	}

	_, ok := d[prop]

	return ok
}

// Plural returns the correct plural form of values based on the numeric property
// prop and the rules of the given language code. Supports int, int64, uint, uint64,
// and float64.
func (d Data) Plural(lang, prop string, values ...string) string {
	if len(values) == 0 || len(d) == 0 {
		return ""
	}

	switch t := d[prop].(type) {
	case int:
		return pluralize.PluralizeSpecial(getPluralizerByLang(lang), t, values...)
	case int64:
		return pluralize.PluralizeSpecial(getPluralizerByLang(lang), t, values...)
	case uint:
		return pluralize.PluralizeSpecial(getPluralizerByLang(lang), t, values...)
	case uint64:
		return pluralize.PluralizeSpecial(getPluralizerByLang(lang), t, values...)
	case float64:
		return pluralize.PluralizeSpecial(getPluralizerByLang(lang), t, values...)
	}

	return values[0]
}

// PrettyNum formats the numeric property prop with thousands separators
// (e.g. 1234567 → "1,234,567")
func (d Data) PrettyNum(prop string) string {
	if !d.Has(prop) {
		return UNKNOWN_VALUE
	}

	return fmtutil.PrettyNum(d[prop])
}

// PrettySize formats the numeric property prop as a human-readable byte size
// (e.g. 1478182 → "1.4 MB"). Supports int, int64, uint, uint64, and float64.
func (d Data) PrettySize(prop string) string {
	if !d.Has(prop) {
		return UNKNOWN_VALUE
	}

	switch t := d[prop].(type) {
	case int:
		return fmtutil.PrettySize(t)
	case int64:
		return fmtutil.PrettySize(t)
	case uint:
		return fmtutil.PrettySize(t)
	case uint64:
		return fmtutil.PrettySize(t)
	case float64:
		return fmtutil.PrettySize(t)
	case nil:
		return UNKNOWN_VALUE
	}

	return fmt.Sprintf("%v", d[prop])
}

// PrettyPerc formats the float64 property prop as a percentage string
// (e.g. 45.31 → "45.3%")
func (d Data) PrettyPerc(prop string) string {
	if !d.Has(prop) {
		return UNKNOWN_VALUE
	}

	t, ok := d[prop].(float64)

	if !ok {
		return UNKNOWN_VALUE
	}

	return fmtutil.PrettyPerc(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// writeWithExtra writes t to wr, appending extra[1] as a suffix when present
func writeWithExtra(wr io.Writer, t Text, extra ...string) {
	switch len(extra) {
	case 0, 1:
		fmt.Fprint(wr, t)
	default:
		fmt.Fprintf(wr, "%s%s", t, extra[1])
	}
}

// copyFieldsData fills empty Text fields in bundles[0] from the first bundle
// further right in the slice that holds a non-zero value; recurses into pointer fields
func copyFieldsData(bundles []reflect.Value) {
	if len(bundles) == 0 {
		return
	}

	for i := range bundles[0].NumField() {
		f := bundles[0].Field(i)
		s := f.Type().String()

		if !f.IsZero() {
			continue
		}

		if s == "i18n.Text" {
			for j := 1; j < len(bundles); j++ {
				jf := bundles[j].Field(i)
				if !jf.IsZero() {
					f.SetString(jf.String())
					break
				}
			}

			continue
		}

		if strings.HasPrefix(s, "*") {
			copyFieldsData(getFields(bundles, i))
		}
	}
}

// getFields returns the dereferenced field at index from each bundle,
// allocating a new zero value for any bundle whose field is nil
func getFields(bundles []reflect.Value, index int) []reflect.Value {
	var result []reflect.Value

	for _, b := range bundles {
		f := b.Field(index)

		if f.IsNil() {
			f.Set(reflect.New(f.Type().Elem()))
		}

		result = append(result, b.Field(index).Elem())
	}

	return result
}

// checkBundle returns an error if any pointer field inside v is nil
func checkBundle(v reflect.Value, parentName string) error {
	if !v.IsValid() {
		return fmt.Errorf("Bundle struct %s is nil", strings.TrimRight(parentName, "."))
	}

	for i := range v.NumField() {
		f := v.Field(i)
		typ := f.Type().String()

		if strings.HasPrefix(typ, "*") {
			err := checkBundle(f.Elem(), parentName+v.Type().Field(i).Name+".")

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// isCompleteBundle returns false and a list of dot-separated field paths
// for every empty Text field found recursively inside v
func isCompleteBundle(v reflect.Value, parentName string) (bool, []string) {
	if !v.IsValid() {
		return false, []string{strings.TrimRight(parentName, ".")}
	}

	var incompleteFields []string

	for i := range v.NumField() {
		f := v.Field(i)
		typ := f.Type().String()

		if typ == "i18n.Text" {
			if f.IsZero() {
				incompleteFields = append(
					incompleteFields,
					parentName+v.Type().Field(i).Name,
				)
			}
		} else if strings.HasPrefix(typ, "*") {
			_, fs := isCompleteBundle(f.Elem(), parentName+v.Type().Field(i).Name+".")

			if len(fs) > 0 {
				incompleteFields = append(incompleteFields, fs...)
			}
		}
	}

	return len(incompleteFields) == 0, incompleteFields
}

// getPluralizerByLang returns the pluralisation function for the given
// BCP 47-style language code (case-insensitive), falling back to
// [DefaultPluralizer] if unknown
func getPluralizerByLang(lang string) pluralize.Pluralizer {
	switch strings.ToUpper(lang) {
	case "ACH":
		return pluralize.Ach
	case "AF":
		return pluralize.Af
	case "AK":
		return pluralize.Ak
	case "AM":
		return pluralize.Am
	case "AN":
		return pluralize.An
	case "ANP":
		return pluralize.Anp
	case "AR":
		return pluralize.Ar
	case "ARN":
		return pluralize.Arn
	case "AS":
		return pluralize.As
	case "AST":
		return pluralize.Ast
	case "AY":
		return pluralize.Ay
	case "AZ":
		return pluralize.Az
	case "BE":
		return pluralize.Be
	case "BG":
		return pluralize.Bg
	case "BN":
		return pluralize.Bn
	case "BO":
		return pluralize.Bo
	case "BR":
		return pluralize.Br
	case "BRX":
		return pluralize.Brx
	case "BS":
		return pluralize.Bs
	case "CA":
		return pluralize.Ca
	case "CGG":
		return pluralize.Cgg
	case "CS":
		return pluralize.Cs
	case "CSB":
		return pluralize.Csb
	case "CY":
		return pluralize.Cy
	case "DA":
		return pluralize.Da
	case "DE":
		return pluralize.De
	case "DOI":
		return pluralize.Doi
	case "DZ":
		return pluralize.Dz
	case "EL":
		return pluralize.El
	case "EN":
		return pluralize.En
	case "EO":
		return pluralize.Eo
	case "ES":
		return pluralize.Es
	case "ESAR":
		return pluralize.EsAR
	case "ET":
		return pluralize.Et
	case "EU":
		return pluralize.Eu
	case "FA":
		return pluralize.Fa
	case "FF":
		return pluralize.Ff
	case "FI":
		return pluralize.Fi
	case "FIL":
		return pluralize.Fil
	case "FO":
		return pluralize.Fo
	case "FR":
		return pluralize.Fr
	case "FUR":
		return pluralize.Fur
	case "FY":
		return pluralize.Fy
	case "GA":
		return pluralize.Ga
	case "GD":
		return pluralize.Gd
	case "GL":
		return pluralize.Gl
	case "GU":
		return pluralize.Gu
	case "GUN":
		return pluralize.Gun
	case "HA":
		return pluralize.Ha
	case "HE":
		return pluralize.He
	case "HI":
		return pluralize.Hi
	case "HNE":
		return pluralize.Hne
	case "HR":
		return pluralize.Hr
	case "HU":
		return pluralize.Hu
	case "HY":
		return pluralize.Hy
	case "IA":
		return pluralize.Ia
	case "ID":
		return pluralize.Id
	case "IS":
		return pluralize.Is
	case "IT":
		return pluralize.It
	case "JA":
		return pluralize.Ja
	case "JBO":
		return pluralize.Jbo
	case "JV":
		return pluralize.Jv
	case "KA":
		return pluralize.Ka
	case "KK":
		return pluralize.Kk
	case "KL":
		return pluralize.Kl
	case "KM":
		return pluralize.Km
	case "KN":
		return pluralize.Kn
	case "KO":
		return pluralize.Ko
	case "KU":
		return pluralize.Ku
	case "KW":
		return pluralize.Kw
	case "KY":
		return pluralize.Ky
	case "LB":
		return pluralize.Lb
	case "LN":
		return pluralize.Ln
	case "LO":
		return pluralize.Lo
	case "LT":
		return pluralize.Lt
	case "LV":
		return pluralize.Lv
	case "MAI":
		return pluralize.Mai
	case "MFE":
		return pluralize.Mfe
	case "MG":
		return pluralize.Mg
	case "MI":
		return pluralize.Mi
	case "MK":
		return pluralize.Mk
	case "ML":
		return pluralize.Ml
	case "MN":
		return pluralize.Mn
	case "MNI":
		return pluralize.Mni
	case "MNK":
		return pluralize.Mnk
	case "MR":
		return pluralize.Mr
	case "MS":
		return pluralize.Ms
	case "MT":
		return pluralize.Mt
	case "MY":
		return pluralize.My
	case "NAH":
		return pluralize.Nah
	case "NAP":
		return pluralize.Nap
	case "NB":
		return pluralize.Nb
	case "NE":
		return pluralize.Ne
	case "NL":
		return pluralize.Nl
	case "NN":
		return pluralize.Nn
	case "NO":
		return pluralize.No
	case "NSO":
		return pluralize.Nso
	case "OC":
		return pluralize.Oc
	case "OR":
		return pluralize.Or
	case "PA":
		return pluralize.Pa
	case "PAP":
		return pluralize.Pap
	case "PL":
		return pluralize.Pl
	case "PMS":
		return pluralize.Pms
	case "PS":
		return pluralize.Ps
	case "PT":
		return pluralize.Pt
	case "PTBR":
		return pluralize.PtBR
	case "RM":
		return pluralize.Rm
	case "RO":
		return pluralize.Ro
	case "RU":
		return pluralize.Ru
	case "RW":
		return pluralize.Rw
	case "SAH":
		return pluralize.Sah
	case "SAT":
		return pluralize.Sat
	case "SCO":
		return pluralize.Sco
	case "SD":
		return pluralize.Sd
	case "SE":
		return pluralize.Se
	case "SI":
		return pluralize.Si
	case "SK":
		return pluralize.Sk
	case "SL":
		return pluralize.Sl
	case "SO":
		return pluralize.So
	case "SON":
		return pluralize.Son
	case "SQ":
		return pluralize.Sq
	case "SR":
		return pluralize.Sr
	case "SU":
		return pluralize.Su
	case "SV":
		return pluralize.Sv
	case "SW":
		return pluralize.Sw
	case "TA":
		return pluralize.Ta
	case "TE":
		return pluralize.Te
	case "TG":
		return pluralize.Tg
	case "TH":
		return pluralize.Th
	case "TI":
		return pluralize.Ti
	case "TK":
		return pluralize.Tk
	case "TR":
		return pluralize.Tr
	case "TT":
		return pluralize.Tt
	case "UG":
		return pluralize.Ug
	case "UK":
		return pluralize.Uk
	case "UR":
		return pluralize.Ur
	case "UZ":
		return pluralize.Uz
	case "VI":
		return pluralize.Vi
	case "WA":
		return pluralize.Wa
	case "WO":
		return pluralize.Wo
	case "YO":
		return pluralize.Yo
	case "ZH":
		return pluralize.Zh
	}

	return pluralize.DefaultPluralizer
}
