package utils

import (
	"reflect"
	"strings"

	"github.com/emvi/null"
)

func isInMapKeys(set map[string]struct{}, search string) bool {
	_, ok := set[search]
	return ok
}

func isItNullable(list []string) bool {
	set := make(map[string]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	return (isInMapKeys(set, "notnull") || isInMapKeys(set, " notnull") || isInMapKeys(set, "not null"))
}

func Validate(i interface{}) ([]FieldError, bool) {
	v := reflect.ValueOf(i)
	isValid := true
	typeOfS := v.Type()
	var errorList = []FieldError{}

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	for i := 0; i < v.NumField(); i++ {
		FieldName := typeOfS.Field(i).Name
		FieldValue := v.Field(i).Interface()
		FieldTag := strings.Split(typeOfS.Field(i).Tag.Get("gorm"), ";")
		nullable := !isItNullable(FieldTag)

		if val, ok := FieldValue.(null.String); ok {
			if !val.Valid && !nullable {
				errorList = append(errorList, FieldError{Field: FieldName})
				isValid = false
			}
		}
	}
	return errorList, isValid
}
