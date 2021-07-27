package filter

import (
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type Filter struct {
	sync.Mutex
	Filepath   string
	ThisUrlKey map[string]byte
}


func (f *Filter) WriteMap(s string){
	f.Lock()
	if _, ok := f.ThisUrlKey[s]; !ok {
		f.ThisUrlKey[s] = 0
	}
	f.Unlock()
}

func (f *Filter) SaveUrlKey() {
	var strSlice []string
	for k := range f.ThisUrlKey{
		strSlice = append(strSlice, k)
	}
	w, err := os.OpenFile(f.Filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil{
		logger.Error(err.Error())
		return
	}
	_, err = w.Write([]byte(strings.Join(strSlice, "\n") + "\n"))
	if err1 := w.Close(); err1 != nil && err == nil {
		err = err1
		logger.Error(err.Error())
		return
	}
	f.ThisUrlKey = map[string]byte{}
}

func (f *Filter) ReadUrlKey() (urlKey map[string]byte) {
	urlKey = make(map[string]byte)
	fi, err := os.Open(f.Filepath)
	if err != nil{
		return
	}
	fd, err := ioutil.ReadAll(fi)
	if err1 := fi.Close(); err1 != nil && err == nil {
		err = err1
		logger.Error(err.Error())
		return
	}
	for _, k := range strings.Split(string(fd), "\n"){
		if _, ok := urlKey[k]; !ok{
			urlKey[k] = 0
		}
	}
	return
}
