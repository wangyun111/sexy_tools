package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

//@Summary        hmac+sha1+base64 加密
//@Param key      加密key
//@Param data     加密数据
func HmacSha1Base64(key, data string) (sign string) {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	signHmac := fmt.Sprintf("%x\n", mac.Sum(nil))
	sign = base64.StdEncoding.EncodeToString([]byte(signHmac))
	return
}
