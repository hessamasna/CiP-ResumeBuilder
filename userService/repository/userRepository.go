package repository

import (
	"fmt"
	"userService/dto"
	"userService/entities"
	"userService/errors"
	"userService/initializers"
)


func CreateUser(userDto dto.UserDto) (*dto.UserDto, *errors.Base_error) {
	// Convert the UserDto to a User entity
	var user entities.User
	err := initializers.Mapper.Map(&user, userDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	// Start a database transaction
	tx := initializers.DB.Begin()

	// Create the user entity in the database
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		errorMessage := "Failed to create user in database"
		return nil, errors.New_internal_error(errorMessage, err).Error
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		errorMessage := "Failed to commit transaction for creating user in database"
		return nil, errors.New_internal_error(errorMessage, err).Error
	}

	// Convert the created User entity back to a UserDto and return it
	var result dto.UserDto
	err = initializers.Mapper.Map(&result, user)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
	}

	return &result, nil
}

func FindUserByEmail(email string) (*dto.UserDto, *errors.Base_error) {
	var user entities.User
	initializers.DB.First(&user, "email = ?", email)
	if user.ID == 0 {
		return nil, nil
	}
	fmt.Println("user is password: " + user.PasswordHash)
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
    initializers.DB.Where("phone_number = ?", phoneNumber).First(&user)
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



func FindUserById(id int) (*dto.UserDto, *errors.Base_error) {
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
