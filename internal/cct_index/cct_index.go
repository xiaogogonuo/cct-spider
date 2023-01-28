package cct_index

import (
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/factory"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/model"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/pkg/text"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/remote"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/eastmoney"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/fx"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/irc"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/nrc"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/ppi100"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/sci"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/sina"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/wuliu"
	"github.com/xiaogogonuo/cct-spider/internal/cct_index/source/xiben"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strings"
	"time"
)

const (
	// SpiderInterval 爬虫运行时间间隔
	SpiderInterval = time.Minute * 30

	// ListenInterval 远程服务器监听时间间隔
	ListenInterval = time.Minute * 5

	// IndexValueGUID 保存历史数据唯一性的字段"ValueGUID"的集合
	// 仅保存非实时数据的ValueGUID
	IndexValueGUID = "indexes.txt"
)

// CCTIndex 城通指标爬虫逻辑
func CCTIndex() {
	// 第一步：将代表历史数据唯一性的字段"ValueGUID"加载到内存
	rows, err := text.ReadFromText(IndexValueGUID)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	uniqueIndexes := map[string]struct{}{}
	for _, row := range rows {
		if _, ok := uniqueIndexes[row]; !ok {
			uniqueIndexes[row] = struct{}{}
		}
	}

	// 第二步：从"诚通指标配置.xlsx"读取指标相关的配置信息
	ics, err := InitIndexConfigFromExcel()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	var indexes []*model.Index

	// 第三步：遍历指标配置表，循环爬取数据
	for _, ic := range ics {
		if !ic.Enable || !ic.IfSpider() {
			continue
		}
		var buffers []*model.Buffer
		switch ic.Adapter {
		case "EastMoney": // 东方财富
			buffers = eastmoney.SpiderEastMoney(ic)
		case "Sina": // 新浪财经
			buffers = sina.SpiderSina(ic)
		case "Sci": // 卓创资讯
			buffers = sci.SpiderSCI(ic)
		case "Fx": // 汇通财经
			buffers = fx.SpiderFx(ic)
		case "XiBen": // 西本资讯
			buffers = xiben.SpiderXiBen(ic)
		case "WuLiu": // 中国物流
			buffers = wuliu.SpiderWuLiu(ic)
		case "Irc": // 银保监会
			buffers = irc.SpiderIrc(ic)
		case "PPI100": //  生意社
			buffers = ppi100.SpiderPPI100(ic)
		case "Nrc": // 中华人民共和国国家发展和改革委员会
			buffers = nrc.SpiderNrc(ic)
		}
		idx := factory.Manufacture(ic, buffers)
		indexes = append(indexes, idx...)
	}

	// 第四步：过滤重复指标
	var realTimeIndexes []*model.Index // 代表了实时指标，直接发送到服务器
	var noneRealIndexes []*model.Index // 代表非实时指标，需要做唯一性校验
	for _, index := range indexes {
		if strings.Contains(index.TargetValue, ",") {
			realTimeIndexes = append(realTimeIndexes, index)
			continue
		}
		if _, ok := uniqueIndexes[index.ValueGUID]; !ok {
			noneRealIndexes = append(noneRealIndexes, index)
		}
	}

	// 第五步：将新数据发送到远程服务器
	_ = remote.Push(realTimeIndexes)
	u := remote.Push(noneRealIndexes)

	// 第六步：将新记录表中的内容追加到indexes.txt
	var indexForAppend []string
	for _, v := range u {
		if _, ok := uniqueIndexes[v]; !ok {
			indexForAppend = append(indexForAppend, v)
		}
	}

	text.AppendToText(IndexValueGUID, indexForAppend...)

	time.Sleep(SpiderInterval)
}

func RunApplication() {
	listener := make(chan struct{})
	go RemoteListen(listener)
	for {
		select {
		case <-listener:
			return
		default:
			CCTIndex()
		}
	}
}
