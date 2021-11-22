package indicator

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/marco/v2/code"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/marco/v2/pkg/net/http/request"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/marco/v2/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/marco/v2/pkg/time/date"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"strings"
	"time"
)

type Model struct {
	Zone          string
	Period        string
	IndicatorName string
	IndicatorInfo map[string]string
}

// downloaded returns history year data
func (m Model) downloaded() (row [][]string) {
	targetCode := m.IndicatorInfo["TargetCode"]
	switch m.Zone + m.Period {
	case code.Marco + code.PeriodTypeYear:
		row = mysql.Query(fmt.Sprintf(MarcoYearSQL, table, targetCode))
	case code.Marco + code.PeriodTypeSeason:
		row = mysql.Query(fmt.Sprintf(MarcoSeasonSQL, table, targetCode))
	case code.Marco + code.PeriodTypeMonth:
		row = mysql.Query(fmt.Sprintf(MarcoMonthSQL, table, targetCode))
	case code.Province + code.PeriodTypeYear:
		row = mysql.Query(fmt.Sprintf(ProvinceYearSQL, table, targetCode))
	case code.Province + code.PeriodTypeSeason:
		row = mysql.Query(fmt.Sprintf(ProvinceSeasonSQL, table, targetCode))
	case code.Province + code.PeriodTypeMonth:
		row = mysql.Query(fmt.Sprintf(ProvinceMonthSQL, table, targetCode))
	}
	return
}

func (m Model) param(dateRegion, sourceTargetCode string) (p request.Param) {
	switch m.Period {
	case code.PeriodTypeYear:
		switch m.Zone {
		case code.Marco:
			p = request.MacroYear(dateRegion)
		case code.Province:
			p = request.ProvinceYear(dateRegion, sourceTargetCode)
		}
	case code.PeriodTypeSeason:
		switch m.Zone {
		case code.Marco:
			p = request.MacroSeason(dateRegion)
		case code.Province:
			p = request.ProvinceSeason(dateRegion, sourceTargetCode)
		}
	case code.PeriodTypeMonth:
		switch m.Zone {
		case code.Marco:
			p = request.MacroMonth(dateRegion)
		case code.Province:
			p = request.ProvinceMonth(dateRegion, sourceTargetCode)
		}
	}
	return
}

func (m Model) cn() (c string) {
	switch m.Zone + m.Period {
	case code.Marco + code.PeriodTypeYear:
		c = code.MarcoAnnualDataCode
	case code.Marco + code.PeriodTypeSeason:
		c = code.MarcoQuarterlyDataCode
	case code.Marco + code.PeriodTypeMonth:
		c = code.MarcoMonthlyDataCode
	case code.Province + code.PeriodTypeYear:
		c = code.ProvinceAnnualDataCode
	case code.Province + code.PeriodTypeSeason:
		c = code.ProvinceQuarterlyDataCode
	case code.Province + code.PeriodTypeMonth:
		c = code.ProvinceMonthlyDataCode
	}
	return
}

func (m Model) entrypoint() {
	dst := m.downloaded()
	src := date.SinceStartYear(m.IndicatorInfo["StartYear"])
	difference := date.Difference(src, dst)
	dateRegion := strings.Join(difference, ",")
	sourceTargetCode := m.IndicatorInfo["SourceTargetCode"]
	mParam := m.param(dateRegion, sourceTargetCode)
	result := response.Crawl(m.cn(), sourceTargetCode, mParam)
	if result == nil || len(result) == 0 {
		//logger.Info(fmt.Sprintf("%s has no data to update", m.IndicatorName))
		return
	}
	data := Constructor{
		Period: m.Period,
		IndicatorName: m.IndicatorName,
		IndicatorInfo: m.IndicatorInfo,
		Respond: result,
	}.Construct()
	//logger.Info(m.IndicatorName, logger.Field("updating rows: ", len(data)))
	Dump(data)
}

func Start() {
	for indicatorName, indicatorInfo := range code.Indicator {
		for _, period := range strings.Split(indicatorInfo["PeriodType"], ",") {
			for _, zone := range strings.Split(indicatorInfo["Zone"], ",") {
				Model{
					Zone: zone,
					Period: period,
					IndicatorName: indicatorName,
					IndicatorInfo: indicatorInfo,
				}.entrypoint()
				time.Sleep(time.Second * 10)
			}
		}
	}
}