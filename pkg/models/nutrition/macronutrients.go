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
	Protein       Nutrient // 蛋白質
	Fat           Nutrient // 脂質
	Carbohydrate  Nutrient // 炭水化物
	TotalCalories float64  // 合計カロリー
}

// 栄養素の量を受け取り、構造体を生成する（コンストラクタ相当）
func NewMacronutrients(protein, fat, carbohydrate float64) Macronutrients {
	// 小数点第1位で四捨五入
	return Macronutrients{
		Protein: Nutrient{ // 蛋白質のデータ
			Amount: utils.Rounding(protein, 1),
		},
		Fat: Nutrient{ // 脂質のデータ
			Amount: utils.Rounding(fat, 1),
		},
		Carbohydrate: Nutrient{ // 炭水化物のデータ
			Amount: utils.Rounding(carbohydrate, 1),
		},
	}
}

// 各栄養素のカロリーを計算する
func (m *Macronutrients) CalorieCalculation() {
	m.Protein.Calories = utils.Rounding(m.Protein.Amount*ProteinCaloriesPerGram, 0)
	m.Fat.Calories = utils.Rounding(m.Fat.Amount*FatCaloriesPerGram, 0)
	m.Carbohydrate.Calories = utils.Rounding(m.Carbohydrate.Amount*CarbohydrateCaloriesPerGram, 0)
}

// 栄養素のカロリーの合計を計算する
func (m *Macronutrients) TotalCaloriesCalculation() {
	m.TotalCalories = utils.Rounding(m.Protein.Calories+m.Fat.Calories+m.Carbohydrate.Calories, 0)
}
