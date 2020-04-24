package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	size      = 16
	secretKey = "akfw*&TGdsfnbi*^Wt"
)

// PasswordVerify 校验密码是否有效
func PasswordVerify(dbPassword string, md5Password string) (bool, error) {

	has := md5V(md5Password + secretKey)
	if (has) == dbPassword {
		return true, nil
	}
	return false, nil
}

// // cryptoPassword 设置的密码 加密
// func cryptoPassword(password string) string  {
// 	md5 := md5V(password)
// 	return md5V(md5+secret_key)
// }

// CryptoPassword 设置的密码 加密
func CryptoPassword(password string) string {
	return md5V(password + secretKey)
}

func md5V(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
