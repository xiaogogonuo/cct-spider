package main

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/excel"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
)

func main() {
	rows := excel.ReadExcel("./诚通指标配置.xlsx", "T_DMAA_BASE_TARGET")
	for _, row := range rows[12:20] {
		code := row[4]
		mm := md5.MD5(code)
		fmt.Println(mm)
	}
}
