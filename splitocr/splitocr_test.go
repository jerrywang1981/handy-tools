package splitocr

import "testing"

func TestSplitOcr(t *testing.T) {
	filename := "/Users/jerry/repo/imgs_bak"
	err := SplitOcr(filename, "result")
	if err != nil {
		t.Error(err)
	}
}
