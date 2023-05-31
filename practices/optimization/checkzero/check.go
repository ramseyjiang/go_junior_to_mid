package checkzero

import (
	"reflect"
)

func isNotEmpty[T any](value T) bool {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return false
	}
	switch v.Kind() {
	case reflect.Map, reflect.Slice, reflect.Chan:
		return v.Len() > 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !reflect.DeepEqual(v.Index(i).Interface(), reflect.Zero(v.Index(i).Type()).Interface()) {
				return true
			}
		}
		return false
	default:
		zero := reflect.Zero(v.Type()).Interface()
		return !reflect.DeepEqual(value, zero)
	}
}
