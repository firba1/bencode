package bencode

import (
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) []byte {
	value := reflect.ValueOf(v)
	return convertValue(value)
}

func convertValue(value reflect.Value) []byte {
	switch value.Type().Kind() {
	case reflect.Int:
		return convertInt(value.Interface().(int))
	case reflect.String:
		return convertString(value.Interface().(string))
	case reflect.Slice:
		return convertSlice(value)
	case reflect.Struct:
		return convertDict(value)
	}
	return []byte{}
}

func convertInt(i int) []byte {
	return []byte("i" + strconv.Itoa(i) + "e")
}

func convertString(s string) []byte {
	return []byte(fmt.Sprintf("%v:%v", len([]byte(s)), s))
}

func convertSlice(value reflect.Value) (representation []byte) {
	representation = append(representation, 'l')
	for i := 0; i < value.Len(); i++ {
		valueRepresentation := convertValue(value.Index(i))
		representation = append(representation, valueRepresentation...)
	}
	representation = append(representation, 'e')
	return
}

func convertDict(value reflect.Value) (representation []byte) {
	representation = append(representation, 'd')
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if field.PkgPath != "" {
			continue
		}
		key := convertString(field.Name)
		representation = append(representation, key...)
		value := convertValue(value.Field(i))
		representation = append(representation, value...)
	}
	representation = append(representation, 'e')
	return
}
