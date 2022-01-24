package posts

import (
	"time"

	"github.com/jackc/pgx"
)

//Post is the struct representing the model for a post
type Post struct {
	ID        int64      `json:"-"`
	Slug      string     `json:"slug"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

//PostRepository interface is the interface representing the post repository
type PostRepository interface {
	NewPostRepository(db *pgx.ConnPool) *PostRepository
	Create(post *Post) (*Post, error)
	Update(post *Post) (*Post, error)
	Delete(id int) error
	Get(slug string) (*Post, error)
	GetAll() ([]Post, error)
}
