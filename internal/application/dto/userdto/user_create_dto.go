package userdto

import "fmt"

type UserCreateDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UserCreateDTO) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("%s is required", "name")
	}
	return nil
}
