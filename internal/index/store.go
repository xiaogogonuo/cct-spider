package index

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"math"
	"strings"
	"sync"
)

func questionMark() string {
	var mark []string
	for i := 0; i < FieldNum; i++ {
		mark = append(mark, "?")
	}
	return strings.Join(mark, ",")
}

var dataString string

func init() {
	dataString = "(" + questionMark() + ")"
}

var batchSize = 10

func batchDump(data []Field, wg *sync.WaitGroup) {
	defer wg.Done()
	length := len(data)
	epoch := int(math.Ceil(float64(length) / float64(batchSize)))
	for i := 0; i < epoch; i++ {
		if batchSize*(i+1) < length {
			batchData := data[i*batchSize : (i+1)*batchSize]
			dump(batchData)
		} else {
			batchData := data[i*batchSize:]
			dump(batchData)
		}
	}
}

func dump(data []Field) {
	var multiQuestionMark []string
	for i := 0; i < len(data); i++ {
		multiQuestionMark = append(multiQuestionMark, dataString)
	}
	multiQuestionMarkString := strings.Join(multiQuestionMark, ",")
	sql := `INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s`
	sql = fmt.Sprintf(sql, Table, FieldStr, multiQuestionMarkString, FieldSVS)
	var dumpData []interface{}
	for _, f := range data {
		values := f.GetValues()
		dumpData = append(dumpData, values...)
	}
	mysql.Transaction(sql, dumpData...)
}
