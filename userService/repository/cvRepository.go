package repository


import (
	"userService/dto"
	"userService/entities"
	"userService/errors"
	"userService/initializers"
)


func AddCv(cvDto dto.CVDto) (*dto.CVDto, *errors.Base_error) {
	var cvEntity entities.CV
	err := initializers.Mapper.Map(&cvEntity, cvDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	result := initializers.DB.Create(&cvEntity)
	err = result.Error

	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to add user to database", err).Error
	}
	return &cvDto, nil
}