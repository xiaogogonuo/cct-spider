package economics

import (
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"os"
	"strings"
)

var (
	table               string
	concurrency         int
	configPath          string
	requestConcurrency  int
	responseConcurrency int
)

func init() {
	dir, _ := os.Getwd()
	cp := strings.Join([]string{dir, "configs/economics/meta"}, "/")
	yamlConfig := config.Config{
		ConfigType: "yaml",
		ConfigPath: cp,
		ConfigName: "config",
	}
	v, err := yamlConfig.NewConfig()
	if err != nil {
		panic(err)
	}
	table = v.GetString("table")
	concurrency = v.GetInt("concurrency")
	configPath = v.GetString("configPath")
	requestConcurrency = v.GetInt("requestConcurrency")
	responseConcurrency = v.GetInt("responseConcurrency")
}
