package time

import (
	"time"
)

//1. 后一个参数在前一个参数 after,befer,equal。如果是一个参数,前面参数为当前时间
//2. 注意传入时间格式跟layout格式一样
const (
	Layout = "2006-01-02 15:04:05"
)

//比较传入时间跟当前时间
func CompareNowTime(timeStr string) (has bool, msg string) {
	msg = "failed"
	t1, err := time.Parse(Layout, timeStr)
	if err != nil {
		msg = err.Error()
		return
	}
	has = true
	nowTime := time.Now()
	if nowTime.After(t1) {
		msg = "after"
	} else if nowTime.Before(t1) {
		msg = "before"
	} else {
		msg = "equal"
	}
	return
}

//比较传入时间大小
func CompareBothTime(timeOne, timeTwo string) (has bool, msg string) {
	msg = "failed"
	t1, err := time.Parse(Layout, timeOne)
	if err != nil {
		msg = err.Error()
		return
	}
	t2, err := time.Parse(Layout, timeTwo)
	if err != nil {
		msg = err.Error()
		return
	}
	has = true
	if t1.After(t2) {
		msg = "after"
	} else if t1.Before(t2) {
		msg = "before"
	} else {
		msg = "equal"
	}
	return
}

//比较前后参数是否在中间值之间
func CompareAmongTime(timeOne, timeTwo, timeThree string) (has bool, msg string) {
	msg = "success"
	t1, err := time.Parse(Layout, timeOne)
	if err != nil {
		msg = err.Error()
		return
	}
	t2, err := time.Parse(Layout, timeTwo)
	if err != nil {
		msg = err.Error()
		return
	}
	t3, err := time.Parse(Layout, timeThree)
	has = true
	if err != nil {
		msg = err.Error()
		return
	}
	if t1.After(t2) {
		msg = "noStart"
		return
	}
	if t2.After(t3) {
		msg = "inFinish"
		return
	}
	return
}

//融合前三方法使用,参数为不定值
func CompareTime(timeTime ...time.Time) (has bool, msg string) {
	counter := len(timeTime)
	msg = "failed"
	nowTime := time.Now()
	switch counter {
	case 0:
		msg = "zero"
	case 1:
		t1 := timeTime[0]
		has = true
		if nowTime.After(t1) {
			msg = "after"
		} else if nowTime.Before(t1) {
			msg = "before"
		} else {
			msg = "equal"
		}
	case 2:
		t1 := timeTime[0]
		t2 := timeTime[1]
		has = true
		if t1.After(t2) {
			msg = "after"
		} else if t1.Before(t2) {
			msg = "before"
		} else {
			msg = "equal"
		}
	case 3:
		t1 := timeTime[0]
		t2 := timeTime[1]
		t3 := timeTime[2]
		has = true
		if t1.After(t2) {
			msg = "noStart"
		} else if t2.After(t3) {
			msg = "inFinish"
		} else {
			msg = "among"
		}
	default:
		msg = "too"
	}
	return
}

//获取当前时间,可自定义样式
func GetNowTime(l ...string) (nowTimeStr string) {
	nowTime := time.Now()
	layout := Layout
	if len(l) == 1 {
		layout = l[0]
	}
	nowTimeStr = nowTime.Format(layout)
	return
}

//解析为时间
func GetParseTime(layout, parseTime string) (has bool, msg string, resultTime time.Time) {
	has = true
	msg = "success"
	if layout == "" {
		layout = Layout
	}
	resultTime, err := time.Parse(layout, parseTime)
	if err != nil {
		has = false
		msg = err.Error()
	}
	return
}

//获取传入时间添加years,months,days后的时间,如果传入时间为空,则获取当前时间添加后的时间
//如果加时分秒,用Add()即可
func GetAddDayTime(layout string, years, months, days int, timeStr ...string) (has bool, msg string, resultTime time.Time) {
	has = true
	msg = "success"
	if layout == "" {
		layout = Layout
	}
	if len(timeStr) == 1 {
		t, err := time.Parse(layout, timeStr[0])
		if err != nil {
			has = false
			msg = err.Error()
		} else {
			resultTime = t.AddDate(years, months, days)
		}
		return
	}
	resultTime = time.Now().AddDate(years, months, days)
	return
}

//自定义格式化时间参数
func GetFormatTime(layout string, paramsTime time.Time) (formatTime string) {
	if layout == "" {
		layout = Layout
	}
	formatTime = paramsTime.Format(layout)
	return
}

//获取今天剩余秒数
func GetTodayLastSecond() time.Duration {
	today := GetNowTime("2006-01-02") + " 23:59:59"
	end, _ := time.ParseInLocation(Layout, today, time.Local)
	return time.Duration(end.Unix()-time.Now().Local().Unix()) * time.Second
}
