package eastmoney

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"strings"
)

// eastMoneyEconomicTargetGDPPipeline 宏观指标`国内生产总值同比增长`数据处理管道
func eastMoneyEconomicTargetGDPPipeline(table []string, dateIndex, valueIndex int) (responses []model.Response) {
	for _, row := range table {
		var response model.Response
		cut := strings.Split(row, ",")
		date := strings.ReplaceAll(cut[dateIndex], "-", "")
		switch date[4:6] {
		case "01", "02", "03":
			response.Date = date[:4] + "Q1"
		case "04", "05", "06":
			response.Date = date[:4] + "Q2"
		case "07", "08", "09":
			response.Date = date[:4] + "Q3"
		case "10", "11", "12":
			response.Date = date[:4] + "Q4"
		}
		response.TargetValue = cut[valueIndex]
		responses = append(responses, response)
	}
	return
}

// eastMoneyEconomicTargetMostPipeline 大部分的宏观指标数据处理管道
// dateIndex 日期索引
// valueIndex 值索引
// offset 偏移量，6代表取日期前六位，8代表取日期前八位
func eastMoneyEconomicTargetMostPipeline(table []string, dateIndex, valueIndex, offset int) (responses []model.Response) {
	for _, row := range table {
		var response model.Response
		cut := strings.Split(row, ",")
		response.Date = strings.ReplaceAll(cut[dateIndex], "-", "")[:offset]
		response.TargetValue = cut[valueIndex]
		responses = append(responses, response)
	}
	return
}

// eastMoneyIndustryTargetMostPipeline 大部分的行业指标数据处理管道
// offset 偏移量，6代表取日期前六位，8代表取日期前八位
func eastMoneyIndustryTargetMostPipeline(frontend eastMoneyIndustryTarget, offset int) (responses []model.Response) {
	for _, data := range frontend.Result.Data {
		var response model.Response
		response.Date = strings.ReplaceAll(data.ReportDate, "-", "")[:offset]
		response.TargetValue = fmt.Sprintf("%.0f", data.IndicatorValue)
		responses = append(responses, response)
	}
	return
}
