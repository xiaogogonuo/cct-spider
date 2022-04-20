package target

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/trigger"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/txt"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/pbc"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
)

// NewTargetSelenium 新建指标爬虫任务
func NewTargetSelenium() {
	// 1、第一步将历史数据唯一编码加载到内存
	var unique = map[string]struct{}{}
	for _, id := range txt.LoadFromTxt("target_selenium.txt") {
		if _, ok := unique[id]; !ok {
			unique[id] = struct{}{}
		}
	}

	// 2、第二步从Excel读取指标相关的配置信息
	ecs := ExcelSetUp()

	var targets []model.DataBase
	// 3、第三步遍历指标配置表，循环爬取数据
	for _, ec := range ecs {
		if !ec.Enable || !trigger.Trigger(ec.SpiderTime) {
			continue
		}
		switch ec.Adapter {
		case "PBCTarget": // 爬取中国人民银行指标
			responses := pbc.SpiderPBCTarget(ec.SourceTargetCodeSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		}
	}

	// 4、第四步过滤重复指标
	var nonRealtimeTargets []model.DataBase // 代表非实时指标，需要做唯一性校验
	var realtimeTargets []model.DataBase    // 代表实时指标，直接发送到数据库服务器
	for _, tar := range targets {
		if strings.Contains(tar.TargetValue, ",") { // 实时指标值包含多个数据，以逗号分隔
			realtimeTargets = append(realtimeTargets, tar)
			continue
		}
		if _, ok := unique[tar.TargetGUID]; !ok {
			nonRealtimeTargets = append(nonRealtimeTargets, tar)
		}
	}

	// 5、第五步将新数据发送到远程服务器
	//_ = push(realtimeTargets)
	//targetForUpdate := push(nonRealtimeTargets)

	logger.Info(fmt.Sprintf("%d条实时指标更新成功", len(realtimeTargets)))
	logger.Info(fmt.Sprintf("%d条非实时指标更新成功", len(nonRealtimeTargets)))

	// 6、第七步将新记录表中的内容追加到target.txt
	//var targetForAppend []string
	//for _, update := range targetForUpdate {
	//	if _, ok := unique[update]; !ok {
	//		targetForAppend = append(targetForAppend, update)
	//	}
	//}
	//txt.Append2Txt("target_selenium.txt", targetForAppend...)
}

func RunTargetSelenium() {
	NewTargetSelenium()
}
