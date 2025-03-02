package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ken344/calcal/pkg/models/data"
	"github.com/ken344/calcal/pkg/models/nutrition"
	"github.com/olekukonko/tablewriter"
)

const dataFilePath = "../../data/meals.csv"

func main() {

	// コマンドライン引数の取得
	// 最低ひとつの引数（食品名）が必要
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go \"ご飯,納豆\"")
		os.Exit(1)
	}

	// 引数を食品名として取得
	mealsInput := os.Args[1]
	// fmt.Printf("mealsInput: %v %T\n", mealsInput, mealsInput)

	// 引数をカンマ区切りで分割し、スライスに格納
	meals := strings.Split(mealsInput, ",")
	// fmt.Printf("meals: %v %T\n", meals, meals)

	// CSVファイルの読み込み(map型に変換)
	mealDB, err := data.LoadCSV(dataFilePath)
	if err != nil {
		fmt.Println("CSVファイルの読み込みに失敗しました:", err)
		os.Exit(1)
	}
	// fmt.Printf("mealDB: %v %T\n", mealDB, mealDB)

	// 食品のデータを集約するスライスを作成
	var mealDatas []nutrition.MealData

	// 食品のデータを送受信するチャネルを作成
	mealChan := make(chan nutrition.MealData, len(meals))

	// sync.WaitGroupを作成ß
	wg := sync.WaitGroup{}
	wg.Add(len(meals))

	// 食品名をキーにして栄養素を取得
	for _, meal := range meals {
		go func(meal string) {
			defer wg.Done()

			// fmt.Printf("meal: %v %T\n", meal, meal)

			// 食品名をキーにして栄養素を取得
			nutrient, exists := mealDB[meal]
			if !exists {
				// fmt.Printf("食品名「%s」は存在しません\n", meal)
				mealChan <- nutrition.MealData{MealName: meal, Info: "DBに存在しません"}
				return
			}

			// 栄養素の構造体を生成
			mealTotal := nutrition.NewMacronutrients(nutrient.Protein, nutrient.Fat, nutrient.Carbohydrate)
			// 栄養素ごとのカロリーの計算
			mealTotal.CalorieCalculation()
			// 合計カロリーの計算
			mealTotal.TotalCaloriesCalculation()

			// 構造体をチャネルに送信
			mealChan <- nutrition.MealData{MealName: meal, MacronutrientsData: mealTotal, Info: "-"}
		}(meal)

	}

	// sync.WaitGroupを待機
	wg.Wait()
	// チャネルを閉じる
	close(mealChan)

	// チャネルから構造体を取得し、スライスに追加
	for meal := range mealChan {
		mealDatas = append(mealDatas, meal)
	}

	// 食品のデータを集約
	sumMealData := nutrition.MealAggregation(mealDatas)
	// fmt.Printf("sumMealData: %v %T\n", sumMealData, sumMealData)

	// 食品のデータを表示
	meal_table := tablewriter.NewWriter(os.Stdout)
	meal_table.SetHeader([]string{"食品名", "蛋白質\n(g)", "蛋白質\n(kcal)", "脂質\n(g)", "脂質\n(kcal)", "炭水化物\n(g)", "炭水化物\n(kcal)", "合計カロリー\n(kcal)", "info"})
	for _, meal := range mealDatas {
		meal_table.Append([]string{
			meal.MealName,
			fmt.Sprintf("%.1f", meal.MacronutrientsData.Protein.Amount),
			fmt.Sprintf("%.0f", meal.MacronutrientsData.Protein.Calories),
			fmt.Sprintf("%.1f", meal.MacronutrientsData.Fat.Amount),
			fmt.Sprintf("%.0f", meal.MacronutrientsData.Fat.Calories),
			fmt.Sprintf("%.1f", meal.MacronutrientsData.Carbohydrate.Amount),
			fmt.Sprintf("%.0f", meal.MacronutrientsData.Carbohydrate.Calories),
			fmt.Sprintf("%.0f", meal.MacronutrientsData.TotalCalories),
			meal.Info,
		})
	}

	// 食品と合計値の間に空行を追加
	meal_table.Append([]string{"", "", "", "", "", "", "", "", ""})

	// 合計行を追加
	meal_table.Append([]string{
		"総計",
		fmt.Sprintf("%.1f", sumMealData.Protein.Amount),
		fmt.Sprintf("%.0f", sumMealData.Protein.Calories),
		fmt.Sprintf("%.1f", sumMealData.Fat.Amount),
		fmt.Sprintf("%.0f", sumMealData.Fat.Calories),
		fmt.Sprintf("%.1f", sumMealData.Carbohydrate.Amount),
		fmt.Sprintf("%.0f", sumMealData.Carbohydrate.Calories),
		fmt.Sprintf("%.0f", sumMealData.TotalCalories),
		"-",
	})

	// 表を表示
	meal_table.Render()

}
