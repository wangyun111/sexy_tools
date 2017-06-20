package tools

import (
	"math/rand"
	"time"
)

//获得区间的随机数[0,num)
func CreateMathRandInt64(num int64) (resultNum int64) {
	if num > 1 {
		source := rand.NewSource(time.Now().UnixNano()) //创建一个伪随机资源
		r := rand.New(source)                           //产生一个随机数值
		resultNum = r.Int63n(num)                       //产生一个区间随机数
	}
	return
}

//获得一个(0,1)之间的随机float64
func CreateMathRandFloat64() (resultNum float64) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	resultNum = r.Float64()
	return
}

//rand.Perm(n) 返回一个伪随机[0,n）切片
