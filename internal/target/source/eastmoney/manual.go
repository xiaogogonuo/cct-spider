package eastmoney

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ManualHandleUpDown 手动计算涨跌幅
// 注意：ResponseArray中每一个response的Date字段值必须是8位，形如：20220415
func ManualHandleUpDown(array model.ResponseArray) (responses []model.Response) {
	sort.Sort(array) // 根据日期字段做降序排列，供下面做涨跌幅计算
	for i := 0; i < len(array)-1; i++ {
		var response model.Response
		response.Date = array[i].Date
		updateTime := response.Date[:4] + "-" + response.Date[4:6] + "-" + response.Date[6:8] + " " +
			strconv.FormatInt(int64(time.Now().Hour()), 10) + ":" +
			strconv.FormatInt(int64(time.Now().Minute()), 10) + ":" +
			strconv.FormatInt(int64(time.Now().Second()), 10)
		todayValue := fmt.Sprintf("%.4f", array[i].TargetValue)
		yesterdayValue := fmt.Sprintf("%.4f", array[i+1].TargetValue)
		var upDown, upDownPercent string
		if todayValue == "" || todayValue == "0.0000" || yesterdayValue == "" || yesterdayValue == "0.0000" {
			upDown = ""
			upDownPercent = ""
		} else {
			delta := array[i].TargetValue-array[i+1].TargetValue
			upDown = fmt.Sprintf("%.2f", delta)
			upDownPercent = fmt.Sprintf("%.2f%s", (delta)/array[i+1].TargetValue*100, "%")
		}
		response.TargetValue = strings.Join([]string{
			todayValue,     // 现价
			upDown,         // 涨跌
			upDownPercent,  // 涨跌幅
			"",             // 最高
			"",             // 最低
			yesterdayValue, // 昨收
			updateTime,     // 更新时间
		}, ",")
		responses = append(responses, response)
	}
	return
}
