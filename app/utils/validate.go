package utils

import (
	"fmt"
	"net/http"
	"net/mail"
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

func getValidationType(validationTagList []string) string {
	set := make(map[string]string)
	for _, v := range validationTagList {
		types := strings.Split(v, ":")
		if len(types) > 1 {
			set[types[0]] = types[1]
		}
	}
	if _, ok := set["type"]; ok {
		return set["type"]
	} else {
		return "none"
	}

}

type fieldData struct {
	name          string
	value         interface{}
	dbTag         []string
	validationTag []string
	nullable      bool
}

func getfieldData(v reflect.Value, typeOfS reflect.StructField) fieldData {
	FieldDBTag := strings.Split(typeOfS.Tag.Get("gorm"), ";")
	nullable := !isItNullable(FieldDBTag)
	return fieldData{
		name:          typeOfS.Name,
		value:         v.Interface(),
		dbTag:         FieldDBTag,
		validationTag: strings.Split(typeOfS.Tag.Get("validation"), ";"),
		nullable:      nullable,
	}
}

func validateNullable(fieldData fieldData) (FieldError, bool) {
	isValid := true
	errorField := FieldError{}
	if val, ok := fieldData.value.(null.String); ok {
		if (!val.Valid || val.String == "") && !fieldData.nullable {
			isValid = false
		}
	}
	if val, ok := fieldData.value.(null.Float64); ok {
		if (!val.Valid) && !fieldData.nullable {
			isValid = false
		}
	}
	if val, ok := fieldData.value.(null.Int64); ok {
		if (!val.Valid) && !fieldData.nullable {
			isValid = false
		}
	}
	if val, ok := fieldData.value.(null.Time); ok {
		if (!val.Valid) && !fieldData.nullable {
			isValid = false
		}
	}
	if !isValid {
		errorField = FieldError{
			Field:   fieldData.name,
			Message: fmt.Sprintf("%s can´t be null", strings.ToLower(fieldData.name)),
		}
	}
	return errorField, isValid
}

func getStringValue(i interface{}) (string, bool) {
	if val, ok := i.(null.String); ok {
		return val.String, ok
	}
	if val, ok := i.(string); ok {
		return val, ok
	}
	return "", false

}

func validateType(fieldData fieldData) (FieldError, bool) {
	if validationType := getValidationType(fieldData.validationTag); validationType != "none" {
		if validationType == "email" {
			email, isString := getStringValue(fieldData.value)
			if !isString {
				return FieldError{
					Field:   fieldData.name,
					Message: fmt.Sprintf("%s must be a string", strings.ToLower(fieldData.name)),
				}, false
			} else if _, err := mail.ParseAddress(email); err != nil {
				return FieldError{
					Field:   fieldData.name,
					Message: fmt.Sprintf("%s isn´t a valid email", strings.ToLower(fieldData.name)),
				}, false
			}
			return FieldError{}, true
		}
	}
	return FieldError{}, true
}

func Validate(w http.ResponseWriter, source string, i interface{}) bool {
	v := reflect.ValueOf(i)
	isNullValid := true
	isTypeValid := true
	typeOfS := v.Type()
	var errorNullList = []FieldError{}
	var errorTypeList = []FieldError{}

	for i := 0; i < v.NumField(); i++ {
		fieldData := getfieldData(v.Field(i), typeOfS.Field(i))
		if err, ok := validateNullable(fieldData); !ok {
			errorNullList = append(errorNullList, err)
			isNullValid = false
		}
		if err, ok := validateType(fieldData); !ok {
			errorTypeList = append(errorTypeList, err)
			isTypeValid = false
		}
	}
	if !isNullValid {
		DisplayFieldErrors(w, source, errorNullList)
		return false
	}
	if !isTypeValid {
		DisplayFieldErrors(w, source, errorTypeList)
		return false
	}
	return true
}
