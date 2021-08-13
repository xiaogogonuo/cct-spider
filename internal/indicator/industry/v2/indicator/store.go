package indicator

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/indicator/industry/v2/pkg/response"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"math"
	"strings"
)

func questionMark() string {
	var mark []string
	for i := 0; i < response.FieldNum; i++ {
		mark = append(mark, "?")
	}
	return strings.Join(mark, ",")
}

var dataString string

func init() {
	dataString = "(" + questionMark() + ")"
}

var batchSize = 1000

func batchDump(data []response.Field) {
	length := len(data)
	epoch := int(math.Ceil(float64(length) / float64(batchSize)))
	for i := 0; i < epoch; i++ {
		if batchSize * (i + 1) < length {
			batchData := data[i*batchSize:(i+1)*batchSize]
			dump(batchData)
		} else {
			batchData := data[i*batchSize:]
			dump(batchData)
		}
	}
}

func dump(data []response.Field) {
	var multiQuestionMark []string
	for i := 0; i < len(data); i++ {
		multiQuestionMark = append(multiQuestionMark, dataString)
	}
	multiQuestionMarkString := strings.Join(multiQuestionMark, ",")
	sql := `INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s`
	sql = fmt.Sprintf(sql, table, response.FieldStr, multiQuestionMarkString, response.FieldSVS)
	var dumpData []interface{}
	for _, f := range data {
		values := f.GetValues()
		dumpData = append(dumpData, values...)
	}
	mysql.Transaction(sql, dumpData...)
}
