package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/application/usecase/useruc"
)

func ShowUser(usecase useruc.UserGetUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idparam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idparam)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("erro on parse id to string: %s", err))
		}
		response := usecase.Execute(r.Context(), id)
		if response.IsFail() {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("erro on get user: %s", response.Error()))
		}
		responseOk(w, http.StatusOK, fmt.Sprintf("showing a user, id: %d", id), response.Value())
	}
}
func IndexUser(usecase useruc.UserListUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := usecase.Execute(r.Context())
		if response.IsFail() {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("error on create user: %s", response.Error()))
			return
		}
		responseOk(w, http.StatusOK, "show users", response.Value())
	}
}

func StoreUser(usecase useruc.UserCreateUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto userdto.UserCreateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("error on create user: %s", err))
			return
		}
		result_uc := usecase.Execute(r.Context(), dto)
		if result_uc.IsFail() {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("error on create user: %s", result_uc.Error()))
			return
		}
		responseOk(w, http.StatusCreated, "user created successfully", result_uc.Value())
	}
}

func UpdateUser(usecase useruc.UserUpdateUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto userdto.UserUpdateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("error on update user: %s", err))
			return
		}
		response := usecase.Execute(r.Context(), dto)
		if response.IsFail() {
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("error on update user: %s", response.Error()))
			return
		}
		responseOk(w, http.StatusOK, "user updated successfully", response.Value())
	}
}
