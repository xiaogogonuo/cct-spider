package insertdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/findmap"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/subString"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"net/http"
	"strings"
	"time"
)

type DataInfo struct {
	DBName     string
	PolicyCode string
	PolicyName string
}

var rep *request.Request

func init() {
	rep = &request.Request{
		Url:    "http://106.37.165.121/inf/chengtong/dm/be/policyNewsInfo/saveRequest", // TODO: changeable
		Method: http.MethodPost,
	}
}

func (di *DataInfo) InsertIntoSQL(f *filter.Filter, message <-chan *callback.Message) {
	var (
		t                               = time.Now().Format("20060102")
		preamble, epilogue, oneQuoteSql = GetInsertBaseSQLCode(&callback.SqlValues{}, di.DBName)
		region                          string
		regionCode                      string
		quotes                          []string
		insertValues                    []interface{}
		beginLen                        = len(preamble) + len(epilogue)
		pullServer                      []callback.SqlValues
	)
	for mes := range message {
		tLen := len(mes.Title)
		sLen := len(mes.Summary)
		if (tLen == 0 && sLen == 0) || tLen+sLen < 15 {
			continue
		}
		if mes.Date == "" || mes.Date > t {
			mes.Date = "10000101"
		}
		if len(mes.Summary) > 65535 {
			n, _ := subString.RuneIndex([]byte(mes.Summary), 65535/3)
			mes.Summary = mes.Summary[:n]
		}
		regionPat, regionMap := findmap.RegionRuntime()
		if r := findmap.FindOne(regionPat, mes.Summary); r != "" {
			region = regionMap[r][0]
			regionCode = regionMap[r][1]
		}
		guid := md5.MD5(mes.Url)
		sqlValues := &callback.SqlValues{
			NEWS_GUID:        guid,
			NEWS_TITLE:       mes.Title,
			NEWS_TS:          mes.Date,
			NEWS_URL:         mes.Url,
			NEWS_SOURCE:      mes.Source,
			NEWS_SOURCE_CODE: mes.SourceCode,
			NEWS_SUMMARY:     mes.Summary,
			POLICY_TYPE:      di.PolicyCode,
			POLICY_TYPE_NAME: di.PolicyName,
			REGION_NAME:      region,
			REGION_CODE:      regionCode,
			IS_CONTROL:       "N",
			IS_INVEST:        "N",
			IS_DEPOSIT:       "N",
			IS_FUND:          "N",
			IS_STOCK:         "N",
			IS_FINANCE:       "N",
			IS_INDUSTRY:      "N",
			IS_CAPITAL:       "N",
			NEWS_GYS_CODE:    "90",
			NEWS_GYS_NAME:    "爬虫",
		}
		f.WriteMap(guid)
		v, l := GetQuotesAndValues(sqlValues)
		if beginLen+l+len(oneQuoteSql) < 500000 {
			insertValues = append(insertValues, v...)
			pullServer = append(pullServer, *sqlValues)
			quotes = append(quotes, oneQuoteSql)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			pullService(pullServer)
			f.SaveUrlKey()
			insertValues = append([]interface{}{}, v...)
			pullServer = append([]callback.SqlValues{}, *sqlValues)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	if len(insertValues) == 0 {
		return
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	mysql.Transaction(SQl, insertValues...)
	pullService(pullServer)
	f.SaveUrlKey()
}

func pullService(info []callback.SqlValues) {
	postData := map[string][]callback.SqlValues{"data": info}
	m, _ := json.Marshal(postData)
	rep.Body = bytes.NewReader(m)
	_, err := rep.Visit()
	if err != nil {
		return
	}

}
