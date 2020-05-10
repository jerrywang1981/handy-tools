package splitocr

import "testing"

func TestSplitOcr(t *testing.T) {
	filename := "/home/ocr/repo/img"
	result := "/home/ocr/repo/result"
	err := SplitOcr(filename, result)
	if err != nil {
		t.Error(err)
	}
}
