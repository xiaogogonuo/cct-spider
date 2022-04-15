package calculator

import "math"

// Round 四舍五入算法
// precision：保留小数点后多少位
func Round(val float64, precision int) float64 {
	if precision == 0 {
		return math.Round(val)
	}
	p := math.Pow10(precision)
	if precision < 0 {
		return math.Floor(val*p+0.5) * math.Pow10(-precision)
	}
	return math.Floor(val*p+0.5) / p
}
