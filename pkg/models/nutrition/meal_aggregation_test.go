package nutrition

import (
	"reflect"
	"testing"
)

func TestMealAggregation(t *testing.T) {
	tests := []struct {
		name     string
		mealData []MealData
		want     Macronutrients
	}{
		{
			name: "single meal",
			mealData: []MealData{
				{
					MealName: "Chicken Breast",
					MacronutrientsData: Macronutrients{
						Protein:       Nutrient{Amount: 31, Calories: 124},
						Fat:           Nutrient{Amount: 3.6, Calories: 32.4},
						Carbohydrate:  Nutrient{Amount: 0, Calories: 0},
						TotalCalories: 156.4,
					},
					Info: "Grilled",
				},
			},
			want: Macronutrients{
				Protein:       Nutrient{Amount: 31, Calories: 124},
				Fat:           Nutrient{Amount: 3.6, Calories: 32.4},
				Carbohydrate:  Nutrient{Amount: 0, Calories: 0},
				TotalCalories: 156.4,
			},
		},
		{
			name: "multiple meals",
			mealData: []MealData{
				{
					MealName: "Chicken Breast",
					MacronutrientsData: Macronutrients{
						Protein:       Nutrient{Amount: 31, Calories: 124},
						Fat:           Nutrient{Amount: 3.6, Calories: 32.4},
						Carbohydrate:  Nutrient{Amount: 0, Calories: 0},
						TotalCalories: 156.4,
					},
					Info: "Grilled",
				},
				{
					MealName: "Brown Rice",
					MacronutrientsData: Macronutrients{
						Protein:       Nutrient{Amount: 2.6, Calories: 10.4},
						Fat:           Nutrient{Amount: 0.9, Calories: 8.1},
						Carbohydrate:  Nutrient{Amount: 45, Calories: 180},
						TotalCalories: 198.5,
					},
					Info: "Cooked",
				},
			},
			want: Macronutrients{
				Protein:       Nutrient{Amount: 33.6, Calories: 134.4},
				Fat:           Nutrient{Amount: 4.5, Calories: 40.5},
				Carbohydrate:  Nutrient{Amount: 45, Calories: 180},
				TotalCalories: 354.9,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MealAggregation(tt.mealData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MealAggregation() = %v, want %v", got, tt.want)
			}
		})
	}
}
