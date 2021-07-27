package findmap

import "regexp"

func FindOne(pat, s string) string {
	reg := regexp.MustCompile(pat)
	v := reg.FindStringSubmatch(s)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func FindAll(pat, s string) (v []string) {
	reg := regexp.MustCompile(pat)
	list := reg.FindAllStringSubmatch(s, -1)
	if len(list) == 0 {
		return
	}
	for _, r := range list{
		v = append(v, r[0])
	}
	return
}
