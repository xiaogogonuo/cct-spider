package main

import (
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/internal/indserver/api/v1/logistics"
	"github.com/xiaogogonuo/cct-spider/internal/indserver/api/v1/ppi"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/insertdb"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"sync"
)

var (
	indV      *viper.Viper
	filt      *filter.Filter
	dataInfo  *insertdb.DataInfo
	urlKeyMap map[string]byte
)

func indConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/ind",
	}
	v, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func init() {
	indV = indConfig()
	filt = &filter.Filter{
		Filepath:   "urlKey.txt",
		ThisUrlKey: make(map[string]byte),
	}
	urlKeyMap = filt.ReadUrlKey()
	dataInfo = &insertdb.DataInfo{
		DBName:     "t_dmbe_policy_news_info",
		PolicyCode: "30",
		PolicyName: "行业政策",
	}
}

func ministries() {
	wg := &sync.WaitGroup{}
	limitChan := make(chan struct{}, 400)
	urlChannel := make(chan *callback.UrlChan, 10000)   // url请求池
	infoChannel := make(chan *callback.InfoChan, 10000) // info请求池
	errChannel := make(chan *callback.InfoChan)         // 异常池
	message := make(chan *callback.Message)             // 数据池
	save := dataInfo.InsertIntoSQL                      // 保存数据的函数

	wg.Add(2)
	go ppi.GetPageUrlList(indV.GetString("造纸协会"), urlChannel, wg)
	go logistics.GetPageUrlList(indV.GetString("物流协会"), urlChannel, wg)

	go func() {
		for v := range urlChannel {
			if _, ok := urlKeyMap[md5.MD5(v.Url)]; ok {
				//logger.Info("Obtained, no need to update", logger.Field("url", v.Url))
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
				//logger.Info("Obtained, no need to update", logger.Field("url", v.Url))
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
	ministries()
}
