package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ken344/calcal/pkg/models/data"
	"github.com/ken344/calcal/pkg/models/nutrition"
)

const dataFilePath = "../../data/meals.csv"

// mealData構造体（食品の名前と栄養素を格納）
type mealData struct {
	mealName           string
	macronutrientsData nutrition.Macronutrients
}

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

	// 引数をカンマ区切りで分割し、スライスに格納
	meals := strings.Split(mealsInput, ",")
	fmt.Printf("meals: %v %T\n", meals, meals)

	// CSVファイルの読み込み(map型に変換)
	mealDB, err := data.LoadCSV(dataFilePath)
	if err != nil {
		fmt.Println("CSVファイルの読み込みに失敗しました:", err)
		os.Exit(1)
	}
	fmt.Printf("mealDB: %v %T\n", mealDB, mealDB)

	// 食品のデータを集約するスライスを作成
	mealDatas := []mealData{}
	// 食品のデータを送受信するチャネルを作成
	mealChan := make(chan mealData, len(meals))

	// sync.WaitGroupを作成ß
	wg := sync.WaitGroup{}
	wg.Add(len(meals))

	// 食品名をキーにして栄養素を取得
	for _, meal := range meals {
		go func(meal string) {
			fmt.Printf("meal: %v %T\n", meal, meal)
			// 食品名をキーにして栄養素を取得
			nutrient, exists := mealDB[meal]
			if !exists {
				fmt.Printf("食品名「%s」は存在しません\n", meal)
				// fmt.Println("食品名「%s」は存在しません", meal)
				wg.Done()
			}
			// 栄養素の表示
			// fmt.Printf("食品名: %s\n", meal)
			// fmt.Printf("蛋白質: %.1fg\n", nutrient.Protein)
			// fmt.Printf("脂質: %.1fg\n", nutrient.Fat)
			// fmt.Printf("炭水化物: %.1fg\n", nutrient.Carbohydrate)
			// fmt.Print("--------------------\n")

			// 栄養素の合計とカロリーを計算
			mealTotal := nutrition.NewMacronutrients(nutrient.Protein, nutrient.Fat, nutrient.Carbohydrate)

			fmt.Printf("蛋白質: %.1fg (%.0fkcal)\n", mealTotal.Protein.Amount, mealTotal.Protein.Calories)
			fmt.Printf("脂質: %.1fg (%.0fkcal)\n", mealTotal.Fat.Amount, mealTotal.Fat.Calories)
			fmt.Printf("炭水化物: %.1fg (%.0fkcal)\n", mealTotal.Carbohydrate.Amount, mealTotal.Carbohydrate.Calories)
			fmt.Printf("合計カロリー: %.0fkcal\n", mealTotal.TotalCalories())
			fmt.Print("--------------------\n")

			// 構造体をチャネルに送信
			mealChan <- mealData{mealName: meal, macronutrientsData: mealTotal}
			wg.Done()
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

	fmt.Printf("mealDatas: %v %T\n", mealDatas, mealDatas)

}
