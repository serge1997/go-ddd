package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/application/usecase/postuc"
	"github.com/serge1997/go-ddd/internal/shared/domain"
)

func ListByAuthor(usecase postuc.PostListByAuthorUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramsAuthorId := chi.URLParam(r, "authorId")
		authorId, err := strconv.Atoi(paramsAuthorId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("erro on parse id to string: %s", err))
		}
		if authorId == 0 {
			//author from jwt token
			authorId = 28
		}
		result := usecase.Execute(r.Context(), authorId)
		if result.IsFail() {
			respondWithError(w, domain.StatusFromError(result.Error()), result.Error())
			return
		}
		responseOk(w, http.StatusOK, "List posts by author", result.Value())
	}
}
func StorePost(usecase postuc.PostCreateUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := parseRequestBody[postdto.PostCreateDTO](r)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		result := usecase.Execute(r.Context(), dto)
		if result.IsFail() {
			respondWithError(w, http.StatusInternalServerError, result.Error())
			return
		}
		responseOk(w, http.StatusCreated, "Post created successfully", result.Value())
	}
}

func ShowPost(usecase postuc.PostGetUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("erro on parse id to string: %s", err))
		}
		result := usecase.Execute(r.Context(), id)
		if result.IsFail() {
			respondWithError(w, domain.StatusFromError(result.Error()), result.Error())
			return
		}
		responseOk(w, http.StatusOK, "showing post", result.Value())
	}
}

func UpdatePost(usecase postuc.PostUpdateUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto, err := parseRequestBody[postdto.PostUpdateDTO](r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}
		result := usecase.Execute(r.Context(), dto)
		if result.IsFail() {
			respondWithError(w, domain.StatusFromError(result.Error()), result.Error())
			return
		}
		responseOk(w, http.StatusOK, "post updated successfully", result.Value())
	}
}
