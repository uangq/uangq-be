package hash

import (
	"golang.org/x/crypto/bcrypt"
	"uangq-be/common/logger"
)

func ComparePasswordBcrypt(passwordHash []byte, pass string) bool {
	passByte := []byte(pass)
	err := bcrypt.CompareHashAndPassword(passwordHash, passByte)
	if err != nil {
		logger.Log("Bcrypt return error. " + err.Error())
		return false
	}
	return true
}
