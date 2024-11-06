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
		attribute := val.Type().Field(i).Tag.Get("attr")
		if attribute == "" {
			continue
		}
		attribute = " " + attribute
		if value.Kind() == reflect.Bool {
			if value.Bool() {
				attr += fmt.Sprintf("%s", attribute)
			}
		} else if value.Kind() == reflect.String {
			if value.String() != "" {
				attr += fmt.Sprintf("%s=\"%s\"", attribute, value.String())
			}
		} else if value.Kind() == reflect.Int {
			attr += fmt.Sprintf("%s=\"%d\"", attribute, value.Int())
		}
	}
	return attr
}
