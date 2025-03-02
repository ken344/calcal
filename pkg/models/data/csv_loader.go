package data

import (
	_ "embed"
	"encoding/csv"
	"strconv"
	"strings"
)

// Nutrient 構造体（栄養素データ）
type Nutrient struct {
	Protein      float64
	Fat          float64
	Carbohydrate float64
}

// 以下のコメントは、embedディレクティブを使って埋め込まれたCSVデータを取得するためのコードです。
// meals.csvファイルを埋め込む設定なので、変更はしないようにしてください。
//
//go:embed meals.csv
var mealsCSV string // CSVデータが文字列として埋め込まれる

// CSVファイルを読み込み、食品名をキーにしたマップを返す
func LoadCSV() (map[string]Nutrient, error) {

	// // CSVデータを表示（enbedデバッグ用）
	// fmt.Println("Embedded CSV Data:")
	// fmt.Println(mealsCSV)

	// CSVデータをパース
	reader := csv.NewReader(strings.NewReader(mealsCSV))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// 食品名をキーにしたマップを作成
	data := make(map[string]Nutrient)

	// ヘッダー行はスキップし、データ行のみmapに格納
	for i, record := range records {
		if i == 0 {
			continue // ヘッダー行をスキップ
		}

		// 栄養素データを構造体に格納：文字列からfloat64に変換
		protein, _ := strconv.ParseFloat(record[1], 64)
		fat, _ := strconv.ParseFloat(record[2], 64)
		carbohydrate, _ := strconv.ParseFloat(record[3], 64)

		// 食品名(record[0])をキーにして、栄養素をmapに格納
		data[record[0]] = Nutrient{
			Protein:      protein,
			Fat:          fat,
			Carbohydrate: carbohydrate,
		}
	}

	return data, nil
}
