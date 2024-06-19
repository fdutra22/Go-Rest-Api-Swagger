package main

import (
	"net/http"

	"github.com/fdutra22/go-rest-api-swagger/exemplo/api"
	_ "github.com/fdutra22/go-rest-api-swagger/exemplo/docs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	Exemplo de cadastro de usuario em GoLang com Swagger.
//	@description.markdown
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@tag.name	admin
//	@tag.description.markdown

//	@BasePath	/v2

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/v2/admin/user/", api.ListUsers).Methods("GET")
	router.HandleFunc("/v2/admin/user/{id}", api.GetUser).Methods("GET")
	router.HandleFunc("/v2/admin/user/", api.AddUser).Methods("POST")
	router.HandleFunc("/v2/admin/user/{id}", api.UpdateUser).Methods("PUT")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)

}
