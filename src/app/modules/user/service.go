package user

type Service interface {
	RegisterUser(input UserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return 	&service{repo}
}
