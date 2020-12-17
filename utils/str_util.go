package utils

import (
	"strconv"
	"strings"
)

func ToInt(val string) (int, error) {
	return strconv.Atoi(val)
}

func ToInt64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

func ToBool(val string) (bool, error) {
	return strconv.ParseBool(val)
}

func ToUInt(val string) (uint, error) {
	temp, err := strconv.Atoi(val)
	if err != nil {
		return uint(0), err
	}
	return uint(temp), nil
}

func ToFloat(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func UIntToString(val uint) string {
	return strconv.FormatUint(uint64(val), 10)
}

func Float64ToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func Split(val string, ch rune) []string {
	values := strings.FieldsFunc(val, func(r rune) bool {
		return r == ch
	})

	return values
}

func SplitToFloat(val string, ch rune) []float64 {
	var ret []float64
	values := Split(val, ch)
	for _, val := range values {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

func SplitToInt(val string, ch rune) []int64 {
	var ret []int64
	values := Split(val, ch)
	for _, val := range values {
		if v, err := strconv.ParseInt(val, 10, 64); err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

func Trim(val string) string {
	return strings.Trim(val, " ")
}
