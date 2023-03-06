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

	err := validate_cv_name(cvDto.Title)
	if err != nil {
		return nil, err
	}

	err = validate_font_size(cvDto.FontSize)
	if err != nil {
		return nil, err
	}

	err = validate_template_number(cvDto.TemplateNumber)
	if err != nil {
		return nil, err
	}

	err = validate_cv_font_family(cvDto.FontFamily)
	if err != nil {
		return nil, err
	}

	err = validate_cv_color(cvDto.Color)
	if err != nil {
		return nil, err
	}

	userDto, error := repository.FindUserById(int(cvDto.UserID))
	if error != nil {
		return nil, error
	}
	if userDto == nil {
		error_message := fmt.Sprintf("User with ID : %v  does not exist.", cvDto.UserID)
		return nil, errors.New_entity_does_not_exist_error(error_message, nil).Error
	}
	_, err = repository.AddCv(cvDto)
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

func Get_cv_by_id(id int) (*dto.CVDto, *errors.Base_error) {
	if id == 0 {
		error_message := "User ID must be specified"
		return nil, errors.New_Invalid_request_error(error_message, nil).Error
	}
	cvDto, error := repository.GetCvById(id)

	if error != nil || cvDto == nil {
		return nil, error
	}

	return cvDto, nil
}

func UpdateCv(cvDto dto.CVDto) *errors.Base_error {
	err := validate_cv_id(cvDto.ID)
	if err != nil {
		return err
	}

	err = validate_cv_name(cvDto.Title)
	if err != nil {
		return err
	}

	err = validate_font_size(cvDto.FontSize)
	if err != nil {
		return err
	}

	err = validate_template_number(cvDto.TemplateNumber)
	if err != nil {
		return err
	}

	err = validate_cv_font_family(cvDto.FontFamily)
	if err != nil {
		return err
	}

	err = validate_cv_color(cvDto.Color)
	if err != nil {
		return err
	}

	_, error := repository.UpdateCv(cvDto)
	if error != nil {
		return error
	}
	return nil
}

func DeleteCv(id int) *errors.Base_error {
	if id == 0 {
		error_message := "User ID must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	error := repository.DeleteCvByID(id)
	if error != nil {
		return error
	}
	return nil
}

func validate_cv_id(id int) *errors.Base_error {
	if id == 0 {
		error_message := "User ID must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	return nil
}

func validate_cv_name(name string) *errors.Base_error {
	if name == "" {
		error_message := "Cv Name must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	return nil
}

func validate_font_size(font_size int) *errors.Base_error {
	if font_size == 0 {
		error_message := "Font size must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	return nil
}

func validate_cv_font_family(font_family string) *errors.Base_error {
	if font_family == "" {
		error_message := "Font Family must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	return nil
}

func validate_cv_color(color string) *errors.Base_error {
	if color == "" {
		error_message := "Color must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	return nil
}

func validate_template_number(templateNumber int) *errors.Base_error {
	if templateNumber == 0 {
		error_message := "templateNumber must be specified"
		return errors.New_Invalid_request_error(error_message, nil).Error
	}
	return nil
}
