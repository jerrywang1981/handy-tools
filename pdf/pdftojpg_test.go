package pdf

import (
	"fmt"
	"testing"
)

func TestConvertPdfsToJpg(t *testing.T) {
	pdfFileName := "./testdata"
	outputFolder := "./testdata"
	if err := ConvertPdfsToJpg(pdfFileName, outputFolder); err != nil {
		fmt.Println(err)
	}

}
