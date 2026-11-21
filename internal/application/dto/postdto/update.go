package postdto

type PostUpdateDTO struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int    `json:"author_id"`
}
