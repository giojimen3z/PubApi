package builder

import "github.com/PubApi/cmd/api/app/domain/model"

type BeerDataBuilder struct {
	beerId   int64
	name     string
	brewery  string
	country  string
	price    float64
	currency string
}

func NewBikeDataBuilder() *BeerDataBuilder {
	return &BeerDataBuilder{
		beerId:   1,
		name:     "Golden",
		brewery:  "Kross",
		country:  "Chile",
		price:    10.5,
		currency: "EUR",
	}
}

func (builder *BeerDataBuilder) Build() model.Beer {
	return model.Beer{
		BeerId:   builder.beerId,
		Name:     builder.name,
		Brewery:  builder.brewery,
		Country:  builder.country,
		Price:    builder.price,
		Currency: builder.currency,
	}
}