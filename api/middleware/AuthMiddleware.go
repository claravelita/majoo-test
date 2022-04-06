package middleware

import (
	"errors"
	"github.com/claravelita/majoo-test/common/helper"
	"github.com/labstack/echo/v4"
	"os"
)

func CheckJWTAndCheckAuthCustomer(token string, c echo.Context) (valid bool, err error) {
	valid = true
	authLibrary, errInit := helper.NewAuthHelper(os.Getenv("SECRET_AUTH"))
	if errInit != nil {
		return false, errInit
	}

	tokenClaims, errParse := authLibrary.ParseJWT(token)
	if errParse != nil {
		return false, errParse
	}
	if tokenClaims.UserID == "" {
		return false, errors.New("invalid user")
	}

	c.Set("user_id", tokenClaims.UserID)
	c.Set("user_name", tokenClaims.UserName)
	c.Set("name", tokenClaims.Name)
	return
}
