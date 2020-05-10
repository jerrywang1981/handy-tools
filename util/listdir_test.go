package util

import (
	"log"
	"os/user"
	"path/filepath"
	"testing"
)

func TestListDir(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	folderName := filepath.Join(user.HomeDir, "Downloads")
	files, err := ListDir(folderName, "pdf")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(files))
	for _, v := range files {
		t.Log(v)
	}
}
