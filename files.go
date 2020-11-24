package testy

import (
	"io/ioutil"
	"os"
)

// CopyFile copies file
func CopyFile(src string, dest string) (err error) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		return err
	}

	return nil
}

// FileSize returns the file size
func FileSize(name string) (s int64, err error) {
	f, err := os.Open(name)
	if err != nil {
		return s, err
	}

	fs, err := f.Stat()
	if err != nil {
		return s, err
	}

	s = fs.Size()

	err = f.Close()
	if err != nil {
		return s, err
	}

	return s, nil
}
