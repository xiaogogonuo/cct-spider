package index

import (
	"encoding/json"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"math"
	"strings"
	"sync"
	"time"
)

// RunIndex24 24小时刷新的数据不录入内部服务器，直接发送到生产机
func RunIndex24() {
	var wg sync.WaitGroup
	var configs Configs
	if err := json.Unmarshal([]byte(ConfigString), &configs); err != nil {
		logger.Fatal(err.Error())
		return
	}
	for _, config := range configs {
		runTimes := strings.Split(config.RunTime, "~")
		runTimeL, _ := time.Parse("15:04", runTimes[0])
		runTimeR, _ := time.Parse("15:04", runTimes[1])
		if runTimeR.Hour() - runTimeL.Hour() != 0 {
			continue
		}
		rowRespond := config.routingDistribution()
		data := config.construct(rowRespond, true) // 推送给java服务器
		logger.Info(config.Name, logger.Field("updating rows: ", len(data)))
		length := len(data)
		epoch := int(math.Ceil(float64(length) / float64(batchSize)))
		wg.Add(epoch)
		for i := 0; i < epoch; i++ {
			if batchSize*(i+1) < length {
				batchData := data[i*batchSize : (i+1)*batchSize]
				go send(webService, batchData, &wg)
			} else {
				batchData := data[i*batchSize:]
				go send(webService, batchData, &wg)
			}
		}
	}
	wg.Wait()
}