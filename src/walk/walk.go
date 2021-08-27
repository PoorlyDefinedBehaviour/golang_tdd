package walk

import "reflect"

func fields(x interface{}) []reflect.Value {
	value := reflect.ValueOf(x)

	out := make([]reflect.Value, 0, value.NumField())

	for i := 0; i < value.NumField(); i++ {
		out = append(out, value.Field(i))
	}

	return out
}

func Struct(x interface{}, f func(string)) {
	for _, field := range fields(x) {
		f(field.String())
	}
}
