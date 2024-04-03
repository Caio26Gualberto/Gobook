package security

import "golang.org/x/crypto/bcrypt"

func Hash(passowrd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)
}

func CheckPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
