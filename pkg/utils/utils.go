package utils

import (
	"github.com/fatih/structs"
	"reflect"
)

func RemoveEmptyMapFields(m map[string]interface{}) {
	for k, v := range m {
		if reflect.ValueOf(v).IsZero() {
			delete(m, k)
		}
	}
}

func RemoveEmptyFields(s any) map[string]interface{} {
	m := structs.Map(s)
	RemoveEmptyMapFields(m)
	return m
}
