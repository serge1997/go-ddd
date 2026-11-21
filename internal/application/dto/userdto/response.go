package userdto

type UserResponseDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserReponseDTOMinified struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
