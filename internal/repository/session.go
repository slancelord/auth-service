package repository

import "auth-service/internal/model"

type Session interface {
	Save(session model.Session) error
}
