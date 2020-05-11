package util

import (
	"io/ioutil"
	"os"
	"strings"
)

// list all files in the directory, you can also specify the suffix
// so you can get the files of that type only
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = []string{} //make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)

	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}
