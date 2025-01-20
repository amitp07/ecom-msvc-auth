package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amitp07/ecom-msvc-auth/internal/data"
	"github.com/go-chi/chi/v5"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		app.logger.Printf("Error: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = app.models.UserModel.Insert(&user)

	if err != nil {
		app.logger.Printf("Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Successfully created user with id %d", user.Id) + "\n"))

}

func (app *application) login(w http.ResponseWriter, r *http.Request) {

	type userType struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var user userType

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.logger.Printf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	dbuser, err := app.models.UserModel.FindByUsername(user.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		app.logger.Printf("%s", err.Error())
		return
	}

	msg := "Either Username or Password is not correct"
	if user.Username != dbuser.Username {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(msg))
		return
	}
	if user.Password != dbuser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(dbuser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)

}

func (app *application) findById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		app.logger.Printf("Error: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	user, err := app.models.UserModel.FindById(id)

	if err != nil {
		app.logger.Printf("Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(user)

	if err != nil {
		app.logger.Printf("Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", js)
}
