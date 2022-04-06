package dto

import "github.com/golang-jwt/jwt"

type (
	authPackage struct {
		JwtKey string
	}

	TokenClaims struct {
		UserID   string `json:"user_id"`
		Name     string `json:"name"`
		UserName string `json:"user_name"`
		jwt.StandardClaims
	}
)
