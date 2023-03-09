package model

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type PostUsecase interface {
	GetAll() ([]Post, error)
	GetByID(id int) (Post, error)
	Insert(p *Post) (Post, error)
	Update(p *Post) (Post, error)
	Delete(id int) error
}

type PostRepository interface {
	GetAll() ([]Post, error)
	GetByID(id int) (Post, error)
	Insert(p *Post) (Post, error)
	Update(p *Post) (Post, error)
	Delete(id int) error
}
