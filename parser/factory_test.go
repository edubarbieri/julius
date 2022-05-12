package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnParseInstace(t *testing.T) {
	parser, err := GetParser("https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx?p=43220590919200000569650010001391029248818480%7C2%7C1%7C02%7C92.15%7CB2AAB3108BD015FEFE28B2BDD1D979EB1B50D5EB%7C1%7C0F9E3DB707C319564C1C03E8838794A4CA4CA905")
	assert.Nil(t, err)
	assert.IsType(t, NfeRs{}, parser)
}
