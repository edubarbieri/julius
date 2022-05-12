package entity

import "time"

type Nfe struct {
	Url       string
	AccessKey string
	Date      time.Time
	StoreName string
	StoreCnpj string
	Items     []NfeItem
	Total     float64
	Discount  float64
}

type NfeItem struct {
	Description   string
	Quantity      float64
	UnitOfMeasure string
	UnityPrice    float64
	TotalPrice    float64
}
