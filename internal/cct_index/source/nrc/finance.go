package nrc

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/api"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/arithmetic"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/downloader"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/safeguard"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/poster"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// Pages 爬取的页数
	// 第一次爬取时将10页全部爬取下来
	// 以后就修改为前5页即可
	Pages = 10
)

// IndexInfo 待爬取的指标信息
// {2020 12 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202101/t20210128_1266089.html 2020年社会融资规模增量统计数据报告}
// {2020 09 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202010/t20201023_1248817.html 2020年前三季度社会融资规模增量统计数据报告}
// {2020 06 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202007/t20200730_1234986.html 2020年上半年社会融资规模增量统计数据报告}
// {2020 04 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202005/t20200527_1229010.html 2020年4月社会融资规模增量统计数据报告}
// {2020 03 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202004/t20200422_1226365.html 2020年一季度社会融资规模增量统计数据报告}
// {2020 12 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202101/t20210128_1266085.html 2020年全年存贷款情况}
// {2020 09 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202010/t20201023_1248814.html 2020年前三季度存贷款情况}
// {2021 06 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202107/t20210727_1291628.html 2021年上半年存贷款情况}
// {2021 03 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202104/t20210425_1277266.html 2021年一季度存贷款情况}
// {2021 03 https://www.ndrc.gov.cn/fggz/fgzh/gnjjjc/hbjr/202104/t20210425_1277266.html 2021年一季度存贷款情况}
type IndexInfo struct {
	Year       string
	Month      string
	DateTime   time.Time
	DateString string
	Href       string
	Title      string
	Value      float64
}

// 适用：
// 三个指标均计算的是月度同比
// - 人民币存款余额增速
// - 人民币贷款余额增速
// - 社会融资规模新增
func nrcFinance(ic *model.IndexConfig) (buffers []*model.Buffer) {
	var url string
	var indexInfos []IndexInfo
	for i := 0; i <= Pages; i++ {
		if i == 0 {
			url = api.MonetaryFinance
		} else {
			url = strings.ReplaceAll(api.MonetaryFinance, "index", fmt.Sprintf("index_%d", i))
		}

		body, err := downloader.SimpleGet(url)

		if err != nil {
			if !safeguard.IsNetworkNormal() {
				logger.Error("请检查服务器的网络是否能联通外网")
				return
			}
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		dom.Find("ul[class='u-list'] a").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Attr("href")
			title, _ := selection.Attr("title")
			if href == "" || title == "" {
				go poster.Poster(ic)
				return
			}
			if strings.Contains(title, "存贷款") || strings.Contains(title, "社会融资规模增量") {
				hrefFull := strings.ReplaceAll(api.MonetaryFinance, "index.html", href[2:])
				t, err := time.Parse("200601", href[2:8])
				if err != nil {
					go poster.Poster(ic)
					return
				}
				t = t.AddDate(0, -1, 0) // 获取一个月前的时间
				ii := IndexInfo{
					Year:       fmt.Sprintf("%d", t.Year()),
					Month:      fmt.Sprintf("%02d", t.Month()),
					DateTime:   t,
					DateString: fmt.Sprintf("%d%02d", t.Year(), t.Month()),
					Href:       hrefFull,
					Title:      title,
				}
				indexInfos = append(indexInfos, ii)
			}
		})
	}
	switch ic.TargetCode {
	case "HG00120":
		buffers = socialFinancingScale(ic, indexInfos)
	case "HG00118", "HG00119":
		buffers = rmbBalance(ic, indexInfos)
	}

	return
}

