package user

import "github.com/serge1997/go-ddd/internal/application/dto/userdto"

func ToResponse(u *User) userdto.UserResponseDTO {
	return userdto.UserResponseDTO{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email.email,
	}
}

func ToMinifiedResponse(u *User) userdto.UserReponseDTOMinified {
	return userdto.UserReponseDTOMinified{
		Id:   u.Id,
		Name: u.Name,
	}
}

func ToResponseList(users []*User) []userdto.UserResponseDTO {
	var result []userdto.UserResponseDTO
	for _, user := range users {
		result = append(result, ToResponse(user))
	}
	return result
}
