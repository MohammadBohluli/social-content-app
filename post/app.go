package post

import "github.com/MohammadBohluli/social-content-app/repository/psql"

type App struct {
	Handler    Handler
	Service    Service
	Validator  Validator
	Repository Repo
}

func New(conn *psql.DB) App {

	handler := NewPostHandler()
	repository := NewPostRepo(conn)
	validator := NewPostValidator()
	service := NewPostService(repository, validator)

	return App{
		Handler:    handler,
		Service:    service,
		Validator:  validator,
		Repository: repository,
	}
}
