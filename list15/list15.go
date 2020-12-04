package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	err := Walk(".", func(b []byte, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	})
	fmt.Println(err)
}

func Walk(dir string, f func(b []byte, err error) error) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			b, err := ioutil.ReadFile(path)
			if err := f(b, err); err != nil {
				return err
			}
			continue
		}
		if err := Walk(path, f); err != nil {
			return err
		}
	}
	return nil
}
