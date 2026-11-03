package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/serge1997/go-ddd/internal/application/usecase/postuc"
	"github.com/serge1997/go-ddd/internal/controller"
)

func PostRouter(r *chi.Mux, postUseCases postuc.UseCase) {
	ch := chi.NewRouter()
	r.Mount("/posts", postRoutes(ch, postUseCases))
}
func postRoutes(r *chi.Mux, postUseCases postuc.UseCase) http.Handler {
	r.Post("/", controller.StorePost(postUseCases.CreateUsease))
	r.Get("/listByAuthor/{authorId?}", controller.ListByAuthor(postUseCases.ListByAuthor))
	r.Get("/{id}", controller.ShowPost(postUseCases.GetUseCase))
	r.Put("/{id}", controller.UpdatePost(postUseCases.UpdateUseCase))
	return r
}
