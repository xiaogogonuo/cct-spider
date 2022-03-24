package sasac

// 国资委

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/group/urls"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const domain = "http://www.sasac.gov.cn/"

func downloader(url string) (body []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

type News struct {
	URL   string   // 新闻链接
	Date  string   // 新闻日期
	Body  string   // 新闻正文
	Title string   // 新闻标题
	Image [][]byte // 新闻图片
	Plate string   // 新闻所属板块
}

// 从每个板块的第一页提取总页数
func getMaxPageFromFirstPage(html string) (maxPage int64, err error) {
	reg := regexp.MustCompile("maxPageNum = ([0-9]+);")
	match := reg.FindStringSubmatch(html)[1]
	return strconv.ParseInt(match, 10, 64)
}

// 从每个板块的第一页提取pag id
func getPagIDFromFirstPage(html string) (pagID int64, err error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}
	var id string
	dom.Find("td[class='pages']").Each(func(i int, selection *goquery.Selection) {
		id, _ = selection.Attr("id")
	})
	if id == "" || !strings.Contains(id, "_") {
		return
	}
	id = strings.Split(id, "_")[1]
	return strconv.ParseInt(id, 10, 64)
}

// 从每个板块的第一页提取新闻URL、新闻标题
func getNewsURLAndNewsTitleFromFirstPage(html, plate string) (news []News, err error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}
	dom.Find("div[class='zsy_conlist'] li a").Each(func(i int, selection *goquery.Selection) {
		var ns News
		href, _ := selection.Attr("href")
		if strings.HasPrefix(href, "http") {
			ns.URL = href
		} else {
			ns.URL = domain + strings.Replace(href, "../", "", -1)
		}
		ns.Title = selection.Text()
		ns.Plate = plate
		news = append(news, ns)
	})
	return
}

type Pager struct {
	URL   string
	Plate string
}

// 对每个板块构造从第二页到最后一页的请求
func structurePageURLFromSecondToLast(pagID, maxPage int64, plate, base string) (pages []Pager) {
	for i := int(maxPage) - 1; i > int(maxPage) - 1 - HowManyPages; i-- {
		next := fmt.Sprintf("_%d_%d.html", pagID, i)
		next = strings.Replace(base, ".html", next, -1)
		var pager Pager
		pager.URL = next
		pager.Plate = plate
		pages = append(pages, pager)
	}
	return
}

// 获取新闻详情
func getNewsDetail(wg *sync.WaitGroup, semaphore chan struct{}, ch chan News, news News) {
	defer wg.Done()
	defer func() { <-semaphore }()
	body, err := downloader(news.URL)
	if err != nil {
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return
	}
	var ns News
	var content []string
	dom.Find("div>p").Each(func(i int, selection *goquery.Selection) {
		content = append(content, selection.Text())
	})
	dom.Find("div>p img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		var imageURL string
		if strings.Contains(src, "http") {
			imageURL = src
		} else if strings.HasPrefix(src, "../") {
			imageURL = domain + strings.Replace(src, "../", "", -1)
		}
		body, e := downloader(imageURL)
		if e != nil {
			logger.Error(e.Error())
		} else {
			ns.Image = append(ns.Image, body)
		}
	})

	dom.Find("meta[name=publishdate]").Each(func(i int, selection *goquery.Selection) {
		attr, ok := selection.Attr("content")
		if ok {
			attr = attr[:10]
			ns.Date = attr
		}
	})
	dom.Find("meta[name='firstpublishedtime']").Each(func(i int, selection *goquery.Selection) {
		attr, ok := selection.Attr("content")
		if ok {
			attr = attr[:10]
			ns.Date = attr
		}
	})

	ns.Title = news.Title
	ns.Plate = news.Plate
	ns.URL = news.URL
	ns.Body = strings.Join(content, " ")
	ch <- ns
}

type ND struct {
	mu         sync.Mutex
	newsSimple []News
	newsDetail []News
}

// 从每个板块的第二页到最后一页提取新闻URL、新闻标题
func (nd *ND) getNewsURLAndNewsTitleFromNextPage(wg *sync.WaitGroup, semaphore chan struct{}, pager Pager) {
	defer wg.Done()
	defer func() { <-semaphore }()
	body, err := downloader(pager.URL)
	if err != nil {
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return
	}
	dom.Find("li a").Each(func(i int, selection *goquery.Selection) {
		var ns News
		href, _ := selection.Attr("href")
		if strings.HasPrefix(href, "http") {
			ns.URL = href
		} else {
			ns.URL = domain + strings.Replace(href, "../", "", -1)
		}
		title := selection.Text()
		ns.Title = title
		ns.Plate = pager.Plate
		nd.mu.Lock()
		nd.newsSimple = append(nd.newsSimple, ns)
		nd.mu.Unlock()
	})
}

type Filter struct {
	name string
	file *os.File
	line map[string]bool
}

func NewFilter(name string) *Filter {
	return &Filter{name: name, line: make(map[string]bool)}
}

