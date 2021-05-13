package main

import (
	"log"

	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"go.mongodb.org/mongo-driver/mongo"
)

func addCategories(conn *mongo.Client) {
	categories := []map[string]interface{}{
		{
			"label":            "해물탕",
			"name":             "Hamultang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Hamultang.png",
			"order":            11,
		},
		{
			"label":            "육개장",
			"name":             "Yukgyejang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Yukgyejang.png",
			"order":            6,
		},
		{
			"label":            "마라탕",
			"name":             "Maratang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Maratang.png",
			"order":            7,
		},
		{
			"label":            "된장찌개",
			"name":             "Duonjangzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Duonjangzzigye.png",
			"order":            2,
		},
		{
			"label":            "김치찌개",
			"name":             "Kimchizzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kimchizzigye.png",
			"order":            0,
		},
		{
			"label":            "감바스",
			"name":             "Gambas",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gambas.png",
			"order":            21,
		},
		{
			"label":            "기타 국/탕/전골",
			"name":             "Etcjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etcjeongol.png",
			"order":            15,
		},
		{
			"label":            "스테이크",
			"name":             "Steak",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Steak.png",
			"order":            18,
		},
		{
			"label":            "고기",
			"name":             "Gogi",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gogi.png",
			"order":            19,
		},
		{
			"label":            "어묵탕",
			"name":             "Umooktang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Umooktang.png",
			"order":            8,
		},
		{
			"label":            "추어탕",
			"name":             "Churtang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Churtang.png",
			"order":            14,
		},
		{
			"label":            "덮밥/비빔밥",
			"name":             "Bibbimbap",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bibbimbap.png",
			"order":            16,
		},
		{
			"label":            "곱창전골",
			"name":             "Gobchangjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gobchangjeongol.png",
			"order":            9,
		},
		{
			"label":            "청국장",
			"name":             "Chunggukjang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Chunggukjang.png",
			"order":            5,
		},
		{
			"label":            "부대찌개",
			"name":             "Budaezzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Budaezzigye.png",
			"order":            3,
		},
		{
			"label":            "기타 요리",
			"name":             "Etc",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etc.png",
			"order":            23,
		},
		{
			"label":            "알탕",
			"name":             "Altang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Altang.png",
			"order":            11,
		},
		{
			"label":            "면",
			"name":             "Myun",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Myun.png",
			"order":            17,
		},
		{
			"label":            "밀푀유나베",
			"name":             "Millefeuille",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Millefeuille.png",
			"order":            10,
		},
		{
			"label":            "우거지탕",
			"name":             "Uguzytang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Uguzytang.png",
			"order":            13,
		},
		{
			"label":            "분식",
			"name":             "Bunsik",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bunsik.png",
			"order":            22,
		},
		{
			"label":            "파스타",
			"name":             "Pasta",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Pasta.png",
			"order":            20,
		},
		{
			"label":            "순두부찌개",
			"name":             "Sundubuzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
			"order":            1,
		},
		{
			"label":            "콩비지찌개",
			"name":             "Kongbeasyzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kongbeasyzzigye.png",
			"order":            4,
		},
	}

	cats := make([]interface{}, 0)
	for _, result := range categories {
		cats = append(cats, result)
	}
	log.Println(cats)
	ret, _ := infra.AddCategories(conn, cats)
	log.Println(ret)
}

func addBrands(conn *mongo.Client) {
	brands := []map[string]interface{}{
		{
			"name":          "맛수러움",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/matsuruum.png",
			"order":         16,
		},
		{
			"name":          "프레시지",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/fresheasy.png",
			"order":         0,
		},
		{
			"name":          "프레시밀",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/freshmeal.png",
			"order":         2,
		},
		{
			"name":          "푸드어셈블",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/foodassemble.png",
			"order":         3,
		},
		{
			"name":          "자연맛남",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/naturefood.png",
			"order":         13,
		},
		{
			"name":          "얌테이블",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/yamtable.png",
			"order":         11,
		},
		{
			"name":          "CJ 이츠웰",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/itswell.png",
			"order":         10,
		},
		{
			"name":          "쿡솜씨",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/cooksomssi.png",
			"order":         7,
		},
		{
			"name":          "피코크",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/peacock.png",
			"order":         15,
		},
		{
			"name":          "마이셰프",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/mychef.png",
			"order":         1,
		},
		{
			"name":          "에슐리",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/ashley.png",
			"order":         6,
		},
		{
			"name":          "파우즈",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/pause.png",
			"order":         12,
		},
		{
			"name":          "앙트레",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/entree.png",
			"order":         4,
		},
		{
			"name":          "올쿡",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/allcook.png",
			"order":         8,
		},
		{
			"name":          "모노키친",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/monokitchen.png",
			"order":         14,
		},
		{
			"name":          "닥터키친",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/drkitchen.png",
			"order":         9,
		},
		{
			"name":          "파파쿡",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/papacook.png",
			"order":         5,
		},
		{
			"name":          "심플리쿡",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
			"order":         17,
		},
		{
			"name":          "아로이키친",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
			"order":         18,
		},
		{
			"name":          "테이스티나인",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
			"order":         19,
		},
	}
	brandsInterface := make([]interface{}, 0)
	for _, result := range brands {
		brandsInterface = append(brandsInterface, result)
	}

	ret, _ := infra.AddBrands(conn, brandsInterface)
	log.Println(ret)

}

func main() {
	conf := config.GetConfiguration()
	conn := infra.MongoConn(conf)
	addCategories(conn)
	addBrands(conn)
}
