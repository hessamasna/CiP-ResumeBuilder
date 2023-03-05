package controller

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"
	"userService/dto"
	"userService/errors"
	"userService/service"
	"userService/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var userDto dto.UserDto
	//validate user
	userDto = *validate_user(c, userDto)
	if &userDto == nil {
		return
	}

	fmt.Print("user pass is : " + userDto.PasswordHash)

	if len(userDto.PhoneNumber) == 0 {
		error := errors.New_Invalid_request_error("Phone number cannot be empty.", nil).Error
		result := dto.Create_http_response(error.Error_code, nil, error)
		c.JSON(error.Error_code, gin.H{
			"result": result,
		})
		return
	}
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userDto.PasswordHash), 10)

	if err != nil {
		error := errors.New_hashing_error(err).Error
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	userDto.PasswordHash = string(hash)

	//add user
	_, error := service.AddUser(userDto)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func Signin(c *gin.Context) {

	var userDto dto.UserDto
	userDto = *validate_user(c, userDto)
	if &userDto == nil {
		return
	}
	result, err := service.Get_user_by_email(userDto.Email)
	if err != nil {
		c.JSON(err.Error_code, gin.H{
			"result": dto.Create_http_response(err.Error_code, nil, err),
		})

		return
	}
	error := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(userDto.PasswordHash))

	if error != nil {
		c.JSON(400, gin.H{
			"result": dto.Create_http_response(
				400,
				nil,
				errors.New_Invalid_request_error("Invalid email or password", nil).Error),
		})
		return
	}
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": result.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	accessTokenString, error := access_token.SignedString([]byte(os.Getenv("SECRET")))

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(
				500,
				nil,
				errors.New_internal_error("Failed to create  access token", error).Error),
		})

		return
	}

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": result.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	refreshTokenString, error := refresh_token.SignedString([]byte(os.Getenv("SECRET")))

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(
				500,
				nil,
				errors.New_internal_error("Failed to create  access token", error).Error),
		})

		return
	}

	c.SetCookie("access_token", accessTokenString, 60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshTokenString, 12*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", 60, "/", "localhost", false, false)

	c.JSON(200, gin.H{
		"result": dto.Create_http_response(
			200,
			map[string]string{"access_token": accessTokenString, "refresh_token": refreshTokenString},
			nil),
	})
}

func RefreshAccessToken(c *gin.Context) {
	// message := "could not refresh access token"
	cookie, err := c.Cookie("refresh_token")
	if len(cookie) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(
				400,
				nil,
				errors.New_Invalid_request_error("Failed to read  refresh token", nil).Error),
		})

		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	isValid, err_message := util.ValidateToken(cookie)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err_message})
		return
	}

	claims, err := util.ExtractToken(cookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err_message})
		return
	}
	// sub := claims["sub"].(string)
	sub := int(claims["sub"].(float64))
	// subInt, err := strconv.Atoi(sub)
	
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err_message})
	// 	return
	// }

	// access_token, err := createAccessToken(subInt)
	access_token, err := createAccessToken(sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(
				500,
				nil,
				errors.New_internal_error("Failed to create  access token", err).Error),
		})

		return
	}

	c.SetCookie("access_token", access_token, 60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", 12*60, "/", "localhost", false, false)

	c.JSON(200, gin.H{
		"result": dto.Create_http_response(
			200,
			map[string]string{"access_token": access_token},
			nil),
	})

}

func Logout(c *gin.Context) {

	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	// id, err_message := ValidateToken(headers)

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func validate_user(c *gin.Context, userDto dto.UserDto) *dto.UserDto {
	//read user from request and validate json
	err := c.Bind(&userDto)
	if err != nil {
		error := errors.New_convert_error(err).Error
		result := dto.Create_http_response(error.Error_code, nil, error)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result,
		})

		return nil
	}
	//check if email is empty
	if len(userDto.Email) == 0 {
		error := errors.New_Invalid_request_error("Email cannot be empty.", nil).Error
		result := dto.Create_http_response(error.Error_code, nil, error)
		c.JSON(error.Error_code, gin.H{
			"result": result,
		})
		return nil
	}

	//Check email regex
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(userDto.Email) {
		error := errors.New_Invalid_request_error("Invalid email address", nil).Error
		result := dto.Create_http_response(error.Error_code, nil, error)
		c.JSON(error.Error_code, gin.H{
			"result": result,
		})
		return nil
	}

	//check if password is empty
	if len(userDto.PasswordHash) == 0 {
		error := errors.New_Invalid_request_error("Password cannot be empty", nil).Error
		result := dto.Create_http_response(error.Error_code, nil, error)
		c.JSON(error.Error_code, gin.H{
			"result": result,
		})
		return nil
	}

	return &userDto

}

func createAccessToken(id int) (string, error) {
	expireTime := time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expireTime,
		"sub": id,
	})
	tokenString, error := token.SignedString([]byte(os.Getenv("SECRET")))
	if error != nil {
		return "", error
	}
	return tokenString, nil
}
