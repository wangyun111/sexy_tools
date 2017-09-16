package tools

import (
	"regexp"
	"strconv"
	"sync"
	"time"
)

var orderNowTime string
var orderCount int64

func GetOrderNo() (orderNo string) {
	nowTimeStr := time.Now().Format("060102150405")
	if orderNowTime == "" && orderNowTime == nowTimeStr {
		orderCount = 1
	} else {
		orderCount++
		orderNowTime = nowTimeStr
	}
	nowTimeLong, err := strconv.ParseInt(nowTimeStr, 10, 64)
	if err == nil {
		nowTimeLong = nowTimeLong*10000 + orderCount
		orderNo = strconv.FormatInt(nowTimeLong, 10)
	} else {
		orderNo = ""
	}
	waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
	go Afunction(int(orderCount))
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	return

}

var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

var num = 0
var data = ""

func GetOrderCode() string {
	t := time.Now()
	s := t.Format("060102150405")
	if data == "" || data != s {
		data = s
		num = 0
	}
	num = num + 1
	a, _ := strconv.Atoi(data)
	orderNum := a*10000 + num
	orderCode := strconv.Itoa(orderNum)
	waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
	go Afunction(num)

	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	return orderCode
}

func ValidatePhone(mobileNum string) bool {
	Regular := "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0-9])|(17[0-9]))\\d{8}$"
	reg := regexp.MustCompile(Regular)
	return reg.MatchString(mobileNum)
}
