package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func IsFile(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFiles(dir string) ([]os.FileInfo, error) {
	if !IsDir(dir) {
		return nil, fmt.Errorf("path is not a dir: %s", dir)
	}

	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fls, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}
	return fls, nil
}

func GetFilesSortByMTime(dir string) ([]os.FileInfo, error) {
	fls, err := GetFiles(dir)
	if err != nil {
		return nil, err
	}

	sort.Slice(fls, func(i, j int) bool {
		return fls[i].ModTime().Before(fls[j].ModTime())
	})

	return fls, nil
}

func GetFilesSortByName(dir string)([]os.FileInfo, error) {
	fls, err := GetFiles(dir)
	if err != nil {
		return nil, err
	}

	sort.Slice(fls, func(i, j int) bool {
		return strings.Compare(fls[i].Name(), fls[j].Name()) < 0
	})

	return fls, nil
}

func WritePid(pth string) error {
	file, err := os.OpenFile(pth, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	if _, err = file.WriteString(strconv.Itoa(os.Getpid())); err != nil {
		return err
	}
	return nil
}

func DelPid(pth string) error {
	return os.Remove(pth)
}

func IsFiles(fs []string) (bool, error) {
	if len(fs) == 0 {
		return false, nil
	}
	for _, v := range fs {
		m, err := filepath.Glob(v)
		if err != nil {
			return false, err
		}
		if len(m) > 0 {
			return true, nil
		}
	}
	return false, nil
}

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
