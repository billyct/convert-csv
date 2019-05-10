package main

import (
	csv "github.com/billyct/convert-csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(basePath())
	checkError(err)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
			err = csv.Convert(basePath(file.Name()))
			if err != nil {
				panic(err)
			}
		}
	}
}

// https://blog.csdn.net/skh2015java/article/details/78515002
func basePath(path ...string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	dir = strings.Replace(dir, "\\", "/", -1) //将\替换成/

	elem := append([]string{dir}, path...)

	return filepath.Join(elem...)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
