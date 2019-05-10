package convert_csv

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

// 读取文件内容后，按照原来的 CSV 文件路径生成文件
func Convert(path string) error {
	b, err := read(path)
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = out.Write(b)

	return err
}

// 读取 CSV 内容，已将 GBK 转换成 UTF-8
func read(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	if !utf8.Valid(content) {
		content, err = simplifiedchinese.GBK.NewDecoder().Bytes(content)
		if err != nil {
			return []byte{}, err
		}
	}

	return content, nil
}
