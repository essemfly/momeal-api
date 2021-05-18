package src

import (
	"time"
)

type BenefitsEntity struct {
	DiscountedSalePrice                 int `json:"discountedSalePrice"`
	MobileDiscountedSalePrice           int `json:"mobileDiscountedSalePrice"`
	MobileSellerImmediateDiscountAmount int `json:"mobileSellerImmediateDiscountAmount"`
	DiscountedRatio                     int `json:"discountRatio"`
	MobileDiscountedRatio               int `json:"mobileDiscountedRatio"`
}

type SaleAmountEntity struct {
	CumulationSaleCount int `json:"cumulationSaleCount"`
	RecentSaleCount     int `json:"recentSaleCount"`
}

type ReviewAmountEntity struct {
	TotalReviewCount   int     `json:"totalReviewCount"`
	AverageReviewScore float32 `json:"averageReviewScore"`
}

type DeliveryInfoEntity struct {
	BaseFee               int    `json:"baseFee"`
	FreeConditionalAmount int    `json:"freeConditionalAmount"`
	DeliveryType          string `json:"deliveryAttributeType"`
}

type DetailEntity struct {
	DetailContentText string `json:"detailContentText"`
}

type SmartstoreProductEntity struct {
	Name           string             `json:"name"`
	Imageurl       string             `json:"representativeImageUrl"`
	NaverProductId int                `json:"id"`
	NaverProductNo int                `json:"productNo"`
	SalePrice      int                `json:"salePrice"`
	SaleType       string             `json:"saleType"`
	Benefits       BenefitsEntity     `json:"benefitsView"`
	SaleAmount     SaleAmountEntity   `json:"saleAmount"`
	ReviewAmount   ReviewAmountEntity `json:"reviewAmount"`
	DeliveryInfo   DeliveryInfoEntity `json:"productDeliveryInfo"`
	Detail         DetailEntity       `json:"detailContents"`
	FreeDelivery   bool               `json:"freeDelivery"`
	TodayDelivery  bool               `json:"todayDelivery"`
}

type SmartstoreResponseParser struct {
	Products   []SmartstoreProductEntity `json:"simpleProducts"`
	TotalCount int                       `json:"totalCount"`
	Page       string                    `json:"page"`
}

type NaverProductEntity struct {
	Name           string          `json:"productName"`
	Title          string          `json:"productTitle"`
	TitleOrg       string          `json:"productTitleOrg"`
	AttributeValue string          `json:"attributeValue"`
	CharacterValue string          `json:"characterValue"`
	ImageUrl       string          `json:"imageUrl"`
	LowPrice       int             `json:"lowPrice"`
	Price          string          `json:"price"`
	PriceUnit      string          `json:"priceUnit"`
	Maker          string          `json:"maker"`
	Brand          string          `json:"brand"`
	Rank           int             `json:"rank"`
	Category1Name  string          `json:"category1Name"`
	Category2Name  string          `json:"category2Name"`
	Category3Name  string          `json:"category3Name"`
	Category4Name  string          `json:"category4Name"`
	MallName       string          `json:"mallName"`
	MallNameOrg    string          `json:"mallNameOrg"`
	MallProductUrl string          `json:"mallProductUrl"`
	MallProdMblUrl string          `json:"mallProdMblUrl"`
	DeliveryFee    string          `json:"dlvryCont"`
	PurchaseCount  int             `json:"purchaseCnt"`
	ReviewCount    int             `json:"reviewCountSum"`
	KeepCount      int             `json:"keepCnt"`
	MallInfo       NaverMallEntity `json:"mallInfoCache"`
	OpenDate       string          `json:"openDate"`
	CreatedAt      time.Time       `json:"created_at,omitempty" bson:"created_at"`
	MomilBrand     string          `bson:"momilbrand"`
	NaverProductId string          `json:"id" bson:"naver_product_id"`
}

type NaverMallEntity struct {
	Name        string            `json:"name"`
	BizAddr     string            `'json:"bizplBaseAddr"`
	BizNo       string            `json:"businessNo"`
	Description string            `json:"mallIntroduction"`
	LogoUrl     map[string]string `json:"mallLogos"`
	CreatedAt   time.Time         `json:"created_at,omitempty" bson:"created_at"`
}
