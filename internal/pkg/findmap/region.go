package findmap

import (
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"regexp"
	"strings"
	"sync"
)

var (
	regionPat string
	regionMap map[string][]string
	onceRegion sync.Once
)

func setUpRegion() {
	regionPat, regionMap = getRegionPat()
}

func RegionRuntime() (string, map[string][]string){
	onceRegion.Do(setUpRegion)
	return regionPat, regionMap
}

func getRegionPat() (regionPat string, regionMap map[string][]string){
	var regionReg []string
	regionMap = make(map[string][]string)
	sqlCode := "SELECT REGION_NAME, REGION_CODE FROM t_dmaa_regioin_code"
	for _, region := range mysql.Query(sqlCode) {
		reg := regexp.MustCompile("省|市|自治区|自治州|特别行政区|区|盟")
		r := reg.ReplaceAllString(region[0], "")
		regionReg = append(regionReg, r + ")")
		if _, ok := regionMap[r]; !ok {
			regionMap[r] = region
		}
	}
	regionPat = "(" + strings.Join(regionReg, "|(")
	return
}

