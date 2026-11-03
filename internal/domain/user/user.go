package user

import (
	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

type User struct {
	Id    int
	Name  string
	Email *Email
}

func New(dto userdto.UserCreateDTO) *User {
	return &User{Name: dto.Name, Email: NewEmailVo(&dto.Email)}
}

func ToResponse(u *User) userdto.UserResponseDTO {
	return userdto.UserResponseDTO{}
}

func ToResponseCollection(users []*User) []userdto.UserResponseDTO {
	var result []userdto.UserResponseDTO
	for _, user := range users {
		result = append(result, ToResponse(user))
	}
	return result
}

func (u *User) Update(dto userdto.UserUpdateDTO) {
	u.Name = dto.Name
	u.Email = NewEmailVo(&dto.Email)
}

func (u *User) Is(us *User) bool {
	return u.Id == us.Id
}
