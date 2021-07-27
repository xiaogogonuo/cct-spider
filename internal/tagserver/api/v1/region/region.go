package region

import (
	"github.com/xiaogogonuo/cct-spider/internal/pkg/findmap"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/store"
	"sync"
)

func GetRegion(n *store.PolicyNewsOrg, wg *sync.WaitGroup)  {
	defer wg.Done()
	n.RegionMap = make(map[string]int)
	regionPat, _ := findmap.RegionRuntime()
	regions := findmap.FindAll(regionPat, n.NEWS_SUMMARY)
	n.RegionMap["sum"] = len(regions)
	for _, r := range regions{
		if _, ok := n.RegionMap[r]; !ok{
			n.RegionMap[r] = 1
		}else {
			n.RegionMap[r] += 1
		}
	}
}