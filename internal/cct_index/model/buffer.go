package model

import "time"

type Buffer struct {
	Date        string
	TargetValue string
	RegionCode  string
	RegionName  string
}

type BufferTF struct {
	Date        time.Time
	TargetValue float64
	RegionCode  string
	RegionName  string
}

type BufferTFList []BufferTF

func (b BufferTFList) Len() int {
	return len(b)
}

func (b BufferTFList) Less(i, j int) bool {
	return b[i].Date.String() > b[j].Date.String() // 按照日期降序
}

func (b BufferTFList) Swap(i, j int) {
	b[i].Date, b[j].Date = b[j].Date, b[i].Date
}
