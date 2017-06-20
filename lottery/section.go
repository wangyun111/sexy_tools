package section

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type GiftInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Url    string `json:"url"`
}

const (
	giftZero = "奖品已抽完"
)

//测试参数
// [{"id":2,"name":"10元减免券","number":3000,"url":"http://www.baidu.com"},
// {"id":2,"name":"20元减免券","number":3000,"url":"http://www.baidu.com"},
// {"id":3,"name":"30元减免券","number":2000,"url":"http://www.baidu.com"},
// {"id":4,"name":"10元话费","number":1500,"url":"http://www.baidu.com"},
// {"id":5,"name":"20元话费","number":500,"url":"http://www.baidu.com"},
//区间概率抽奖,计算数量,存数据库
func SectionLottery(jsonStr string) (has bool, msg, surplusGifts string, lotteryGift *GiftInfo) {
	has = true
	msg = "success"
	var gifts []GiftInfo
	var gift GiftInfo
	if err := json.Unmarshal([]byte(jsonStr), &gifts); err != nil {
		has = false
		msg = err.Error()
		return
	}
	var sumNum, rendNum, indexNum int
	for _, gift = range gifts {
		sumNum += gift.Number
	}
	if sumNum == 0 {
		has = false
		msg = giftZero
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rendNum = r.Intn(sumNum) + 1
	for i := 0; i < len(gifts); i++ {
		gift = gifts[i]
		nowNum := gift.Number
		if rendNum > indexNum && rendNum <= nowNum+indexNum {
			gift.Number = nowNum - 1
			gifts[i] = gift
			lotteryGift = &gift
			break
		} else {
			indexNum += nowNum
		}
	}
	giftsStr, _ := json.Marshal(&gifts)
	surplusGifts = string(giftsStr)
	return
}

//只计算概率,不计数量
func RatioLottery(jsonStr string) (has bool, msg string, lotteryGift *GiftInfo) {
	has = true
	msg = "success"
	var gifts []GiftInfo
	var gift GiftInfo
	if err := json.Unmarshal([]byte(jsonStr), &gifts); err != nil {
		has = false
		msg = err.Error()
		return
	}
	var sumNum, rendNum, indexNum int
	sumNum = 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rendNum = r.Intn(sumNum) + 1
	for i := 0; i < len(gifts); i++ {
		gift = gifts[i]
		nowNum := gift.Number
		if rendNum > indexNum && rendNum <= nowNum+indexNum {
			gifts[i] = gift
			lotteryGift = &gift
			break
		} else {
			indexNum += nowNum
		}
	}
	return
}
