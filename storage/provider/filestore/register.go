package filestore

import "api-gmr/storage"

func init() {
	storage.Register("file", NewStore())
}
