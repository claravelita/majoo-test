package helper

import (
	"github.com/claravelita/majoo-test/common"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type AuthHelper interface {
	ParseJWT(accessToken string) (claims dto.TokenClaims, err error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	CreateJWTTokenLogin(users, name, username string) (token string, err error)
}

func NewAuthHelper(jwtKey string) (AuthHelper, error) {
	return &authHelper{JwtKey: jwtKey}, nil
}

func (a authHelper) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), common.DefaultCost)
	return string(bytes), err
}

func (a authHelper) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a authHelper) CreateJWTTokenLogin(users, name, username string) (token string, err error) {
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

func (a authHelper) ParseJWT(accessToken string) (claims dto.TokenClaims, err error) {
	_, err = jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_AUTH")), nil
	})
	return
}
