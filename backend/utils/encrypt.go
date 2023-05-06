package utils

import "golang.org/x/crypto/bcrypt"

// EncryptPassword 用 bcrypt 算法对密码进行加密
func EncryptPassword(password string) (passwordDigest string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil { // 密码加密错误
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 验证密码, 成功则返回true, 失败返回false
func CheckPassword(password string, passwordDigest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(password))
	return err == nil
}
