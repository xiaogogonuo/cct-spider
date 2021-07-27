package company

import (
	"github.com/xiaogogonuo/cct-spider/internal/pkg/findmap"
	"github.com/xiaogogonuo/cct-spider/internal/tagserver/store"
	"sync"
)

func GetCompany(n *store.PolicyNewsOrg, wg *sync.WaitGroup)  {
	defer wg.Done()
	n.CompanyMap = make(map[string]int)
	companyPat, _ := findmap.CompanyRuntime()
	company := findmap.FindAll(companyPat, n.NEWS_SUMMARY)
	n.CompanyMap["sum"] = len(company)
	for _, r := range company{
		if _, ok := n.CompanyMap[r]; !ok{
			n.CompanyMap[r] = 1
		}else {
			n.CompanyMap[r] += 1
		}
	}
}


