package main

import (
	"fmt"
	"os"
	"strings"
)

const filePath = "../../data/meals.csv"

func main() {

	// コマンドライン引数の取得
	// 最低ひとつの引数（食品名）が必要
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go \"ご飯,納豆\"")
		os.Exit(1)
	}

	// 引数を食品名として取得
	mealsInput := os.Args[1]
	fmt.Printf("mealsInput: %v %T\n", mealsInput, mealsInput)

	// カンマ区切りで分割
	meals := strings.Split(mealsInput, ",")
	fmt.Printf("meals: %v %T\n", meals, meals)

	// // CSVファイルの読み込み
	// mealData, err := data.LoadCSV(filePath)
	// if err != nil {
	// 	fmt.Println("CSVファイルの読み込みに失敗しました:", err)
	// 	os.Exit(1)
	// }

	// 食品名をキーにして栄養素を取得
	// nutrient, exists := mealData[meal]
	// if !exists {
	// 	fmt.Println("指定された食品名は存在しません")
	// 	os.Exit(1)
	// }

	// // 栄養素の表示
	// fmt.Printf("食品名: %s\n", meal)
	// fmt.Printf("蛋白質: %.1fg\n", nutrient.Protein)
	// fmt.Printf("脂質: %.1fg\n", nutrient.Fat)
	// fmt.Printf("炭水化物: %.1fg\n", nutrient.Carbohydrate)

}
