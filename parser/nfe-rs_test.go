package parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldParseNfe(t *testing.T) {
	parser := NewNfeRs()
	nfe := parser.Parse("https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx?p=43220490919200000569650080000895589133243560%7C2%7C1%7C28%7C146.62%7C65798DBD06DB772568CBBFBADC16EC2F9EB78145%7C1%7CB83112B3D6582EC49DF1837606A226254B1A0F83")
	assert.Equal(t, "HIPERFACIL ATACAREJO", nfe.StoreName)
	assert.Equal(t, "90.919.200/0005-69", nfe.StoreCnpj)
	assert.Equal(t, "4322 0490 9192 0000 0569 6500 8000 0895 5891 3324 3560", nfe.AccessKey)
	assert.Equal(t, 0.0, nfe.Discount)
	assert.Equal(t, 146.62, nfe.Total)
	assert.Equal(t, 20, len(nfe.Items))
	tz, _ := time.LoadLocation("America/Sao_Paulo")
	date, _ := time.ParseInLocation("02/01/2006 15:04:05", "28/04/2022 18:26:27", tz)
	assert.Equal(t, date, nfe.Date)
	assert.Equal(t, "BANANA PRATA kg", nfe.Items[0].Description)
	assert.Equal(t, 1.02, nfe.Items[0].Quantity)
	assert.Equal(t, "KG", nfe.Items[0].UnitOfMeasure)
	assert.Equal(t, 6.49, nfe.Items[0].UnityPrice)
	assert.Equal(t, 6.62, nfe.Items[0].TotalPrice)
}

func TestShouldParseLongerNfe(t *testing.T) {
	parser := NewNfeRs()
	nfe := parser.Parse("https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx?p=43220597320451006693650070002348661205910352%7C2%7C1%7C1%7CC05358F4DBE545C66285D92ACE0131D629E75FA4")
	assert.Equal(t, "Cooperativa Triticola Sarandi Ltda", nfe.StoreName)
	assert.Equal(t, "97.320.451/0066-93", nfe.StoreCnpj)
	assert.Equal(t, "4322 0597 3204 5100 6693 6500 7000 2348 6612 0591 0352", nfe.AccessKey)
	assert.Equal(t, 0.0, nfe.Discount)
	assert.Equal(t, 410.65, nfe.Total)
	assert.Equal(t, 52, len(nfe.Items))
	tz, _ := time.LoadLocation("America/Sao_Paulo")
	date, _ := time.ParseInLocation("02/01/2006 15:04:05", "07/05/2022 11:40:16", tz)
	assert.Equal(t, date, nfe.Date)
	assert.Equal(t, "LIMP.LIMPOL 500ML PERF.FRESCOR DO OCEANO", nfe.Items[0].Description)
	assert.Equal(t, 1.0, nfe.Items[0].Quantity)
	assert.Equal(t, "FR", nfe.Items[0].UnitOfMeasure)
	assert.Equal(t, 4.79, nfe.Items[0].UnityPrice)
	assert.Equal(t, 4.79, nfe.Items[0].TotalPrice)
}
