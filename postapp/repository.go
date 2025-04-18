package postapp

import (
	"context"
	"fmt"

	"github.com/MohammadBohluli/social-content-app/pkg/errormessage"
	"github.com/MohammadBohluli/social-content-app/pkg/richerror"
	"github.com/MohammadBohluli/social-content-app/repository/psql"
	"github.com/jackc/pgtype"
)

type Repo struct {
	psqlDB *psql.DB
}

func NewRepo(conn *psql.DB) Repo {
	return Repo{psqlDB: conn}
}

func (r Repo) GetPost() {}

func (r Repo) GetAllPost() {}

func (r Repo) CreatePost(ctx context.Context, p Post) (Post, error) {
	const op = "repository.CreatePost"

	query := `
		INSERT INTO posts (content, title, tags, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			title,
			user_id,
			content,
			tags,
			created_at,
			updated_at;
	`

	stmt, err := r.psqlDB.Conn().PrepareContext(ctx, query)
	if err != nil {
		return Post{}, richerror.New(op).
			WithMessage(errormessage.ErrorMsgCantPrepareStatement).
			WithKind(richerror.KindUnexpected).
			WithErr(err)
	}
	defer stmt.Close()

	var post Post
	var tags pgtype.VarcharArray
	err = stmt.QueryRowContext(ctx, p.Content, p.Title, p.Tags, p.UserID).
		Scan(
			&post.ID,
			&post.Title,
			&post.UserID,
			&post.Content,
			&tags,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
	if err != nil {
		return Post{}, richerror.New(op).
			WithMessage(errormessage.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected).
			WithErr(err)
	}

	post.Tags, err = convertTextArray(tags)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func (r Repo) UpdatePost() {}

func (r Repo) DeletePost() {}

func convertTextArray(arr pgtype.VarcharArray) ([]string, error) {
	var result []string
	if err := arr.AssignTo(&result); err != nil {
		return nil, fmt.Errorf("failed to convert text array: %w", err)
	}
	return result, nil
}
