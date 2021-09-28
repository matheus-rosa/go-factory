package go_factory

import (
	"reflect"
)

type Factory struct {
	opts   *Options
	fields Fields
}

type Options struct {
	BaseFactory func() map[string]interface{}
	// gorm *DB
}

func NewFactory(opts *Options) *Factory {
	return &Factory{opts: opts}
}

func (f Factory) Create(model string, fields Fields) interface{} {
	f.fields = fields
	stub := f.opts.BaseFactory()[model]
	fun := reflect.ValueOf(stub)

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(f)
	res := fun.Call(in)

	return res[0].Interface()
}
