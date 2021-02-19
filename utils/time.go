package utils

import (
	"time"
)

var btime string

func Nstime() time.Time {
	k := time.Now()
	if btime == "" {
		return k
	}
	d, err := time.ParseDuration(btime)
	if err != nil {
		return k
	}
	return k.Add(d)
}

func NstimeUnix() int64 {
	return Nstime().Unix()
}

func NstimeUnixNano() int64 {
	return Nstime().UnixNano()
}

func NstimeString() string {
	return Nstime().Format("2006-01-02 15:04:05")
}

//NstimeHour00
func NstimeHour00() string {
	return Nstime().Format("2006-01-02 15")
}

func NstimeString_() string {
	return Nstime().Format("20060102150405")
}

//TsTimeYmd
func TsTimeYmd(t string) string {
	k := time.Now()
	if t != "" {
		d, _ := time.ParseDuration(t)
		k = time.Now().Add(d)
	}
	return k.Format("2006-01-02")
}

//TsTimeDate
func TsTimeDate(t string) string {
	k := time.Now()
	if t != "" {
		d, _ := time.ParseDuration(t)
		k = time.Now().Add(d)
	}
	return k.Format("20060102")
}

func Ystime() time.Time {
	btime := "-24h"
	k := time.Now()
	d, err := time.ParseDuration(btime)
	if err != nil {
		return k
	}
	return k.Add(d)
}

func Ytimeymd() string {
	return Ystime().Format("2006-01-02")
}

//GetNsTimeHour返回当前小时
func GetNsTimeHour() int {
	return time.Now().Hour()
}

func Ytimeday() string {
	return Ystime().Format("20060102")
}

func UnixToTime(u int64) string {
	tm := time.Unix(u, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//UnixToFormat
func UnixToFormat(u int64, s string) string {
	tm := time.Unix(u, 0)
	return tm.Format(s)
}

//TimeStrToUnix 字符串转时间戳
func TimeStrToUnix(stime string) int64 {
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", stime, time.Local)
	return stamp.Unix()
}

//TimeSub
func TimeSub(stime, etime string) int64 {
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", stime, time.Local)
	etamp, _ := time.ParseInLocation("2006-01-02 15:04:05", etime, time.Local)
	return etamp.Unix() - stamp.Unix()
}

//GetUnix 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

//AddTimeHours 时间加减
func AddTimeHours(ddate, dhour string) string {
	atime, _ := time.Parse("2006-01-02 15:04:05", ddate)
	d, _ := time.ParseDuration(dhour)
	return atime.Add(d).Format("2006-01-02 15:04:05")
}
