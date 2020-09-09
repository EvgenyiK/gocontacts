package main

import (
	"fmt"
	"gocontacts/app"
	"gocontacts/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

/*
Мы создаём новый объект Router, подключаем наше промежуточное программное обеспечение JWT auth,
 используя функцию маршрутизатора Use(), а затем приступаем к прослушиванию входящих запросов.
*/
