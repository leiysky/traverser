package traverser

import (
	"fmt"
	"reflect"
)

func Traverse(i interface{}) {
	v := reflect.ValueOf(i)
	name := v.Type().Name()
	if v.Kind() == reflect.Ptr {
		name = v.Elem().Type().Name()
	}
	fmt.Println("Ready to traverse:", name)
	traverse(name, v)
}

func traverse(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		return
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < v.Len(); idx++ {
			traverse(fmt.Sprintf("%s[%d]", path, idx), v.Index(idx))
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			traverse(fmt.Sprintf("%s[%v]", path, k), v.MapIndex(k))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			traverse(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil", path)
		} else {
			traverse(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	default:
		fmt.Println(path, "=", v)
	}
}
