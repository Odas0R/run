package fs

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

const (
	DefaultPerms = 0600
)

func Touch(path string) error {
	myfile, err := os.Create(path)
	if err != nil {
		return err
	}

	if err := myfile.Close(); err != nil {
		return err
	}

	return nil
}

func Cat(text string, path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(text); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func Find(path string) []string {
	var files []string
	err := fs.WalkDir(os.DirFS(path), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return files
}

// list all directories full paths, recursively, in a given path with a given depth limit (0 for no limit)
func FindDir(path string, depth int) []string {
	var dirs []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, path+"/"+file.Name())
			if depth > 0 {
				dirs = append(dirs, FindDir(path+"/"+file.Name(), depth-1)...)
			}
		}
	}
	return dirs
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}

func NotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, fs.ErrNotExist)
}
