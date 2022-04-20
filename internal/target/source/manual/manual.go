package manual

import (
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/excel"
	"sort"
)

type Man [][]string

func (m Man) Len() int {
	return len(m)
}

func (m Man) Less(i, j int) bool {
	return m[i][0] > m[j][0]
}

func (m Man) Swap(i, j int)  {
	m[i], m[j] = m[j], m[i]
}

// SpiderManualTarget 网页复制数据，手动计算，并无真正的下载操作，暂用
// 适用指标：
// - 人民币贷款余额增速月度、人民币存款余额增速月度、社会融资规模新增月度(三个指标均计算的同比)
func SpiderManualTarget(targetCode string) (responses []model.Response) {
	rows := make(Man, 0)
	switch targetCode {
	case "HG00118":
		rows = excel.ReadExcel("诚通指标配置.xlsx", "人民币存款")[1:]
	case "HG00119":
		rows = excel.ReadExcel("诚通指标配置.xlsx", "人民币贷款")[1:]
	case "HG00120":
		rows = excel.ReadExcel("诚通指标配置.xlsx", "社会融资规模增量")[1:]
	}
	sort.Sort(rows)
	responses = manualPipeline(rows)
	return
}
