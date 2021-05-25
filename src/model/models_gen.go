// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CrawlingRecord struct {
	NewProducts int       `json:"newproducts"`
	OutProducts int       `json:"outproducts"`
	Date        time.Time `json:"date"`
}

type Brand struct {
	ID                  string `json:"ID" bson:"_id,omitempty"`
	Name                string `json:"name" bson:"name"`
	Order               int    `json:"order" bson:"order"`
	Description         string `json:"description" bson:"description,omitempty"`
	Brandimageurl       string `json:"brandimageurl" bson:"brandimageurl"`
	Onmain              bool   `json:"onmain" bson:"onmain"`
	CrawlingUrl         string `json:"CrawlingUrl" bson:"crawlingurl,omitempty"`
	CrawlFrom           string `json:"crawlform" bson:"crawlfrom"`
	SmartstoreBrandName string `json:"smartstorebrandname" bson:"smartstorebrandname"`
}

type Category struct {
	ID               string       `json:"ID" bson:"_id,omitempty"`
	Label            string       `json:"label"`
	Name             CategoryEnum `json:"name"`
	Order            int          `json:"order"`
	Categoryimageurl string       `json:"categoryimageurl"`
	Onmain           bool         `json:"onmain"`
}

type Product struct {
	ID              string    `json:"ID" bson:"_id,omitempty"`
	Name            string    `json:"name"`
	Imageurl        string    `json:"imageurl"`
	Price           int       `json:"price" bson:"price"`
	Discountedprice int       `json:"discountedprice"`
	Brand           *Brand    `json:"brand" bson:"brand"`
	Producturl      string    `json:"mallproducturl" bson:"producturl"`
	Deliveryfee     string    `json:"deliveryfee"`
	Category        *Category `json:"category" bson:"category"`
	Purchasecount   int       `json:"purchasecount"`
	Reviewcount     int       `json:"reviewcount"`
	Reviewscore     float64   `json:"reviewscore"`
	Mallname        string    `json:"mallname"`
	Originalid      string    `json:"originalid"`
	Soldout         bool      `json:"soldout"`
	Removed         bool      `json:"removed"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
	IsNew           bool      `json:"isnew"`
}

type ProductsInput struct {
	Offset   int           `json:"offset"`
	Limit    int           `json:"limit"`
	Category *CategoryEnum `json:"category"`
	Brand    *string       `json:"brand"`
	Search   *string       `json:"search"`
}

type CategoryEnum string

const (
	CategoryEnumHamultang       CategoryEnum = "Hamultang"
	CategoryEnumYukgyejang      CategoryEnum = "Yukgyejang"
	CategoryEnumMaratang        CategoryEnum = "Maratang"
	CategoryEnumDuonjangzzigye  CategoryEnum = "Duonjangzzigye"
	CategoryEnumKimchizzigye    CategoryEnum = "Kimchizzigye"
	CategoryEnumGambas          CategoryEnum = "Gambas"
	CategoryEnumEtcjeongol      CategoryEnum = "Etcjeongol"
	CategoryEnumSteak           CategoryEnum = "Steak"
	CategoryEnumGogi            CategoryEnum = "Gogi"
	CategoryEnumUmooktang       CategoryEnum = "Umooktang"
	CategoryEnumChurtang        CategoryEnum = "Churtang"
	CategoryEnumBibbimbap       CategoryEnum = "Bibbimbap"
	CategoryEnumGobchangjeongol CategoryEnum = "Gobchangjeongol"
	CategoryEnumChunggukjang    CategoryEnum = "Chunggukjang"
	CategoryEnumBudaezzigye     CategoryEnum = "Budaezzigye"
	CategoryEnumEtc             CategoryEnum = "Etc"
	CategoryEnumAltang          CategoryEnum = "Altang"
	CategoryEnumMyun            CategoryEnum = "Myun"
	CategoryEnumMillefeuille    CategoryEnum = "Millefeuille"
	CategoryEnumUguzytang       CategoryEnum = "Uguzytang"
	CategoryEnumBunsik          CategoryEnum = "Bunsik"
	CategoryEnumPasta           CategoryEnum = "Pasta"
	CategoryEnumSundubuzzigye   CategoryEnum = "Sundubuzzigye"
	CategoryEnumKongbeasyzzigye CategoryEnum = "Kongbeasyzzigye"
	CategoryEnumBokumzzim       CategoryEnum = "Bokumzzim"
	CategoryEnumJjagle          CategoryEnum = "Jjagle"
	CategoryEnumShabshab        CategoryEnum = "Shabshab"
	CategoryEnumBulgogijeongol  CategoryEnum = "Bulgogijeongol"
	CategoryEnumDakkalbi        CategoryEnum = "Dakkalbi"
	CategoryEnumBokum           CategoryEnum = "Bokum"
	CategoryEnumZzim            CategoryEnum = "Zzim"
	CategoryEnumDonkkas         CategoryEnum = "Donkkas"
)

var AllCategoryEnum = []CategoryEnum{
	CategoryEnumHamultang,
	CategoryEnumYukgyejang,
	CategoryEnumMaratang,
	CategoryEnumDuonjangzzigye,
	CategoryEnumKimchizzigye,
	CategoryEnumGambas,
	CategoryEnumEtcjeongol,
	CategoryEnumSteak,
	CategoryEnumGogi,
	CategoryEnumUmooktang,
	CategoryEnumChurtang,
	CategoryEnumBibbimbap,
	CategoryEnumGobchangjeongol,
	CategoryEnumChunggukjang,
	CategoryEnumBudaezzigye,
	CategoryEnumEtc,
	CategoryEnumAltang,
	CategoryEnumMyun,
	CategoryEnumMillefeuille,
	CategoryEnumUguzytang,
	CategoryEnumBunsik,
	CategoryEnumPasta,
	CategoryEnumSundubuzzigye,
	CategoryEnumKongbeasyzzigye,
	CategoryEnumBokumzzim,
	CategoryEnumJjagle,
	CategoryEnumShabshab,
	CategoryEnumBulgogijeongol,
	CategoryEnumDakkalbi,
	CategoryEnumBokum,
	CategoryEnumZzim,
	CategoryEnumDonkkas,
}

func (e CategoryEnum) IsValid() bool {
	switch e {
	case CategoryEnumHamultang, CategoryEnumYukgyejang, CategoryEnumMaratang, CategoryEnumDuonjangzzigye, CategoryEnumKimchizzigye, CategoryEnumGambas, CategoryEnumEtcjeongol, CategoryEnumSteak, CategoryEnumGogi, CategoryEnumUmooktang, CategoryEnumChurtang, CategoryEnumBibbimbap, CategoryEnumGobchangjeongol, CategoryEnumChunggukjang, CategoryEnumBudaezzigye, CategoryEnumEtc, CategoryEnumAltang, CategoryEnumMyun, CategoryEnumMillefeuille, CategoryEnumUguzytang, CategoryEnumBunsik, CategoryEnumPasta, CategoryEnumSundubuzzigye, CategoryEnumKongbeasyzzigye, CategoryEnumBokumzzim, CategoryEnumJjagle, CategoryEnumShabshab, CategoryEnumBulgogijeongol, CategoryEnumDakkalbi, CategoryEnumBokum, CategoryEnumZzim, CategoryEnumDonkkas:
		return true
	}
	return false
}

func (e CategoryEnum) String() string {
	return string(e)
}

func (e *CategoryEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CategoryEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CategoryEnum", str)
	}
	return nil
}

func (e CategoryEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
