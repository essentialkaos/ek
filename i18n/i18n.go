// Package i18n provides provides methods and structs for internationalization
package i18n

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/pluralize"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// UNKNOWN_VALUE contains string representation of unknown value
const UNKNOWN_VALUE = "???"

// ////////////////////////////////////////////////////////////////////////////////// //

// String is a simple string
type String string

// Data can be used for storing data for templates
type Data map[string]any

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilBundle is returned when one or more provided bundles is nil
	ErrNilBundle = fmt.Errorf("One or more provided bundles is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Fallback copies values between bundles to use it as a fallback
func Fallback(bundles ...any) (any, error) {
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
		} else if bundleType != bv.Type().Name() {
			return nil, fmt.Errorf(
				"Given bundles have different types (%s, %s)",
				bundleType, bv.Type().Name(),
			)
		}
	}

	copyFieldsData(bundleVars)

	return bundles[len(bundles)-1], nil
}

// ValidateBundle validates bundle for problems
func ValidateBundle(bundle any) error {
	return checkBundle(reflect.Indirect(reflect.ValueOf(bundle)), "")
}

// IsComplete checks if given bundle is complete and returns slice with
// empty fields
func IsComplete(bundle any) (bool, []string) {
	if bundle == nil {
		return false, nil
	}

	return isCompleteBundle(reflect.Indirect(reflect.ValueOf(bundle)), "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// With uses String as a templates and applies payload from given data to it
func (s String) With(data any) string {
	if s == "" || data == nil ||
		!strings.Contains(string(s), `{{`) ||
		!strings.Contains(string(s), `}}`) {
		return string(s)
	}

	t, err := template.New("").Parse(string(s))

	if err != nil {
		return string(s)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)

	if err != nil {
		return string(s)
	}

	return buf.String()
}

// Add adds prefix and/or suffix to result string
func (s String) Add(prefix, suffix string) string {
	return prefix + string(s) + suffix
}

// String converts String to string type
func (s String) String() string {
	return string(s)
}

// S is shortcut for String
func (s String) S() string {
	return string(s)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Plural prints plural form of word based on language and value
//
// Note that this method only supports int, int64, uint, uint64, and float64
func (d Data) Plural(lang, prop string, values ...string) string {
	if len(values) == 0 {
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

// PrettyNum formats number to "pretty" form (e.g 1234567 -> 1,234,567)
func (d Data) PrettyNum(prop string) string {
	if d[prop] == nil {
		return UNKNOWN_VALUE
	}

	return fmtutil.PrettyNum(d[prop])
}

// PrettySize formats value to "pretty" size (e.g 1478182 -> 1.34 Mb)
//
// Note that this method only supports int, int64, uint, uint64, and float64
func (d Data) PrettySize(prop string) string {
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

	return fmt.Sprintf("%s", d[prop])
}

// PrettySize formats value to "pretty" size (e.g 1478182 -> 1.34 Mb)
//
// Note that this method only supports float64
func (d Data) PrettyPerc(prop string) string {
	t, ok := d[prop].(float64)

	if !ok {
		return UNKNOWN_VALUE
	}

	return fmtutil.PrettyPerc(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// copyFieldsData copy fields data between bundles
func copyFieldsData(bundles []reflect.Value) {
	for i := range bundles[0].NumField() {
		f := bundles[0].Field(i)
		s := f.Type().String()

		if !f.IsZero() {
			continue
		}

		if s == "i18n.String" {
		DEFLOOP:
			for j := 1; j < len(bundles); j++ {
				jf := bundles[j].Field(i)
				if !jf.IsZero() {
					f.SetString(jf.String())
					break DEFLOOP
				}
			}

			continue
		}

		if strings.HasPrefix(s, "*") {
			copyFieldsData(getFields(bundles, i))
		}
	}
}

// getFields returns slice with fields values
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

// checkBundle checks bundle for nil structs
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

// isCompleteBundle checks given object for empty fields
func isCompleteBundle(v reflect.Value, parentName string) (bool, []string) {
	if !v.IsValid() {
		return false, []string{strings.TrimRight(parentName, ".")}
	}

	var incompleteFields []string

	for i := range v.NumField() {
		f := v.Field(i)
		typ := f.Type().String()

		if typ == "i18n.String" {
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

// getPluralizerByLang returns pluralization function by language name
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
