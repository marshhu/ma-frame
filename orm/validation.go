package orm

import (
	"errors"
	"reflect"
)

// Validate 校验
// slice/slicePtr/struct/structPtr
func Validate(value interface{}, vtype string, isZero bool) error {
	var msg string = "错误的类型"
	var pass bool

	kind := reflect.TypeOf(value).Kind()

	if vtype == "structPtr" && kind == reflect.Ptr {
		valueKind := reflect.ValueOf(value).Elem()
		if valueKind.Kind() == reflect.Struct {
			id := valueKind.FieldByName("ID").Interface()
			if isZero == false {
				pass = checkIDNotZero(id)
				if pass == false {
					msg = "ID不能为0"
				}
			} else {
				pass = checkIDZero(id)
				msg = "ID必须为0"
			}
		}
	}

	if vtype == "slice" && kind == reflect.Slice {
		raws := reflect.ValueOf(value)

		if raws.Len() == 0 {
			return errors.New("没有记录")
		}

		val := reflect.ValueOf(reflect.ValueOf(value).Index(0).Interface())
		id := val.FieldByName("ID").Interface()
		if isZero == false {
			pass = checkIDNotZero(id)
			if pass == false {
				msg = "ID不能为0"
			}
		} else {
			pass = checkIDZero(id)
			msg = "ID必须为0"
		}
	}

	if pass == false {
		panic(msg)
	}

	return nil
}

func checkIDNotZero(id interface{}) bool {
	var pass bool
	switch v := id.(type) {
	case int:
		if v > 0 {
			pass = true
		}
	case int8:
		if v > 0 {
			pass = true
		}
	case int16:
		if v > 0 {
			pass = true
		}
	case int32:
		if v > 0 {
			pass = true
		}
	case int64:
		if v > 0 {
			pass = true
		}
	case uint:
		if v > 0 {
			pass = true
		}
	case uint8:
		if v > 0 {
			pass = true
		}
	case uint16:
		if v > 0 {
			pass = true
		}
	case uint32:
		if v > 0 {
			pass = true
		}
	case uint64:
		if v > 0 {
			pass = true
		}
	default:
		pass = false
	}

	return pass
}

func checkIDZero(id interface{}) bool {
	var pass bool
	switch v := id.(type) {
	case int:
		if v == 0 {
			pass = true
		}
	case int8:
		if v == 0 {
			pass = true
		}
	case int16:
		if v == 0 {
			pass = true
		}
	case int32:
		if v == 0 {
			pass = true
		}
	case int64:
		if v == 0 {
			pass = true
		}
	case uint:
		if v == 0 {
			pass = true
		}
	case uint8:
		if v == 0 {
			pass = true
		}
	case uint16:
		if v == 0 {
			pass = true
		}
	case uint32:
		if v == 0 {
			pass = true
		}
	case uint64:
		if v == 0 {
			pass = true
		}
	default:
		pass = false
	}

	return pass
}
