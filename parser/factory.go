package parser

import (
	"fmt"
	"strings"

	"github.com/edubarbieri/julius/entity"
)

type Parser interface {
	Parse(url string) entity.Nfe
}

func GetParser(nfeUrl string) (Parser, error) {
	if strings.Contains(nfeUrl, "sefaz.rs") {
		return NewNfeRs(), nil
	}
	return nil, fmt.Errorf("dont have parser implementation for url %s", nfeUrl)
}
