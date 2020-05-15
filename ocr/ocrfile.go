package ocr

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/jerrywang1981/handy-tools/util"
	"github.com/otiai10/gosseract/v2"
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
	folderName, err := filepath.Abs(foldername)
	if err != nil {
		log.Fatal(err)
		return result
	}
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

	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage(lang)
	for _, file := range fileNames {
		client.SetImage(path.Join(folderName, file))
		text, err := client.Text()
		result = append(result, &OcrResult{FileName: file, Text: text, Error: err})
	}

	return result
}
