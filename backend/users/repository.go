package users

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

// UserRepo is the struct representing the model for a post
type UserRepo struct {
	db *pgx.ConnPool
}

// NewUserRepo returns a new instance of UserRepo
func NewUserRepo(db *pgx.ConnPool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Create creates a new user
func (u *UserRepo) Create(user *User) (*User, error) {
	var createdUser User

	createdAt := time.Now()
	user.CreatedAt = &createdAt
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user.Password = string(encryptedPassword)
	user.Slug = uuid.New().String()
	row := u.db.QueryRow(`INSERT INTO "user" (slug, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING *`, user.Slug, user.Name, user.Email, user.Password, user.CreatedAt)

	err = row.Scan(&createdUser.ID, &createdUser.Slug, &createdUser.Name, &createdUser.Email, &createdUser.Password, &createdUser.CreatedAt, &createdUser.DeletedAt, &createdUser.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &createdUser, nil
}

// Update updates a user
func (u *UserRepo) Update(user *User) (*User, error) {
	return nil, nil
}

// Delete deletes a user
func (u *UserRepo) Delete(slug string) error {
	deletedAt := time.Now()

	_, err := u.db.Exec(`UPDATE "user" SET deleted_at = $1 where slug = $2`, &deletedAt, slug)
	if err != nil {
		log.Println("Failed to delete post: ", slug)
		return err
	}

	return nil
}

// Get returns a user
func (u *UserRepo) Get(slug string) (*User, error) {
	var user User

	row := u.db.QueryRow(`SELECT * FROM "user" WHERE slug = $1`, slug)

	err := row.Scan(&user.ID, &user.Slug, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.DeletedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetByEmail(email string) (*User, error) {
	var user User

	row := u.db.QueryRow(`SELECT * FROM "user" WHERE email = $1`, email)

	err := row.Scan(&user.ID, &user.Slug, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.DeletedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

// GetAll returns all users
func (u *UserRepo) GetAll() ([]User, error) {
	var users []User

	rows, err := u.db.Query(`SELECT * FROM "user"`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Slug, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.DeletedAt, &user.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
