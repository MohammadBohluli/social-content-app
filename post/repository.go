package post

import "github.com/MohammadBohluli/social-content-app/repository/psql"

type Repo struct {
	conn *psql.DB
}

func NewRepo(conn *psql.DB) Repo {
	return Repo{conn: conn}
}

func (r Repo) GetPost() {}

func (r Repo) GetAllPost() {}

func (r Repo) CreatePost() {}

func (r Repo) UpdatePost() {}

func (r Repo) DeletePost() {}
