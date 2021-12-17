package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := `
var yesterdayPrice = "9477.5000";
        var hangq = "https://api-q.fx678img.com/";
        var histor = "https://api-q.fx678img.com/histories.php";
        var symbol = "https://api-q.fx678img.com/getQuote.php";
        var quotes = "https://api-q.fx678img.com/getQuote.php";
        var backgroundurl = "https://quote.fx678.com/assets/img";`
	re := regexp.MustCompile("var yesterdayPrice = \"(.*)\";")
	slice := re.FindAllStringSubmatch(s, -1)[0][1]
	fmt.Println(slice)
}
