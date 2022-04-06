package helper

import (
	"github.com/claravelita/majoo-test/common"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), common.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateJWTTokenLogin(users, name, username string) (token string, err error) {
	tokenClaims := jwt.MapClaims{}
	tokenClaims["user_id"] = users
	tokenClaims["name"] = name
	tokenClaims["user_name"] = username

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	token, err = claim.SignedString([]byte(os.Getenv("SECRET_AUTH")))
	if err != nil {
		return "", err
	}
	return
}
