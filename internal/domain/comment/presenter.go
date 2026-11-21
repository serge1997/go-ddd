package comment

import (
	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

func ToResponseFromRaw(r commentdto.Raw) commentdto.Response {
	return commentdto.Response{
		Id:        r.Id,
		Content:   r.Content,
		CreatedAt: r.CreatedAt,
		Author:    userdto.UserReponseDTOMinified{Id: r.Id, Name: r.AuthorName},
	}
}

func ToResponseList(rows []commentdto.Raw) []commentdto.Response {
	var rs []commentdto.Response
	for _, row := range rows {
		rs = append(rs, ToResponseFromRaw(row))
	}
	return rs
}
