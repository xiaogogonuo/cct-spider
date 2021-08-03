package v1

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/v1/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"strings"
)

var Table = "t_dmaa_base_target_value"

func questionMark() string {
	var mark []string
	for i := 0; i < response.TargetValueFieldNum; i++ {
		mark = append(mark, "?")
	}
	return strings.Join(mark, ",")
}

var dataString string

func init() {
	dataString = "(" + questionMark() + ")"
}

func Dump(data []response.TargetValue) {
	var multiQuestionMark []string
	for i := 0; i < len(data); i++ {
		multiQuestionMark = append(multiQuestionMark, dataString)
	}
	multiQuestionMarkString := strings.Join(multiQuestionMark, ",")
	sql := `INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s`
    sql = fmt.Sprintf(sql, Table, response.TargetValueFieldStr, multiQuestionMarkString, response.TargetValueFieldSVS)
    var dumpData []interface{}
    for _, tv := range data {
    	values := tv.GetValues()
    	dumpData = append(dumpData, values...)
	}
	mysql.Transaction(sql, dumpData...)
}