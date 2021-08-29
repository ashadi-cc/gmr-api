package service

import (
	"api-gmr/config"
	"api-gmr/storage"
	"api-gmr/storage/driver"
	"api-gmr/store"
	"api-gmr/store/provider"

	//register drivers
	_ "api-gmr/storage/provider/filestore"
	_ "api-gmr/store/repository/mysql"
)

//load repo driver instance by given app driver configuration
func repo() provider.Driver {
	driver, err := store.Open(config.GetApp().DbDriver)
	if err != nil {
		panic(err)
	}
	return driver
}

func fstorage() driver.Driver {
	driver, err := storage.NewDriver(config.GetApp().StorageDriver)
	if err != nil {
		panic(err)
	}
	return driver
}
