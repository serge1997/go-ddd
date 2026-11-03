package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/serge1997/go-ddd/internal/application/usecase/useruc"
	"github.com/serge1997/go-ddd/internal/controller"
)

func UserRouter(r *chi.Mux, userUseCases useruc.UseCase) {
	ch := chi.NewRouter()
	r.Mount("/users", userRoutes(ch, userUseCases))
}

func userRoutes(r *chi.Mux, userUseCases useruc.UseCase) http.Handler {
	r.Post("/", controller.StoreUser(userUseCases.CreateUseCase))
	r.Get("/", controller.IndexUser(userUseCases.ListUseCase))
	r.Get("/{id}", controller.ShowUser(userUseCases.GetUseCase))
	r.Put("/{id}", controller.UpdateUser(userUseCases.UpdateUseCase))
	return r
}
