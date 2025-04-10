package postapp

type Repository interface {
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type Service struct {
	repo      Repository
	validator Validator
}

func NewService(r Repository, v Validator) Service {
	return Service{
		repo:      r,
		validator: v,
	}
}

func (s Service) GetPost() {}

func (s Service) GetAllPost() {}

func (s Service) CreatePost() {}

func (s Service) UpdatePost() {}

func (s Service) DeletePost() {}
