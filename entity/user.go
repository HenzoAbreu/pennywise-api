package entity

type User struct {
	ID           int64  `db:"user_id`
	UUID         string `db:"uuid"`
	FullName     string `db:"name"`
	Email        string `db:"email"`
	Phone        string `db:"phone"`
	CPF          string `db:"cpf"`
	Password     string `db:"password"`
	PasswordSalt string `db:"password_salt"`
}
