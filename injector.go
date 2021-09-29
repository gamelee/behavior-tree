package behavoir_tree

import (
	"fmt"
	"reflect"
)

type injector struct {
	values map[reflect.Type]reflect.Value
}

func NewInjector() *injector {
	i := new(injector)
	i.values = make(map[reflect.Type]reflect.Value)
	return i
}

func (injector *injector) Invoke(f interface{}) ([]reflect.Value, error) {
	rt := reflect.TypeOf(f)
	args := make([]reflect.Value, rt.NumIn())
	for i := 0; i < rt.NumIn(); i++ {
		val := injector.get(rt.In(i))
		if !val.IsValid() {
			return nil, fmt.Errorf("value not found for type %s", rt.In(i).String())
		}
		args[i] = val
	}
	return reflect.ValueOf(f).Call(args), nil
}

func (injector *injector) Apply(val interface{}) error {
	rv := reflect.ValueOf(val)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		if fv.CanSet() && rt.Field(i).Tag.Get("inject") == "true" {
			v := injector.get(fv.Type())
			if !v.IsValid() {
				return fmt.Errorf("value not found for type %s", rt.Field(i).Name)
			}
			fv.Set(v)
		}
	}
	return nil
}


func (injector *injector) get(rt reflect.Type) reflect.Value {
	val := injector.values[rt]
	if val.IsValid() {
		return val
	}
	if rt.Kind() == reflect.Interface {
		for k, v := range injector.values {
			if k.Implements(rt) {
				val = v
				break
			}
		}
	}
	return val
}
