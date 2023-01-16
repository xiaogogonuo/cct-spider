package factory

import (
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
)

// Manufacture 对指标配置信息和指标爬取信息进行组合加工
func Manufacture(ic *model.IndexConfig, buffers []*model.Buffer) (indexes []*model.Index) {
	for _, buffer := range buffers {
		index := &model.Index{}
		index.ValueGUID = md5.MD5(ic.TargetCode + buffer.Date + buffer.RegionCode)
		index.TargetGUID = md5.MD5(ic.TargetCode)
		index.TargetCode = ic.TargetCode
		index.TargetName = ic.TargetName
		index.TargetNameEN = ic.TargetNameEn
		index.DataSourceCode = ic.DataSourceCode
		index.DataSourceName = ic.DataSourceName
		index.SourceTargetCode = ic.SourceTargetCode
		index.RegionCode = buffer.RegionCode
		index.RegionName = buffer.RegionName
		index.IsQuantity = "Y"
		index.UnitType = ic.UnitType
		index.UnitName = ic.UnitName
		index.PeriodType = ic.PeriodType
		index.PeriodName = ic.PeriodName
		index.AcctSetter(buffer.Date)
		index.TargetValue = buffer.TargetValue
		indexes = append(indexes, index)
	}
	return
}
