package filestore

import "api-gmr/storage"

const driverName = "file"

type FileDriver struct {
	*fstorage
}

func (f FileDriver) GetName() string {
	return driverName
}

func init() {
	storage.Register(driverName, &FileDriver{
		fstorage: &fstorage{},
	})
}
