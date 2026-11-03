package userdto

type UserUpdateDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}
