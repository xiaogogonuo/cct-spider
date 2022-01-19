package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/insertdb"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"net/http"
	"strings"
	"sync"
)

var rep *request.Request

// 生产机
var regionAPI = "http://106.37.165.121/inf/chengtong/py/sy/newsRegionLabel/saveRequest"
var companyAPI = "http://106.37.165.121/inf/chengtong/py/sy/newsCompanyLabel/saveRequest"
var industryAPI = "http://106.37.165.121/inf/chengtong/py/sy/newsIndustryLabel/saveRequest"
var updateNewsAPI = "http://106.37.165.121/inf/chengtong/py/sy/policyNewsInfo/updateRequest"
var maxByte = 50000

func init() {
	rep = &request.Request{
		Method: http.MethodPost,
		Header: map[string]string{"Content-Type": "application/json"},
	}
}

func InsertRegion(newsRegionChan <-chan *NewsRegion, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		quotes                          []string
		insertValues                    []interface{}
		preamble, epilogue, oneQuoteSql = insertdb.GetInsertBaseSQLCode(&NewsRegion{}, "t_dmbe_news_region_label")
		beginLen                        = len(preamble) + len(epilogue)
		regionServer                    []NewsRegion
	)

	for region := range newsRegionChan {
		if len(region.EMOTION_INDICATOR) == 0{
			continue
		}
		v, l := insertdb.GetQuotesAndValues(region)
		if beginLen+l+len(oneQuoteSql) < maxByte {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			regionServer = append(regionServer, *region)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			//pullService(regionServer)
			//mysql.Transaction(SQl, insertValues...)
			recursion(regionServer, SQl, insertValues, 0)
			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
			regionServer = append([]NewsRegion{}, *region)
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	//pullService(regionServer)
	//mysql.Transaction(SQl, insertValues...)
	recursion(regionServer, SQl, insertValues, 0)


}

func InsertCompany(newsCompanyChan <-chan *NewsCompany, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		quotes                          []string
		insertValues                    []interface{}
		preamble, epilogue, oneQuoteSql = insertdb.GetInsertBaseSQLCode(&NewsCompany{}, "t_dmbe_news_company_label")
		beginLen                        = len(preamble) + len(epilogue)
		companyServer                    []NewsCompany
	)

	for company := range newsCompanyChan {
		if len(company.EMOTION_INDICATOR) == 0{
			continue
		}
		v, l := insertdb.GetQuotesAndValues(company)
		if beginLen+l+len(oneQuoteSql) < maxByte {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			companyServer = append(companyServer, *company)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			//pullService(companyServer)
			//mysql.Transaction(SQl, insertValues...)
			recursion(companyServer, SQl, insertValues, 0)

			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			companyServer = append([]NewsCompany{}, *company)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	//pullService(companyServer)
	//mysql.Transaction(SQl, insertValues...)
	recursion(companyServer, SQl, insertValues, 0)


}

func InsertIndustry(newsIndustryChan <-chan *NewsIndustry, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		quotes                          []string
		insertValues                    []interface{}
		preamble, epilogue, oneQuoteSql = insertdb.GetInsertBaseSQLCode(&NewsIndustry{}, "t_dmbe_news_industry_label")
		beginLen                        = len(preamble) + len(epilogue)
		industryServer                  []NewsIndustry
	)

	for industry := range newsIndustryChan {
		if len(industry.EMOTION_INDICATOR) == 0{
			continue
		}
		v, l := insertdb.GetQuotesAndValues(industry)
		if beginLen+l+len(oneQuoteSql) < maxByte {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			industryServer = append(industryServer, *industry)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			//pullService(industryServer)
			//mysql.Transaction(SQl, insertValues...)
			recursion(industryServer, SQl, insertValues, 0)

			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			industryServer = append([]NewsIndustry{}, *industry)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	//pullService(industryServer)
	//mysql.Transaction(SQl, insertValues...)
	recursion(industryServer, SQl, insertValues, 0)

}

func UpdateNews(f *filter.Filter, newsChan <-chan *PolicyNews, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		idList                           []string
		sqlCode                          string
		updateFields, epilogue, fieldLen = insertdb.GetUpdateBaseSQLCode(&PolicyNews{})
		beginLen                         = len(epilogue)
		newsServer                       []PolicyNews

	)
	sumLen := 0
	newsValue := make([]string, fieldLen)
	for news := range newsChan {
		if len(news.EMOTION_INDICATOR) == 0{
			continue
		}
		updateValues := insertdb.GetWhenAndThen(news)
		f.WriteMap(news.NEWS_GUID)
		if sumLen+beginLen+len(idList)*len(news.NEWS_GUID) < maxByte {
			idList = append(idList, fmt.Sprintf(`'%s'`, news.NEWS_GUID))
			for i := 0; i < fieldLen; i++ {
				updateFields[i] = append(updateFields[i], updateValues[i])
				sumLen += len(updateValues[i])
			}
			newsServer = append(newsServer, *news)

		} else {
			for index, data := range updateFields {
				newsValue[index] = strings.Join(data, ` `)
			}
			sqlCode = fmt.Sprintf(`UPDATE %s SET %s END %s (%s)`, `t_dmbe_policy_news_info`,
				strings.Join(newsValue, ` END, `), epilogue, strings.Join(idList, ", "))
			//pullService(newsServer)
			//mysql.Transaction(sqlCode)
			//f.SaveUrlKey()
			recursionSaveIdKey(newsServer, sqlCode, f, 0)
			sumLen = 0
			idList = []string{}
			newsValue = make([]string, fieldLen)
			newsServer = append([]PolicyNews{}, *news)
			updateFields, epilogue, fieldLen = insertdb.GetUpdateBaseSQLCode(&PolicyNews{})

		}
	}
	if len(updateFields) == 0 {
		return
	}
	for index, data := range updateFields {
		newsValue[index] = strings.Join(data, ` `)
	}
	sqlCode = fmt.Sprintf(`UPDATE %s SET %s END %s (%s)`, `t_dmbe_policy_news_info`,
		strings.Join(newsValue, ` END, `), epilogue, strings.Join(idList, ", "))
	//pullService(newsServer)
	//mysql.Transaction(sqlCode)
	//f.SaveUrlKey()
	recursionSaveIdKey(newsServer, sqlCode, f, 0)

}


func pullService(info interface{}) bool {
	var m []byte
	switch info.(type) {
	case []NewsRegion:
		postData := map[string][]NewsRegion{"data": info.([]NewsRegion)}
		m, _ = json.Marshal(postData)
		rep.Url = regionAPI
	case []NewsCompany:
		postData := map[string][]NewsCompany{"data": info.([]NewsCompany)}
		m, _ = json.Marshal(postData)
		rep.Url = companyAPI

	case []NewsIndustry:
		postData := map[string][]NewsIndustry{"data": info.([]NewsIndustry)}
		m, _ = json.Marshal(postData)
		rep.Url = industryAPI

	case []PolicyNews:
		postData := map[string][]PolicyNews{"data": info.([]PolicyNews)}
		m, _ = json.Marshal(postData)
		rep.Url = updateNewsAPI
	}
	rep.Body = bytes.NewReader(m)
	b, err := rep.Visit()
	if err != nil {
		return false
	}
	fmt.Println(string(b))
	return true
}

func recursionSaveIdKey(info interface{}, SQl string, f *filter.Filter, reqNum int) {
	reqNum ++
	if pullService(info) {
		mysql.Transaction(SQl)
		f.SaveUrlKey()
		return
	} else {
		if reqNum >= 5 {
			return
		}
		recursionSaveIdKey(info, SQl, f, reqNum)
	}
}

func recursion(info interface{}, SQl string, insertValues []interface{}, reqNum int) {
	reqNum ++
	if pullService(info) {
		mysql.Transaction(SQl, insertValues...)
		return
	} else {
		if reqNum >= 5 {
			return
		}
		recursion(info, SQl, insertValues, reqNum)
	}
}