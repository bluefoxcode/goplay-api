package util

import (
	"net/http"
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	dbInfo *sqlx.DB
	mutex  sync.RWMutex
)

// StoreDB stores the database connection settigns so controller functions can access them safely.
func StoreDB(db *sqlx.DB) {
	mutex.Lock()
	dbInfo = db
	mutex.Unlock()
}

// info structures the application settings.
type Info struct {
	W  http.ResponseWriter
	R  *http.Request
	DB *sqlx.DB
}

// Context returns the application settings.
func Context(w http.ResponseWriter, r *http.Request) Info {
	mutex.RLock()
	i := Info{
		W:  w,
		R:  r,
		DB: dbInfo,
	}
	mutex.RUnlock()
	return i
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
