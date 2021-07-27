package parse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/urlprocess"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"regexp"
	"strconv"
	"strings"
)

type Parse struct {
	Html            string
	BaseUrl         string
	Source          string
	SourceCode      string
	UrlSelector     string
	TitleSelector   string
	TextSelector    string
	DateSelector    string
	PageNumSelector string
	Suffix          string
	DomainName      string
}

func (p *Parse) GetTextByParseHtml() (title string, info []string, date string) {
	var titleL []string
	var key map[string]byte
	key = make(map[string]byte)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if p.TitleSelector != "" {
		dom.Find(p.TitleSelector).Each(func(i int, selection *goquery.Selection) {
			if selection.Text() != "" {
				titleL = append(titleL, selection.Text())
			}
		})
		title = strings.Join(titleL, "")
	}
	if p.TextSelector != "" {
		dom.Find(p.TextSelector).Each(func(i int, selection *goquery.Selection) {
			s := selection.Text()
			if _, ok := key[s]; s != "" && !ok {
				info = append(info, s)
				key[s] = 0
			}
		})
	}
	if p.DateSelector != "" {
		dom.Find(p.DateSelector).Each(func(i int, selection *goquery.Selection) {
			t := ""
			timeReg := regexp.MustCompile("([0-9]{4})[\\-\\/年]([0-9]{1,2})[\\-\\/月]([0-9]{1,2})")
			timeStr := timeReg.FindStringSubmatch(selection.Text())
			if len(timeStr) == 0 {
				return
			}
			for j, d := range timeStr[1:] {

				if j >= 1 && len(d) == 1 {
					t += "0" + d
				} else {
					t += d
				}
			}
			date = t
		})
	}
	return
}

func (p *Parse) GetOneUrlByParseHtml(attrName string) (src string, b bool) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dom.Find(p.UrlSelector).Each(func(i int, selection *goquery.Selection) {
		src, b = selection.Attr(attrName)
	})
	return
}

func (p *Parse) GetAllUrlByParseHtml(attrName string) (hrefList []string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dom.Find(p.UrlSelector).Each(func(i int, selection *goquery.Selection) {
		href, b := selection.Attr(attrName)
		if !b || href == "" {
			logger.Warn(fmt.Sprintf("b :%v, href: %s\n", b, href))
			return
		}
		if strings.Contains(href, "http") || strings.Contains(href, "https") {
			hrefList = append(hrefList, urlprocess.UrlJoint(href, p.Suffix))
		} else {
			hrefList = append(hrefList, urlprocess.UrlJoint(p.BaseUrl, href+p.Suffix))
		}

	})
	return
}

func (p *Parse) GetPageNum(r string) (num int) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dom.Find(p.PageNumSelector).Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" {
			reg := regexp.MustCompile(r)
			numReg := regexp.MustCompile("[0-9]+")
			numStr := reg.FindString(selection.Text())
			numStr = numReg.FindString(numStr)
			if numStr == "" {
				return
			}
			num, err = strconv.Atoi(numStr)
			if err != nil {
				logger.Error(err.Error())
			}
		}
	})
	return
}

func (p *Parse) GetCountAndSize(countR string, sizeR string) (count int, size int) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dom.Find(p.PageNumSelector).Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		if s != "" {
			countReg := regexp.MustCompile(countR)
			sizeReg := regexp.MustCompile(sizeR)
			numReg := regexp.MustCompile("[0-9]+")
			countStr := countReg.FindString(s)
			countStr = numReg.FindString(countStr)
			sizeStr := sizeReg.FindString(s)
			sizeStr = numReg.FindString(sizeStr)
			if countStr != "" {
				count, err = strconv.Atoi(countStr)
				if err != nil {
					logger.Error(err.Error())
				}
			} else {
				logger.Warn("count is nil")
			}

			if sizeStr != "" {
				size, err = strconv.Atoi(sizeStr)
				if err != nil {
					logger.Error(err.Error())
				}
			} else {
				logger.Warn("size is nil")
			}
		}
	})
	return
}
