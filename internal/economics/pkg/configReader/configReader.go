package configReader

import (
	"github.com/xiaogogonuo/cct-spider/internal/economics/pkg/internal/fileReader"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"os"
	"strings"
)

type EconomicsConfig struct {
	Case                  string
	TargetCode            string
	TargetName            string
	TargetNameEN          string
	DataSourceCode        string
	DataSourceName        string
	SourceTargetCode      string
	SourceTargetCodeTable string
	IsQuantity            string
	UnitType              string
	UnitName              string
	PeriodType            string
	PeriodName            string
}

// ReadConfig 遍历internal/economics/configs下的yml配置文件
func ReadConfig(configPath string) (economicsConfig []EconomicsConfig) {
	dir, _ := os.Getwd()
	cp := strings.Join([]string{dir, configPath}, "/")
	files, err := fileReader.RecursionDir(cp)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		lastIndex := strings.LastIndex(file, "/")
		yamlConfig := config.Config{
			ConfigType: "yaml",
			ConfigPath: file[:lastIndex],
			ConfigName: file[lastIndex+1:],
		}
		v, err := yamlConfig.NewConfig()
		if err != nil {
			panic(err)
		}
		var economics EconomicsConfig
		economics.Case = v.GetString("Case")
		economics.TargetCode = v.GetString("TargetCode")
		economics.TargetName = v.GetString("TargetName")
		economics.TargetNameEN = v.GetString("TargetNameEN")
		economics.DataSourceCode = v.GetString("DataSourceCode")
		economics.DataSourceName = v.GetString("DataSourceName")
		economics.SourceTargetCode = v.GetString("SourceTargetCode")
		economics.SourceTargetCodeTable = v.GetString("SourceTargetCodeTable")
		economics.IsQuantity = v.GetString("IsQuantity")
		economics.UnitType = v.GetString("UnitType")
		economics.UnitName = v.GetString("UnitName")
		economics.PeriodType = v.GetString("PeriodType")
		economics.PeriodName = v.GetString("PeriodName")
		economicsConfig = append(economicsConfig, economics)
	}
	return
}