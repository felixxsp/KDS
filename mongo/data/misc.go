package myMongo

type OrderModifier struct {
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

type Customer struct {
	Name    string `json:"name" bson:"name"`
	Phone   Phone  `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
}

type Phone struct {
	CountryCode string `json:"country_code" bson:"country_code"`
	Number      string `json:"number" bson:"number"`
}
