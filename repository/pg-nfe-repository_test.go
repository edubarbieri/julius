package repository

import (
	"testing"

	"github.com/edubarbieri/julius/parser"

	"github.com/stretchr/testify/assert"
)

const (
	NfeUrl = "https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx?p=43220597320451006693650070002348661205910352%7C2%7C1%7C1%7CC05358F4DBE545C66285D92ACE0131D629E75FA4"
)

func TestShouldSaveNfeRepository(t *testing.T) {
	repository, err := NewPgNfeRepository("postgres://nfe:p@ssword@localhost/nfe")
	urlParser, err := parser.GetParser(NfeUrl)
	nfe := urlParser.Parse(NfeUrl)
	_, err = repository.Save(nfe)
	assert.Nil(t, err)
}
