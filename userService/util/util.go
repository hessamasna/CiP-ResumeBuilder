package util

import (
	"fmt"
	"os"
	"time"
	"userService/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExtractToken(tokenStr string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

func ExtractSubFromToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("could not parse claims")
	}
	subFloat64, ok := claims["sub"].(float64)
	if !ok {
		return 0, fmt.Errorf("could not parse sub")
	}
	sub := int(subFloat64)
	return sub, nil
}

func ValidateToken(tokenStr string) (bool, string) {
	claims, err := ExtractToken(tokenStr)
	if err != nil {
		return false, err.Error()
	}
	exp := int64(claims["exp"].(float64))
	if exp < time.Now().Unix() {
		return false, errors.New_token_expired_exception("token is expired", nil).Error.Message
	}
	return true, ""
}

func Check_if_is_login(c *gin.Context, token_name string) *errors.Base_error {
	cookie, err := c.Cookie(token_name)
	if len(cookie) == 0 {
		error_message := "Failed to read " + token_name
		return errors.New_authorization_error(error_message, nil).Error
	}

	if err != nil {
		return errors.New_authorization_error(err.Error(), err).Error
	}

	isValid, err_message := ValidateToken(cookie)
	if !isValid {
		return errors.New_authorization_error(err_message, nil).Error

	}
	return nil
}
