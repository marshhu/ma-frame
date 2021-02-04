package utils

import (
	"math"
	"time"
)

const (
	AUTO_TIME_ZONE  = true
	DATE_LAYOUT     = "2006-01-02"
	TIME_LAYOUT     = "15:04:05"
	DATETIME_LAYOUT = "2006-01-02 15:04:05"
)

func ToTime(val string, format string) time.Time {
	return timeParse(val, format)
}

func TimeToString(val time.Time, format string) string {
	return timeFormat(val, format)
}

//获取传入的时间所在月份的第一天，即某月第一天的0点
func GetMothBegin(val time.Time) time.Time {
	val = val.AddDate(0, 0, -val.Day()+1)
	return GetZeroTime(val)
}

//获取某一天的0点时间
func GetZeroTime(val time.Time) time.Time {
	return time.Date(val.Year(), val.Month(), val.Day(), 0, 0, 0, 0, val.Location())
}

func timeParse(val string, layout string) time.Time {
	if s, err := time.Parse(layout, val); err == nil {
		if AUTO_TIME_ZONE {
			var duration time.Duration
			_, offset := time.Now().Zone()
			if offset > 0 {
				tz := math.Ceil(float64(offset / 3600))
				if tz > 0 {
					duration = time.Duration(-tz)
				} else {
					duration = time.Duration(tz)
				}
			}
			return s.Add(time.Hour * duration).Local()
		}

		return s
	} else {
		panic(err)
	}
}

func timeFormat(val time.Time, layout string) string {
	if val.IsZero() {
		return ""
	}
	return val.Format(layout)
}
