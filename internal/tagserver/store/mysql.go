package store

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/insertdb"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"strings"
	"sync"
)

func InsertRegion(newsRegionChan <-chan *NewsRegion, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		quotes                          []string
		insertValues                    []interface{}
		preamble, epilogue, oneQuoteSql = insertdb.GetInsertBaseSQLCode(&NewsRegion{}, "t_dmbe_news_region_label")
		beginLen                        = len(preamble) + len(epilogue)
	)

	for region := range newsRegionChan {
		v, l := insertdb.GetQuotesAndValues(region)
		if beginLen+l+len(oneQuoteSql) < 500000 {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	mysql.Transaction(SQl, insertValues...)
}

func InsertCompany(newsCompanyChan <-chan *NewsCompany, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		quotes                          []string
		insertValues                    []interface{}
		preamble, epilogue, oneQuoteSql = insertdb.GetInsertBaseSQLCode(&NewsCompany{}, "t_dmbe_news_company_label")
		beginLen                        = len(preamble) + len(epilogue)
	)

	for company := range newsCompanyChan {
		v, l := insertdb.GetQuotesAndValues(company)
		if beginLen+l+len(oneQuoteSql) < 500000 {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	mysql.Transaction(SQl, insertValues...)
}

func InsertIndustry(newsIndustryChan <-chan *NewsIndustry, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		quotes                          []string
		insertValues                    []interface{}
		preamble, epilogue, oneQuoteSql = insertdb.GetInsertBaseSQLCode(&NewsIndustry{}, "t_dmbe_news_industry_label")
		beginLen                        = len(preamble) + len(epilogue)
	)

	for industry := range newsIndustryChan {
		v, l := insertdb.GetQuotesAndValues(industry)
		if beginLen+l+len(oneQuoteSql) < 500000 {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	mysql.Transaction(SQl, insertValues...)
}

func UpdateNews(f *filter.Filter, newsChan <-chan *PolicyNews, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		idList                           []string
		sqlCode                          string
		updateFields, epilogue, fieldLen = insertdb.GetUpdateBaseSQLCode(&PolicyNews{})
		beginLen                         = len(epilogue)
	)
	sumLen := 0
	newsValue := make([]string, fieldLen)
	for news := range newsChan {
		updateValues := insertdb.GetWhenAndThen(news)
		f.WriteMap(news.NEWS_GUID)
		if sumLen+beginLen+len(idList)*len(news.NEWS_GUID) < 500000 {
			idList = append(idList, fmt.Sprintf(`'%s'`, news.NEWS_GUID))
			for i := 0; i < fieldLen; i++ {
				updateFields[i] = append(updateFields[i], updateValues[i])
				sumLen += len(updateValues[i])
			}

		} else {
			for index, data := range updateFields {
				newsValue[index] = strings.Join(data, ` `)
			}
			sqlCode = fmt.Sprintf(`UPDATE %s SET %s END %s (%s)`, `t_dmbe_policy_news_info`,
				strings.Join(newsValue, ` END, `), epilogue, strings.Join(idList, ", "))
			mysql.Transaction(sqlCode)
			f.SaveUrlKey()
			sumLen = 0
			idList = []string{}
			newsValue = make([]string, fieldLen)
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
	f.SaveUrlKey()
}
