package similarity

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func batchUpdate(n []News) {
	var data []string
	var id []string
	for _, v := range n {
		s := fmt.Sprintf("WHEN '%s' THEN '%s'", v.NewsGUID, time.Now().Format("2006-01-02 15:04:05"))
		data = append(data, s)
		id = append(id, fmt.Sprintf("'%s'", v.NewsGUID))
	}
	dataString := strings.Join(data, " ")
	idString := strings.Join(id, ",")
	SQL := fmt.Sprintf("UPDATE %s SET DELETE_DATE = CASE NEWS_GUID %s END WHERE NEWS_GUID IN (%s);", TABLE, dataString, idString)
	Transaction(SQL)
}

const batchSize = 1000

func update(n []News) {
	epoch := int(math.Ceil(float64(len(n)) / float64(batchSize)))
	for i := 0; i < epoch; i++ {
		if batchSize*(i+1) < len(n) {
			batchData := n[i*batchSize : (i+1)*batchSize]
			batchUpdate(batchData)
		} else {
			batchData := n[i*batchSize:]
			batchUpdate(batchData)
		}
	}
}
