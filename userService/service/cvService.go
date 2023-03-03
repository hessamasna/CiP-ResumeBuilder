package service

import (
	"fmt"
	"userService/dto"
	"userService/errors"
	"userService/repository"
)

func AddCv(cvDto dto.CVDto) (*dto.CVDto, *errors.Base_error) {
	// if cvDto.ID != 0 {
	// 	cv, error := repository.FindCVById(cvDto.ID)
	// 	if error != nil {
	// 		return nil, error
	// 	}
	// 	if cv != nil {
	// 		error_message := fmt.Sprintf("Cv with ID : %v already exists", cv.ID)
	// 		return nil, errors.New_entity_already_exist_error(error_message, nil).Error
	// 	}
	// }
	cvDto.ID = 0

	if cvDto.User_id == 0  {
		error_message := "User ID must be specified"
		return nil , errors.New_Invalid_request_error(error_message , nil).Error
	}
	userDto , error := repository.FindUserById(int(cvDto.User_id))
	if error != nil {
		return nil, error
	}
	if userDto == nil {
		error_message := fmt.Sprintf("User with ID : %v  does not exist."  , cvDto.User_id)
		return nil , errors.New_entity_does_not_exist_error(error_message , nil).Error
	}
	_, err := repository.AddCv(cvDto)
	if err != nil {
		//TODO log error
		return nil, err
	}

	return &cvDto, nil
}
