package repo

import "io"

//FileRepo base store methods interface
type FileRepo interface {
	//Store store file by given filename. returns file location
	Store(f io.Reader, fileName string) (string, error)
}
