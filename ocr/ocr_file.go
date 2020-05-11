package ocr

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/jerrywang1981/handy-tools/util"
	"github.com/otiai10/gosseract"
)

var (
	fileExt = []string{".jpg", ".png"}
)

type OcrResult struct {
	FileName string
	Text     string
	Error    error
}

func RecognizeFile(filename, lang string) (string, error) {
	file, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage(lang)
	client.SetImage(file)
	text, err := client.Text()
	return text, err
}

// RecognizeFolder ocr the files in the folder
func RecognizeFolder(foldername, lang string) []*OcrResult {
	result := []*OcrResult{}
	files, err := ioutil.ReadDir(foldername)
	if err != nil {
		log.Fatal(err)
		return result
	}
	fileNames := []string{}
	for _, file := range files {
		if !file.IsDir() &&
			util.Contains(fileExt, strings.ToLower(filepath.Ext(file.Name()))) {
			fileNames = append(fileNames, file.Name())
		}
	}
	count := len(fileNames)
	if count == 0 {
		return result
	}
	ch := make(chan *OcrResult)
	for _, name := range fileNames {
		go func(filename string) {
			text, err := RecognizeFile(filename, lang)
			ch <- &OcrResult{FileName: filename, Text: text, Error: err}
		}(name)
	}

	for i := 0; i < count; i++ {
		result = append(result, <-ch)
	}

	return result
}
