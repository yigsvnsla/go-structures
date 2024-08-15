package utils

import (
	"fmt"
	"reflect"
)

func PrintStruct(v interface{}) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// Asegúrate de que v sea un struct
	if rv.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct")
		return
	}

	fmt.Println("Struct fields and values:")
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}
