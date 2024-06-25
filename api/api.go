package api

import (
	"net/http"
	"time"
)

// User example
type User struct {
	ID       int64
	Email    string
	Password string
}

// UsersCollection example
type UsersCollection []User

// Error example
type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

// ListUsers example
//
//	@Summary	Lista de usuarios
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}	api.UsersCollection	"ok"
//	@Router		/admin/user/ [get]
func ListUsers(w http.ResponseWriter, r *http.Request) {

}

// GetUser example
//
//	@Summary	Busca usuario por id
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"User Id"
//	@Success	200	{object}	api.User
//	@Failure	400	{object}	api.APIError	"ID Necessário!!"
//	@Failure	404	{object}	api.APIError	"ID não encontrado"
//	@Router		/admin/user/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// AddUser example
//
//	@Summary	Novo usuario
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		message	body		api.User		true	"User Data"
//	@Failure	400	{object}	api.APIError	"ID Necessário!!"
//	@Failure	404	{object}	api.APIError	"ID não encontrado""
//	@Router		/admin/user/ [post]
func AddUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser example
//
//	@Summary	Atualiza usuario
//	@Tags		admin
//	@Accept		json
//	@Produce	json
//	@Param		message	body		api.User		true	"Dados usuário"
//	@Success	200		{string}	string			"ok"
//	@Failure	400	{object}	api.APIError	"ID Necessário!!"
//	@Failure	404	{object}	api.APIError	"ID não encontrado"
//	@Router		/admin/user/ [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
