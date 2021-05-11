package main

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"lessbutter.co/mealkit/config"
	product "lessbutter.co/mealkit/domains"
	"lessbutter.co/mealkit/external"
)

func addCategories(conn *mongo.Client) {
	categories := []map[string]string{
		{
			"name":             "해물탕",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Hamultang.png",
		},
		{
			"name":             "육계장",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Yukgyejang.png",
		},
		{
			"name":             "마라탕",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Maratang.png",
		},
		{
			"name":             "된장찌개",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Duonjangzzigye.png",
		},
		{
			"name":             "김치찌개",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kimchizzigye.png",
		},
		{
			"name":             "감바스",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gambas.png",
		},
		{
			"name":             "기타 국/탕/전골",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etcjeongol.png",
		},
		{
			"name":             "스테이크",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Steak.png",
		},
		{
			"name":             "고기",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gogi.png",
		},
		{
			"name":             "어묵탕",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Umooktang.png",
		},
		{
			"name":             "추어탕",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Churtang.png",
		},
		{
			"name":             "덮밥/비빔밥",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bibbimbap.png",
		},
		{
			"name":             "곱창전골",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gobchangjeongol.png",
		},
		{
			"name":             "청국장",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Chunggukjang.png",
		},
		{
			"name":             "부대찌개",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Budaezzigye.png",
		},
		{
			"name":             "기타 요리",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etc.png",
		},
		{
			"name":             "알탕",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Altang.png",
		},
		{
			"name":             "면",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Myun.png",
		},
		{
			"name":             "밀푀유나베",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Millefeuille.png",
		},
		{
			"name":             "우거지탕",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Uguzytang.png",
		},
		{
			"name":             "분식",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bunsik.png",
		},
		{
			"name":             "파스타",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Pasta.png",
		},
		{
			"name":             "순두부찌개",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
		},
		{
			"name":             "콩비지찌개",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kongbeasyzzigye.png",
		},
	}

	cats := make([]interface{}, 0)
	for _, result := range categories {
		cats = append(cats, result)
	}
	log.Println(cats)
	ret, _ := product.AddCategories(conn, cats)
	log.Println(ret)
}

// func addBrands(conn *mongo.Client) {
// 	brands := []map[string]string{
// 		{
// 			"name": "맛수러움",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/matsuruum.png",
// 		},
// 		{
// 			"name":          "프레시지",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/fresheasy.png",
// 		},
// 		{
// 			"name": "프레시몰",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/freshmeal.png",
// 		},
// 		{
// 			"name":          "푸드어셈블",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/foodassemble.png",
// 		},
// 		{
// 			"name": "네이쳐푸드",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/naturefood.png",
// 		},
// 		{
// 			"name":          "얌테이블",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/yamtable.png",
// 		},
// 		{
// 			"name":          "이츠웰",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/itswell.png",
// 		},
// 		{
// 			"name": "쿡솜씨",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/cooksomssi.png",
// 		},
// 		{
// 			"name": "피콕",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/peacock.png",
// 		},
// 		{
// 			"name": "마이셰프",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/mychef.png",
// 		},
// 		{
// 			"name": "에슐리",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/ashley.png",
// 		},
// 		{
// 			"name":          "파우즈",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/pause.png",
// 		},
// 		{
// 			"name":          "앙트레",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/entree.png",
// 		},
// 		{
// 			"name": "올쿡",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/allcook.png",
// 		},
// 		{
// 			"name": "모노키친",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/monokitchen.png",
// 		},
// 		{
// 			"name": "닥터키친",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/drkitchen.png",
// 		},
// 		{
// 			"name":          "파파쿡",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/papacook.png",
// 		},
// 		{
// 			"name": "기타",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
// 		},
// 	}
// 	brandsInterface := make([]interface{}, 0)
// 	for _, result := range brands {
// 		brandsInterface = append(brandsInterface, result)
// 	}

// 	ret, _ := product.AddCategories(conn, brandsInterface)
// 	log.Println(ret)

// }

func main() {
	conf := config.GetConfiguration()
	conn := external.MongoConn(conf)
	addCategories(conn)
	// addBrands(conn)

}
