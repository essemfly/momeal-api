package main

import (
	"log"

	"github.com/lessbutter/mealkit/config"
	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func addCategories(conn *mongo.Client) {
	categories := []map[string]interface{}{
		{
			"label":            "해물탕",
			"name":             "Hamultang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Hamultang.png",
			"order":            11,
			"onmain":           false,
		},
		{
			"label":            "육개장",
			"name":             "Yukgyejang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Yukgyejang.png",
			"order":            6,
			"onmain":           false,
		},
		{
			"label":            "마라탕",
			"name":             "Maratang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Maratang.png",
			"order":            7,
			"onmain":           true,
		},
		{
			"label":            "된장찌개",
			"name":             "Duonjangzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Duonjangzzigye.png",
			"order":            2,
			"onmain":           true,
		},
		{
			"label":            "김치찌개",
			"name":             "Kimchizzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kimchizzigye.png",
			"order":            0,
			"onmain":           true,
		},
		{
			"label":            "감바스",
			"name":             "Gambas",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gambas.png",
			"order":            21,
			"onmain":           true,
		},
		{
			"label":            "기타 국/탕/전골",
			"name":             "Etcjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etcjeongol.png",
			"order":            15,
			"onmain":           false,
		},
		{
			"label":            "스테이크",
			"name":             "Steak",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Steak.png",
			"order":            18,
			"onmain":           true,
		},
		{
			"label":            "고기",
			"name":             "Gogi",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gogi.png",
			"order":            19,
			"onmain":           true,
		},
		{
			"label":            "어묵탕",
			"name":             "Umooktang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Umooktang.png",
			"order":            8,
			"onmain":           true,
		},
		{
			"label":            "덮밥/비빔밥",
			"name":             "Bibbimbap",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bibbimbap.png",
			"order":            16,
			"onmain":           true,
		},
		{
			"label":            "곱창전골",
			"name":             "Gobchangjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gobchangjeongol.png",
			"order":            9,
			"onmain":           false,
		},
		{
			"label":            "청국장",
			"name":             "Chunggukjang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Chunggukjang.png",
			"order":            5,
			"onmain":           false,
		},
		{
			"label":            "부대찌개",
			"name":             "Budaezzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Budaezzigye.png",
			"order":            3,
			"onmain":           true,
		},
		{
			"label":            "기타 요리",
			"name":             "Etc",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etc.png",
			"order":            24,
			"onmain":           false,
		},
		{
			"label":            "알탕",
			"name":             "Altang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Altang.png",
			"order":            11,
			"onmain":           false,
		},
		{
			"label":            "면",
			"name":             "Myun",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Myun.png",
			"order":            17,
			"onmain":           true,
		},
		{
			"label":            "밀푀유나베",
			"name":             "Millefeuille",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Millefeuille.png",
			"order":            10,
			"onmain":           true,
		},
		{
			"label":            "우거지탕",
			"name":             "Uguzytang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Uguzytang.png",
			"order":            13,
			"onmain":           false,
		},
		{
			"label":            "분식",
			"name":             "Bunsik",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bunsik.png",
			"order":            22,
			"onmain":           true,
		},
		{
			"label":            "파스타",
			"name":             "Pasta",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Pasta.png",
			"order":            20,
			"onmain":           true,
		},
		{
			"label":            "순두부찌개",
			"name":             "Sundubuzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
			"order":            1,
			"onmain":           true,
		},
		{
			"label":            "볶음, 찜요리",
			"name":             "Bokumzzim",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
			"order":            23,
			"onmain":           true,
		},
	}

	cats := make([]interface{}, 0)
	for _, result := range categories {
		cats = append(cats, result)
	}
	ret, _ := infra.AddCategories(conn, cats)
	log.Println(ret)
}

func addBrands(conn *mongo.Client) {
	brands := []model.Brand{
		{
			Name:                "프레시지",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/fresheasy.png",
			Order:               0,
			Onmain:              true,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "fresheasy",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/500230195/categories/36e7f87fa6e24e19b30fd5b0c2fae4bb/products?categoryId=36e7f87fa6e24e19b30fd5b0c2fae4bb&categorySearchType=DISPCATG",
		},
		{
			Name:                "마이셰프",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/mychef.png",
			Order:               1,
			Onmain:              true,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "recipebox_mychef",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/500239930/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "프레시밀",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/freshmeal.png",
			Order:               2,
			Onmain:              true,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "fresh-meal",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/101066654/categories/488cad1747f34b01b13a1021cf19230c/products?categoryId=488cad1747f34b01b13a1021cf19230c&categorySearchType=DISPCATG",
		},
		{
			Name:                "푸드어셈블",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/foodassemble.png",
			Order:               3,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "foodasb",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/100029405/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "앙트레",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/entree.png",
			Order:               4,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "entree_organico",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/100009290/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "파파쿡",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/papacook.png",
			Order:               5,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "friendsfood",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/500289590/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "애슐리",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/ashley.png",
			Order:               6,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "elandparkfood",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/100002411/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "쿡솜씨",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/cooksomssi.png",
			Order:               7,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "cooksomssi",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/100179980/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "올쿡",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/allcook.png",
			Order:               8,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "allcook_store",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/100885144/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "파우즈",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/pause.png",
			Order:               12,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "pausemealkit",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/101089550/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "맛수러움",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/matsuruum.png",
			Order:               16,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "bangpofishery",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/500149336/categories/a015ed45aa724cb39516a0def00a1bdb/products?categoryId=a015ed45aa724cb39516a0def00a1bdb&categorySearchType=DISPCATG",
		},
		{
			Name:                "자연맛남",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/naturefood.png",
			Order:               13,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "ifoodmall",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/500057665/categories/92ed12de0bbd41b99eb9a7bd15a06ab7/products?categoryId=92ed12de0bbd41b99eb9a7bd15a06ab7&categorySearchType=DISPCATG",
		},
		{
			Name:                "아로이키친",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/rama.png",
			Order:               18,
			Onmain:              false,
			CrawlFrom:           "smartstore",
			SmartstoreBrandName: "rama",
			CrawlingUrl:         "https://smartstore.naver.com/i/v1/stores/100875480/categories/ALL/products?categoryId=ALL&categorySearchType=DISPCATG",
		},
		{
			Name:                "쿡킷",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
			Order:               18,
			Onmain:              false,
			CrawlFrom:           "self",
			SmartstoreBrandName: "",
			CrawlingUrl:         "https://www.cjcookit.com/pc/menu/menuProdList.json",
		},
		{
			Name:                "피코크",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/peacock.png",
			Order:               15,
			Onmain:              false,
			CrawlFrom:           "self",
			SmartstoreBrandName: "",
			CrawlingUrl:         "http://emart.ssg.com/specialStore/ssgpeacock/ajaxSubItemList.ssg?isCornrItemShow=false&pageSize=80&ctgId=6000073847",
		},
		{
			Name:                "얌테이블",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/yamtable.png",
			Order:               11,
			Onmain:              false,
			CrawlFrom:           "self",
			SmartstoreBrandName: "",
			CrawlingUrl:         "",
		},
		{
			Name:                "CJ 이츠웰",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/itswell.png",
			Order:               10,
			Onmain:              false,
			CrawlFrom:           "",
			SmartstoreBrandName: "",
			CrawlingUrl:         "",
		},
		{
			Name:                "모노키친",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/monokitchen.png",
			Order:               14,
			Onmain:              false,
			CrawlFrom:           "",
			SmartstoreBrandName: "",
			CrawlingUrl:         "",
		},
		{
			Name:                "닥터키친",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/drkitchen.png",
			Order:               9,
			Onmain:              false,
			CrawlFrom:           "",
			SmartstoreBrandName: "",
			CrawlingUrl:         "",
		},
		{
			Name:                "심플리쿡",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/simplycook.png",
			Order:               17,
			Onmain:              false,
			CrawlFrom:           "",
			SmartstoreBrandName: "",
			CrawlingUrl:         "",
		},
		{
			Name:                "테이스티나인",
			Brandimageurl:       "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/tasty9.png",
			Order:               19,
			Onmain:              false,
			CrawlFrom:           "",
			SmartstoreBrandName: "",
			CrawlingUrl:         "",
		},
	}

	infra.AddBrands(conn, brands)
}

func main() {
	conf := config.GetConfiguration()
	conn := infra.MongoConn(conf)
	// addCategories(conn)
	addBrands(conn)
}
