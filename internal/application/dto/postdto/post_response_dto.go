package postdto

import (
	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

type PostResponseDTO struct {
	Id      int                            `json:"id"`
	Title   string                         `json:"title"`
	Content string                         `json:"content"`
	Author  userdto.UserReponseDTOMinified `json:"author"`
}
