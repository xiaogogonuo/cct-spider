package main

import (
	"fmt"
	"time"
)

func main() {
	s := "2017-06-01 00:00:00"
	t, err := time.Parse("2006-01-02 03:04:05", s)
	//fmt.Println(int(t.Month()))
	fmt.Println(err)
	//x := "xxx"
	//x += "yyy"
	fmt.Println(fmt.Sprintf("%d%02d", t.Year(), t.Month()))
}
