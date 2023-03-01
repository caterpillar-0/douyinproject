package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// password encryption
func PasswordHash(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

// password verification
func PasswordValid(hashPassword string, rawPassword string) bool {
	byteHash := []byte(hashPassword)
	byteRaw := []byte(rawPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, byteRaw)
	return err == nil
}

var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

// 获取token，使用NewWithClaims
func GenToken(userid uint) (string, error) {
	claims := MyCustomClaims{
		UserID: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//toeknstring是签名
	tokenstring, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

// 从token中解析
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// 验证token是否过期/有效
func ValidToken(tokenString string) (*MyCustomClaims, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return nil, errors.New("expired token")
	}
	return claims, nil
}

func String2uint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}
