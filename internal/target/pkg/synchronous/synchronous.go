package synchronous

import (
	"bufio"
	"encoding/csv"
	"github.com/xiaogogonuo/cct-spider/internal/target/pkg/txt"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"io"
	"os"
	"strings"
)

// HistoryDataEncrypt 为爬虫服务器和数据库服务器同步做数据md5加密操作
func HistoryDataEncrypt(filepath string) {
	csvFile, _ := os.Open(filepath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var ids []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		targetCode := strings.Trim(line[0], "\n ")
		regionCode := strings.Trim(line[1], "\n ")
		year := strings.Trim(line[2], "\n ")
		season := strings.Trim(line[3], "\n ")
		month := strings.Trim(line[4], "\n ")
		date := strings.Trim(line[5], "\n ")
		var id string
		if len(date) == 8 {
			id = md5.MD5(targetCode + date + regionCode)
		} else {
			id = md5.MD5(targetCode + year + season + month + regionCode)
		}
		ids = append(ids, id)
	}
	txt.Append2Txt("target.txt", ids...)
}
