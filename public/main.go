package main

import (
	"authuser/app/controllers"
	"authuser/app/controllers/api"
	"fmt"

	"github.com/go-martini/martini"
)

// Точка входа приложения
func main() {

	fmt.Println("Listening on port :8080")

	// Создание экземпляра фреймворка "martini" */
	m := martini.Classic()

	/* Регистрация наши маршруты */
	/*
		m.Get("/", func() string {
			return "Hello"
		})
	*/

	m.Get("/", controllers.Hello)

	m.Group("/api", func(r martini.Router) {

		m.Post("/user/register", api.RegisterUser)
		m.Post("/auth/login", api.LoginUser)
		m.Get("/currentuser", api.GetCurrentUser)
		// confirm user /user/confirm?token=werereetwwrwqwwwrreyyertweYZxssgsfwr, /user/confirm/{token}
		// try if can write like this /user/confirm and after ger param token, /user/confirm/:token
		m.Get("/user/confirm", api.Confirm)
		m.Post("/user/new", api.CreateNewUser)
		m.Post("/user/:id", api.GetUser)
		m.Put("/user/:id/edit", api.EditUser)
		m.Delete("/user/:id/delete", api.DeleteUser)
		m.Get("/user/:id/restore", api.RestoreUser)
	})

	// Слушаем в порте 8080, or run by default m.Run()
	m.RunOnAddr(":8080")
}

/*
func main() {

	fmt.Println("Listening on port :8080")

	// Создаем новый роутер
	m := martini.Classic()


	// Тестирование martini при вызове анонимной функции (anonymous function)
	m.Get("/", func() string {
			return "Точка входа приложения!"
	})

	// m.Run()
	m.RunOnAddr(":8080")
}
*/
