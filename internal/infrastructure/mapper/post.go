package mapper

import (
	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
)

func ToPostRaw(m *model.PostModel) *postdto.Row {
	return &postdto.Row{
		Id:         m.Id,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
		AuthorId:   m.AuthorId,
		AuthorName: m.Author.Name,
	}
}
