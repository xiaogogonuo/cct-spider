package sentiment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/store"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"net/http"
	"sort"
	"strings"
	"sync"
)

var (
	rep  *request.Request
	sMap map[string]string
)

func init() {
	rep = &request.Request{
		Url:     "http://127.0.0.1:9090/clsSentiment",
		Method:  http.MethodPost,
		//Timeout: time.Second * 2,
	}
	sMap = make(map[string]string)
	sMap["0"] = "中性"
	sMap["1"] = "正面"
	sMap["2"] = "负面"
}

type M [][]float64

func (m M) Len() int {
	return len(m)
}

func (m M) Less(i, j int) bool {
	return m[i][1] > m[j][1]
}

func (m M) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func GetSentiment(n *store.PolicyNewsOrg, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		m     M
		sList []string
	)
	rep.Body = bytes.NewBuffer([]byte(n.NEWS_SUMMARY))
	b, err := rep.Visit()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		logger.Error(err.Error())
	}
	sort.Sort(m)
	s := fmt.Sprintf("%v", m[0][0])
	for _, v := range m {
		sList = append(sList, fmt.Sprintf(`%v=%f`, v[0], v[1]))
	}
	n.EMOTION_INDICATOR = s
	n.EMOTION_INDICATOR_NAME = sMap[s]
	n.EMOTION_WEIGHT = fmt.Sprintf(`%v`, m[0][1])
	n.EMOTION_DETAIL = "{" + strings.Join(sList, ", ") + "}"
}
