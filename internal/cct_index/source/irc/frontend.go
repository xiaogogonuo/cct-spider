package irc

import "regexp"

type Irc struct {
	Data struct {
		Lists []struct {
			BuildDate   string   `json:"builddate"`
			PublishDate string   `json:"publishDate"`
			DocTitle    string   `json:"docTitle"`
			ItemId      string   `json:"itemId"`
			DocId       string   `json:"docId"`
			Keyword     []string `json:"keyword"`
		} `json:"lists"`
	} `json:"data"`
}

// GetDate 从标题提取日期
func (c Irc) GetDate(docTitle string) string {
	pattern := "([0-9]{4})年([一二三四])季度"
	reg := regexp.MustCompile(pattern)
	matched := reg.FindAllStringSubmatch(docTitle, -1)
	if len(matched) > 0 {
		if len(matched[0]) > 2 {
			year := matched[0][1]
			var season string
			switch matched[0][2] {
			case "一":
				season = "Q1"
			case "二":
				season = "Q2"
			case "三":
				season = "Q3"
			case "四":
				season = "Q4"
			}
			return year + season
		}
	}
	return ""
}

// GetTargetValue 从内容提取指标值
func (c Irc) GetTargetValue(content string) string {
	pattern := "[0-9]{4}年[一二三四]季度末，商业银行.*不良贷款余额(.*)万亿元，.*不良贷款率"
	reg := regexp.MustCompile(pattern)
	matched := reg.FindAllStringSubmatch(content, -1)
	if len(matched) > 0 {
		if len(matched[0]) > 1 {
			return matched[0][1]
		}
	}
	return ""
}
