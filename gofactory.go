package go_factory

import (
	"gorm.io/gorm"
	"reflect"
)

type Factory struct {
	opts   *Options
	fields Fields
}

type Options struct {
	BaseFactory func() map[string]interface{}
	Gorm        *gorm.DB
}

func NewFactory(opts *Options) *Factory {
	return &Factory{opts: opts}
}

func (f Factory) Create(model string, fields ...Fields) interface{} {
	if len(fields) > 0 {
		f.fields = fields[0]
	} else {
		f.fields = Fields{}
	}

	stub := f.opts.BaseFactory()[model]
	fun := reflect.ValueOf(stub)

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(f)
	res := fun.Call(in)

	return res[0].Interface()
}

func (f Factory) CreateN(model string, output interface{}, fields ...Fields) {
	v := reflect.ValueOf(output)
	n := v.Elem().Len()
	fieldsLen := len(fields)

	for i := 0; i < n; i++ {
		var modelCreated interface{}
		if i < fieldsLen {
			modelCreated = f.Create(model, fields[i])
		} else {
			modelCreated = f.Create(model)
		}

		modelCreatedRfl := reflect.ValueOf(modelCreated)
		if v.Elem().Index(i).Kind() == reflect.Ptr {
			v.Elem().Index(i).Set(modelCreatedRfl)
			continue
		}

		currentStruct := v.Elem().Index(i)
		for j := 0; j < currentStruct.NumField(); j++ {
			currentStruct.Field(j).Set(modelCreatedRfl.Elem().Field(j))
		}
	}
}

func (f Factory) Insert(model string, fields ...Fields) interface{} {
	m := f.Create(model, fields...)

	return f.opts.Gorm.Create(m).Error
}

func (f Factory) InsertN(model string, output interface{}, fields ...Fields) interface{} {
	f.CreateN(model, output, fields...)
	return f.opts.Gorm.Create(output).Error
}
