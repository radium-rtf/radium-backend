package utils

import "reflect"

func RemoveEmptyMapFields(m map[string]interface{}) {
	for k, v := range m {
		if reflect.ValueOf(v).IsZero() {
			delete(m, k)
		}
	}
}
