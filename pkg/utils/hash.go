package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func IsCorrectPassword(PWD, hashPWD string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPWD), []byte(PWD))
	return err == nil
}
