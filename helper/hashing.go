package helper

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	password := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	return err == nil
}
