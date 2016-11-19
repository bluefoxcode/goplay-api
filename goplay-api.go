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

	// var err error
	// db, err = sql.Open("postgres", info.DatabaseURL)
	// if err != nil {
	// 	log.Fatalf("Error opening database: %q", err)
	// }
	// initDB()

	// router := mux.NewRouter()
	// heroes := controllers.Heroes{}
	// router.HandleFunc("/heroes", index).Methods("GET")
	// router.HandleFunc("/heroes", create).Methods("POST")

	n := negroni.New(negroni.NewLogger())
	// n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router.Instance())

	n.Run(":" + info.Port)
}
