package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Dump(i interface{}) {
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Struct, reflect.Slice, reflect.Map:
		b, _ := json.Marshal(i)
		fmt.Println(string(b))
	default:
		fmt.Println(i)
	}
}
