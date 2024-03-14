package decoders

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	dateLayout     = "2006-01-02T15:04:05Z"
	base           = 10
	bitSize        = 64
	sliceSplitChar = ","
)

var (
	decoders   map[reflect.Kind]decodeField
	intTypes   = []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	uintTypes  = []reflect.Kind{reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}
	floatTypes = []reflect.Kind{reflect.Float32, reflect.Float64}
)

type decodeField func(field reflect.Value, value string) error

func init() {
	decoders = make(map[reflect.Kind]decodeField)
	for _, v := range intTypes {
		decoders[v] = decodeInt
	}
	for _, v := range uintTypes {
		decoders[v] = decodeUInt
	}
	for _, v := range floatTypes {
		decoders[v] = decodeFloat
	}
	decoders[reflect.Bool] = decodeBool
	decoders[reflect.Struct] = decodeStruct
	decoders[reflect.String] = decodeString
	decoders[reflect.Slice] = decodeSlice
}

func DecodeField(field reflect.Value, value string) error {
	if f, ok := decoders[field.Kind()]; ok {
		return f(field, value)
	}
	return errors.ErrUnsupported
}

func decodeString(field reflect.Value, value string) error {
	if field.CanSet() {
		field.SetString(value)
	}
	return nil
}

func decodeInt(field reflect.Value, value string) error {
	intValue, err := strconv.ParseInt(value, base, bitSize)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetInt(intValue)
	}
	return nil
}

func decodeUInt(field reflect.Value, value string) error {
	uintValue, err := strconv.ParseUint(value, base, bitSize)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetUint(uintValue)
	}
	return nil
}

func decodeFloat(field reflect.Value, value string) error {
	floatValue, err := strconv.ParseFloat(value, bitSize)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetFloat(floatValue)
	}
	return nil
}

func decodeBool(field reflect.Value, value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	if field.CanSet() {
		field.SetBool(boolValue)
	}
	return nil
}

func decodeStruct(field reflect.Value, value string) error {
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

func decodeSlice(field reflect.Value, value string) error {
	values := strings.Split(value, sliceSplitChar)
	if len(values) != 0 {
		t := field.Type().Elem()
		s := reflect.MakeSlice(reflect.SliceOf(t), 0, len(values))
		for _, v := range values {
			f := reflect.New(t).Elem()
			if err := DecodeField(f, v); err != nil {
				return err
			}
			s = reflect.Append(s, f)
		}
		field.Set(s)
	}
	return nil
}
