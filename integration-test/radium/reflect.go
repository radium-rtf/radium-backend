package radium

import (
	"reflect"
)

func FieldsFirstStructEqualSecond(v1, v2 any) bool {
	t2 := reflect.ValueOf(v2)
	t1 := reflect.ValueOf(v1)

	for t2.Kind() == reflect.Pointer {
		t2 = t2.Elem()
	}
	for t1.Kind() == reflect.Pointer {
		t1 = t1.Elem()
	}

	for i := 0; i < t1.NumField(); i++ {
		field := t1.Type().Field(i)

		f2 := t2.FieldByName(field.Name)
		f1 := t1.FieldByName(field.Name)
		if f1.Kind() == reflect.Slice {
			continue
		}
		if !f1.Equal(f2) {
			return false
		}
	}

	return true
}
