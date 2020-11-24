package testy

import (
	"io/ioutil"
	"os"
)

func copyFile(src string, dest string) (err error) {
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

func fileSize(name string) (s int64, err error) {
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
