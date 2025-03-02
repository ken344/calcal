package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ken344/calcal/pkg/models/nutrition"
)

func main() {

	// オプション用フラグの定義
	p := flag.Float64("p", 0, "蛋白質の量 (g)")
	f := flag.Float64("f", 0, "脂質の量 (g)")
	c := flag.Float64("c", 0, "炭水化物の量 (g)")

	// フラグの解析
	flag.Parse()

	// フラグが1つも指定されていない場合はヘルプを表示して終了
	if flag.NFlag() == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// 構造体の生成
	meal := nutrition.NewMacronutrients(*p, *f, *c)

	// カロリーの計算
	meal.CalorieCalculation()

	// 合計カロリーの計算
	meal.TotalCaloriesCalculation()

	// 結果を出力
	fmt.Printf("蛋白質: %.1fg (%.0fkcal)\n", meal.Protein.Amount, meal.Protein.Calories)
	fmt.Printf("脂質: %.1fg (%.0fkcal)\n", meal.Fat.Amount, meal.Fat.Calories)
	fmt.Printf("炭水化物: %.1fg (%.0fkcal)\n", meal.Carbohydrate.Amount, meal.Carbohydrate.Calories)
	fmt.Printf("合計カロリー: %.0fkcal\n", meal.TotalCalories)
}
