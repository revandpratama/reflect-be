package helper

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {

	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(b), err
}

func ValidatePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}
