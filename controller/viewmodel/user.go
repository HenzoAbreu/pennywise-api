package viewmodel

import "pennywise-api/entity"

type UserResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	CPF      string `json:"cpf"`
	UUID     string `json:"uuid"`
}

func GenerateGetUserResponse(user entity.User) UserResponse {
	return UserResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		CPF:      user.CPF,
		UUID:     user.UUID,
	}
}
