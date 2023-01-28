package ipaddr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var Website = map[string]string{
	"卓创资讯_造纸行业价格指数|原油价格指数": "https://index.sci99.com/api/zh-cn/dataitem/datavalue",
	"新浪财经_人民币汇率|美元指数":      "https://vip.stock.finance.sina.com.cn/forex/api/jsonp.php/_/NewForexService.getDayKLine?symbol=CNYUSD",
	"凤凰网财经_国债指数":           "http://app.finance.ifeng.com/hq/stock_daily.php?code=sh000012&begin_day=2021-01-01&end_day=2021-12-31",
	"东方财富网_宏观指标":           "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GJZB&sty=ZGZB&p=2&ps=10&mkt=0",
	"东方财富网_行业指标":           "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=5&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE&filter=(INDICATOR_ID%3D%22EMI00107664%22)",
	"东方财富网_上海银行间同业拆放利率":    "http://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99228&cu=cny&type=009023&p=1",
	"中华人民共和国工业和信息化部_":      "https://www.miit.gov.cn/search-front-server/api/search/info?websiteid=110000000000000&scope=basic&q=&pg=15&cateid=57&pos=title_text%2Cinfocontent%2Ctitlepy&selectFields=title,url,&group=distinct&level=6&sortFields=%5B%7B%22name%22%3A%22deploytime%22%2C%22type%22%3A%22desc%22%7D%5D&p=1",
	"国家市场监督管理总局_":          "http://www.samr.gov.cn/zw/wjfb/zdjd/index.html",
	"中华人民共和国生态环境部_":        "http://www.mee.gov.cn/zcwj/",
	"中国银行保险监督管理委员会_928":    "https://www.irc.gov.cn/cbircweb/DocInfo/SelectDocByItemIdAndChild?itemId=928&pageSize=20&pageIndex=1",
	"中国银行保险监督管理委员会_927":    "https://www.irc.gov.cn/cbircweb/DocInfo/SelectDocByItemIdAndChild?itemId=927&pageSize=20&pageIndex=1",
	"中国物流与采购协会_":           "http://www.chinawuliu.com.cn/xsyj/tjsj/",
	"中国造纸协会_":              "http://man.chinappi.org/f2e051061e47815ae6b3c46c692915.msyl?l=dca1f2&t=dca1f2&h=www.chinappi.org&p=1",
	"安徽省人民政府_":             "https://www.ah.gov.cn/public/column/1681?type=6&typeNode=6&nav=0",
	"北京市人民政府_":             "http://www.beijing.gov.cn/zhengce/zhengcefagui/",
	"重庆市人民政府_":             "http://www.cq.gov.cn/zwgk/zfxxgkml/szfwj/",
	"广东省人民政府_":             "http://www.gd.gov.cn/gkmlpt/api/all/5?page=1&sid=2",
	"上海市人民政府_":             "https://www.shanghai.gov.cn/nw12344/",
	"天津市人民政府_":             "http://www.tj.gov.cn/zwgk/szfwj/",
}

// PostData 卓创资讯接口的请求入参
type PostData struct {
	HY    string `json:"hy"`
	Level string `json:"level"`
	Path1 string `json:"path1"`
	Path2 string `json:"path2"`
	Path3 string `json:"path3"`
	Path4 string `json:"path4"`
	Type  string `json:"type"`
}

// getApiDomain
// 获取Api的域名
func getApiDomain(api string) string {
	u, err := url.Parse(api)
	if err != nil {
		return ""
	}
	return u.Host
}

var IPPool map[string]struct{}

func init() {
	IPPool = make(map[string]struct{})
	row := mysql.Query("SELECT IP FROM ipaddr")
	for _, r := range row {
		IPPool[r[0]] = struct{}{}
	}
}

func extractIPAndPortFromError(err error) (ip, port string) {
	errString := err.Error()
	reg := regexp.MustCompile("dial tcp (.*):(.*)?: i/o timeout")
	ret := reg.FindStringSubmatch(errString)
	if len(ret) < 3 || ret == nil {
		return
	}
	ip = ret[1]
	port = ret[2]
	return
}

var sqlData = "INSERT INTO ipaddr (ID, IP, Domain, Name, Port) VALUES ('%s', '%s', '%s', '%s', '%s');"

func postDetect(name, api string, body io.Reader, header map[string]string) {
	trace := &httptrace.ClientTrace{
		ConnectStart: func(_, addr string) {
			if strings.Contains(addr, "[") {
				return
			}
			addrS := strings.Split(addr, ":")
			ip := addrS[0]
			if _, ok := IPPool[ip]; !ok {
				IPPool[ip] = struct{}{}
				port := addrS[1]
				sql := fmt.Sprintf(sqlData, md5.MD5(ip), ip, getApiDomain(api), name, port)
				mysql.Transaction(sql)
			}
		},
	}
	req, err := http.NewRequest(http.MethodPost, api, body)
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	req = req.WithContext(httptrace.WithClientTrace(context.Background(), trace))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ip, port := extractIPAndPortFromError(err)
		if ip == "" {
			return
		}
		if _, ok := IPPool[ip]; !ok {
			IPPool[ip] = struct{}{}
			sql := fmt.Sprintf(sqlData, md5.MD5(ip), ip, getApiDomain(api), name, port)
			mysql.Transaction(sql)
		}
		return
	}
	defer resp.Body.Close()
}

func getDetect(name, api string) {
	trace := &httptrace.ClientTrace{
		ConnectStart: func(_, addr string) {
			if strings.Contains(addr, "[") {
				return
			}
			addrS := strings.Split(addr, ":")
			ip := addrS[0]
			if _, ok := IPPool[ip]; !ok {
				IPPool[ip] = struct{}{}
				port := addrS[1]
				sql := fmt.Sprintf(sqlData, md5.MD5(ip), ip, getApiDomain(api), name, port)
				mysql.Transaction(sql)
			}
		},
	}
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return
	}
	req = req.WithContext(httptrace.WithClientTrace(context.Background(), trace))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ip, port := extractIPAndPortFromError(err)
		if ip == "" {
			return
		}
		if _, ok := IPPool[ip]; !ok {
			IPPool[ip] = struct{}{}
			sql := fmt.Sprintf(sqlData, md5.MD5(ip), ip, getApiDomain(api), name, port)
			mysql.Transaction(sql)
		}
		return
	}
	defer resp.Body.Close()
}

func Run() {
	header := map[string]string{"Content-Type": "application/json"}
	for {
		for k, api := range Website {
			name := strings.Split(k, "_")[0]
			switch name {
			case "卓创资讯":
				postData := PostData{
					HY:    "造纸",
					Level: "0",
					Path1: "造纸行业价格指数",
					Path2: "",
					Path3: "",
					Path4: "",
					Type:  "2",
				}
				m, _ := json.Marshal(postData)
				postDetect(name, api, bytes.NewReader(m), header)
			case "中国外汇交易中心":
				postDetect(name, api, nil, header)
			default:
				getDetect(name, api)
			}
		}
		time.Sleep(time.Hour)
	}
}
