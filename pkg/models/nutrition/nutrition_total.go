package nutrition

import (
	"fmt"
	"os"
	"strings"

	"github.com/ken344/calcal/pkg/models/data"
)

const filePath = "../../data/meals.csv"

func SumNutrients(mealsInput string) Nutrient {

	// CSVファイルの読み込み
	mealDB, err := data.LoadCSV(filePath)
	if err != nil {
		fmt.Println("CSVファイルの読み込みに失敗しました:", err)
		os.Exit(1)
	}

	// カンマ区切りの文字列を分割
	meals := strings.Split(mealsInput, ",")

	


	// 食品名をキーにして栄養素を取得
	nutrient, exists := mealDB[meal]
	if !exists {
		fmt.Println("指定された食品名は存在しません")
		os.Exit(1)
	}

	// 食品名をキーにして栄養素を取得
	for i, meal := range meals {

	}

	// 栄養素のスライスを作成
	var nutrients []Nutrient
	for _, meal := range meals {

	}

	// 食品名をキーにして栄養素を取得
	nutrient, exists := mealData[meal]
	if !exists {
		fmt.Println("指定された食品名は存在しません")
		os.Exit(1)
	}

	// 栄養素の表示
	fmt.Printf("食品名: %s\n", meal)
	fmt.Printf("蛋白質: %.1fg\n", nutrient.Protein)
	fmt.Printf("脂質: %.1fg\n", nutrient.Fat)
	fmt.Printf("炭水化物: %.1fg\n", nutrient.Carbohydrate)

	var sum Nutrient
	for _, n := range nutrients {
		sum.Protein += n.Protein
		sum.Fat += n.Fat
		sum.Carbohydrate += n.Carbohydrate
	}
	return sum
}

