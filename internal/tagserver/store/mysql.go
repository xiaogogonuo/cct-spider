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
var regionAPI = "http://106.37.165.121/inf/dm/be/policyNewsInfo/saveRequest"
var companyAPI = "http://106.37.165.121/inf/dm/be/policyNewsInfo/saveRequest"
var industryAPI = "http://106.37.165.121/inf/dm/be/policyNewsInfo/saveRequest"
var maxByte = 500000

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
		v, l := insertdb.GetQuotesAndValues(region)
		if beginLen+l+len(oneQuoteSql) < maxByte {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			regionServer = append(regionServer, *region)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			pullService(regionServer)
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
	mysql.Transaction(SQl, insertValues...)
	pullService(regionServer)
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
		v, l := insertdb.GetQuotesAndValues(company)
		if beginLen+l+len(oneQuoteSql) < maxByte {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			companyServer = append(companyServer, *company)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			pullService(companyServer)
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
	mysql.Transaction(SQl, insertValues...)
	pullService(companyServer)
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
		v, l := insertdb.GetQuotesAndValues(industry)
		if beginLen+l+len(oneQuoteSql) < maxByte {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			industryServer = append(industryServer, *industry)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			pullService(industryServer)
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
	mysql.Transaction(SQl, insertValues...)
	pullService(industryServer)
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
			mysql.Transaction(sqlCode)
			pullService(newsServer)
			f.SaveUrlKey()
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
	mysql.Transaction(sqlCode)
	pullService(newsServer)
	f.SaveUrlKey()
}


func pullService(info interface{}) {
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
	}
	rep.Body = bytes.NewReader(m)
	b, err := rep.Visit()
	if err != nil {
		return
	}
	fmt.Println(string(b))
}