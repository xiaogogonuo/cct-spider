package arithmetic

import "strings"

func TrimZero(str string) string {
	trim := strings.TrimRight(str, "0")
	if strings.HasSuffix(trim, ".") {
		trim += "0"
	}
	return trim
}
