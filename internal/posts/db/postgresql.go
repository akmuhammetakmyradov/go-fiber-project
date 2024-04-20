package postsdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/akmuhammetakmyradov/test/internal/posts"
	"github.com/akmuhammetakmyradov/test/pkg/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) posts.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUser(ctx context.Context, login string) (posts.User, error) {
	var result posts.User
	randomString := utils.GenerateRandomString(5)
	login = fmt.Sprintf("$%s$%s$%s$", randomString, login, randomString)

	query := `
		SELECT 
			id, name, login, password, type
		FROM users
		WHERE login = ` + login + `
 `
	err := r.db.QueryRow(ctx, query).Scan(&result.ID,
		&result.Name, &result.Login, &result.Password, &result.Type)

	if err != nil {
		fmt.Println("err in auth GetUser repo:", err)
		return result, err
	}

	return result, nil
}

func (r *repository) CreateUser(ctx context.Context, user posts.UserDTO) (posts.ID, error) {
	var result posts.ID

	query := `
		INSERT INTO 
			users (name, login, password, type)
			VALUES ($1, $2, $3, $4)
		RETURNING id;
			`
	err := r.db.QueryRow(ctx, query, user.Name, user.Login, user.Password,
		user.Type).Scan(&result.ID)

	if err != nil {
		fmt.Println("err in posts CreateUser repo:", err)
		return result, err
	}

	return result, nil
}

func (r *repository) CreatePost(ctx context.Context, user posts.PostDTO) (posts.ID, error) {
	var result posts.ID

	query := `INSERT INTO posts (header, text) VALUES ($1, $2) RETURNING id;`

	err := r.db.QueryRow(ctx, query, user.Header, user.Text).Scan(&result.ID)

	if err != nil {
		fmt.Println("err in posts CreatePost repo:", err)
		return result, err
	}

	return result, nil
}

func (r *repository) DeletePost(ctx context.Context, postID int) error {

	query := `DELETE FROM posts WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, postID)

	if err != nil {
		fmt.Println("err in posts DeletePost repo:", err)
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("not row effected")
	}

	return nil
}

func (r *repository) GetPosts(ctx context.Context, limit, page int) ([]posts.Post, error) {
	var result []posts.Post

	query := `
		SELECT
			id, header, text
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, limit, page)

	if err != nil {
		fmt.Println("err in posts GetPosts repo:", err)
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		post := posts.Post{}

		err = rows.Scan(&post.ID, &post.Header, &post.Text)
		if err != nil {
			fmt.Println("err in posts GetPosts query:", err)
			return result, err
		}

		result = append(result, post)
	}

	return result, nil
}

func (r *repository) GetPost(ctx context.Context, postID int) (posts.Post, error) {
	var result posts.Post

	query := `
		SELECT
			id, header, text
		FROM posts
		WHERE id = $1
		ORDER BY created_at DESC`

	err := r.db.QueryRow(ctx, query, postID).Scan(&result.ID, &result.Header, &result.Text)

	if err != nil {
		fmt.Println("err in posts GetPost repo:", err)
		return result, err
	}

	return result, nil
}
