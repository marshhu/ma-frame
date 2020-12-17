package utils

import "regexp"

const (
	mobile   string = `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$` //手机
	email    string = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`       //邮箱
	datetime string = `^([0-9]{4}[-/][0-9]{2}[-/][0-9]{2})( [0-9]{2}:[0-9]{2}(:[0-9]{2})?)?$`
	username string = `^[a-zA-Z]{1}([a-zA-Z0-9]|[._-]){3,15}$`                                                       //4到16位（字母，数字，下滑线，减号）
	password string = `^.*(?=.{6,16})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*? ]).*$`                            //密码强度正则：最少6位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符
	id       string = `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$` //身份证
)

func MatchMobile(str string) bool {
	reg := regexp.MustCompile(mobile)
	return reg.MatchString(str)
}

func MatchEmail(str string) bool {
	reg := regexp.MustCompile(email)
	return reg.MatchString(str)
}

func MatchDatetime(str string) bool {
	reg := regexp.MustCompile(datetime)
	return reg.MatchString(str)
}

func MatchUserName(str string) bool {
	reg := regexp.MustCompile(username)
	return reg.MatchString(str)
}

func MatchPassword(str string) bool {
	reg := regexp.MustCompile(password)
	return reg.MatchString(str)
}

func MatchID(str string) bool {
	reg := regexp.MustCompile(id)
	return reg.MatchString(str)
}
