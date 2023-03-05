package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"userService/dto"
	"userService/errors"
	"userService/service"

	"github.com/gin-gonic/gin"
)

func CreateCv(c *gin.Context) {
	var cvDto dto.CVDto
	cvDto = *validate_cv(c, cvDto)

	_, error := service.AddCv(cvDto)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func GetCvsByUserId(c *gin.Context) {
	// user_id := c.Param("user_id")
	idString := c.Param("user_id")
	user_id, err := strconv.Atoi(idString)
	if err != nil {
		base_error := errors.New_Invalid_request_error("user id must be an integer.", err).Error
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(http.StatusBadRequest, nil, base_error),
		})

		return
	}
	result, error := service.Get_cvs_by_user_id(user_id)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})

}

func GetCvById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		base_error := errors.New_Invalid_request_error("user id must be an integer.", err).Error
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(http.StatusBadRequest, nil, base_error),
		})

		return
	}

	result, error := service.Get_cv_by_id(id)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})

}

func UpdateCv(c *gin.Context) {
	var cvDto dto.CVDto
	cvDto = *validate_cv(c, cvDto)
	fmt.Printf("id is %d\n" , cvDto.ID)

	error := service.UpdateCv(cvDto)
	if error != nil {
		c.JSON(error.Error_code, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}

func validate_cv(c *gin.Context, cvDto dto.CVDto) *dto.CVDto {
	err := c.Bind(&cvDto)
	if err != nil {
		error := errors.New_convert_error(err).Error
		result := dto.Create_http_response(error.Error_code, nil, error)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result,
		})

		return nil
	}

	return &cvDto
}
