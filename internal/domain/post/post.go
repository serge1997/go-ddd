package post

import (
	"errors"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/domain/user"
)

var ErrNotAuthor = errors.New("user is not the author of the post")

type Post struct {
	Id      int
	Title   string
	Content string
	Author  *user.User
}

func New(dto postdto.PostCreateDTO, author *user.User) *Post {
	return &Post{Title: dto.Title, Content: dto.Content, Author: author}
}

func ToResponse(post *Post) postdto.PostResponseDTO {
	return postdto.PostResponseDTO{Id: post.Id, Title: post.Title, Content: post.Content, Author: user.ToMinifiedResponse(post.Author)}
}

func ToResponseCollection(posts []*Post) []postdto.PostResponseDTO {
	var result []postdto.PostResponseDTO
	for _, post := range posts {
		result = append(result, ToResponse(post))
	}
	return result
}

func (p *Post) IsAuthor(u *user.User) bool {
	return p.Author.Is(u)
}

func (p *Post) Update(dto postdto.PostUpdateDTO) {
	p.Title = dto.Title
	p.Content = dto.Content
}
