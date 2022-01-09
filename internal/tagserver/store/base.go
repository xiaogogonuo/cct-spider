package store

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/findmap"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strconv"
	"strings"
	"sync"
)

var (
	regionLabels = `{"labelCode":"70","labelName":"地区标签","code":"%s","name":"%s"}`
	companyLabels = `{"labelCode":"10","labelName":"公司标签","code":"%s","name":"%s"}`
	industryLabels = `{"labelCode":"20","labelName":"行业标签","code":"%s","name":"%s"}`
)


func QueryData(newsOrgChan chan<- *PolicyNewsOrg, wg *sync.WaitGroup) {
	defer wg.Done()
	sqlCode := "SELECT NEWS_GUID, NEWS_ID, NEWS_TS, NEWS_SUMMARY FROM t_dmbe_policy_news_info;"
	for _, newData := range mysql.Query(sqlCode) {
			id, err := strconv.Atoi(newData[1])
			if err != nil{
				logger.Error(err.Error())
				return
			}
			newsOrgChan <- &PolicyNewsOrg{

				NEWS_GUID:    newData[0],
				NEWS_ID:      id,
				NEWS_TS:      newData[2],
				NEWS_SUMMARY: newData[3],
			}

	}
}


func AssembleData(n *PolicyNewsOrg, news chan<- *PolicyNews, newsRegion chan<- *NewsRegion,
	newsCompany chan<- *NewsCompany, newsIndustry chan<- *NewsIndustry, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		wg2 = sync.WaitGroup{}
		regionList, companyList, industryList, newsList []string
		induLabs, compLabs, regiLabs string
	)
	_, regionMap := findmap.RegionRuntime()
	_, companyMap := findmap.CompanyRuntime()
	wg2.Add(3)
	go func() {
		defer wg2.Done()
		if len(n.RegionMap) == 0 {
			return
		}
		for r, count := range n.RegionMap{
			if r == "sum"{
				continue
			}
			region := regionMap[r]
			regionList = append(regionList, fmt.Sprintf(regionLabels, region[1], region[0]))
			relevance := float64(count) / float64(n.RegionMap["sum"])
			newsRegion <- &NewsRegion{
				REGION_LABEL_GUID: md5.MD5(n.NEWS_GUID + region[1]),
				NEWS_GUID:         n.NEWS_GUID,
				REGION_CODE:       region[1],
				REGION_NAME:       region[0],
				ENGLISH_NAME:      "",
				NEWS_ID:           n.NEWS_ID,
				NEWS_TS:           n.NEWS_TS,
				RELEVANCE:         fmt.Sprintf("%v", relevance),
				EMOTION_INDICATOR: n.EMOTION_INDICATOR,
				EMOTION_WEIGHT:    n.EMOTION_WEIGHT,
				EMOTION_DETAIL:    n.EMOTION_DETAIL,
			}
		}
		regiLabs = strings.Join(regionList, ",")
	}()
	go func() {
		defer wg2.Done()
		if len(n.CompanyMap) == 0{
			return
		}
		for c, count := range n.CompanyMap{
			if c == "sum"{
				continue
			}
			comp := companyMap[c]
			companyList = append(companyList, fmt.Sprintf(companyLabels, comp.CompGuid, comp.Name))
			relevance := float64(count) / float64(n.CompanyMap["sum"])
			newsCompany <- &NewsCompany{
				COMP_LABEL_GUID:   md5.MD5(n.NEWS_GUID + comp.CompGuid),
				NEWS_GUID:         n.NEWS_GUID,
				STOCK_CODE:        comp.StockCode,
				CREDIT_CODE:       comp.CreditCode,
				COMP_GUID:         comp.CompGuid,
				COMPANY_ID:        comp.CompanyId,
				CHINESE_NAME:      comp.Name,
				ENGLISH_NAME:      comp.NameEN,
				NEWS_ID:           n.NEWS_ID,
				NEWS_TS:           n.NEWS_TS,
				RELEVANCE:         fmt.Sprintf("%v", relevance),
				EMOTION_INDICATOR: n.EMOTION_INDICATOR,
				EMOTION_WEIGHT:    n.EMOTION_WEIGHT,
				EMOTION_DETAIL:    n.EMOTION_DETAIL,
			}
		}
		compLabs = strings.Join(companyList, ",")
	}()

	go func() {
		defer wg2.Done()
		if len(n.IndustryMap) == 0{
			return
		}
		for k, r := range n.IndustryMap{
			nameAndCode := strings.Split(k, "|")
			industryList = append(industryList, fmt.Sprintf(industryLabels, nameAndCode[0], nameAndCode[1]))
			newsIndustry <- &NewsIndustry{
				INDUSTRY_LABEL_GUID: md5.MD5(n.NEWS_GUID + nameAndCode[0]),
				NEWS_GUID:           n.NEWS_GUID,
				INDUSTRY_CODE:       nameAndCode[0],
				INDUSTRY_NAME:       nameAndCode[1],
				ENGLISH_NAME:        nameAndCode[2],
				NEWS_ID:             n.NEWS_ID,
				NEWS_TS:             n.NEWS_TS,
				RELEVANCE:           fmt.Sprintf("%v", r),
				EMOTION_INDICATOR:   n.EMOTION_INDICATOR,
				EMOTION_WEIGHT:      n.EMOTION_WEIGHT,
				EMOTION_DETAIL:      n.EMOTION_DETAIL,
			}
		}
		induLabs = strings.Join(industryList, ",")
	}()
	wg2.Wait()

	policyNews := &PolicyNews{
		NEWS_GUID:              n.NEWS_GUID,
		EMOTION_INDICATOR:      n.EMOTION_INDICATOR,
		EMOTION_INDICATOR_NAME: n.EMOTION_INDICATOR_NAME,
		EMOTION_WEIGHT:         n.EMOTION_WEIGHT,
		EMOTION_DETAIL:         n.EMOTION_DETAIL,
	}
	if len(regiLabs) != 0 {
		policyNews.REGION_LABELS = fmt.Sprintf(`[%s]`, regiLabs)
		newsList = append(newsList, regiLabs)
	}
	if len(compLabs) != 0 {
		policyNews.COMPANY_LABELS = fmt.Sprintf(`[%s]`, compLabs)
		newsList = append(newsList, compLabs)
	}
	if len(induLabs) != 0 {
		policyNews.INDUSTRY_LABELS = fmt.Sprintf(`[%s]`, induLabs)
		newsList = append(newsList, induLabs)
	}
	policyNews.NEWS_LABELS = fmt.Sprintf(`[%s]`, strings.Join(newsList, ","))

	news <- policyNews
}
