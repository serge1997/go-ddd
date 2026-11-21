package postdto

import (
	"time"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

type Row struct {
	Id         int
	Title      string
	Content    string
	AuthorId   int
	AuthorName string
	CreatedAt  time.Time
}

type PostResponseDTO struct {
	Id      int                            `json:"id"`
	Title   string                         `json:"title"`
	Content string                         `json:"content"`
	Author  userdto.UserReponseDTOMinified `json:"author"`
}
