package main

import (
	"log"

	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"go.mongodb.org/mongo-driver/mongo"
)

func addCategories(conn *mongo.Client) {
	categories := []map[string]string{
		{
			"label":            "해물탕",
			"name":             "Hamultang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Hamultang.png",
		},
		{
			"label":            "육개장",
			"name":             "Yukgyejang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Yukgyejang.png",
		},
		{
			"label":            "마라탕",
			"name":             "Maratang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Maratang.png",
		},
		{
			"label":            "된장찌개",
			"name":             "Duonjangzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Duonjangzzigye.png",
		},
		{
			"label":            "김치찌개",
			"name":             "Kimchizzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kimchizzigye.png",
		},
		{
			"label":            "감바스",
			"name":             "Gambas",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gambas.png",
		},
		{
			"label":            "기타 국/탕/전골",
			"name":             "Etcjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etcjeongol.png",
		},
		{
			"label":            "스테이크",
			"name":             "Steak",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Steak.png",
		},
		{
			"label":            "고기",
			"name":             "Gogi",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gogi.png",
		},
		{
			"label":            "어묵탕",
			"name":             "Umooktang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Umooktang.png",
		},
		{
			"label":            "추어탕",
			"name":             "Churtang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Churtang.png",
		},
		{
			"label":            "덮밥/비빔밥",
			"name":             "Bibbimbap",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bibbimbap.png",
		},
		{
			"label":            "곱창전골",
			"name":             "Gobchangjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gobchangjeongol.png",
		},
		{
			"label":            "청국장",
			"name":             "Chunggukjang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Chunggukjang.png",
		},
		{
			"label":            "부대찌개",
			"name":             "Budaezzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Budaezzigye.png",
		},
		{
			"label":            "기타 요리",
			"name":             "Etc",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etc.png",
		},
		{
			"label":            "알탕",
			"name":             "Altang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Altang.png",
		},
		{
			"label":            "면",
			"name":             "Myun",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Myun.png",
		},
		{
			"label":            "밀푀유나베",
			"name":             "Millefeuille",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Millefeuille.png",
		},
		{
			"label":            "우거지탕",
			"name":             "Uguzytang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Uguzytang.png",
		},
		{
			"label":            "분식",
			"name":             "Bunsik",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bunsik.png",
		},
		{
			"label":            "파스타",
			"name":             "Pasta",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Pasta.png",
		},
		{
			"label":            "순두부찌개",
			"name":             "Sundubuzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
		},
		{
			"label":            "콩비지찌개",
			"name":             "Kongbeasyzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kongbeasyzzigye.png",
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
	brands := []map[string]string{
		{
			"name": "맛수러움",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/matsuruum.png",
		},
		{
			"name":          "프레시지",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/fresheasy.png",
		},
		{
			"name": "프레시몰",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/freshmeal.png",
		},
		{
			"name":          "푸드어셈블",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/foodassemble.png",
		},
		{
			"name": "네이쳐푸드",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/naturefood.png",
		},
		{
			"name":          "얌테이블",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/yamtable.png",
		},
		{
			"name":          "이츠웰",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/itswell.png",
		},
		{
			"name": "쿡솜씨",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/cooksomssi.png",
		},
		{
			"name": "피콕",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/peacock.png",
		},
		{
			"name": "마이셰프",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/mychef.png",
		},
		{
			"name": "에슐리",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/ashley.png",
		},
		{
			"name":          "파우즈",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/pause.png",
		},
		{
			"name":          "앙트레",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/entree.png",
		},
		{
			"name": "올쿡",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/allcook.png",
		},
		{
			"name": "모노키친",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/monokitchen.png",
		},
		{
			"name": "닥터키친",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/drkitchen.png",
		},
		{
			"name":          "파파쿡",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/papacook.png",
		},
		{
			"name": "기타",
			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
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
