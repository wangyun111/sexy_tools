package tools

import (
	"strconv"
	"strings"
)

//截取小数点后几位
func SubFloatToString(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if m == 0 {
		return newn[0]
	}
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

//截取小数点后几位
func GetSeveralFloatString(f float64, m int) float64 {
	newn := SubFloatToString(f, m)
	newf, _ := strconv.ParseFloat(newn, 64)
	return newf
}
