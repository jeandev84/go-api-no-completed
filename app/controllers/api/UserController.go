package api

import (
	"net/http"
)

// Получить текущего пользователя (Me)
var GetCurrentUser = func(w http.ResponseWriter, r *http.Request) string {
	return "Current user"
}

// Регистрация пользователя
var RegisterUser = func(w http.ResponseWriter, r *http.Request) {

}

// Подтверждения пользователя
var Confirm = func(w http.ResponseWriter, r *http.Request) {

}

// Создание нового пользователя
var CreateNewUser = func(w http.ResponseWriter, r *http.Request) {

}

// Найти пользователя по идентификатору
var GetUser = func(w http.ResponseWriter, r *http.Request) {

}

// Редактирование пользователя
var EditUser = func(w http.ResponseWriter, r *http.Request) {

}

// Удаление пользователя
var DeleteUser = func(w http.ResponseWriter, r *http.Request) {

}

// Восстановление пользователя
var RestoreUser = func(w http.ResponseWriter, r *http.Request) {

}
