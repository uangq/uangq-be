package user

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"uangq-be/common/hash"
	"uangq-be/common/logger"
	"uangq-be/common/token"
)

func Login(username, password string) (bool, User, string, string) {
	isValid, msg, u := isLoginValid(username, password)
	if !isValid {
		return false, User{}, "", msg
	}

	token := token.GenerateJWT(u.Username)

	return true, u, token, ""
}

func isLoginValid(username, pass string) (bool, string, User) {
	u, err := GetUserByUsername(username)

	if err != nil && err != sql.ErrNoRows {
		logger.Log("Failed to get user." + err.Error())
		return false, "Failed to get user", User{}
	} else if err == sql.ErrNoRows {
		return false, "Can't find user", User{}
	}

	if hash.ComparePasswordBcrypt(u.PasswordHash, pass) {
		return true, "", u
	}

	return false, "Wrong password", User{}
}

func Signup(username, password string) (bool, User, string, string) {
	valid, msg := isSignupValid(username)
	if !valid {
		return false, User{}, "", msg
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logger.Log("Failed to hash password." + err.Error())
		return false, User{}, "", "Failed to hash password"
	}

	u, err := CreateUser(username, string(hashedPassword))
	if err != nil {
		logger.Log("Failed to create user." + err.Error())
		return false, User{}, "", "Failed to create user"
	}
	u.Username = username

	return true, u, token.GenerateJWT(u.Username), ""
}

func isSignupValid(username string) (bool, string) {
	_, err := GetUserByUsername(username)

	if err != nil && err != sql.ErrNoRows {
		logger.Log("Failed to get user." + err.Error())
		return false, "Failed to check"
	}

	if err == sql.ErrNoRows {
		return true, ""
	}

	return false, "User already exists"
}
