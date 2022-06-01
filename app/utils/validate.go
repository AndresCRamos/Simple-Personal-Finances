package utils

import (
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"strings"
	"sync"

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
	id            int
	name          string
	value         interface{}
	dbTag         []string
	validationTag []string
	nullable      bool
}

func getfieldData(id int, v reflect.Value, typeOfS reflect.StructField) fieldData {
	FieldDBTag := strings.Split(typeOfS.Tag.Get("gorm"), ";")
	nullable := !isItNullable(FieldDBTag)
	return fieldData{
		id:            id,
		name:          typeOfS.Name,
		value:         v.Interface(),
		dbTag:         FieldDBTag,
		validationTag: strings.Split(typeOfS.Tag.Get("validation"), ";"),
		nullable:      nullable,
	}
}

func validateNullable(fieldData fieldData, ChanNullErr chan<- FieldError, wg *sync.WaitGroup) {
	defer wg.Done()
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
			ID:      fieldData.id,
			Field:   fieldData.name,
			Message: fmt.Sprintf("%s can´t be null", strings.ToLower(fieldData.name)),
		}
		ChanNullErr <- errorField
	}
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

func validateType(fieldData fieldData, ChanTypeErr chan<- FieldError, wg *sync.WaitGroup) {
	defer wg.Done()
	if validationType := getValidationType(fieldData.validationTag); validationType != "none" {
		if validationType == "email" {
			email, isString := getStringValue(fieldData.value)
			if !isString {
				ChanTypeErr <- FieldError{
					ID:      fieldData.id,
					Field:   fieldData.name,
					Message: fmt.Sprintf("%s must be a string", strings.ToLower(fieldData.name)),
				}
			} else if _, err := mail.ParseAddress(email); err != nil {
				ChanTypeErr <- FieldError{
					ID:      fieldData.id,
					Field:   fieldData.name,
					Message: fmt.Sprintf("%s isn´t a valid email", strings.ToLower(fieldData.name)),
				}
			}
		}
	}
}

func validateField(validateFieldData fieldData, ChanNullErr chan<- FieldError, ChanTypeErr chan<- FieldError, wg *sync.WaitGroup) {
	go validateNullable(validateFieldData, ChanNullErr, wg)
	go validateType(validateFieldData, ChanTypeErr, wg)
}

func Validate(w http.ResponseWriter, source string, in interface{}) bool {
	v := reflect.ValueOf(in)
	typeOfS := v.Type()
	ChanNullErr := make(chan FieldError, v.NumField())
	ChanTypeErr := make(chan FieldError, v.NumField())
	wg := new(sync.WaitGroup)
	wg.Add(v.NumField() * 2)
	for i := 0; i < v.NumField(); i++ {
		fieldData := getfieldData(i, v.Field(i), typeOfS.Field(i))
		go validateField(fieldData, ChanNullErr, ChanTypeErr, wg)
	}
	wg.Wait()
	if len(ChanNullErr) > 0 {
		len := len(ChanNullErr)
		errNullList := make([]FieldError, len)
		for i := 0; i < len; i++ {
			err := <-ChanNullErr
			errNullList[i] = err
		}
		defer DisplayFieldErrors(w, source, errNullList)
		return false
	}
	if len(ChanTypeErr) != 0 {
		errTypeList := make([]FieldError, len(ChanTypeErr))
		DisplayFieldErrors(w, source, errTypeList)
		return false
	}
	return true
}
