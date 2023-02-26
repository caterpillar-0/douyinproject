package utils

import "golang.org/x/crypto/bcrypt"

//password encryption
func PasswordHash(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

//password verification
func PasswordValid(hashPassword string, rawPassword string) bool {
	byteHash := []byte(hashPassword)
	byteRaw := []byte(rawPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, byteRaw)
	return err == nil
}
