package scs

import (
	"reflect"
)

// 判断是否是零值
func (self *scs) IsBlank(v interface{}) bool {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func (self *scs) ToSlice(v interface{}) []interface{} {
	var params []interface{}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		if value.Len() > 0 {
			for i := 0; i < value.Len(); i++ {
				params = append(params, value.Index(i).Interface())
			}
		}
	case reflect.String:
		fallthrough
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fallthrough
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Ptr:
		params = append(params, value.Interface())
	}
	return params
}
