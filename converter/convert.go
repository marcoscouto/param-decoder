package converter

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

const (
	dateLayout = "2006-01-02T15:04:05Z"
	base       = 10
	bitSize    = 64
)

var (
	resolvers  map[reflect.Kind]func(field reflect.Value, value string) error
	intTypes   = []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	uintTypes  = []reflect.Kind{reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}
	floatTypes = []reflect.Kind{reflect.Float32, reflect.Float64}
)

func init() {
	resolvers = make(map[reflect.Kind]func(field reflect.Value, value string) error)
	for _, v := range intTypes {
		resolvers[v] = resolveInt
	}
	for _, v := range uintTypes {
		resolvers[v] = resolveUInt
	}
	for _, v := range floatTypes {
		resolvers[v] = resolveFloat
	}
	resolvers[reflect.Bool] = resolveBool
	resolvers[reflect.Struct] = resolveStruct
	resolvers[reflect.String] = resolveString
}

func Resolve(field reflect.Value, value string) error {
	if f, ok := resolvers[field.Kind()]; ok {
		return f(field, value)
	}
	return errors.ErrUnsupported
}

var resolveString = func(field reflect.Value, value string) error {
	if field.CanSet() {
		field.SetString(value)
	}
	return nil
}

var resolveInt = func(field reflect.Value, value string) error {
	intValue, err := strconv.ParseInt(value, base, bitSize)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetInt(intValue)
	}
	return nil
}

var resolveUInt = func(field reflect.Value, value string) error {
	uintValue, err := strconv.ParseUint(value, base, bitSize)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetUint(uintValue)
	}
	return nil
}

var resolveFloat = func(field reflect.Value, value string) error {
	floatValue, err := strconv.ParseFloat(value, bitSize)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetFloat(floatValue)
	}
	return nil
}

var resolveBool = func(field reflect.Value, value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetBool(boolValue)
	}
	return nil
}

var resolveStruct = func(field reflect.Value, value string) error {
	if field.Type() == reflect.TypeOf(time.Time{}) {
		t, err := time.Parse(dateLayout, value)
		if err != nil {
			return err
		}
		if field.CanSet() {
			field.Set(reflect.ValueOf(t))
		}
	}
	return nil
}
