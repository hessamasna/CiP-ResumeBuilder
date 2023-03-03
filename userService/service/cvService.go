package service

import (
	"fmt"
	"userService/dto"
	"userService/errors"
	"userService/repository"
)

func AddCv(cvDto dto.CVDto) (*dto.CVDto, *errors.Base_error) {
	cvDto.ID = 0

	if cvDto.UserID == 0 {
		error_message := "User ID must be specified"
		return nil, errors.New_Invalid_request_error(error_message, nil).Error
	}
	userDto, error := repository.FindUserById(int(cvDto.UserID))
	if error != nil {
		return nil, error
	}
	if userDto == nil {
		error_message := fmt.Sprintf("User with ID : %v  does not exist.", cvDto.UserID)
		return nil, errors.New_entity_does_not_exist_error(error_message, nil).Error
	}
	_, err := repository.AddCv(cvDto)
	if err != nil {
		//TODO log error
		return nil, err
	}

	return &cvDto, nil
}

func Get_cvs_by_user_id(id int) ([]dto.CVDto, *errors.Base_error) {
	if id == 0 {
		error_message := "User ID must be specified"
		return nil, errors.New_Invalid_request_error(error_message, nil).Error
	}
	userDto, error := repository.FindUserById(id)
	if error != nil {
		return nil, error
	}
	if userDto == nil {
		error_message := fmt.Sprintf("User with ID : %v  does not exist.", id)
		return nil, errors.New_entity_does_not_exist_error(error_message, nil).Error
	}
	return repository.GetCVsByUserId(id)
}
