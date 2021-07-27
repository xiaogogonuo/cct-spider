package main

import (
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/api/v1/anhui"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/api/v1/beijing"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/api/v1/chongqing"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/api/v1/guangzhou"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/api/v1/shanghai"
	"github.com/xiaogogonuo/cct-spider/internal/govserver/api/v1/tianjin"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/insertdb"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"sync"
)

var (
	govV      *viper.Viper
	filt      *filter.Filter
	dataInfo  *insertdb.DataInfo
	urlKeyMap map[string]byte
)

func govConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/gov",
	}
	v, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func init() {
	govV = govConfig()
	filt = &filter.Filter{
		Filepath:   "urlKey.txt",
		ThisUrlKey: make(map[string]byte),
	}
	urlKeyMap = filt.ReadUrlKey()
	dataInfo = &insertdb.DataInfo{
		DBName:     "t_dmbe_policy_news_info",
		PolicyCode: "20",
		PolicyName: "地方政策",
	}
}

func government() {
	wg := &sync.WaitGroup{}
	limitChan := make(chan struct{}, 500)
	urlChannel := make(chan *callback.UrlChan, 5000)   // url请求池
	infoChannel := make(chan *callback.InfoChan, 100000) // info请求池
	errChannel := make(chan *callback.InfoChan)         // 异常池
	message := make(chan *callback.Message)             // 数据池
	save := dataInfo.InsertIntoSQL                      // 保存数据的函数

	wg.Add(6)
	go anhui.GetPageUrlList(govV.GetString("安徽"), urlChannel, wg)
	go beijing.GetPageUrlList(govV.GetString("北京"), urlChannel, wg)
	go chongqing.GetFirstUrl(govV.GetString("重庆"), urlChannel, wg)
	go guangzhou.GetPageUrlList(govV.GetString("广州"), urlChannel, wg)
	go shanghai.GetPageUrlList(govV.GetString("上海"), urlChannel, wg)
	go tianjin.GetPageUrlList(govV.GetString("天津"), urlChannel, wg)



	go func() {
		for v := range urlChannel {
			if _, ok := urlKeyMap[md5.MD5(v.Url)]; ok {
				logger.Info("Obtained, no need to update", logger.Field("url", v.Url))
				continue
			}
			limitChan <- struct{}{}
			wg.Add(1)
			go func(v *callback.UrlChan) {
				v.GetUrlFunc(urlChannel, infoChannel, wg)
				<-limitChan
			}(v)

		}
	}()
	go func() {
		for v := range infoChannel {
			if _, ok := urlKeyMap[md5.MD5(v.Url)]; ok {
				logger.Info("Obtained, no need to update", logger.Field("url", v.Url))
				continue
			}
			limitChan <- struct{}{}
			wg.Add(1)
			go func(v *callback.InfoChan) {
				v.GetInfoFunc(errChannel, message, wg)
				<-limitChan
			}(v)
		}
	}()
	go func() {
		wg.Wait()
		close(urlChannel)
		close(infoChannel)
		close(message)
	}()
	save(filt, message)
}

func main() {
	government()
}
