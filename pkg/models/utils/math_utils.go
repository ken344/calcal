package utils

import (
	"math"
)

// 指定した小数点の位で四捨五入する関数
func Rounding(f float64, prec int) float64 {

	// 四捨五入（10のprec乗を掛けた値で四捨五入した後、10のprec乗で割る）
	factor := math.Pow(10, float64(prec))
	rounded := math.Round(f*factor) / factor

	return rounded
}
