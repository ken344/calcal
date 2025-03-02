package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ken344/calcal/pkg/models/nutrition"
	"github.com/olekukonko/tablewriter"
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

	// 栄養素ごとのカロリーの計算
	meal.CalorieCalculation()

	// 合計カロリーの計算
	meal.TotalCaloriesCalculation()

	// 結果を出力
	nutritionTable := tablewriter.NewWriter(os.Stdout)
	nutritionTable.SetHeader([]string{"栄養素", "量\n (g)", "カロリー \n(kcal)"})
	nutritionTable.Append([]string{"蛋白質", fmt.Sprintf("%.1f", meal.Protein.Amount), fmt.Sprintf("%.0f", meal.Protein.Calories)})
	nutritionTable.Append([]string{"脂質", fmt.Sprintf("%.1f", meal.Fat.Amount), fmt.Sprintf("%.0f", meal.Fat.Calories)})
	nutritionTable.Append([]string{"炭水化物", fmt.Sprintf("%.1f", meal.Carbohydrate.Amount), fmt.Sprintf("%.0f", meal.Carbohydrate.Calories)})
	nutritionTable.Append([]string{"", "", ""})
	nutritionTable.Append([]string{"合計カロリー", "", fmt.Sprintf("%.0f", meal.TotalCalories)})
	nutritionTable.Render()

}
