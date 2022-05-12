package repository

import "github.com/edubarbieri/julius/entity"

type NfeRepository interface {
	Save(entity.Nfe) (entity.Nfe, error)
}
