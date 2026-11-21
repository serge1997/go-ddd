package controller

import (
	"net/http"

	"github.com/serge1997/go-ddd/internal/application/usecase/commentuc"
)

func ListCommentByPost(useCase commentuc.ListByPostUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
