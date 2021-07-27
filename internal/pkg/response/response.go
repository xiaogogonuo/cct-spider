package response

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
)

type PR struct {
	Request request.Request
	Parse   parse.Parse
}

func (pr *PR) GetPageUrl(attrName string) (hrefList []string) {
	html, err := pr.Request.VisitString()
	if err != nil{
		return
	}
	pr.Parse.Html = html
	hrefList = pr.Parse.GetAllUrlByParseHtml(attrName)
	return hrefList
}

func (pr *PR) GetHtmlInfo() (message *callback.Message) {
	message = &callback.Message{}
	var info []string
	if !strings.Contains(pr.Request.Url, pr.Parse.DomainName) {
		logger.Warn(fmt.Sprintf("域名：%s 网址：%s 域名不存在\n", pr.Parse.DomainName, pr.Request.Url))
		return
	}
	if !strings.Contains(pr.Request.Url, pr.Parse.Suffix) {
		logger.Warn(fmt.Sprintf("后缀：%s 网址：%s 后缀不存在\n", pr.Parse.Suffix, pr.Request.Url))
		return
	}
	html, err := pr.Request.VisitString()
	if err != nil{
		return
	}
	pr.Parse.Html = html
	title, content, date := pr.Parse.GetTextByParseHtml()
	info = append(info, content...)
	message = &callback.Message{
		Url:        pr.Request.Url,
		Title:      _replace(title),
		Summary:    _replace(strings.Join(info, "")),
		Source:     pr.Parse.Source,
		SourceCode: pr.Parse.SourceCode,
		Date:       date,
	}
	return
}

func (pr *PR) GetPageNum(r string) (num int) {
	html, err := pr.Request.VisitString()
	if err != nil{
		return
	}
	pr.Parse.Html = html
	num = pr.Parse.GetPageNum(r)
	return
}

func (pr *PR) GetCountAndSize(countR string, sizeR string) (count int, size int) {
	html, err := pr.Request.VisitString()
	if err != nil{
		return
	}
	pr.Parse.Html = html
	count, size = pr.Parse.GetCountAndSize(countR, sizeR)
	return
}

func _replace(s string) string {
	return strings.Replace(s, `'`, `"`, -1)
}
