package date

import (
	"strconv"
	"time"
)

// SinceStartYear returns years from start year to current year
func SinceStartYear(start string) (years []string) {
	startInt, _ := strconv.ParseInt(start, 10, 64)
	for i := startInt; i <= int64(time.Now().Year()); i++ {
		years = append(years, strconv.FormatInt(i, 10))
	}
	return
}

// Difference returns different year within src and dst(including current no matter how)
func Difference(src []string, dst [][]string) (years []string) {
	md := make(map[string]struct{})
	for _, array := range dst {
		md[array[0]] = struct{}{}
	}
	ms := make(map[string]struct{})
	for _, year := range src {
		if _, ok := md[year]; !ok {
			ms[year] = struct{}{}
		}
	}
	ms[strconv.FormatInt(int64(time.Now().Year()), 10)] = struct{}{} // add current year
	for year := range ms {
		years = append(years, year)
	}
	return
}