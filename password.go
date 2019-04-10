package goutlis

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 计算密码hash值
func HashPassword(pwd string) (string, error) {
	pwdBytes := []byte(pwd)
	hashPwd, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

// VerifyPassword 验证密码
func VerifyPassword(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
