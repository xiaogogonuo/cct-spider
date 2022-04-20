package manual

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strconv"
	"time"
)

func manualPipeline(m Man) (responses []model.Response) {
	uniqueDate := map[time.Time]float64{}
	for _, row := range m {
		date := row[0]
		value := row[1]
		dateTime, err := time.Parse("200601", date)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if _, ok := uniqueDate[dateTime]; !ok {
			uniqueDate[dateTime] = valueFloat64
		}
	}
	for _, row := range m {
		date := row[0]
		value := row[1]
		dateTime, err := time.Parse("200601", date)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		lastYearSameMonth := dateTime.AddDate(-1, 0 , 0) // 获取去年同一个月
		if _, ok := uniqueDate[lastYearSameMonth]; ok {
			var response model.Response
			percent := (valueFloat64 - uniqueDate[lastYearSameMonth]) / uniqueDate[lastYearSameMonth] * 100
			response.Date = date
			response.TargetValue = fmt.Sprintf("%.2f", percent)
			responses = append(responses, response)
		}
	}
	return
}
