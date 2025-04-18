package postapp

import (
	"github.com/MohammadBohluli/social-content-app/repository/psql"
	"github.com/go-playground/validator/v10"
)

type App struct {
	Handler    Handler
	Service    Service
	Repository Repo
	Validator  *validator.Validate
}

func New(conn *psql.DB, v *validator.Validate) App {

	repository := NewRepo(conn)
	service := NewService(repository, v)
	handler := NewHandler(service)

	return App{
		Handler:    handler,
		Service:    service,
		Repository: repository,
		Validator:  v,
	}
}
