package tools

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/koron/go-dproxy"
	"strconv"
	"strings"
)

// 解析json字符串
func JsonParse(str string) interface{} {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	var v interface{}
	err := json.Unmarshal([]byte(str), &v)
	if err == nil {
		dpr := dproxy.New(v)
		return dpr
	}
	return err
}

//截取小数点后两位
func GetTwoFloat(f float64) float64 {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	i := strings.Index(s, ".")
	if i == -1 {
		return f
	}
	d := s[i:len(s)]
	beego.Info(d, len(d))
	if len(d) < 3 {
		s = s[0 : i+2]
	} else {
		s = s[0 : i+3]
	}
	a, _ := strconv.ParseFloat(s, 64)
	return a
}
