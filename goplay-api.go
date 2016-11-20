package main

import (
	"log"
	"runtime"

	"github.com/bluefoxcode/goplay-api/lib/boot"
	"github.com/bluefoxcode/goplay-api/lib/router"

	"github.com/urfave/negroni"
)

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	info := boot.LoadConfig()
	boot.RegisterServices(info)

	n := negroni.New(negroni.NewLogger())
	// n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router.Instance())

	n.Run(":" + info.Port)
}
