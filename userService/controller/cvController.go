package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"userService/dto"
	"userService/errors"
	"userService/service"
	"userService/util"
)

func CreateCv(c *gin.Context) {
	//check if user is login
	err := util.Check_if_is_login(c, "access_token")
	if err != nil {
		c.JSON(err.Error_code, gin.H{
			"result": dto.Create_http_response(err.Error_code, nil, err),
		})
		return
	}
	//read cvDto
	var cvDto dto.CVDto
	cvDto = *validate_cv(c, cvDto)
	//validate authorization
	err = util.CheckCurrentUserHasAccess(c , int(cvDto.UserID))
	if err != nil {
		c.JSON(err.Error_code, gin.H{
			"result": dto.Create_http_response(err.Error_code, nil, err),
		})
		return
	}
	// adding cv 
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
	
	
	err1 := util.Check_if_is_login(c, "access_token")
	if err1 != nil {
		c.JSON(err1.Error_code, gin.H{
			"result": dto.Create_http_response(err1.Error_code, nil, err1),
		})
		return 
	}
	idString := c.Param("user_id")
	user_id, err := strconv.Atoi(idString)

	if err != nil {
		base_error := errors.New_Invalid_request_error("user id must be an integer.", err).Error
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(http.StatusBadRequest, nil, base_error),
		})

		return
	}

	err2 := util.CheckCurrentUserHasAccess(c , user_id)
	if err2 != nil {
		c.JSON(err2.Error_code, gin.H{
			"result": dto.Create_http_response(err2.Error_code, nil, err2),
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
	err1 := util.Check_if_is_login(c, "access_token")
	if err1 != nil {
		c.JSON(err1.Error_code, gin.H{
			"result": dto.Create_http_response(err1.Error_code, nil, err1),
		})
		return 
	}
	var cvDto dto.CVDto
	cvDto = *validate_cv(c, cvDto)

	error := service.UpdateCv(cvDto)
	if error != nil {
		c.JSON(error.Error_code, gin.H{
			"result": dto.Create_http_response(error.Error_code, nil, error),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}

func DeleteCv(c *gin.Context) {
	err1 := util.Check_if_is_login(c, "access_token")
	if err1 != nil {
		c.JSON(err1.Error_code, gin.H{
			"result": dto.Create_http_response(err1.Error_code, nil, err1),
		})
		return 
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		base_error := errors.New_Invalid_request_error("user id must be an integer.", err).Error
		c.JSON(http.StatusBadRequest, gin.H{
			"result": dto.Create_http_response(http.StatusBadRequest, nil, base_error),
		})

		return
	}

	error := service.DeleteCv(id)
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
