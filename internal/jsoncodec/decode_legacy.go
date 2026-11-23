//go:build !go1.27 || !goexperiment.jsonv2

package jsoncodec

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode/utf16"
	"unicode/utf8"
)

func Unmarshal(data []byte, dst any) error       { return decode(data, dst, false) }
func UnmarshalStrict(data []byte, dst any) error { return decode(data, dst, true) }

// encoding/json v1 accepts duplicate names and matches struct fields regardless
// of case. Check the original JSON, including ignored members and RawMessage
// values, before handing it to v1. Only copy the input if a name would otherwise
// bind to a differently cased struct field.
func decode(data []byte, dst any, strict bool) error {
	if !json.Valid(data) || !utf8.Valid(data) || !validSurrogates(data) {
		return errors.New("invalid UTF-8 JSON")
	}
	r := reader{Decoder: json.NewDecoder(bytes.NewReader(data)), data: data, strict: strict}
	r.UseNumber()
	if err := r.value(reflect.TypeOf(dst)); err != nil {
		return err
	}
	if r.copied != 0 {
		r.filtered.Write(data[r.copied:])
		data = r.filtered.Bytes()
	}
	return json.Unmarshal(data, dst)
}

// Called after json.Valid, so every unicode escape has four hex digits.
func validSurrogates(data []byte) bool {
	for {
		i := bytes.IndexByte(data, '\\')
		if i < 0 {
			return true
		}
		data = data[i+1:]
		if data[0] != 'u' {
			data = data[1:]
			continue
		}
		code, _ := strconv.ParseUint(string(data[1:5]), 16, 16)
		data = data[5:]
		if !utf16.IsSurrogate(rune(code)) {
			continue
		}
		if code >= 0xdc00 || len(data) < 6 || data[0] != '\\' || data[1] != 'u' {
			return false
		}
		low, _ := strconv.ParseUint(string(data[2:6]), 16, 16)
		if low < 0xdc00 || low > 0xdfff {
			return false
		}
		data = data[6:]
	}
}

type reader struct {
	*json.Decoder
	data     []byte
	strict   bool
	filtered bytes.Buffer
	copied   int64
}

var unmarshaler = reflect.TypeFor[json.Unmarshaler]()

func (r *reader) value(t reflect.Type) error {
	token, err := r.Token()
	if err != nil {
		return err
	}
	delim, ok := token.(json.Delim)
	if !ok {
		return nil
	}
	for t != nil && t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	// Custom unmarshalers (including RawMessage) own their field names.
	if t != nil && reflect.PointerTo(t).Implements(unmarshaler) {
		t = nil
	}
	var elem reflect.Type
	var fields map[string]member
	if t != nil {
		switch t.Kind() {
		case reflect.Struct:
			fields = structMembers(t)
		case reflect.Map, reflect.Array, reflect.Slice:
			elem = t.Elem()
		}
	}
	var seen map[string]struct{}
	if delim == '{' {
		seen = make(map[string]struct{})
	}
	for r.More() {
		fieldType := elem
		if delim == '{' {
			start := r.InputOffset()
			key, err := r.Token()
			if err != nil {
				return err
			}
			name := key.(string) // json.Valid already established object syntax.
			if _, duplicate := seen[name]; duplicate {
				return errors.New("duplicate JSON object member")
			}
			seen[name] = struct{}{}
			if fields != nil {
				field, ok := fields[name]
				fieldType = field.typ
				if !ok {
					if r.strict {
						return fmt.Errorf("unknown JSON field %q", name)
					}
					// A NUL is not permitted in a Go struct tag name. Mask a
					// case-only match without changing valid unknown fields.
					for known := range fields {
						if strings.EqualFold(name, known) {
							end := r.InputOffset()
							quote := bytes.IndexByte(r.data[start:end], '"')
							r.filtered.Write(r.data[r.copied : start+int64(quote)])
							r.filtered.WriteString(`"\u0000"`)
							r.copied = end
							break
						}
					}
				}
			}
		}
		if err := r.value(fieldType); err != nil {
			return err
		}
	}
	_, err = r.Token()
	return err
}

type member struct {
	typ    reflect.Type
	depth  int
	tagged bool
}

var memberCache sync.Map // reflect.Type -> map[JSON name]member, immutable after publication

// Follow the standard embedded-field rules: the shallowest field wins, with
// tagged fields preferred at the same depth. Equal candidates are ambiguous.
func structMembers(t reflect.Type) map[string]member {
	if cached, ok := memberCache.Load(t); ok {
		return cached.(map[string]member)
	}
	fields := make(map[string]member)
	visiting := make(map[reflect.Type]bool)
	var visit func(reflect.Type, int)
	visit = func(st reflect.Type, depth int) {
		if visiting[st] {
			return
		}
		visiting[st] = true
		defer delete(visiting, st)
		for i := 0; i < st.NumField(); i++ {
			f := st.Field(i)
			ft := f.Type
			if ft.Kind() == reflect.Pointer {
				ft = ft.Elem()
			}
			tag := f.Tag.Get("json")
			if tag == "-" || (!f.IsExported() && !(f.Anonymous && ft.Kind() == reflect.Struct)) {
				continue
			}
			name, _, _ := strings.Cut(tag, ",")
			if f.Anonymous && ft.Kind() == reflect.Struct && name == "" {
				visit(ft, depth+1)
				continue
			}
			tagged := name != ""
			if !tagged {
				name = f.Name
			}
			old, ok := fields[name]
			switch {
			case !ok || depth < old.depth || (depth == old.depth && tagged && !old.tagged):
				fields[name] = member{f.Type, depth, tagged}
			case depth == old.depth && tagged == old.tagged:
				old.typ = nil
				fields[name] = old
			}
		}
	}
	visit(t, 0)
	for name, field := range fields {
		if field.typ == nil {
			delete(fields, name)
		}
	}
	actual, _ := memberCache.LoadOrStore(t, fields)
	return actual.(map[string]member)
}
