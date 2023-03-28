package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if err := CreateNestedDirectory(basePath); err != nil {
		return nil, err
	}
	return os.Create(path)
}

func CreateNestedDirectory(path string) error {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		log.Errorf("can't create folder, %s", err)
	}
	return err
}
