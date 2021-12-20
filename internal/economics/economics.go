package economics

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/economics/core"
	"github.com/xiaogogonuo/cct-spider/internal/economics/distribution/request"
	"github.com/xiaogogonuo/cct-spider/internal/economics/distribution/response"
	"github.com/xiaogogonuo/cct-spider/internal/economics/pkg/configReader"
	"sync"
	"sync/atomic"
)

type Economic struct {
	*sync.WaitGroup
	semaphore chan struct{}

	//RequestChannel chan *core.Request
	//RespondChannel chan backend.BackEnd
}

func (e *Economic) newRequest(req core.Sender) {
	e.semaphore <- struct{}{}
	req.Send()
	<-e.semaphore
	e.Done()
}

// agent 分发到对应的网络
func (e *Economic) agent(c configReader.EconomicsConfig) {
	switch c.Case {
	case "financeSinaMac":
		u := request.NewFinanceSinaURL()
		u.Event = c.SourceTargetCode
		req := core.NewRequest(u.ToURL())
		req.Parse = response.FinanceSinaMacParse
		req.Meta = c
		response.RequestChannel <- req
	}
}

func newEconomic() *Economic {
	return &Economic{
		&sync.WaitGroup{},
		make(chan struct{}, concurrency),
	}
}

func RunEconomics() {
	configs := configReader.ReadConfig(configPath)
	e := newEconomic()
	e.Add(1)
	go func() {
		defer e.Done()
		for {
			req := <-response.RequestChannel
			if req == nil {
				return
			}
			e.Add(1)
			go e.newRequest(req)
		}
	}()
	e.Add(1)
	go func() {
		defer e.Done()
		for _, c := range configs {
			e.agent(c)
		}
	}()
	go func() {
		for {
			if int(atomic.LoadUint64(&response.Stop)) == len(configs) {
				close(response.RespondChannel)
				close(response.RequestChannel)
				return
			}
		}
	}()
	for back := range response.RespondChannel {
		fmt.Println(back)
	}
	e.Wait()
}
