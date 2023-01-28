package eastmoney

import (
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
)

// SpiderEastMoney 东方财富
func SpiderEastMoney(ic *model.IndexConfig) (buffers []*model.Buffer) {
	switch ic.TargetCode {
	// 东方财富 - 经济数据 - 国内生产总值
	case "HG00001", "HG00098":
		buffers = eastMoneyEconomicGDP(ic)
	// 东方财富 - 经济数据 - 工业增加值
	case "HG00016", "HG00017":
		buffers = eastMoneyEconomicIAV(ic)
	// 东方财富 - 经济数据 - 社会消费品零售总额
	case "HG00027", "HG00028", "HG00029", "HG00030":
		buffers = eastMoneyEconomicXFP(ic)
	// 东方财富 - 经济数据 - 货币供应量
	case "HG00006", "HG00007":
		buffers = eastMoneyEconomicM2(ic)
	// 东方财富 - 经济数据 - 居民消费价格指数
	case "HG00003", "HG00004", "HG00088":
		buffers = eastMoneyEconomicCPI(ic)
	// 东方财富 - 经济数据 - 采购经理人指数
	case "HG00020":
		buffers = eastMoneyEconomicPMI(ic)
	// 东方财富 - 经济数据 - 工业品出厂价格指数
	case "HG00023", "HG00089":
		buffers = eastMoneyEconomicPPI(ic)
	// 东方财富 - 经济数据 - 存款准备金率
	case "HG00066":
		buffers = eastMoneyEconomicZBJ(ic)
	// 东方财富 - 经济数据 - 海关进出口
	case "HG00065":
		buffers = eastMoneyEconomicJCK(ic)
	// 东方财富 - 经济数据 - 利率调整
	case "HY00007", "HY00011":
		buffers = eastMoneyEconomicLL(ic)
	// 东方财富 - 经济数据 - 外汇和黄金储备
	case "HG00090", "HG00091":
		buffers = eastMoneyEconomicWH(ic)
	// 东方财富 - 拆借利率 - ShiBor、LiBor
	case "HY00006", "HY00012", "HY00013", "HY00014", "HY00015", "HY00016", "HY00017", "HY00018":
		buffers = eastMoneyLendingRate(ic)
	// 东方财富 - 行业指数 - 物流业景气指数、中国大宗商品指数、波罗的海干散货指数、美原油指数、费城半导体指数
	case "HY00001", "HY00002", "HY00003", "HG00110", "HG00111":
		buffers = eastMoneyIndustry(ic)
	// 东方财富 - 行情中心 - 全球指数
	case "HY00005", "HG00099", "HG00100", "HG00101", "HG00102", "HG00103", "HG00104", "HG00105", "HG00106", "HG00107",
		"HG00108", "HG00109", "HG00112", "HG00113", "HG00114", "HG00115", "HG00116", "HG00117":
		buffers = eastMoneyGlobal(ic)
	// 东方财富 - 数据中心 - 经济数据 - 中美国债收益率
	case "HG00062":
		buffers = eastMoneyEconomicNationalDebt(ic)
	}
	return
}
