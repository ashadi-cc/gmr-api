package service

import (
	"api-gmr/config"
	"api-gmr/store"
	"api-gmr/store/provider"

	//register repo driver
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
