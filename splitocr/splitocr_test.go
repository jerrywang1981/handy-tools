package splitocr

import "testing"

func TestSplitOcr(t *testing.T) {
	filename := "/home/ocr/output1"
	result := "/home/ocr/result"
	err := SplitOcr(filename, result)
	if err != nil {
		t.Error(err)
	}
}
