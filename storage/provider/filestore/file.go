package filestore

import (
	"api-gmr/config"
	"api-gmr/storage/repo"
	"io"
	"os"
	"path/filepath"
)

type fstorage struct{}

//NewStore returns new fstorage instance
func NewStore() repo.FileRepo {
	return &fstorage{}
}

//Store implementing FileRepo.Store
func (s fstorage) Store(f io.Reader, filename string) (string, error) {
	targetDir, err := filepath.Abs(config.GetApp().BaseImageDir)
	if err != nil {
		return "", err
	}

	fileLocation := filepath.Join(targetDir, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, f); err != nil {
		return "", err
	}

	return fileLocation, nil
}
