package text

import (
	"bufio"
	"os"
	"strings"
)

// ReadFromText 从文本文件读取所有行到内存
func ReadFromText(filename string) (rows []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Trim(scanner.Text(), "\n")
		if row != "" {
			rows = append(rows, row)
		}
	}
	return
}

// AppendToText 追加字符串到文本文件尾部
func AppendToText(filename string, str ...string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	write := bufio.NewWriter(file)
	for _, s := range str {
		_, _ = write.WriteString(s + "\n")
	}
	_ = write.Flush()
	_ = file.Close()
	return
}
