package mapper

import (
	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
)

func ToCommentRaw(m *model.CommentModel) *commentdto.Raw {
	return &commentdto.Raw{
		Id:         m.Id,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
		AuthorId:   m.AuthorId,
		AuthorName: m.Author.Name,
	}
}
