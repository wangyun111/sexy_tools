package crypto

import (
	"encoding/base64"
)

func Base64Encrypt(byteStr []byte) (encStr string) {
	encStr = base64.StdEncoding.EncodeToString(byteStr)
	return
}

func Base64Decrypt(encStr string) (desStr []byte, err error) {
	desStr, err = base64.StdEncoding.DecodeString(encStr)
	return
}
