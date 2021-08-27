package store

import (
	"api-gmr/store/provider"
	"fmt"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]provider.Driver)
)

//Register register the driver
func Register(name string, driver provider.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("repo: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("repo: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

//Open the driver
func Open(driverName string) (provider.Driver, error) {
	driversMu.RLock()
	driveri, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown driver %q (forgotten import?)", driverName)
	}
	return driveri, nil
}
