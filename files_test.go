package testy

import (
	"os"
	"testing"
)

func Test_copyFile(t *testing.T) {
	t.Run("existing file", func(t *testing.T) {
		err := CopyFile("testdata/file.txt", "testdata/tmp.txt")
		if err != nil {
			t.Error("Can not copy file")
		}

		err = os.Remove("testdata/tmp.txt")
		if err != nil {
			t.Error("Can not delete file")
		}
	})

	t.Run("not existing file", func(t *testing.T) {
		err := CopyFile("testdata/unknown.txt", "testdata/tmp.txt")
		if err == nil {
			t.Error("No error")
		}
	})
}

func Test_fileSize(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		s, err := FileSize("testdata/file.txt")
		if err != nil {
			t.Fatal("Can not get file size")
		}

		if s != 8 {
			t.Error(s, "!=", 8)
		}
	})

	t.Run("not existing", func(t *testing.T) {
		_, err := FileSize("testdata/unknown.txt")
		if err == nil {
			t.Fatal("No error")
		}
	})
}
