package parse

import (
	"fmt"
	"reflect"
)

func GetAttr(data interface{}) string {
	var attr string
	val := reflect.ValueOf(data)
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		rawTag := val.Type().Field(i).Tag.Get("attr")
		tag := " " + rawTag
		if value.Kind() == reflect.Bool {
			if value.Bool() {
				attr += fmt.Sprintf("%s", tag)
			}
		} else if value.Kind() == reflect.String {
			if value.String() != "" {
				attr += fmt.Sprintf("%s=\"%s\"", tag, value.String())
			}
		} else if value.Kind() == reflect.Int {
			attr += fmt.Sprintf("%s=\"%d\"", tag, value.Int())
		}
	}
	return attr
}
