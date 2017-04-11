package tools

import (
	// "fmt"
	"time"
)

func CompareTime(beginTimeStr, endTimeStr string) (err error, str string) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	beginTime, err := time.ParseInLocation(timeLayout, beginTimeStr, loc)
	endTime, err := time.ParseInLocation(timeLayout, endTimeStr, loc)
	// beginTime, err := time.Parse("2006-01-02 15:04:05", beginTimeStr)
	// endTime, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
	if err == nil {
		nowTimeStamp := time.Now().Unix()
		beginTimeStamp := beginTime.Unix()
		endTimeStamp := endTime.Unix()
		if beginTimeStamp > nowTimeStamp {
			return nil, "notStart"
		} else if nowTimeStamp > endTimeStamp {
			return nil, "notFinish"
		} else {
			return nil, ""
		}
	}
	return err, "error"
}

//比较参数时间+天数跟当前时间大小
//返回true,代表小于等于当前时间
func CompareNowTime(intervalDay int64, paramsTime string) bool {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	formatTime, err := time.ParseInLocation(timeLayout, paramsTime, loc)
	if err == nil {
		nowTimeStamp := time.Now().Unix()
		formatTimeStamp := formatTime.Unix() + intervalDay*24*3600
		if formatTimeStamp > nowTimeStamp {
			return true
		}
	}
	return false
}

func GetNowTime() (createTime string) {
	createTime = time.Now().Format("2006-01-02 15:04:05")
	return
}

func GetParseTime(parseTime string) (resultTime string) {
	t, _ := time.Parse("20060102150405", parseTime)
	resultTime = t.Format("2006-01-02 15:04:05")
	return
}
