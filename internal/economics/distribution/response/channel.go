package response

import (
	"github.com/xiaogogonuo/cct-spider/internal/economics/backend"
	"github.com/xiaogogonuo/cct-spider/internal/economics/core"
)

var Stop uint64 = 0
var RequestChannel = make(chan *core.Request)
var RespondChannel = make(chan backend.BackEnd)
