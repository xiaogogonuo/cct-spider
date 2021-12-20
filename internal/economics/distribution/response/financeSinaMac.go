package response

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/economics/backend"
	"github.com/xiaogogonuo/cct-spider/internal/economics/core"
	"github.com/xiaogogonuo/cct-spider/internal/economics/distribution/request"
	"github.com/xiaogogonuo/cct-spider/internal/economics/pkg/decoder"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"sync/atomic"
)

// 新浪财经解析器

// FinanceSinaMac 新浪财经-中国宏观经济数据-响应体结构
type FinanceSinaMac [][]string

// FinanceSinaMacParse 新浪财经-中国宏观经济数据-解析器
// 适用：制造业PMI、
func FinanceSinaMacParse(res *core.Response) {
	body, err := decoder.GBK2UTF8(res.Body)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	body = extractData(body)
	// no next page if body is nil
	if body == nil {
		atomic.AddUint64(&Stop, 1)
		return
	}
	var record FinanceSinaMac
	if err := json.Unmarshal(body, &record); err != nil {
		logger.Error(err.Error())
		return
	}
	// TODO: push data to RespondChannel is different
	FinanceSinaMacPipeline(record)
	// 构造下一页请求
	res.Page++
	u := request.NewFinanceSinaURL()
	u.From = res.Page * 31
	u.Event = res.Meta.SourceTargetCode
	fmt.Println(u.ToURL())
	req := core.NewRequest(u.ToURL())
	req.Page = res.Page
	req.Parse = FinanceSinaMacParse
	req.Meta = res.Meta
	RequestChannel <- req
}

// FinanceSinaMacPipeline 新浪财经-中国宏观经济数据-数据构建管道
func FinanceSinaMacPipeline(record [][]string) {
	for _, r := range record {
		var bk backend.BackEnd
		bk.TargetValue = r[1]
		switch len(r[0]) {
		case 6:
			bk.AcctDate = strings.ReplaceAll(r[0], ".", "0")
		case 7:
			bk.AcctDate = strings.ReplaceAll(r[0], ".", "")
		}
		RespondChannel <- bk
	}
}

// 抽取有效数据
func extractData(body []byte) []byte {
	s := string(body)
	lastCountIndex := strings.LastIndex(s, "count")
	lastDataIndex := strings.LastIndex(s, "data")
	if lastDataIndex > lastCountIndex {
		return []byte(s[lastDataIndex+5:len(s)-4])
	}
	return nil
}