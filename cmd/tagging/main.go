package main

import (
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/api/v1/company"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/api/v1/industry"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/api/v1/region"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/api/v1/sentiment"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/store"
	"sync"
)

func tagging(){
	wg := &sync.WaitGroup{}
	limitChan := make(chan struct{}, 2)
	newsOrgChan := make(chan *store.PolicyNewsOrg)
	newsChan := make(chan *store.PolicyNews)
	newsRegionChan := make(chan *store.NewsRegion)
	newsCompanyChan := make(chan *store.NewsCompany)
	newsIndustryChan := make(chan *store.NewsIndustry)

	filt := &filter.Filter{
		Filepath: "idKey.txt",
		ThisUrlKey: make(map[string]byte),
	}
	idKeyMap := filt.ReadUrlKey()

	wg.Add(1)
	go store.QueryData(newsOrgChan, wg)

	go func() {
		for n := range newsOrgChan {
			if _, ok := idKeyMap[n.NEWS_GUID]; ok {
				//logger.Info("Obtained, no need to update", logger.Field("news_guid", n.NEWS_GUID))
				continue
			}
			limitChan <- struct {}{}
			wg.Add(1)
			go func(n *store.PolicyNewsOrg) {
				wg2 := &sync.WaitGroup{}
				wg2.Add(4)
				go sentiment.GetSentiment(n, wg2)
				go industry.GetIndustry(n, wg2)
				go region.GetRegion(n, wg2)
				go company.GetCompany(n, wg2)
				wg2.Wait()
				wg.Add(1)
				go store.AssembleData(n, newsChan, newsRegionChan, newsCompanyChan, newsIndustryChan, wg)
				wg.Done()
				<- limitChan
			}(n)
		}
	}()

	go func() {
		wg.Wait()
		close(newsOrgChan)
		close(newsChan)
		close(newsRegionChan)
		close(newsCompanyChan)
		close(newsIndustryChan)
	}()

	wg3 := &sync.WaitGroup{}
	wg3.Add(4)
	go store.InsertRegion(newsRegionChan, wg3)
	go store.InsertCompany(newsCompanyChan, wg3)
	go store.InsertIndustry(newsIndustryChan, wg3)
	go store.UpdateNews(filt, newsChan, wg3)
	wg3.Wait()

}

func main() {
	tagging()
}