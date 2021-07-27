package set

import "fmt"

type Setter interface{}

type Set struct {
	Src Setter
}

// Diff return (dst-src) part
// dst from web, src from db
func (s Set) Diff(dst Setter) (diff [][]string, err error) {
	m := make(map[string]string)
	for _, i := range s.Src.([][]string) {
		if _, ok := m[i[0]]; !ok {
			m[i[0]] = i[1]
		}
	}
	c := 0
	for _, i := range dst.([][]string) {
		if _, ok := m[i[0]]; !ok {
			diff = append(diff, i)
			continue
		}
		if m[i[0]] != i[1] {
			c++
		}
	}
	if c > len(dst.([][]string)) / 2. {
		err = fmt.Errorf("dst data fake")
	}
	return
}
