package repo

import (
	"ideal-journey/clients/errors"
	"ideal-journey/clients/logger"
	"ideal-journey/entity"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type AuthenticateRepo struct {
	db *gocql.Session
}

func NewAuthenticateRepo(db *gocql.Session) *AuthenticateRepo {
	return &AuthenticateRepo{
		db: db,
	}
}

const (
	queryFindById = "select id, nonce, fingerprint from user_auth where id = ?"
)

func (r *AuthenticateRepo) FindById(id string) (*entity.UserAuth, errors.RestErr) {
	result := &entity.UserAuth{}
	if err := r.db.Query(queryFindById, id).Scan(&result); err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil
		}
		logger.Error("[Auth]", err)
		return nil, errors.InternalServerError("database error", err)
	}
	return result, nil
}

const (
	querySave = "INSERT INTO user_auth(id, nonce, created) VALUES (?, ?, ?)"
)

func (r *AuthenticateRepo) Create(id string) (*entity.UserAuth, errors.RestErr) {
	created := time.Now().UTC().Unix()
	nonce := uuid.New().String()
	if err := r.db.Query(querySave, id, nonce, created).Exec(); err != nil {
		logger.Error("[Auth]", err)
		return nil, errors.InternalServerError("database error", err)
	}
	logger.Info("[Auth] Create %s", id)
	return &entity.UserAuth{
		Id:    id,
		Nonce: nonce,
	}, nil
}

const (
	queryUpdate = "UPDATE user_auth SET fingerprint = ?, nonce = ? WHERE id = ?"
)

func (r *AuthenticateRepo) Update(user *entity.UserAuth) errors.RestErr {
	nonce := uuid.New().String()
	if err := r.db.Query(queryUpdate, user.Fingerprint, nonce, user.Id).Exec(); err != nil {
		logger.Error("[Auth]", err)
		return errors.InternalServerError("database error", err)
	}
	logger.Info("[Auth] %s", user.Id)

	return nil
}
