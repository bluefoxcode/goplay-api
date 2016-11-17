package boot

import (
	"log"
	"os"
	"runtime"
)

// Info contains application settings.
type Info struct {
	Port        string
	DatabaseURL string
}

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// LoadConfig loads the config object from env vars.
func LoadConfig() *Info {
	config := &Info{}

	config.Port = os.Getenv("PORT")
	config.DatabaseURL = os.Getenv("DATABASE_URL")
	return config
}
