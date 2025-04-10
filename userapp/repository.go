package postapp

import "github.com/MohammadBohluli/social-content-app/repository/psql"

type Repo struct {
	conn *psql.DB
}

func NewRepo(conn *psql.DB) Repo {
	return Repo{conn: conn}
}

func (r Repo) GetUser() {}

func (r Repo) CreateUser() {}

func (r Repo) UpdateUser() {}

func (r Repo) DeleteUser() {}
