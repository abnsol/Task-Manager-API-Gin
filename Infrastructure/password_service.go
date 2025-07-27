package infrastructure

import "golang.org/x/crypto/bcrypt"

func HashPassword(userPassword string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return
}

func CheckPassword(existingUserPassword string, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(existingUserPassword), []byte(userPassword))
}
