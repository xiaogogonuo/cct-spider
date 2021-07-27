package findmap

import (
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"regexp"
	"strings"
	"sync"
)

type comp struct {
	Name       string
	NameEN     string
	StockCode  string
	CreditCode string
	CompGuid   string
	CompanyId  string
}

var (
	compPat string
	compMap map[string]*comp
	once sync.Once
)

func setUpCompany() {
	compPat, compMap = getCompanyPat()
}

func CompanyRuntime() (string, map[string]*comp) {
	once.Do(setUpCompany)
	return compPat, compMap
}

func getCompanyPat() (compPat string, newCM map[string]*comp) {
	var companyReg []string
	newCM = make(map[string]*comp)
	sqlCode := "SELECT COMP_NAME, COMP_NAME_EN, STOCK_CODE, CREDIT_CODE, COMP_GUID  FROM t_dmaa_comp_base;"
	for _, company := range mysql.Query(sqlCode) {
		reg := regexp.MustCompile("公司|总公司|有限公司|有限责任公司|物流中心|流通中心|合作公司|发展公司|商社|贸易公司|批发市场|设备厂|总厂")
		s := strings.Replace(company[0], "(", "*", -1)
		s = strings.Replace(s, ")", "*", -1)
		c := reg.ReplaceAllString(s, "")
		companyReg = append(companyReg, c+")")
		if _, o := newCM[c]; !o {
			newCM[c] = &comp{
				Name:       company[0],
				NameEN:     company[1],
				StockCode:  company[2],
				CreditCode: company[3],
				CompGuid:   company[4],
			}
		}
	}
	compPat = "(" + strings.Join(companyReg, "|(")
	return
}
