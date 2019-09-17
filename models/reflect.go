package models

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ForEachFunc is the function that executes the query builder
type ForEachFunc func(field, operator string, value interface{}) error

// ForEachFilter reflects the struct an executes fn
func ForEachFilter(stru interface{}, fn ForEachFunc) error {
	tagName := "field"
	t := reflect.TypeOf(stru)
	v := reflect.ValueOf(stru)
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)
		if !fv.CanInterface() {
			continue
		}
		if fv.IsNil() { // skip nil values
			continue
		}
		if ft.PkgPath != "" { // unexported
			continue
		}
		var field, option string
		field, option = parseTag(ft.Tag.Get(tagName))
		if field == "-" {
			continue // ignore "-"
		}
		if field == "" {
			field = ft.Name // use field name
		}
		if option == "omitempty" {
			if isEmpty(&fv) {
				continue // skip empty field
			}
		}
		operator, ok := ft.Tag.Lookup("operator")
		if !ok {
			return errors.Errorf("operator tag required in field %s", field)
		}
		value := fv.Interface()
		if err := fn(field, operator, value); err != nil {
			return errors.Wrap(err, "Failed to exec ForeachFunc")
		}
	}
	return nil
}

func toString(fv reflect.Value) interface{} {
	kind := fv.Kind()
	if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
		return strconv.FormatInt(fv.Int(), 10)
	} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
		return strconv.FormatUint(fv.Uint(), 10)
	} else if kind == reflect.Float32 || kind == reflect.Float64 {
		return strconv.FormatFloat(fv.Float(), 'f', 2, 64)
	}
	return nil
}

func isEmpty(v *reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Bool {
		return v.Bool() == false
	} else if reflect.Int < k && k < reflect.Int64 {
		return v.Int() == 0
	} else if reflect.Uint < k && k < reflect.Uintptr {
		return v.Uint() == 0
	} else if k == reflect.Float32 || k == reflect.Float64 {
		return v.Float() == 0
	} else if k == reflect.Array || k == reflect.Map || k == reflect.Slice || k == reflect.String {
		return v.Len() == 0
	} else if k == reflect.Interface || k == reflect.Ptr {
		return v.IsNil()
	}
	return false
}

func parseTag(tag string) (string, string) {
	tags := strings.Split(tag, ",")
	if len(tags) <= 0 {
		return "", ""
	}
	if len(tags) == 1 {
		return tags[0], ""
	}
	return tags[0], tags[1]
}
