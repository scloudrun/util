//
// time.go
//
// Distributed under terms of the MIT license.
//

package util

import (
	"strconv"
	"time"
)

var (
	DefaultDateFormat     = "2006-01-02"
	DefaultDatetimeFormat = "2006-01-02 15:04:05"
)

func TimestampMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampString() string {
	return strconv.FormatInt(Timestamp(), 10)
}

func DatetimeString() string {
	return time.Now().Format(DefaultDatetimeFormat)
}

func DateFormat(t time.Time) string {
	return t.Format(DefaultDateFormat)
}

func DatetimeFormat(t time.Time) string {
	return t.Format(DefaultDatetimeFormat)
}

func DatetimeParse(t string) (time.Time, error) {
	return time.Parse(DefaultDatetimeFormat, t)
}

func DatetimeParseLocal(t string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(DefaultDatetimeFormat, t, loc)
}

func ToDate(t time.Time) time.Time {
	t, _ = time.Parse(DefaultDateFormat, DateFormat(t))
	return t
}
