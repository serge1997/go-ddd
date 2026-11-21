package post

import (
	"errors"
	"time"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
)

var ErrNotAuthor = errors.New("user is not the author of the post")

type Post struct {
	Id        int
	Title     string
	Content   string
	AuthorId  int
	CreatedAt time.Time
}

func New(dto postdto.PostCreateDTO, authorId int) *Post {
	return &Post{Title: dto.Title, Content: dto.Content, AuthorId: authorId}
}

func (p *Post) IsAuthor(authorId int) bool {
	return p.AuthorId == authorId
}

func (p *Post) Update(dto postdto.PostUpdateDTO) {
	p.Title = dto.Title
	p.Content = dto.Content
}

func Restore(r *postdto.Row) *Post {
	return &Post{
		Id:        r.AuthorId,
		Title:     r.Title,
		Content:   r.Content,
		AuthorId:  r.AuthorId,
		CreatedAt: r.CreatedAt,
	}
}
