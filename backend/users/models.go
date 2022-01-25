package users

import (
	"time"

	"github.com/jackc/pgx"
)

//User is the struct representing the model for a post
type User struct {
	ID        int64      `json:"-"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

//UserRepository interface is the interface representing the post repository
type UserRepository interface {
	NewUserRepository(db *pgx.ConnPool) *UserRepository
	Create(post *User) (*User, error)
	Update(post *User) (*User, error)
	Delete(id int) error
	Get(slug string) (*User, error)
	GetAll() ([]User, error)
}
