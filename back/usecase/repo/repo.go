package repo

import "github.com/gocql/gocql"

type Repos struct {
	Authenticate AuthenticateRepo
}

func NewRepositories(db *gocql.Session) *Repos {
	return &Repos{
		Authenticate: *NewAuthenticateRepo(db),
	}
}
