package mapper

import (
	entity "github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
)

func UserToDomain(m model.UserModel) *entity.User {
	return &entity.User{
		Id:    m.Id,
		Name:  m.Name,
		Email: entity.NewEmailVo(m.Email),
	}
}

func UserToModel(user *entity.User) model.UserModel {
	return model.NewUserModel(user)
}

func UserModelsToDomainCollection(users []model.UserModel) []*entity.User {
	var domains []*entity.User
	for _, user := range users {
		domains = append(domains, UserToDomain(user))
	}
	return domains
}
