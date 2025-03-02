package nutrition

// mealData構造体（食品の名前と栄養素を格納）
type MealData struct {
	MealName           string
	MacronutrientsData Macronutrients
	Info               string
}

func MealAggregation(MealDatas []MealData) Macronutrients {
	var sumMealData Macronutrients
	for _, n := range MealDatas {
		sumMealData.Protein.Amount += n.MacronutrientsData.Protein.Amount
		sumMealData.Fat.Amount += n.MacronutrientsData.Fat.Amount
		sumMealData.Carbohydrate.Amount += n.MacronutrientsData.Carbohydrate.Amount
		sumMealData.Protein.Calories += n.MacronutrientsData.Protein.Calories
		sumMealData.Fat.Calories += n.MacronutrientsData.Fat.Calories
		sumMealData.Carbohydrate.Calories += n.MacronutrientsData.Carbohydrate.Calories
		sumMealData.TotalCalories += n.MacronutrientsData.TotalCalories
	}

	return sumMealData
}
