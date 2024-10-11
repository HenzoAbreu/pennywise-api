package data

import (
	database "pennywise-api/db"
	"pennywise-api/entity"

	"github.com/pkg/errors"
)

func Save(user entity.User) (retVal entity.User, error error) {
	query := `
	insert into tb_user
		(name, user_uuid, email, phone, cpf, password, password_salt)
		values (?,?,?,?,?,?,?)
		;
	`

	_, err := database.DB.Exec(query, user.FullName, user.UUID, user.Email, user.Phone, user.CPF, user.Password, user.PasswordSalt)
	if err != nil {
		return retVal, errors.Wrap(err, "failed to register user")
	}
	return user, nil
}
