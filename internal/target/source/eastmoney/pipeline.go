package eastmoney

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/calculator"
	"strconv"
	"strings"
)

// eastMoneyEconomicTargetGDPTBPipeline 宏观指标`国内生产总值同比增长`数据处理管道
func eastMoneyEconomicTargetGDPTBPipeline(table []string, dateIndex, valueIndex int) (responses []model.Response) {
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

// eastMoneyEconomicTargetGDPHBPipeline 宏观指标`国内生产总值环比增长`数据处理管道
func eastMoneyEconomicTargetGDPHBPipeline(table []string, dateIndex, valueIndex int) (responses []model.Response) {
	_responses := eastMoneyEconomicTargetGDPTBPipeline(table, dateIndex, valueIndex)
	var agents []model.ResponseDateStringValueFloat
	// 将指标值字符串转换为浮点数
	for _, response := range _responses {
		var agent model.ResponseDateStringValueFloat
		value, _ := strconv.ParseFloat(response.TargetValue, 64)
		agent.TargetValue = value
		agent.Date = response.Date
		agents = append(agents, agent)
	}
	var agentHB []model.ResponseDateStringValueFloat
	for idx, agent := range agents {
		var res model.ResponseDateStringValueFloat
		switch strings.Contains(agent.Date, "Q1") {
		case true:
			res.Date = agent.Date
			res.TargetValue = agent.TargetValue
			agentHB = append(agentHB, res)
		case false:
			if idx < len(agents)-1 {
				res.Date = agent.Date
				res.TargetValue = agent.TargetValue - agents[idx+1].TargetValue
				agentHB = append(agentHB, res)
			}
		}
	}
	for idx := 0; idx < len(agentHB) - 1; idx++ {
		var response model.Response
		response.Date = agentHB[idx].Date
		hb := (agentHB[idx].TargetValue - agentHB[idx+1].TargetValue)/agentHB[idx+1].TargetValue*100
		response.TargetValue = fmt.Sprintf("%.2f", hb)
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
		response.TargetValue = calculator.KeepDecimal(cut[valueIndex], 2)
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
		response.TargetValue = strconv.FormatFloat(data.IndicatorValue, 'f', 2, 64)
		responses = append(responses, response)
	}
	return
}
