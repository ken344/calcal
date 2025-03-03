package nutrition

import (
	"testing"

	"github.com/ken344/calcal/pkg/models/utils"
)

func TestNewMacronutrients(t *testing.T) {
	tests := []struct {
		name         string
		protein      float64
		fat          float64
		carbohydrate float64
		want         Macronutrients
	}{
		{
			name:         "Test case 1",
			protein:      10.123,
			fat:          20.456,
			carbohydrate: 30.789,
			want: Macronutrients{
				Protein:      Nutrient{Amount: utils.Rounding(10.123, 1)},
				Fat:          Nutrient{Amount: utils.Rounding(20.456, 1)},
				Carbohydrate: Nutrient{Amount: utils.Rounding(30.789, 1)},
			},
		},
		{
			name:         "Test case 2",
			protein:      15.555,
			fat:          25.555,
			carbohydrate: 35.555,
			want: Macronutrients{
				Protein:      Nutrient{Amount: utils.Rounding(15.555, 1)},
				Fat:          Nutrient{Amount: utils.Rounding(25.555, 1)},
				Carbohydrate: Nutrient{Amount: utils.Rounding(35.555, 1)},
			},
		},
		{
			name:         "Test case 3",
			protein:      0,
			fat:          0,
			carbohydrate: 0,
			want: Macronutrients{
				Protein:      Nutrient{Amount: utils.Rounding(0, 1)},
				Fat:          Nutrient{Amount: utils.Rounding(0, 1)},
				Carbohydrate: Nutrient{Amount: utils.Rounding(0, 1)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMacronutrients(tt.protein, tt.fat, tt.carbohydrate); got != tt.want {
				t.Errorf("NewMacronutrients() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalorieCalculation(t *testing.T) {
	tests := []struct {
		name                string
		protein             float64
		fat                 float64
		carbs               float64
		wantProteinCalories float64
		wantFatCalories     float64
		wantCarbCalories    float64
	}{
		{
			name:                "Test case 1",
			protein:             10,
			fat:                 5,
			carbs:               20,
			wantProteinCalories: 40,
			wantFatCalories:     45,
			wantCarbCalories:    80,
		},
		{
			name:                "Test case 2",
			protein:             15.5,
			fat:                 7.3,
			carbs:               25.2,
			wantProteinCalories: 62,
			wantFatCalories:     66,
			wantCarbCalories:    101,
		},
		{
			name:                "Test case 3",
			protein:             0,
			fat:                 0,
			carbs:               0,
			wantProteinCalories: 0,
			wantFatCalories:     0,
			wantCarbCalories:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMacronutrients(tt.protein, tt.fat, tt.carbs)
			m.CalorieCalculation()
			if m.Protein.Calories != tt.wantProteinCalories {
				t.Errorf("Protein.Calories = %v, want %v", m.Protein.Calories, tt.wantProteinCalories)
			}
			if m.Fat.Calories != tt.wantFatCalories {
				t.Errorf("Fat.Calories = %v, want %v", m.Fat.Calories, tt.wantFatCalories)
			}
			if m.Carbohydrate.Calories != tt.wantCarbCalories {
				t.Errorf("Carbohydrate.Calories = %v, want %v", m.Carbohydrate.Calories, tt.wantCarbCalories)
			}
		})
	}
}
func TestTotalCaloriesCalculation(t *testing.T) {
	tests := []struct {
		name              string
		protein           float64
		fat               float64
		carbs             float64
		wantTotalCalories float64
	}{
		{
			name:              "Test case 1",
			protein:           10,
			fat:               5,
			carbs:             20,
			wantTotalCalories: 165,
		},
		{
			name:              "Test case 2",
			protein:           15.5,
			fat:               7.3,
			carbs:             25.2,
			wantTotalCalories: 229,
		},
		{
			name:              "Test case 3",
			protein:           0,
			fat:               0,
			carbs:             0,
			wantTotalCalories: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMacronutrients(tt.protein, tt.fat, tt.carbs)
			m.CalorieCalculation()
			m.TotalCaloriesCalculation()
			if m.TotalCalories != tt.wantTotalCalories {
				t.Errorf("TotalCalories = %v, want %v", m.TotalCalories, tt.wantTotalCalories)
			}
		})
	}
}
