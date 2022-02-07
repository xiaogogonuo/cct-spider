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
	NEWS_GUID              string `json:"newsGuid"`
	EMOTION_INDICATOR      string `json:"emotionIndicator"`
	EMOTION_INDICATOR_NAME string `json:"emotionIndicatorName"`
	EMOTION_WEIGHT         string `json:"emotionWeight"`
	EMOTION_DETAIL         string `json:"emotionDetail"`
	NEWS_LABELS            string `json:"newsLabels"`
	INDUSTRY_LABELS        string `json:"industryLabels"`
	COMPANY_LABELS         string `json:"companyLabels"`
	REGION_LABELS          string `json:"regionLabels"`
	EVENT_LABELS           string `json:"eventLabels"`
}

type NewsRegion struct {
	REGION_LABEL_GUID string `json:"regionLabelGuid"`
	NEWS_GUID         string `json:"newsGuid"`
	REGION_CODE       string `json:"regionCode"`
	REGION_NAME       string `json:"regionName"`
	ENGLISH_NAME      string `json:"englishName"`
	NEWS_ID           int    `json:"newsId"`
	NEWS_TS           string `json:"newsTs"`
	RELEVANCE         string `json:"relevance"`
	EMOTION_INDICATOR string `json:"emotionIndicator"`
	EMOTION_WEIGHT    string `json:"emotionWeight"`
	EMOTION_DETAIL    string `json:"emotionDetail"`
}

type NewsCompany struct {
	COMP_LABEL_GUID   string `json:"compLabelGuid"`
	NEWS_GUID         string `json:"newsGuid"`
	STOCK_CODE        string `json:"stockCode"`
	CREDIT_CODE       string `json:"creditCode"`
	COMP_GUID         string `json:"compGuid"`
	COMPANY_ID        string `json:"companyId"`
	CHINESE_NAME      string `json:"chineseName"`
	ENGLISH_NAME      string `json:"englishName"`
	NEWS_ID           int    `json:"newsId"`
	NEWS_TS           string `json:"newsTs"`
	RELEVANCE         string `json:"relevance"`
	EMOTION_INDICATOR string `json:"emotionIndicator"`
	EMOTION_WEIGHT    string `json:"emotionWeight"`
	EMOTION_DETAIL    string `json:"emotionDetail"`
}

type NewsIndustry struct {
	INDUSTRY_LABEL_GUID string `json:"industryLabelGuid"`
	NEWS_GUID           string `json:"newsGuid"`
	INDUSTRY_CODE       string `json:"industryCode"`
	INDUSTRY_NAME       string `json:"industryName"`
	ENGLISH_NAME        string `json:"englishName"`
	NEWS_ID             int    `json:"newsId"`
	NEWS_TS             string `json:"newsTs"`
	RELEVANCE           string `json:"relevance"`
	EMOTION_INDICATOR   string `json:"emotionIndicator"`
	EMOTION_WEIGHT      string `json:"emotionWeight"`
	EMOTION_DETAIL      string `json:"emotionDetail"`
}
