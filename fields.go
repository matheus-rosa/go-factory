package go_factory

import "time"

type Fields map[string]interface{}

func (f *Factory) GetField(fieldName string) []Fields {
	fields := f.get(fieldName, []Fields{})

	switch fields.(type) {
	case Fields:
		return []Fields{fields.(Fields)}
	default:
		return fields.([]Fields)
	}
}

func (f *Factory) Bool(fieldName string, defaultValue bool) bool {
	return f.get(fieldName, defaultValue).(bool)
}

func (f *Factory) String(fieldName, defaultValue string) string {
	return f.get(fieldName, defaultValue).(string)
}

func (f *Factory) Int(fieldName string, defaultValue int) int {
	return f.get(fieldName, defaultValue).(int)
}

func (f *Factory) Int8(fieldName string, defaultValue int8) int8 {
	return f.get(fieldName, defaultValue).(int8)
}

func (f *Factory) Int16(fieldName string, defaultValue int16) int16 {
	return f.get(fieldName, defaultValue).(int16)
}

func (f *Factory) Int32(fieldName string, defaultValue int32) int32 {
	return f.get(fieldName, defaultValue).(int32)
}

func (f *Factory) Int64(fieldName string, defaultValue int64) int64 {
	return f.get(fieldName, defaultValue).(int64)
}

func (f *Factory) Uint(fieldName string, defaultValue uint) uint {
	return f.get(fieldName, defaultValue).(uint)
}

func (f *Factory) Uint8(fieldName string, defaultValue uint8) uint8 {
	return f.get(fieldName, defaultValue).(uint8)
}

func (f *Factory) Uint16(fieldName string, defaultValue uint16) uint16 {
	return f.get(fieldName, defaultValue).(uint16)
}

func (f *Factory) Uint32(fieldName string, defaultValue uint32) uint32 {
	return f.get(fieldName, defaultValue).(uint32)
}

func (f *Factory) Uint64(fieldName string, defaultValue uint64) uint64 {
	return f.get(fieldName, defaultValue).(uint64)
}

func (f *Factory) UintPtr(fieldName string, defaultValue uintptr) uintptr {
	return f.get(fieldName, defaultValue).(uintptr)
}

func (f *Factory) Byte(fieldName string, defaultValue byte) byte {
	return f.Uint8(fieldName, defaultValue)
}

func (f *Factory) Rune(fieldName string, defaultValue rune) rune {
	return f.Int32(fieldName, defaultValue)
}

func (f *Factory) Float32(fieldName string, defaultValue float32) float32 {
	return f.get(fieldName, defaultValue).(float32)
}

func (f *Factory) Float64(fieldName string, defaultValue float64) float64 {
	return f.get(fieldName, defaultValue).(float64)
}

func (f *Factory) Complex64(fieldName string, defaultValue complex64) complex64 {
	return f.get(fieldName, defaultValue).(complex64)
}

func (f *Factory) Complex128(fieldName string, defaultValue complex128) complex128 {
	return f.get(fieldName, defaultValue).(complex128)
}

func (f *Factory) Time(fieldName string, defaultValue time.Time) time.Time {
	return f.get(fieldName, defaultValue).(time.Time)
}

func (f *Factory) get(fieldName string, defaultValue interface{}) interface{} {
	if _, ok := f.fields[fieldName]; ok {
		return f.fields[fieldName]
	}

	return defaultValue
}
