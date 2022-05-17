package repository

import (
	"context"

	"github.com/edubarbieri/julius/entity"
)

type NfeRepository interface {
	ExistByAccessKey(context.Context, string) (bool, error)
	Save(context.Context, entity.Nfe) (entity.Nfe, error)
}
