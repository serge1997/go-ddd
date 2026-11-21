package commentdto

import (
	"time"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
)

type CommentResponse struct {
	Id        int
	Content   string
	Author    userdto.UserReponseDTOMinified
	CreatedAt time.Time
}
