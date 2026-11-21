package mapper

import (
	"time"

	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/domain/comment"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
)

type CommentRaw struct {
	Id         int
	Content    string
	PostId     int
	AuthorId   int
	AuthorName string
	CreatedAt  time.Time
}

func ToCommentResponse(row CommentRaw) commentdto.CommentResponse {
	return commentdto.CommentResponse{
		Id:      row.Id,
		Content: row.Content,
		Author:  userdto.UserReponseDTOMinified{Id: row.AuthorId, Name: row.AuthorName},
	}
}

func ToCommentResponseList(rows []CommentRaw) []commentdto.CommentResponse {
	var rs []commentdto.CommentResponse
	for _, row := range rows {
		rs = append(rs, ToCommentResponse(row))
	}
	return rs
}

func CommentRowToDomain(row CommentRaw) *comment.Comment {
	return comment.Restore(row.Id, row.Content, row.AuthorId, row.PostId, row.CreatedAt)
}

func CommentResponseFromModel(c *model.CommentModel) *commentdto.CommentResponse {
	return &commentdto.CommentResponse{
		Id:        c.Id,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		Author:    userdto.UserReponseDTOMinified{Id: c.AuthorId, Name: c.Author.Name},
	}
}
