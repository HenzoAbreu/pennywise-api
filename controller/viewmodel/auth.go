package viewmodel

import "pennywise-api/entity"

type SignupRequest struct {
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	Phone    string `json"phone"`
	Email    string `json"email"`
	Password string `json"password"`
}

type SignupResponse struct {
	FullName string `json:"full_name"`
	Email    string `json"email"`
}

func (v SignupRequest) ToEntity() entity.User {
	return entity.User{
		Email:    v.Email,
		Phone:    v.Phone,
		Password: v.Password,
		FullName: v.FullName,
		CPF:      v.CPF,
	}
}

func GenerateSignupResponse(user entity.User) SignupResponse {
	return SignupResponse{
		FullName: user.FullName,
		Email:    user.Email,
	}
}
