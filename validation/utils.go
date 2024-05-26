package validation

import (
	"reflect"
	"time"
)

func isEmpty(val any) bool {
	vt := reflect.ValueOf(val)
	kind := vt.Kind()
	switch kind {
	case reflect.String, reflect.Map, reflect.Slice, reflect.Array:
		return vt.Len() == 0
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Ptr:
		if vt.IsNil() {
			return true
		}
		return isEmpty(vt.Elem().Interface())
	case reflect.Struct:
		v, ok := val.(time.Time)
		if ok && v.IsZero() {
			return true
		}
	}
	return false
}
