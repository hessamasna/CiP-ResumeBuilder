package controller

import (
	"net/http"
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
