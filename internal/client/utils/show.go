package utils

import (
	"fmt"
	"reflect"
)

func ShowResult(s interface{}) {
	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if valueField.String() != "" && typeField.Name != "Binary" {
			fmt.Printf("%s\t : %v\n", typeField.Name, valueField.Interface())
		}
	}
}
