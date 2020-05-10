package pdf

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/jerrywang1981/handy-tools/util"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func ConvertPdfToJpg(pdfFileName, imageFolderName string) error {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	if err := mw.SetResolution(300, 300); err != nil {
		return err
	}

	if err := mw.ReadImage(pdfFileName); err != nil {
		return err
	}

	pdfFileName, _ = filepath.Abs(pdfFileName)
	dir, filename := filepath.Split(pdfFileName)
	filename = strings.TrimSuffix(filename, path.Ext(filename))

	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_OFF); err != nil {
		return err
	}

	if err := mw.SetCompressionQuality(95); err != nil {
		return err
	}

	for i := 0; i < int(mw.CoalesceImages().GetNumberImages()); i++ {
		mw.SetIteratorIndex(i)
		name := path.Join(dir, fmt.Sprintf("%s_%d.jpg", filename, i))
		if err := mw.SetFormat("jpg"); err != nil {
			return err
		}
		if err := mw.WriteImage(name); err != nil {
			return err
		}
	}
	return nil
}

func ConvertPdfsToJpg(pdfFolderName, imageFolderName string) error {
	files, err := util.ListDir(pdfFolderName, "pdf")
	if err != nil {
		return err
	}
	if len(files) < 1 {
		return nil
	}
	for _, f := range files {
		err = ConvertPdfToJpg(f, imageFolderName)
		if err != nil {
			return err
		}
	}
	return nil
}
