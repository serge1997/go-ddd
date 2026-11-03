package app

import (
	"github.com/serge1997/go-ddd/internal/application/usecase/postuc"
	"github.com/serge1997/go-ddd/internal/application/usecase/useruc"
	"github.com/serge1997/go-ddd/internal/database"
	"github.com/serge1997/go-ddd/internal/infrastructure/repository"
	"github.com/serge1997/go-ddd/internal/routes"
)

func (a *app) boostrap() {
	userRepository := repository.NewUserRepository(database.Db)
	postRepostory := repository.NewPostRepository(database.Db)
	userCases := useruc.New(userRepository)
	postCases := postuc.New(postRepostory, userRepository)

	routes.UserRouter(a.Mux, userCases)
	routes.PostRouter(a.Mux, postCases)
}
