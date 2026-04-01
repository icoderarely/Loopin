package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Version   int64     `json:"version"`
	Comments  []Comment `json:"comments"`
}

type PostStore struct {
	db *sql.DB
}

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, postID int64) (*Post, error) {
	query := `
	SELECT id, title, content, user_id, tags, created_at, version
	FROM posts
	WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	post := &Post{}

	err := s.db.QueryRowContext(
		ctx,
		query,
		postID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return post, nil
}

func (s *PostStore) Delete(ctx context.Context, postID int64) error {
	query := `
	DELETE FROM posts WHERE id = $1 RETURNING id;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var deletedID int64
	err := s.db.QueryRowContext(
		ctx,
		query,
		postID,
	).Scan(
		&deletedID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `
	UPDATE posts
	SET title = $1, content = $2, updated_at = NOW(), version = version + 1
	WHERE id = $3 AND version = $4
	RETURNING updated_at, version
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(
		&post.UpdatedAt,
		&post.Version,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	return nil
}
