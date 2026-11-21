package post

import (
	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

func ToResponse(post *postdto.Row) postdto.PostResponseDTO {
	return postdto.PostResponseDTO{Id: post.Id, Title: post.Title, Content: post.Content, Author: userdto.UserReponseDTOMinified{
		Id:   post.AuthorId,
		Name: post.AuthorName,
	}}
}

func ToResponseList(rows []postdto.Row) []postdto.PostResponseDTO {
	var result []postdto.PostResponseDTO
	for _, row := range rows {
		result = append(result, ToResponse(&row))
	}
	return result
}
