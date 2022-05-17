package service

import (
	"context"
	"fmt"

	"github.com/edubarbieri/julius/parser"

	"github.com/edubarbieri/julius/repository"
)

type NfeService struct {
	nfeRepository repository.NfeRepository
}

func NewNfeService(nfeRepository repository.NfeRepository) NfeService {
	return NfeService{
		nfeRepository: nfeRepository,
	}
}

func (s NfeService) SaveNfe(ctx context.Context, url string) error {
	if len(url) == 0 {
		return fmt.Errorf("could not parser empty url")
	}

	nfeParser, err := parser.GetParser(url)
	if err != nil {
		return err
	}

	nfe := nfeParser.Parse(url)
	exists, err := s.nfeRepository.ExistByAccessKey(ctx, nfe.AccessKey)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("nfe allready imported %s", nfe.AccessKey)
	}
	_, err = s.nfeRepository.Save(ctx, nfe)
	return err
}
