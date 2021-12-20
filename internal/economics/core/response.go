package core

import "github.com/xiaogogonuo/cct-spider/internal/economics/pkg/configReader"

type Response struct {
	Page int
	URL  string
	Body []byte
	Meta configReader.EconomicsConfig
}
