package postapp

import "github.com/MohammadBohluli/social-content-app/repository/psql"

type App struct {
	Handler    Handler
	Service    Service
	Validator  Validator
	Repository Repo
}

func New(conn *psql.DB) App {

	handler := NewHandler()
	repository := NewRepo(conn)
	validator := NewValidator()
	service := NewService(repository, validator)

	return App{
		Handler:    handler,
		Service:    service,
		Validator:  validator,
		Repository: repository,
	}
}