// 社会融资规模新增
func socialFinancingScale(ic *model.IndexConfig, iis []IndexInfo) (buffers []*model.Buffer) {
	pool := make(map[string]IndexInfo)
	for _, ii := range iis {
		if strings.Contains(ii.Title, "存贷款") {
			continue
		}
		body, err := downloader.SimpleGet(ii.Href)

		if err != nil {
			if !safeguard.IsNetworkNormal() {
				logger.Error("请检查服务器的网络是否能联通外网")
				return
			}
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		var paras []string
		dom.Find("div[class='article_con article_con_title']").Each(func(i int, selection *goquery.Selection) {
			paras = append(paras, selection.Text())
		})

		para := strings.Join(paras, "")
		para = strings.ReplaceAll(para, " ", "")
		para = strings.ReplaceAll(para, ",", "")
		para = strings.ReplaceAll(para, "，", "")
		para = strings.ReplaceAll(para, "份", "")

		if para == "" {
			go poster.Poster(ic)
			return
		}

		// 实际出现过的案例
		// 2022年11月份社会融资规模增量为1.99万亿元，比上年同期少6109亿元。
		// 2022年10月社会融资规模增量为9079亿元，比上年同期少7097亿元。
		// 9月份,社会融资规模增量为3.53万亿元，比上年同期多6245亿元。
		// 6月份，社会融资规模增量为5.17万亿元，比上年同期多1.47万亿元。
		// 12月份社会融资规模增量为1.59万亿元，比上年同期多33亿元。

		pat := `月社会融资规模增量为([0-9]*\.[0-9]+|[0-9]+).*?(万|千|百|十){0,1}亿元`
		re := regexp.MustCompile(pat)
		matched := re.FindAllStringSubmatch(para, -1)
		if matched == nil || len(matched) == 0 {
			logger.Warn(fmt.Sprintf("%s没有社会融资规模增量，该日期对应的网站：%s", ii.DateString, ii.Href))
			continue
		}
		v := matched[0][1] // 数值
		u := matched[0][2] // 单位
		vc, _ := strconv.ParseFloat(v, 64)
		switch u {
		case "万":
			vc *= 10000
		case "千":
			vc *= 1000
		case "百":
			vc *= 100
		case "十":
			vc *= 10
		}
		ii.Value = vc
		if _, ok := pool[ii.DateString]; !ok {
			pool[ii.DateString] = ii
		}
	}

	for _, indexInfo := range pool {
		curTime := indexInfo.DateTime
		preTime := curTime.AddDate(-1, 0, 0)
		preDate := preTime.Format("200601")
		if _, ok := pool[preDate]; ok {
			buffer := &model.Buffer{}
			buffer.Date = curTime.Format("200601")
			v := (indexInfo.Value - pool[preDate].Value) / pool[preDate].Value * 100
			buffer.TargetValue = arithmetic.TrimZero(fmt.Sprintf("%.2f", v))
			buffers = append(buffers, buffer)
		}
	}

	return
}

func rmbBalance(ic *model.IndexConfig, iis []IndexInfo) (buffers []*model.Buffer) {
	for _, ii := range iis {
		if strings.Contains(ii.Title, "社会融资规模增量") {
			continue
		}

		body, err := downloader.SimpleGet(ii.Href)

		if err != nil {
			if !safeguard.IsNetworkNormal() {
				logger.Error("请检查服务器的网络是否能联通外网")
				return
			}
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Error(err.Error())
			go poster.Poster(ic)
			return
		}

		var paras []string
		dom.Find("div[class='article_con article_con_title']").Each(func(i int, selection *goquery.Selection) {
			paras = append(paras, selection.Text())
		})

		para := strings.Join(paras, "")
		para = strings.ReplaceAll(para, " ", "")
		para = strings.ReplaceAll(para, ",", "，")
		para = strings.ReplaceAll(para, "份，", "份")

		if para == "" {
			go poster.Poster(ic)
			return
		}

		// 月末人民币贷款余额169.37万亿元，同比增长13%，
		// 月末人民币存款余额211.08万亿元，同比增长10.7%，
		// 人民币贷款余额203.54万亿元，同比增长11%，
		// 人民币存款余额246.22万亿元，同比增长10.5%，

		var pattern string
		switch ic.TargetCode {
		case "HG00118":
			pattern = "人民币存款余额(.*?)元，同比增长(.*?)%，"
		case "HG00119":
			pattern = "人民币贷款余额(.*?)元，同比增长(.*?)%，"
		}

		re := regexp.MustCompile(pattern)
		matched := re.FindAllStringSubmatch(para, -1)
		if matched == nil || len(matched) == 0 {
			logger.Warn(fmt.Sprintf("%s没有对应的人民币存款、贷款余额，该日期对应的网站：%s", ii.DateString, ii.Href))
			continue
		}

		buffer := &model.Buffer{}
		buffer.Date = ii.DateString
		buffer.TargetValue = matched[0][2]
		buffers = append(buffers, buffer)
	}

	return
}
