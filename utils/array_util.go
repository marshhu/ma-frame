package utils

import "reflect"

func ArrayContains(array interface{}, val interface{}) (hasItem bool, index int) {
	obj := reflect.TypeOf(array)
	objVal := reflect.ValueOf(array)
	switch obj.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < objVal.Len(); i++ {
			if reflect.DeepEqual(objVal.Index(i).Interface(), val) {
				return true, i
			}
		}
	case reflect.Map:
		{
			mVal := reflect.ValueOf(val)
			mObj := objVal.MapIndex(mVal)
			if mObj.IsValid() {
				for i, v := range objVal.MapKeys() {
					if reflect.DeepEqual(v.Interface(), val) {
						return true, i
					}
				}
			}
		}
	}

	return false, -1
}
