package industry

import (
	"bytes"
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/store"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/http"
	"sort"
	"sync"
)

var rep *request.Request

func init() {
	rep = &request.Request{
		Url:     "http://172.17.0.23:9090/clsIndustry",
		Method:  http.MethodPost,
		Header: map[string]string{"Content-Type": "application/json" },
		//Timeout: time.Second * 2,
	}
}

type M [][]interface{}

func (m M) Len() int {
	return len(m)
}

func (m M) Less(i, j int) bool {
	return m[i][1].(float64) > m[j][1].(float64)
}

func (m M) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func GetIndustry(n *store.PolicyNewsOrg, wg *sync.WaitGroup) {
	defer wg.Done()
	var m M
	rep.Body = bytes.NewBuffer([]byte(n.NEWS_SUMMARY))
	b, err := rep.Visit()
	if err != nil || len(b)== 0 {
		return
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	sort.Sort(m)
	n.IndustryMap = make(map[string]float64)
	for _, v := range m {
		n.IndustryMap[v[0].(string)] = v[1].(float64)
	}
}
