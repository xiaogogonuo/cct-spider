package code

// 指标周期
const (
	PeriodTypeYear      = "10"
	PeriodNameYear      = "年"
	PeriodTypeSeason    = "20"
	PeriodNameSeason    = "季"
	PeriodTypeMonth     = "30"
	PeriodNameMonth     = "月"
	PeriodTypeWeek      = "40"
	PeriodNameWeek      = "周"
	PeriodTypeDay       = "50"
	PeriodNameDay       = "日"
	PeriodTypeReal      = "60"
	PeriodNameReal      = "实时"
	PeriodTypeUnCertain = "90"
	PeriodNameUnCertain = "不定期"
)

var PeriodTypeName = make(map[string]string)
var PeriodNameType = make(map[string]string)

func setUpPeriod() {
	PeriodTypeName[PeriodTypeYear] = PeriodNameYear
	PeriodTypeName[PeriodTypeSeason] = PeriodNameSeason
	PeriodTypeName[PeriodTypeMonth] = PeriodNameMonth
	PeriodTypeName[PeriodTypeWeek] = PeriodNameWeek
	PeriodTypeName[PeriodTypeDay] = PeriodNameDay
	PeriodTypeName[PeriodTypeReal] = PeriodNameReal
	PeriodTypeName[PeriodTypeUnCertain] = PeriodNameUnCertain

	PeriodNameType[PeriodNameYear] = PeriodTypeYear
	PeriodNameType[PeriodNameSeason] = PeriodTypeSeason
	PeriodNameType[PeriodNameMonth] = PeriodTypeMonth
	PeriodNameType[PeriodNameWeek] = PeriodTypeWeek
	PeriodNameType[PeriodNameDay] = PeriodTypeDay
	PeriodNameType[PeriodNameReal] = PeriodTypeReal
	PeriodNameType[PeriodNameUnCertain] = PeriodTypeUnCertain
}