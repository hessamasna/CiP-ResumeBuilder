package repository

import (
	"userService/dto"
	"userService/entities"
	"userService/errors"
	"userService/initializers"
)

func AddUser(userDto dto.UserDto) (*dto.UserDto, *errors.Base_error) {
	var userEntity entities.User
	err := initializers.Mapper.Map(&userEntity, userDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	result := initializers.DB.Create(&userEntity)
	err = result.Error

	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to add user to database", err).Error
	}
	return &userDto, nil
}

func FindUserByEmail(email string) (*dto.UserDto, *errors.Base_error) {
	var user entities.User
	initializers.DB.First(&user, "email = ?", email)
	if user.ID == 0 {
		return nil, nil
	}
	var result dto.UserDto
	err := initializers.Mapper.Map(&result, user)
 	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
	}
	return &result, nil
}


func FindUserByPhoneNumber(phoneNumber string) (*dto.UserDto, *errors.Base_error) {
	var user entities.User
	initializers.DB.First(&user, "phone_number = ?", phoneNumber)
	if user.ID == 0 {
		return nil, nil
	}
	var result dto.UserDto
	err := initializers.Mapper.Map(&result, user)
 	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
	}
	return &result, nil
}

func FindUserById(id int) (*dto.UserDto, *errors.Base_error){
	var user entities.User
	initializers.DB.First(&user, "ID = ?", id)
	if user.ID == 0 {
		return nil, nil
	}
		var result dto.UserDto
	err := initializers.Mapper.Map(&result, user)
 	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
	}
	return &result, nil
}
