package data

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Nutrient 構造体（栄養素データ）
type Nutrient struct {
	Protein      float64
	Fat          float64
	Carbohydrate float64
}

// CSVファイルを読み込み、食品名をキーにしたマップを返す
func LoadCSV(filePath string) (map[string]Nutrient, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
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
