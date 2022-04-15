package target

import (
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"strings"
)

// Generator 数据生成器，将配置文件数据和爬虫数据组合在一起
// ec：每个指标的配置
// responses：每个指标爬取的数据
// targetValueCut：指标值是否要切分
// - 如汇率、股票等数据，因为指标值包含最高、最低、开盘、收盘、涨幅等多个值，故多个值以逗号分隔当作一体发送给Java服务器
// - 如GDP、CPI等数据，因为指标值只包含一个值
func Generator(ec ExcelConfig, responses []model.Response, targetValueCut bool) (data []model.DataBase) {
	for _, response := range responses {
		var db model.DataBase
		db.ValueGUID = md5.MD5(ec.TargetCode + response.Date + response.RegionCode)
		db.TargetGUID = md5.MD5(ec.TargetCode)
		db.TargetCode = ec.TargetCode
		db.TargetName = ec.TargetName
		db.TargetNameEN = ec.TargetNameEn
		db.DataSourceCode = ec.DataSourceCode
		db.DataSourceName = ec.DataSourceName
		db.SourceTargetCode = ec.SourceTargetCode
		db.IsQuantity = ec.IsQuantity
		db.UnitType = ec.UnitType
		db.UnitName = ec.UnitName
		db.RegionCode = response.RegionCode
		db.RegionName = response.RegionName
		db.PeriodType = ec.PeriodType
		db.PeriodName = ec.PeriodName
		db.UpdateAcct(response.Date, ec.PeriodType)
		if targetValueCut {
			db.TargetValue = strings.Split(response.TargetValue, ",")[0]
		} else {
			db.TargetValue = response.TargetValue
		}
		data = append(data, db)
	}
	return 
}
