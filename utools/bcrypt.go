/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : bcrypt.go
@Software: GoLand
*/

package utools

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

// encrypt and decode
// 加密用户密码
func EncryptUserPassword(password string) ([]byte, error) {
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hasePassowrd, nil
}

// 解密用户密码
func DecodeUserPassword(password string, verifyPasswd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(verifyPasswd)); err != nil {
		return false, err
	}
	return true, nil
}

// db 的密码
func EncryptDbPasswd(password string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(password))
	return encoded
}

func DecodeDbPasswd(password string) string {
	decoded, _ := base64.StdEncoding.DecodeString(password)
	return string(decoded)
}
