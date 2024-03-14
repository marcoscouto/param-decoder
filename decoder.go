package decoder

import (
	"net/url"
	"reflect"

	"github.com/marcoscouto/param-decoder/decoders"
)

const queryTag = "query"

func DecodeQueryParams[T any](values url.Values) T {
	return DecodeQueryParamsWithCustomTag[T](values, queryTag)
}

func DecodeQueryParamsWithCustomTag[T any](values url.Values, customTag string) T {
	var t T
	targetValue := reflect.ValueOf(&t).Elem()
	for i := 0; i < targetValue.NumField(); i++ {
		field := targetValue.Field(i)
		param := targetValue.Type().Field(i).Tag.Get(customTag)
		value := values.Get(param)
		if len(value) != 0 {
			decoders.DecodeField(field, value)
		}
	}
	return t
}
