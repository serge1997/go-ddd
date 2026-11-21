package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/serge1997/go-ddd/internal/application/usecase/commentuc"
	"github.com/serge1997/go-ddd/internal/controller"
)

func CommentRouter(r *chi.Mux, useCases commentuc.UseCase) {
	ch := chi.NewRouter()
	r.Mount("/comments", commentRoutes(ch, useCases))
}

func commentRoutes(r *chi.Mux, useCases commentuc.UseCase) http.Handler {
	r.Get("/listByPost/{postId}", controller.ListCommentByPost(useCases.ListByPost))
	return r
}
