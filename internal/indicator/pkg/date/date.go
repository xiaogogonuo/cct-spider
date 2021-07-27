package date

import (
	"errors"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"math"
	"strconv"
	"time"
)

func dateCheck(dateType string, indicatorInfo map[string]string) (err error) {
	if len(indicatorInfo[dateType]) == 0 {
		msg := fmt.Sprintf("%s doesn't set %s", indicatorInfo["TargetCode"], dateType)
		logger.Error(msg)
		err = errors.New(msg)
		return
	}
	return
}

// 季度数字转字母
func numberToCharacter(num int) (character string) {
	switch num {
	case 1:
		character = "A"
	case 2:
		character = "B"
	case 3:
		character = "C"
	case 4:
		character = "D"
	}
	return
}

// 季度字母转数字
func characterToNum(character string) (num int) {
	switch character {
	case "A":
		num = 1
	case "B":
		num = 2
	case "C":
		num = 3
	case "D":
		num = 4
	}
	return
}

func StartToCurrentYear(indicatorInfo map[string]string) (history map[string]struct{}, err error) {
	if err = dateCheck("StartYear", indicatorInfo); err != nil {
		return
	}
	startYear, err := strconv.Atoi(indicatorInfo["StartYear"])
	if err != nil {
		logger.Error(err.Error())
		return
	}
	currentYear := time.Now().Year()
	history = make(map[string]struct{})
	for i := startYear; i < currentYear; i++ {
		history[strconv.Itoa(i)] = struct{}{}
	}
	return
}

func StartToCurrentSeason(indicatorInfo map[string]string) (history map[string]struct{}, err error) {
	if err = dateCheck("StartYear", indicatorInfo); err != nil {
		return
	}
	if err = dateCheck("StartSeason", indicatorInfo); err != nil {
		return
	}
	history = make(map[string]struct{})
	startYear := indicatorInfo["StartYear"]
	startSeason := indicatorInfo["StartSeason"]
	currentYear := time.Now().Year()
	currentSeason := int(math.Ceil(float64(time.Now().Month()) / 3.))
	for i := characterToNum(startSeason); i <= 4; i++ {
		history[fmt.Sprintf("%s%s", startYear, numberToCharacter(i))] = struct{}{}
	}
	startYearInt, err := strconv.Atoi(startYear)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	for i := startYearInt + 1; i < currentYear; i++ {
		for j := 1; j <= 4; j++ {
			history[fmt.Sprintf("%d%s", i, numberToCharacter(j))] = struct{}{}
		}
	}
	for i := 1; i < currentSeason; i++ {
		history[fmt.Sprintf("%d%s", currentYear, numberToCharacter(i))] = struct{}{}
	}
	return
}

func StartToCurrentMonth(indicatorInfo map[string]string) (history map[string]struct{}, err error) {
	if err = dateCheck("StartYear", indicatorInfo); err != nil {
		return
	}
	if err = dateCheck("StartMonth", indicatorInfo); err != nil {
		return
	}
	startYear := indicatorInfo["StartYear"]
	startYearInt, err := strconv.Atoi(startYear)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	startMonth := indicatorInfo["StartMonth"]
	startMonthInt, err := strconv.Atoi(startMonth)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	history = make(map[string]struct{})
	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())
	for i := startMonthInt; i <= 12; i++ {
		if i < 10 {
			history[fmt.Sprintf("%s0%d", startYear, i)] = struct{}{}
		} else {
			history[fmt.Sprintf("%s%d", startYear, i)] = struct{}{}
		}
	}
	for i := startYearInt + 1; i < currentYear; i++ {
		for j := 1; j <= 12; j++ {
			if j < 10 {
				history[fmt.Sprintf("%d0%d", i, j)] = struct{}{}
			} else {
				history[fmt.Sprintf("%d%d", i, j)] = struct{}{}
			}
		}
	}
	for i := 1; i < currentMonth; i++ {
		if i < 10 {
			history[fmt.Sprintf("%d0%d", currentYear, i)] = struct{}{}
		} else {
			history[fmt.Sprintf("%d%d", currentYear, i)] = struct{}{}
		}
	}
	return
}

// Diff return dst - src and src value close to 0
// src from db, dst calculated
func Diff(src [][]string, dst map[string]struct{}) (diff []string) {
	m := make(map[string]string)
	diff = make([]string, 0)
	for _, row := range src {
		if _, ok := m[row[0]]; !ok {
			m[row[0]] = row[1]
		}
	}
	for date := range dst {
		if _, ok := m[date]; !ok {
			diff = append(diff, date)
			continue
		}
		// 对于值为空的数据重新爬取，不管它是不是真的为空
		if m[date] == "" {
			diff = append(diff, date)
			continue
		}
		v, err := strconv.ParseFloat(m[date], 64)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		// 对于值为零的数据重新爬取，不管它真的是不是零
		if v < 0.001 {
			diff = append(diff, date)
		}
	}
	return
}

func SplitDiff(diff []string, block int) (diffBlock [][]string) {
	pieces := int(math.Ceil(float64(len(diff)) / float64(block)))
	diffBlock = make([][]string, 0)
	for i := 0; i < pieces - 1; i++ {
		diffBlock = append(diffBlock, diff[i*block:(i+1)*block])
		diffBlock[i] = diff[i*block:(i+1)*block]
	}
	diffBlock = append(diffBlock, diff[block * (pieces-1):])
	return
}