package postdto

type PostCreateDTO struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int    `json:"author_id"`
}
