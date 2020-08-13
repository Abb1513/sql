/*
@Time    : 2020/8/4
@Author  : Wangcq
@File    : bcrypt_test.go
@Software: GoLand
*/

package utools

import "testing"

func TestBcrypt(t *testing.T) {
	var psw string
	psw = "root"
	res, err := EncryptUserPassword(psw)
	if err != nil {
		t.Log("err, ", err)
	}
	t.Log(string(res))
	var e = "$2a$10$uV3noykM1.aNkh14r3T39.4E62TKebEQ9kYJxYZxiJdmfc0CSCPY2"
	r, err := DecodeUserPassword(e, psw)
	if err != nil {
		t.Log(err)
	}
	t.Log(r)
}
