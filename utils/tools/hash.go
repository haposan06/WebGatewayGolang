package tools

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(plain string) (string, error){
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	if err != nil {
		return "", nil
	}
	return string(result), err
}
