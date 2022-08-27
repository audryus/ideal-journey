package usecase

import "ideal-journey/usecase/repo"

type Dependecies struct {
	Repos *repo.Repos
}
type AuthenticateDependecies struct {
	Repos repo.AuthenticateRepo
	JwtUC *JwtUC
}

type Services struct {
	Authenticate AuthenticateUC
}

func NewServices(deps Dependecies) *Services {
	return &Services{
		Authenticate: *NewAuthenticateUC(&AuthenticateDependecies{
			Repos: deps.Repos.Authenticate,
			JwtUC: NewJwtUC(),
		}),
	}
}
