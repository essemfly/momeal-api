package src

import "time"

type NaverProductEntity struct {
	Name           string          `json:"productName"`
	Title          string          `json:"productTitle"`
	TitleOrg       string          `json:"productTitleOrg"`
	AttributeValue string          `json:"attributeValue"`
	CharacterValue string          `json:"characterValue"`
	ImageUrl       string          `json:"imageUrl"`
	Price          string          `json:"price"`
	PriceUnit      string          `json:"priceUnit"`
	Maker          string          `json:"maker"`
	Brand          string          `json:"brand"`
	Category1Name  string          `json:"category1Name"`
	MallName       string          `json:"mallName"`
	MallNameOrg    string          `json:"mallNameOrg"`
	MallProductUrl string          `json:"mallProductUrl"`
	MallProdMblUrl string          `json:"mallProdMblUrl"`
	DeliveryFee    string          `json:"dlvryCont"`
	PurchaseCount  int             `json:"purchaseCnt"`
	ReviewCount    int             `json:"reviewCountSum"`
	MallInfo       NaverMallEntity `json:"mallInfoCache"`
	OpenDate       string          `json:"openDate"`
	CreatedAt      time.Time       `json:"created_at,omitempty" bson:"created_at"`
}

type NaverMallEntity struct {
	Name        string            `json:"name"`
	BizAddr     string            `'json:"bizplBaseAddr"`
	BizNo       string            `json:"businessNo"`
	Description string            `json:"mallIntroduction"`
	LogoUrl     map[string]string `json:"mallLogos"`
	CreatedAt   time.Time         `json:"created_at,omitempty" bson:"created_at"`
}
