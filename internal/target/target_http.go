package target

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/trigger"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/txt"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/cbirc"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/chinawuliu"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/cni"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/eastmoney"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/fx678"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/manual"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/ppi100"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/sci"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/sina"
	"github.com/xiaogogonuo/cct-spider/internal/target/source/xiben"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
)

// pages：-1 代表爬完所有页数，适用于第一次写数据库
// pages：1  代表后期每次只爬一页，做增量爬取

// NewTarget 新建指标爬虫任务
func NewTarget() {
	// 1、第一步将历史数据唯一编码加载到内存
	var unique = map[string]struct{}{}
	for _, id := range txt.LoadFromTxt("target.txt") {
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
		case "EastMoneyEconomicTarget": // 爬取"东方财富"网站的`宏观指标`
			responses := eastmoney.SpiderEastMoneyEconomicTarget(ec.SourceTargetCodeSpider, ec.TargetCode)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "EastMoneyEconomicTargetBOR": // 爬取"东方财富"网站的`宏观指标-上海银行间同业拆放利率隔夜`
			responses := eastmoney.SpiderEastMoneyEconomicTargetBOR(ec.SourceTargetCodeSpider, 1)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "EastMoneyEconomicTargetCHN10": // 爬取"东方财富"网站的`宏观指标-中债10年期国债到期收益率`
			responses := eastmoney.SpiderEastMoneyEconomicTargetCHN10(1)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "EastMoneyGlobalTarget": // 爬取"东方财富"网站的`全球指数`
			responses := eastmoney.SpiderEastMoneyGlobalTarget(ec.SourceTargetCodeSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "EastMoneyIndustryTarget": // 爬取"东方财富"网站的`行业指标`
			responses := eastmoney.SpiderEastMoneyIndustryTarget(ec.SourceTargetCodeSpider, 1)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "EastMoneyQiHuoTarget": // 爬取"东方财富"网站的`期货指标`
			responses := eastmoney.SpiderEastMoneyQiHuoTarget(ec.SourceTargetCodeSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "SinaEconomicTarget": // 爬取新浪财经宏观指标
			responses := sina.SpiderSinaEconomicTarget(ec.SourceTargetCodeSpider, ec.TargetCode, 1)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "SinaTargetForex": // 爬取"新浪财经"网站的`外汇指标`
			responses := sina.SpiderSinaTargetForex(ec.SourceTargetCodeSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "FxExchangeTarget": // 爬取"汇通财经"网站的外汇、债券、原油、期货、外盘、汇率等指标
			responses := fx678.SpiderFxExchangeTarget(ec.SourceTargetCodeSpider, ec.TargetNameSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "FxExchangeTargetSpecial": // 爬取"汇通财经"网站的外汇、债券、原油、期货、外盘、汇率等指标
			responses := fx678.SpiderFxExchangeTargetSpecial(ec.SourceTargetCodeSpider, ec.SourceTargetCode)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "SCITargetCOI": // 爬取卓创资讯的原油价格指数
			responses := sci.SpiderSCITargetCOI()
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "SCITargetPII": // 爬取卓创资讯的造纸行业价格指数
			responses := sci.SpiderSCITargetPII()
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "CniCNYXTarget": // 爬取国证指数网站的人民币汇率
			responses := cni.SpiderCniCNYXTarget()
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "XiBenTarget": // 爬取西本新干线指标
			responses := xiben.SpiderXiBenTarget(ec.SourceTargetCodeSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "ChinaWuLiuTarget": // 爬取中国物流指标
			responses := chinawuliu.SpiderChinaWuLiuTarget(ec.TargetNameSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "CBIRCTarget": // 爬取银保监会指标
			responses := cbirc.SpiderCBIRCTarget(ec.TargetNameSpider)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "PPI100Target": // 爬取宏观数据网站的宏观指标
			responses := ppi100.SpiderPPI100Target(ec.TargetCode, ec.SourceTargetCodeSpider, 18)
			data := Generator(ec, responses, false)
			targets = append(targets, data...)
		case "ManualTarget": // 网页复制数据，手动计算，并无真正的下载操作，暂用
			responses := manual.SpiderManualTarget(ec.TargetCode)
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
	_ = push(realtimeTargets)
	targetForUpdate := push(nonRealtimeTargets)

	logger.Info(fmt.Sprintf("%d条实时指标更新成功", len(realtimeTargets)))
	logger.Info(fmt.Sprintf("%d条非实时指标更新成功", len(nonRealtimeTargets)))

	// 6、第七步将新记录表中的内容追加到target.txt
	var targetForAppend []string
	for _, update := range targetForUpdate {
		if _, ok := unique[update]; !ok {
			targetForAppend = append(targetForAppend, update)
		}
	}
	txt.Append2Txt("target.txt", targetForAppend...)
}

// RunInterval 默认每隔半小时跑一次
const RunInterval = time.Minute * 30

func RunTarget() {
	for {
		NewTarget()
		time.Sleep(RunInterval)
	}
}

// 对实时数据指标，发送给Java服务器的指标值格式如下：
// 最新价,涨跌,涨跌幅,最高,最低,昨收,更新时间
// 1.3543,-0.0045,-0.33%,1.3604,1.3533,1.3588,2022-04-15 23:28:50
