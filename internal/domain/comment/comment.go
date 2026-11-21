package comment

import (
	"time"
)

type Comment struct {
	Id        int
	Content   string
	AuthorId  int
	PostId    int
	CreatedAt time.Time
}

func Restore(Id int, Content string, AuthorId int, PostId int, CreatedAt time.Time) *Comment {
	return &Comment{Id, Content, AuthorId, PostId, CreatedAt}
}
