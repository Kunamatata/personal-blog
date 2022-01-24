package posts

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

//PostRepo is the struct representing the post repository
type PostRepo struct {
	db *pgx.ConnPool
}

// NewPostRepo returns a new instance of PostRepository
func NewPostRepo(db *pgx.ConnPool) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

//Create creates a new post
func (p *PostRepo) Create(post *Post) (*Post, error) {
	var createdPost Post

	post.Slug = uuid.New().String()
	createdAt := time.Now()
	post.CreatedAt = &createdAt

	row := p.db.QueryRow("INSERT INTO post (slug, title, content, created_at) VALUES ($1, $2, $3, $4) RETURNING *", post.Slug, post.Title, post.Content, post.CreatedAt)

	err := row.Scan(&createdPost.ID, &createdPost.Slug, &createdPost.Title, &createdPost.Content, &createdPost.CreatedAt, &createdPost.DeletedAt, &createdPost.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &createdPost, nil
}

//Update updates a post
func (p *PostRepo) Update(post *Post) (*Post, error) {

	return nil, nil
}

//Delete deletes a post
func (p *PostRepo) Delete(slug string) error {
	deletedAt := time.Now()

	_, err := p.db.Exec("UPDATE post SET deleted_at = $1 WHERE slug = $2", &deletedAt, slug)

	if err != nil {
		log.Println("Failed to delete post: ", slug)
		return err
	}

	return nil
}

//Get returns a post
func (p *PostRepo) Get(slug string) (*Post, error) {
	var post Post
	row := p.db.QueryRow("SELECT * FROM post WHERE slug = $1", slug)

	err := row.Scan(&post.ID, &post.Slug, &post.Title, &post.Content, &post.CreatedAt, &post.DeletedAt, &post.UpdatedAt)
	if err != nil {
		log.Println("Failed to retrieve post")
		return nil, err
	}

	if post.DeletedAt != nil {
		return nil, nil
	}

	return &post, nil
}

//GetAll returns all posts
func (p *PostRepo) GetAll() ([]Post, error) {
	var posts []Post

	rows, err := p.db.Query("SELECT * FROM post where deleted_at is null")
	if err != nil {
		log.Println("Failed to retrieve all posts")
		return nil, err
	}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Slug, &post.Title, &post.Content, &post.CreatedAt, &post.DeletedAt, &post.UpdatedAt)
		if err != nil {
			log.Println("Failed to retrieve post")
		}
		posts = append(posts, post)
	}

	return posts, nil
}
