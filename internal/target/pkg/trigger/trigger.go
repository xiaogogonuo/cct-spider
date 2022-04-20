package trigger

import (
	"strings"
	"time"
)

// Trigger 爬虫运行触发器
// spiderTime 代表用户自定义指标的爬取时间
func Trigger(spiderTime string) bool {
	if spiderTime == "" { // 实时更新的指标
		return true
	}
	curTime := time.Now()
	runTimes := strings.Split(spiderTime, "~")
	runTimeL, _ := time.Parse("15:04", runTimes[0])
	runTimeR, _ := time.Parse("15:04", runTimes[1])
	if curTime.Hour() < runTimeL.Hour() || curTime.Hour() > runTimeR.Hour() {
		return false
	}
	if curTime.Hour() == runTimeL.Hour() && curTime.Minute() < runTimeL.Minute() {
		return false
	}
	if curTime.Hour() == runTimeR.Hour() && curTime.Minute() > runTimeR.Minute() {
		return false
	}
	return true
}
