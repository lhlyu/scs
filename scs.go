package scs

import (
	"bytes"
	"log"
	"reflect"
)

type scs struct {
	base string
	str string
	params []interface{}
}

type noneType = struct {}
var noneValue = struct {}{}


var except = map[reflect.Kind]noneType{
	reflect.Map: noneValue,
	reflect.Chan: noneValue,
	reflect.Func: noneValue,
	reflect.Struct: noneValue,
	reflect.UnsafePointer: noneValue,
}

func New(s string) *scs{
	return &scs{
		base:s,
		str:s,
		params:make([]interface{},0),
	}
}

func (self *scs) Reset(){
	self.str = self.base
	self.params = make([]interface{},0)
}

func (self *scs) GetString() string{
	return self.str
}

func (self *scs) GetParams() []interface{}{
	return self.params
}

func (self *scs) Adds(ss ...string) *scs{
	var buffer bytes.Buffer
	buffer.WriteString(self.GetString())
	for _,s := range ss {
		buffer.WriteString(s)
	}
	self.str = buffer.String()
	return self
}

func (self *scs) AddOp(column string,value interface{},f func(column string,placeholderCount int) string) *scs{
	if value == nil{
		return self
	}
	self.params = append(self.params,value)
	return self.Adds(f(column,1))
}

// 排除零值的字段
func (self *scs) AddOpNotZero(column string,value interface{},f func(column string,placeholderCount int) string) *scs{
	v := reflect.ValueOf(value)
	if _,ok := except[v.Kind()];ok{
		log.Println("second param type is not valid,it's type is ",v.Kind().String())
		return self
	}
	if ok := self.IsBlank(value);ok{
		return self
	}
	self.params = append(self.params,value)
	return self.Adds(f(column,1))
}


func (self *scs) AndBlock(f func()) *scs{
	self.Adds(" and (")
	f()
	self.Adds(") ")
	return self
}

func (self *scs) OrBlock(f func()) *scs{
	self.Adds(" or (")
	f()
	self.Adds(") ")
	return self
}

func (self *scs) ExistsBlock(f func()) *scs{
	self.Adds(" exists (")
	f()
	self.Adds(") ")
	return self
}

func (self *scs) NotExistsBlock(f func()) *scs{
	self.Adds(" not exists (")
	f()
	self.Adds(") ")
	return self
}

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