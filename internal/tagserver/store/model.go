package store

type PolicyNewsOrg struct {
	NEWS_GUID              string
	NEWS_ID                int
	NEWS_TS                string
	NEWS_SUMMARY           string
	EMOTION_INDICATOR      string
	EMOTION_INDICATOR_NAME string
	EMOTION_WEIGHT         string
	EMOTION_DETAIL         string
	RegionMap              map[string]int
	CompanyMap             map[string]int
	IndustryMap            map[string]float64
}

type PolicyNews struct {
	NEWS_GUID              string
	EMOTION_INDICATOR      string
	EMOTION_INDICATOR_NAME string
	EMOTION_WEIGHT         string
	EMOTION_DETAIL         string
	NEWS_LABELS            string
	INDUSTRY_LABELS        string
	COMPANY_LABELS         string
	REGION_LABELS          string
	EVENT_LABELS           string
}

type NewsRegion struct {
	REGION_LABEL_GUID string
	NEWS_GUID         string
	REGION_CODE       string
	REGION_NAME       string
	ENGLISH_NAME      string
	NEWS_ID           int
	NEWS_TS           string
	RELEVANCE         string
	EMOTION_INDICATOR string
	EMOTION_WEIGHT    string
	EMOTION_DETAIL    string
}

type NewsCompany struct {
	COMP_LABEL_GUID   string
	NEWS_GUID         string
	STOCK_CODE        string
	CREDIT_CODE       string
	COMP_GUID         string
	COMPANY_ID        string
	CHINESE_NAME      string
	ENGLISH_NAME      string
	NEWS_ID           int
	NEWS_TS           string
	RELEVANCE         string
	EMOTION_INDICATOR string
	EMOTION_WEIGHT    string
	EMOTION_DETAIL    string
}

type NewsIndustry struct {
	INDUSTRY_LABEL_GUID string
	NEWS_GUID           string
	INDUSTRY_CODE       string
	INDUSTRY_NAME       string
	ENGLISH_NAME        string
	NEWS_ID             int
	NEWS_TS             string
	RELEVANCE           string
	EMOTION_INDICATOR   string
	EMOTION_WEIGHT      string
	EMOTION_DETAIL      string
}
