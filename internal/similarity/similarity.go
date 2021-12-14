package similarity

import (
	"flag"
	"fmt"
	"github.com/antlabs/strsim"
	"sort"
	"strings"
	"sync"
)

const (
	TABLE = "t_dmbe_news_info"
)

// 查询所有新闻日期
func queryNewsTs() [][]string {
	SQL := fmt.Sprintf("SELECT NEWS_TS2 FROM %s ORDER BY NEWS_TS2", TABLE)
	return Query(SQL)
}

// 删除日期重复的数据
func dropRepeatNewsTs(row [][]string) []string {
	setter := make([]string, 0)
	filter := make(map[string]struct{})
	for _, r := range row {
		ts := r[0]
		if _, ok := filter[ts]; !ok {
			filter[ts] = struct{}{}
			setter = append(setter, ts)
		}
	}
	return setter
}

// 根据日期降序排列
func sortNewsTs(row []string) []string {
	sort.Slice(row, func(i, j int) bool {
		return row[i] > row[j] // 降序
	})
	return row
}

// 返回最近的n天
func newsTs(n int) []string {
	row := sortNewsTs(dropRepeatNewsTs(queryNewsTs()))
	if n == -1 || n >= len(row) {
		return row
	}
	return row[:n]
}

type News struct {
	NewsGUID  string
	NewsTitle string
	Deleted   bool
}

func queryOneDayNews(ts string) []News {
	var daily []News
	SQL := fmt.Sprintf("SELECT NEWS_GUID, NEWS_TITLE FROM %s WHERE NEWS_TS2 = '%s'", TABLE, ts)
	for _, r := range Query(SQL) {
		var n News
		n.NewsGUID = r[0]
		n.NewsTitle = r[1]
		daily = append(daily, n)
	}
	return daily
}

// 删除标点符符号
func dropSymbol(s string) string {
	s = strings.ReplaceAll(s, "：", "")
	s = strings.ReplaceAll(s, "？", "")
	s = strings.ReplaceAll(s, "，", "")
	s = strings.ReplaceAll(s, "。", "")
	s = strings.ReplaceAll(s, "“", "")
	s = strings.ReplaceAll(s, "《", "")
	s = strings.ReplaceAll(s, "》", "")
	s = strings.ReplaceAll(s, "【", "")
	s = strings.ReplaceAll(s, "】", "")
	s = strings.ReplaceAll(s, "、", "")
	s = strings.ReplaceAll(s, "！", "")
	s = strings.ReplaceAll(s, "—", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, "@", "")
	s = strings.ReplaceAll(s, "#", "")
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, "%", "")
	s = strings.ReplaceAll(s, "*", "")
	s = strings.ReplaceAll(s, "&", "")
	s = strings.ReplaceAll(s, "^", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "=", "")
	s = strings.ReplaceAll(s, "+", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "?", "")
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, "｜", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "	", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "（", "")
	s = strings.ReplaceAll(s, "）", "")
	return s
}

// 比较一天中所有新闻标题的相似度，返回相似度高的新闻
func combineCompare(n []News, per float64) []News {
	var deletedNews []News
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i].Deleted || n[j].Deleted {
				continue
			}
			if textCompare(n[i], n[j], per) {
				if len(n[i].NewsTitle) > len(n[j].NewsTitle) {
					n[j].Deleted = true
					deletedNews = append(deletedNews, n[j])
				} else {
					n[i].Deleted = true
					deletedNews = append(deletedNews, n[i])
				}
			}
		}
	}
	return deletedNews
}

func textCompare(n1, n2 News, per float64) bool {
	// 1、存在包含关系的
	if strings.Contains(n1.NewsTitle, n2.NewsTitle) || strings.Contains(n2.NewsTitle, n1.NewsTitle) {
		return true
	}
	// 2、相似度超过per的
	if strsim.Compare(dropSymbol(n1.NewsTitle), dropSymbol(n2.NewsTitle)) > per {
		return true
	}
	return false
}

func SimServer() {
	// 第一个参数：存放值的参数地址
	// 第二个参数：命令行参数的名称
	// 第三个参数：命令行不输入时的默认值
	var day int
	flag.IntVar(&day, "day", -1, "")
	var per float64
	flag.Float64Var(&per, "per", 0.75, "")
	flag.Parse()

	var wg sync.WaitGroup
	var semaphore = make(chan struct{}, 10)
	for _, ts := range newsTs(day) {
		wg.Add(1)
		go func(n []News) {
			defer wg.Done()
			semaphore <- struct{}{}
			deletedNews := combineCompare(n, per)
			wg.Add(1)
			go func(deletedNews []News) { // 更新数据库DELETE_DATE字段
				defer wg.Done()
				update(deletedNews)
			}(deletedNews)
			<-semaphore
		}(queryOneDayNews(ts))
	}
	wg.Wait()
}
