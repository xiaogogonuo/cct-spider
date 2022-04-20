package txt

import (
	"bufio"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"os"
	"strings"
)

// Append2Txt 追加内容到文件，如果文件不存在则创建文件
func Append2Txt(filename string, str ...string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	write := bufio.NewWriter(file)
	for _, s := range str {
		_, _ = write.WriteString(s + "\n")
	}
	_ = write.Flush()
	_ = file.Close()
}

// LoadFromTxt 读取文件内容
func LoadFromTxt(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lineText := strings.Trim(scanner.Text(), "\n")
		lines = append(lines, lineText)
	}
	return lines
}