package seeds

import (
	"log"
	"strconv"

	infra "github.com/lessbutter/mealkit/src"
	"github.com/lessbutter/mealkit/src/model"
)

func AddCategories() {
	categories := []model.Category{
		{
			Label:            "김치찌개",
			Name:             "Kimchizzigye",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kimchizzigye.png",
			Order:            0,
			Onmain:           true,
		},
		{
			Label:            "순두부찌개",
			Name:             "Sundubuzzigye",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
			Order:            1,
			Onmain:           true,
		},
		{
			Label:            "된장찌개",
			Name:             "Duonjangzzigye",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Duonjangzzigye.png",
			Order:            2,
			Onmain:           true,
		},
		{
			Label:            "부대찌개",
			Name:             "Budaezzigye",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Budaezzigye.png",
			Order:            3,
			Onmain:           true,
		},
		{
			Label:            "청국장",
			Name:             "Chunggukjang",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Chunggukjang.png",
			Order:            4,
			Onmain:           false,
		},
		{
			Label:            "마라탕",
			Name:             "Maratang",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Maratang.png",
			Order:            5,
			Onmain:           true,
		},
		{
			Label:            "샤브샤브",
			Name:             "Shabshab",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Shabshab.png",
			Order:            6,
			Onmain:           true,
		},
		{
			Label:            "닭갈비",
			Name:             "Dakkalbi",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Dakkalbi.png",
			Order:            7,
			Onmain:           false,
		},
		{
			Label:            "불고기전골",
			Name:             "Bulgogijeongol",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bulgogijeongol.png",
			Order:            8,
			Onmain:           false,
		},
		{
			Label:            "곱창전골",
			Name:             "Gobchangjeongol",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gobchangjeongol.png",
			Order:            9,
			Onmain:           false,
		},
		{
			Label:            "밀푀유나베",
			Name:             "Millefeuille",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Millefeuille.png",
			Order:            10,
			Onmain:           true,
		},
		{
			Label:            "해물탕",
			Name:             "Hamultang",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Hamultang.png",
			Order:            11,
			Onmain:           true,
		},
		{
			Label:            "볶음요리",
			Name:             "Bokum",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bokum.png",
			Order:            12,
			Onmain:           false,
		},
		{
			Label:            "찜요리",
			Name:             "Zzim",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Zzim.png",
			Order:            13,
			Onmain:           false,
		},
		{
			Label:            "돈까스",
			Name:             "Donkkas",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Donkkas.png",
			Order:            14,
			Onmain:           false,
		},
		{
			Label:            "기타 국/탕/전골",
			Name:             "Etcjeongol",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etcjeongol.png",
			Order:            15,
			Onmain:           false,
		},
		{
			Label:            "덮밥/비빔밥",
			Name:             "Bibbimbap",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bibbimbap.png",
			Order:            16,
			Onmain:           true,
		},
		{
			Label:            "면",
			Name:             "Myun",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Myun.png",
			Order:            17,
			Onmain:           true,
		},
		{
			Label:            "스테이크",
			Name:             "Steak",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Steak.png",
			Order:            18,
			Onmain:           true,
		},
		{
			Label:            "고기",
			Name:             "Gogi",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gogi.png",
			Order:            19,
			Onmain:           true,
		},
		{
			Label:            "파스타",
			Name:             "Pasta",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Pasta.png",
			Order:            20,
			Onmain:           true,
		},
		{
			Label:            "감바스",
			Name:             "Gambas",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gambas.png",
			Order:            21,
			Onmain:           true,
		},
		{
			Label:            "분식",
			Name:             "Bunsik",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bunsik.png",
			Order:            22,
			Onmain:           true,
		},
		{
			Label:            "짜글이",
			Name:             "Jjagle",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Jjagle.png",
			Order:            23,
			Onmain:           false,
		},
		{
			Label:            "기타 요리",
			Name:             "Etc",
			Categoryimageurl: "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etc.png",
			Order:            24,
			Onmain:           false,
		},
	}

	infra.AddCategories(categories)
	log.Println(strconv.Itoa(len(categories)) + " Categories Addition Finished")
}
