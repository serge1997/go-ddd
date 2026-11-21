package commentdto

import (
	"time"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

type Raw struct {
	Id         int
	Content    string
	PostId     int
	AuthorId   int
	AuthorName string
	CreatedAt  time.Time
}
type Response struct {
	Id        int
	Content   string
	Author    userdto.UserReponseDTOMinified
	CreatedAt time.Time
}
