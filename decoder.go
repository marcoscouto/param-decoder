package decoder

import (
	"net/url"
	"reflect"

	"github.com/marcoscouto/param-decoder/converter"
)

const queryTag = "query"

type Decoder[T any] interface {
	DecodeQueryParams(values url.Values) T
	DecodeQueryParamsWithCustomTag(values url.Values, customTag string) T
}

type decoder[T any] struct{}

func New[T any]() Decoder[T] {
	return &decoder[T]{}
}

func (d decoder[T]) DecodeQueryParams(values url.Values) T {
	return d.DecodeQueryParamsWithCustomTag(values, queryTag)
}

func (d decoder[T]) DecodeQueryParamsWithCustomTag(values url.Values, customTag string) T {
	var t T
	targetValue := reflect.ValueOf(&t).Elem()
	for i := 0; i < targetValue.NumField(); i++ {
		field := targetValue.Field(i)
		param := targetValue.Type().Field(i).Tag.Get(customTag)
		value := values.Get(param)
		if len(value) != 0 {
			converter.Resolve(field, value)
		}
	}
	return t
}
