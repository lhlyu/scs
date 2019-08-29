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
	slices := self.ToSlice(value)
	self.params = append(self.params,slices...)
	return self.Adds(f(column,len(slices)))
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
	slices := self.ToSlice(value)
	self.params = append(self.params,slices...)
	return self.Adds(f(column,len(slices)))
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

func (self *scs) InBlock(prefix string,f func()) *scs{
	self.Adds(ONE_SPACE,prefix," in (")
	f()
	self.Adds(") ")
	return self
}

func (self *scs) NotInBlock(prefix string,f func()) *scs{
	self.Adds(ONE_SPACE,prefix," not in (")
	f()
	self.Adds(") ")
	return self
}

