package traverser

import (
	"fmt"
	"reflect"
	"runtime"
)

type Filter func(reflect.Value) bool

func IsZeroValue(v reflect.Value) bool {
	defer func() {
		recover()
	}()
	if v.Interface() == reflect.Zero(v.Type()).Interface() {
		return true
	}
	return false
}

func NotZeroValue(v reflect.Value) bool {
	return !IsZeroValue(v)
}

func Traverse(i interface{}, ops ...Filter) {
	v := reflect.ValueOf(i)
	name := v.Type().Name()
	if v.Kind() == reflect.Ptr {
		name = v.Elem().Type().Name()
	}
	fmt.Println("Ready to traverse:", name)
	for _, f := range ops {
		fmt.Println("Filter: ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	}
	traverse(name, v, ops...)
}

func traverse(path string, v reflect.Value, ops ...Filter) {
	switch v.Kind() {
	case reflect.Invalid:
		return
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < v.Len(); idx++ {
			traverse(fmt.Sprintf("%s[%d]", path, idx), v.Index(idx), ops...)
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			traverse(fmt.Sprintf("%s[%v]", path, k), v.MapIndex(k), ops...)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			traverse(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i), ops...)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			traverse(fmt.Sprintf("(*%s)", path), v.Elem(), ops...)
		}
	default:
		for _, f := range ops {
			if !f(v) {
				return
			}
		}
		fmt.Println(path, "=", v)
	}
}
