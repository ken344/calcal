package data

import (
	"reflect"
	"testing"
)

// テスト用のCSVデータ
// テスト用のデータを用いることで、テストに一貫性を持たせることができる
const testCSV = `Meal,Protein,Fat,Carbohydrate
Apple,0.3,0.2,14.0
Banana,1.1,0.3,23.0
Chicken,27.3,3.6,0.0
`

// LoadCSV 関数のテスト
func TestLoadCSV(t *testing.T) {
	// 本番ではenbedにより埋め込まれたCSVデータを使用するが、テスト用にCSVデータを差し替える
	// オリジナルの mealsCSV を保存
	originalCSV := mealsCSV
	// テスト用CSVに置き換える
	mealsCSV = testCSV
	// テスト終了後にオリジナルのCSVに戻す
	defer func() { mealsCSV = originalCSV }()

	expected := map[string]Nutrient{
		"Apple":   {Protein: 0.3, Fat: 0.2, Carbohydrate: 14.0},
		"Banana":  {Protein: 1.1, Fat: 0.3, Carbohydrate: 23.0},
		"Chicken": {Protein: 27.3, Fat: 3.6, Carbohydrate: 0.0},
	}

	// LoadCSV を実行
	result, err := LoadCSV()
	if err != nil {
		t.Fatalf("LoadCSV failed: %v", err)
	}

	// 結果を比較
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

// テスト用の不正なCSVデータのロード。エラーが発生することを確認。
func TestLoadCSV_InvalidData(t *testing.T) {
	invalidCSV := `Meal,Protein,Fat,Carbohydrate
	Apple,abc,0.2,14.0
	Banana,1.1,0.3,xyz
	`
	originalCSV := mealsCSV
	mealsCSV = invalidCSV
	defer func() { mealsCSV = originalCSV }()

	_, err := LoadCSV()
	if err == nil {
		t.Errorf("Expected error for invalid CSV data, but got nil")
	}
}
