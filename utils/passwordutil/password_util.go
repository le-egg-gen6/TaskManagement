package passwordutil

import "golang.org/x/crypto/bcrypt"

func CompareHashedPassword(password string, storePassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(storePassword)) == nil
}

func GenerateEncryptedPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), err
}
