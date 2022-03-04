package group

import (
	"github.com/xiaogogonuo/cct-spider/internal/group/cctgroup"
	"github.com/xiaogogonuo/cct-spider/internal/group/sasac"
)

func EntryPoint()  {
	cctgroup.EntryPoint()
	sasac.EntryPoint()
}
