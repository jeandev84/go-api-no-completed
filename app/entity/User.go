package entity

import (
	"time"
)

/*
Описание свойства
ID uid идентификатор
CreatedAt     time.Time  created_at
UpdatedAt     time.Time  updated_at
DeletedAt     time.Time  deleted_at
Email         string     email (required)
Secret        string     secret (обязательно, храним как хэш, не в открытом виде)
FirstName     string     first_name (обязательно указано при регистрации)
LastName      string     last_name (обязательно указано при регистрации)
MiddleName    string     middle_name (не обязательно указано при регистрации)
Role          string     role (user или admin [далее возможно будут варианты ролей])
isActive      boolean    is_active (признак, что пользователь не заблокирован)
ConfirmToken  string     confirm_token (токен для подтверждения почты)
Confirmed     boolean    confirmed (признак, что почта подтверждена)
*/

// Структура пользователь (Object User)
type User struct {

	// Свойства
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	Secret       string `json:"secret"`
	FirstName    string `json:"password"`
	LastName     string `json:"lastname"`
	MiddleName   string `json:"middlename"`
	Role         string `json:"role"`
	isActive     bool
	ConfirmToken string    `json:"confirm_token"`
	Confirmed    string    `json:"confirmed"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

/*
func CreateNewUser() {
    return &User{}
}
*/
