package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

// KeepDecimal 保留小数点后几位，并四舍五入
func KeepDecimal(oldString string, behind int) (newString string) {
	dotIndex := strings.Index(oldString, ".")
	if dotIndex == -1 {
		return oldString
	}
	afterDotString := oldString[dotIndex+1:]
	if len(afterDotString) <= behind {
		return oldString
	}
	old, err := strconv.ParseFloat(oldString, 64)
	if err != nil {
		return oldString
	}
	oldFloat64 := Round(old, behind)
	return fmt.Sprintf("%." + strconv.Itoa(behind) +  "f", oldFloat64)
}
