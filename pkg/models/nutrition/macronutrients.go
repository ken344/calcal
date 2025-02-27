package nutrition

import (
	"github.com/ken344/calcal/pkg/models/utils"
)

// 栄養素のカロリー（Kcal/g）を定義
const (
	ProteinCaloriesPerGram      = 4 // 蛋白質のグラム当たりカロリー
	FatCaloriesPerGram          = 9 // 脂質のグラム当たりカロリー
	CarbohydrateCaloriesPerGram = 4 // 炭水化物のグラム当たりカロリー
)

// 栄養素の量とカロリーを持つ構造体
type Nutrient struct {
	Amount   float64 // 量 (g)
	Calories float64 // カロリー (kcal)
}

// 主要栄養素をまとめた構造体
type Macronutrients struct {
	Protein      Nutrient
	Fat          Nutrient
	Carbohydrate Nutrient
}

// 栄養素の量を受け取り、自動的にカロリーを計算する関数（コンストラクタ相当）
func NewMacronutrients(protein, fat, carbohydrate float64) Macronutrients {
	// 小数点第1位で四捨五入
	protein = utils.Rounding(protein, 1)
	fat = utils.Rounding(fat, 1)
	carbohydrate = utils.Rounding(carbohydrate, 1)

	return Macronutrients{
		Protein: Nutrient{ // 蛋白質のデータ
			Amount:   protein,
			Calories: utils.Rounding(protein*ProteinCaloriesPerGram, 0),
		},
		Fat: Nutrient{ // 脂質のデータ
			Amount:   fat,
			Calories: utils.Rounding(fat*FatCaloriesPerGram, 0),
		},
		Carbohydrate: Nutrient{ // 炭水化物のデータ
			Amount:   carbohydrate,
			Calories: utils.Rounding(carbohydrate*CarbohydrateCaloriesPerGram, 0),
		},
	}
}

// 栄養素のカロリーの合計を計算する
func (m Macronutrients) TotalCalories() float64 {
	return utils.Rounding(m.Protein.Calories+m.Fat.Calories+m.Carbohydrate.Calories, 0)
}
