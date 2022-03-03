package cctgroup

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// 中国诚通控股集团有限公司-新闻中心

const domain = "http://www.cctgroup.com.cn"

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

type PagerIndex struct {
	Plate string // 新闻所属板块
	Index string // 新闻页码链接
}

type C struct {
	mu         sync.Mutex
	wg         *sync.WaitGroup
	stop       chan struct{}
	semaphore  chan struct{}
	pageIndex  chan PagerIndex
	news       chan News
	newsCenter []News
}

func (c *C) newsDetail(n News) {
	defer c.wg.Done()
	body, err := downloader(n.URL)
	if err != nil {
		<-c.semaphore
		logger.Error(err.Error())
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		<-c.semaphore
		logger.Error(err.Error())
		return
	}
	var ns News
	var content []string
	dom.Find("div[class='w1000']").Each(func(i int, selection *goquery.Selection) {
		selection.Find("p").Each(func(i int, selection *goquery.Selection) {
			content = append(content, selection.Text())
		})
		selection.Find("img").Each(func(i int, selection *goquery.Selection) {
			src, _ := selection.Attr("src")
			var imageURL string
			if strings.Contains(src, "http") {
				imageURL = src
			} else {
				imageURL = domain + src
			}
			body, e := downloader(imageURL)
			if e != nil {
				logger.Error(e.Error())
			} else {
				ns.Image = append(ns.Image, body)
			}
		})
	})
	dom.Find("meta[name='createDate']").Each(func(i int, selection *goquery.Selection) {
		date, _ := selection.Attr("content")
		ns.Date = strings.Split(date, " ")[0]
	})
	ns.Title = n.Title
	ns.Plate = n.Plate
	ns.URL = n.URL
	ns.Body = strings.Join(content, " ")
	c.mu.Lock()
	c.newsCenter = append(c.newsCenter, ns)
	c.mu.Unlock()
	<-c.semaphore
}

// 获取新闻链接
func (c *C) getNews(pi PagerIndex, dom *goquery.Document) {
	dom.Find("a[istitle='true']").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		if strings.Contains(href, "http") {
			return
		}
		title, _ := selection.Attr("title")
		var n News
		n.URL = domain + href
		n.Title = title
		n.Plate = pi.Plate
		c.news <- n
	})
}

// 获取下一页
func (c *C) getNextPage(pi PagerIndex, dom *goquery.Document) {
	dom.Find("a[title='下一页']").Each(func(i int, selection *goquery.Selection) {
		tagName, _ := selection.Attr("tagname")
		// 全爬：[NEXTPAGE]
		// 爬前3页(不包括第3页)：-3.html
		if strings.Contains(tagName, fmt.Sprintf("-%d.html", HowManyPages)) {
			// 停止爬取板块
			c.stop <- struct{}{}
		} else {
			// 加入到页码通道
			var nextPagerIndex PagerIndex
			nextPagerIndex.Plate = pi.Plate
			nextPagerIndex.Index = domain + tagName
			c.pageIndex <- nextPagerIndex
		}
	})
}

// 从每一页中提取新闻链接和下一页
func (c *C) getNewsList(pi PagerIndex) {
	defer c.wg.Done()
	body, err := downloader(pi.Index)
	if err != nil {
		<-c.semaphore
		logger.Error(err.Error())
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		<-c.semaphore
		logger.Error(err.Error())
		return
	}
	c.getNews(pi, dom)
	c.getNextPage(pi, dom)
	<-c.semaphore
}

func (c *C) GetNewsList() {
	defer c.wg.Done()
	for pi := range c.pageIndex {
		c.semaphore <- struct{}{}
		c.wg.Add(1)
		go c.getNewsList(pi)
	}
	<-c.semaphore
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

var HowManyPages = 2

func EntryPoint() {
	filter := NewFilter("cctNewsURL.txt")
	filter.Open()
	filter.ReadAll()

	startURL := map[string]string{
		"集团动态": "http://www.cctgroup.com.cn/cctgroup/xwzx/jtdt/index.html",
		"国资动态": "http://www.cctgroup.com.cn/cctgroup/xwzx/gzdt4/index.html",
		"媒体聚焦": "http://www.cctgroup.com.cn/cctgroup/xwzx/mtjj/index.html",
		"集团公告": "http://www.cctgroup.com.cn/cctgroup/xwzx/jtgg/index.html",
		"出资企业": "http://www.cctgroup.com.cn/cctgroup/xwzx/czqy/index.html",
	}
	c := C{
		wg:        &sync.WaitGroup{},
		stop:      make(chan struct{}, len(startURL)),
		semaphore: make(chan struct{}, 10),
		pageIndex: make(chan PagerIndex, len(startURL)),
		news:      make(chan News),
	}
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			if len(c.stop) == len(startURL) {
				close(c.news)
				close(c.pageIndex)
				break
			}
		}
	}()
	for k, v := range startURL {
		c.pageIndex <- PagerIndex{k, v}
	}
	c.semaphore <- struct{}{}
	c.wg.Add(1)
	go c.GetNewsList()
	for n := range c.news {
		c.semaphore <- struct{}{}
		c.wg.Add(1)
		go c.newsDetail(n)
	}
	c.wg.Wait()
	var groupNewsList []GroupNews
	var newNews []string
	for _, ns := range c.newsCenter {
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
		gn.NewsSource = "中国诚通控股集团有限公司"
		gn.NewsSourceCode = "WEB_02019"
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
		send(webService, gn)
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
