package similarity

import (
	"flag"
	"fmt"
	"github.com/antlabs/strsim"
	"sort"
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

func combineCompare(n []News) []News {
	var deletedNews []News
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i].Deleted || n[j].Deleted {
				continue
			}
			if textCompare(n[i], n[j], 0.75) {
				if len(n[i].NewsTitle) > len(n[j].NewsTitle) {
					n[j].Deleted = true
					deletedNews = append(deletedNews, n[j])
				}
			}
		}
	}
	return deletedNews
}

func textCompare(n1, n2 News, per float64) bool {
	if strsim.Compare(n1.NewsTitle, n2.NewsTitle) > per {
		return true
	}
	return false
}

func SimServer() {
	var day int
	flag.IntVar(&day, //第一个参数：存放值的参数地址
		"day", //第二个参数：命令行参数的名称
		-1,    //第三个参数：命令行不输入时的默认值
		"")
	flag.Parse()

	var wg sync.WaitGroup
	var semaphore = make(chan struct{}, 10)
	for _, ts := range newsTs(day) {
		wg.Add(1)
		go func(n []News) {
			defer wg.Done()
			semaphore <- struct{}{}
			deletedNews := combineCompare(n)
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
