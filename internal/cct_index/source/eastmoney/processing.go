package eastmoney

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"strings"
)

// 数据加工处理

// 补充扩展计算出拆借利率的涨跌幅等信息
func expansion(bufferTFList model.BufferTFList) (buffers []*model.Buffer) {
	for i := 0; i < len(bufferTFList)-1; i++ {
		cur := fmt.Sprintf("%.4f", bufferTFList[i+0].TargetValue) // 今天的拆借利率
		pre := fmt.Sprintf("%.4f", bufferTFList[i+1].TargetValue) // 昨天的拆借利率

		// 计算涨跌值和涨跌幅
		var (
			upDownValue = "" // 涨跌额
			upDownRange = "" // 涨跌幅
		)
		if cur == "0.0000" || pre == "0.0000" {
			upDownValue = ""
			upDownRange = ""
		} else {
			delta := bufferTFList[i].TargetValue - bufferTFList[i+1].TargetValue
			upDownValue = fmt.Sprintf("%.2f", delta)
			upDownRange = fmt.Sprintf("%.2f%s", (delta)/bufferTFList[i+1].TargetValue*100, "%")
		}

		buffer := &model.Buffer{}
		buffer.Date = bufferTFList[i].Date.Format("20060102")
		buffer.TargetValue = strings.Join([]string{
			cur,         // 现价
			upDownValue, // 涨跌额
			upDownRange, // 涨跌幅
			"",          // 最高
			"",          // 最低
			pre,         // 昨收
			bufferTFList[i].Date.Format("2006-01-02 03:04:05"), // 更新时间
		}, ",")
		buffers = append(buffers, buffer)
	}
	return
}
