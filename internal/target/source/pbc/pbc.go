package pbc

import (
	"github.com/tebeka/selenium"
	"github.com/xiaogogonuo/cct-spider/internal/target/model"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/simulator"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
)

// SpiderPBCTarget 爬取中国人民银行指标
// 适用指标：
// - 7天逆回购利率
//   • 页面展示接口：http://www.pbc.gov.cn/zhengcehuobisi/125207/125213/125431/125475/4532152/index.html
//   • 数据获取接口：http://www.pbc.gov.cn/zhengcehuobisi/125207/125213/125431/125475/4532152/index.html
func SpiderPBCTarget(SourceTargetCodeSpider string) (responses []model.Response) {
	switch SourceTargetCodeSpider {
	case "RRR": // 7天逆回购利率
		responses = spiderReverseRepoRate()
	}
	_ = simulator.WebDriver.Close()
	_ = simulator.WebDriver.Quit()
	_ = simulator.DriverService.Stop()
	return
}

// APIReverseRepoRate 7天逆回购利率接口
var APIReverseRepoRate = "http://www.pbc.gov.cn/zhengcehuobisi/125207/125213/125431/index.html"

func spiderReverseRepoRate() (responses []model.Response) {
	simulator.InitPBCSimulator()
	if simulator.DriverService == nil || simulator.WebDriver == nil {
		logger.Error("init selenium fail")
		return
	}
	if err := simulator.WebDriver.Get(APIReverseRepoRate); err != nil {
		logger.Error(err.Error())
		return
	}
	time.Sleep(time.Second)
	// 公开市场业务交易公告
	elements, err := simulator.WebDriver.FindElements(selenium.ByXPATH, "//td[@class='unline']/a")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	// 提取公开市场业务交易公告链接
	var hrefs []string
	for _, element := range elements {
		href, err := element.GetAttribute("href")
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		text, err := element.Text()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if strings.Contains(text, "公开市场业务交易公告") {
			hrefs = append(hrefs, href)
		}
	}
	for _, href := range hrefs {
		if err := simulator.WebDriver.Get(href); err != nil {
			logger.Error(err.Error())
			continue
		}
		time.Sleep(time.Second)
		// 获取日期
		element, err := simulator.WebDriver.FindElement(selenium.ByCSSSelector, "#shijian")
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		date, err := element.Text()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		// 获取指标值
		element, err = simulator.WebDriver.FindElement(selenium.ByXPATH, "//div[@id='zoom']//tbody/tr[2]/td[3]//span")
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		value, err := element.Text()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		var response model.Response
		response.Date = strings.ReplaceAll(date, "-", "")[:8]
		response.TargetValue = strings.ReplaceAll(value, "%", "")
		responses = append(responses, response)
	}
	return
}
