package fileReader

import (
	"io/ioutil"
)

// RecursionDir 递归获取指定目录下的所有文件名
func RecursionDir(dir string) (files []string, err error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, fi := range fis {
		fullName := dir + "/" + fi.Name()
		if fi.IsDir() {
			temp, e := RecursionDir(fullName)
			if e != nil {
				return nil, e
			}
			files = append(files, temp...)
		} else {
			files = append(files, fullName)
		}
	}
	return
}