func (f *Filter) Open() {
	file, err := os.OpenFile(f.name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	f.file = file
}

func (f *Filter) ReadAll() {
	reader := bufio.NewReader(f.file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		f.line[strings.Trim(line, "\n")] = true
	}
}

func (f *Filter) AppendNewLine(newline []string) {
	writer := bufio.NewWriter(f.file)
	for _, line := range newline {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		_ = writer.Flush()
	}
}

var HowManyPages = 1

func EntryPoint() {
	filter := NewFilter("sasNewsURL.txt")
	filter.Open()
	filter.ReadAll()

	// 新闻池
	var newsPool []News
	// 页码池：存放每个板块从第二页到最后一页的请求
	var pagePool []Pager

	startURL := map[string]string{
		"时政要闻": "http://www.sasac.gov.cn/n2588025/n2643309/index.html",
		"领导活动": "http://www.sasac.gov.cn/n2588025/n2643314/index.html",
		"新闻":   "http://www.sasac.gov.cn/n2588025/n2588119/index.html",
		"评论":   "http://www.sasac.gov.cn/n2588025/n2588134/index.html",
		"央企联播": "http://www.sasac.gov.cn/n2588025/n2588124/index.html",
		"地方扫描": "http://www.sasac.gov.cn/n2588025/n2588129/index.html",
		"媒体观察": "http://www.sasac.gov.cn/n2588025/n2588139/index.html",
	}

	for k, v := range startURL {
		// 下载每个板块的第一页
		body, err := downloader(v)
		if err != nil {
			logger.Error(fmt.Sprintf("%s: %s", k, err.Error()))
			continue
		}

		// 从每个板块的第一页提取总页数
		maxPage, err := getMaxPageFromFirstPage(string(body))
		if err != nil {
			logger.Error(fmt.Sprintf("%s: not find max page", k))
			continue
		}

		// 从每个板块的第一页提取pag id
		pagID, err := getPagIDFromFirstPage(string(body))
		if err != nil {
			logger.Error(fmt.Sprintf("%s: %s", k, err.Error()))
			continue
		}

		// 从每个板块的第一页提取新闻URL、新闻标题
		news, err := getNewsURLAndNewsTitleFromFirstPage(string(body), k)
		if err != nil {
			logger.Error(fmt.Sprintf("%s: %s", k, err.Error()))
			continue
		}
		// 将从每个板块的第一页提取新闻URL、新闻标题添加到新闻池
		newsPool = append(newsPool, news...)

		// 对每个板块构造从第二页到最后一页的请求
		pages := structurePageURLFromSecondToLast(pagID, maxPage, k, v)
		// 将对每个板块构造从第二页到最后一页的请求加入页码池
		pagePool = append(pagePool, pages...)
	}

	var nd = &ND{}
	nd.newsSimple = append(nd.newsSimple, newsPool...)

	var wg = &sync.WaitGroup{}
	var newsDetail = make(chan News)
	var semaphorePage = make(chan struct{}, 10)
	var semaphoreNews = make(chan struct{}, 10)

	go func() {
		wg.Wait()
		close(newsDetail)
	}()

	// 并发请求各板块的每一页，提取新闻URL和新闻标题
	for _, pager := range pagePool {
		semaphorePage <- struct{}{}
		wg.Add(1)
		go nd.getNewsURLAndNewsTitleFromNextPage(wg, semaphorePage, pager)
	}

	// 并发请求新闻详情页
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, news := range nd.newsSimple {
			semaphoreNews <- struct{}{}
			wg.Add(1)
			go getNewsDetail(wg, semaphoreNews, newsDetail, news)
		}
	}()

	var groupNewsList []GroupNews
	var newNews []string
	for ns := range newsDetail {
		if _, ok := filter.line[ns.URL]; ok {
			continue
		}
		filter.line[ns.URL] = true
		newNews = append(newNews, ns.URL)

		var gn GroupNews
		gn.NewsGuid = md5.MD5(ns.URL)
		gn.NewsTitle = ns.Title
		gn.NewsTs = ns.Date
		gn.NewsUrl = ns.URL
		gn.NewsSource = "国务院国有资产监督管理委员会"
		gn.NewsSourceCode = "WEB_01024"
		gn.NewsSummary = ns.Body
		gn.PolicyType = "10"
		gn.PolicyTypeName = "国家政策"
		gn.NewsGysCode = "90"
		gn.NewsGysName = "爬虫"
		gn.NewsId = 0
		gn.Image = ns.Image
		groupNewsList = append(groupNewsList, gn)
	}
	filter.AppendNewLine(newNews)
	for _, gn := range groupNewsList {
		if _, ok := urls.URLS[gn.NewsUrl]; !ok {
			send(webService, gn)
		}
	}
}

var webService = "http://106.37.165.121/inf/chengtong/py/sy/groupNewsInfo/saveGroupNewsInfo"

func send(api string, data GroupNews) {
	postData := map[string][]GroupNews{"data": {data}}
	m, _ := json.Marshal(postData)
	req, err := http.NewRequest(http.MethodPost, api, bytes.NewReader(m))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	logger.Info(string(b))
}
