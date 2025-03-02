package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

// Nutrient 構造体（栄養素データ）
type Nutrient struct {
	Protein      float64
	Fat          float64
	Carbohydrate float64
}

// CSVファイルのパス(DB相当データ)
const dataFilePath = "./meals.csv"

// CSVファイルのパスを取得
func getDataFilePath(csvFilePath string) (string, error) {
	// 実行中のファイルのパスを取得
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current file path")
	}

	// 実行中のファイルのディレクトリを取得
	mainDir := filepath.Dir(filename)

	// CSVファイルのパスを生成して返す
	return filepath.Join(mainDir, csvFilePath), nil
}

// CSVファイルを読み込み、食品名をキーにしたマップを返す
func LoadCSV() (map[string]Nutrient, error) {

	// CSVファイルのパスを取得
	filePath, err := getDataFilePath(dataFilePath)
	if err != nil {
		return nil, err
	}

	// CSVファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// CSVファイルを読み込む
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
