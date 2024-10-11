package viewmodel

import (
	"fmt"
	"pennywise-api/entity"
	"regexp"
	"strings"

	"github.com/paemuri/brdoc"
	"github.com/pkg/errors"
)

type SignupRequest struct {
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
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

func (v SignupRequest) Normalize() (retVal SignupRequest) {
	retVal = v

	numbersOnlyRegex := regexp.MustCompile(`[^0-9]+`)
	retVal.Phone = numbersOnlyRegex.ReplaceAllString(v.Phone, "")
	retVal.CPF = numbersOnlyRegex.ReplaceAllString(v.CPF, "")
	retVal.Email = strings.ToLower(v.Email)
	retVal.FullName = strings.Title(strings.ToLower(v.FullName))

	return retVal
}

func (v SignupRequest) Validate() (err error) {
	fieldsErr := []string{}

	if v.Email == "" {
		fieldsErr = append(fieldsErr, "email")
	}

	if v.Password == "" {
		fieldsErr = append(fieldsErr, "password")
	}
	if v.Phone == "" {
		fieldsErr = append(fieldsErr, "phone")
	}

	if v.FullName == "" {
		fieldsErr = append(fieldsErr, "name")
	}

	if len(fieldsErr) > 0 {
		return errors.New(fmt.Sprintf("%s cannot be empty", fieldsErr))
	}

	numbersOnlyRegex := regexp.MustCompile(`[^0-9]+`)
	cleanPhoneNumber := numbersOnlyRegex.ReplaceAllString(v.Phone, "")
	if len(cleanPhoneNumber) != 11 {
		return errors.New("invalid phone number")
	}

	spllitedFullName := strings.Split(v.FullName, " ")
	if len(spllitedFullName) < 2 {
		return errors.New("invalid name")
	}

	for _, namePart := range spllitedFullName {
		if len(namePart) < 2 {
			return errors.New("invalid name")
		}
	}

	anyNumberRegex := regexp.MustCompile(`[0-9]+`)
	if anyNumberRegex.MatchString(v.FullName) {
		return errors.New("invalid phone number")
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(v.Email) {
		return errors.New("invalid email")
	}

	cleanCPF := numbersOnlyRegex.ReplaceAllString(v.CPF, "")
	if !brdoc.IsCPF(cleanCPF) {
		return errors.New("invalid cpf")
	}

	return nil
}
