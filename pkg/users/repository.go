package users

import (
	"github.com/Thal-e/go-template/pkg/models"
	"github.com/google/uuid"
)

type Repository interface {
	Find(id *uuid.UUID) (*models.Users, error)
}
