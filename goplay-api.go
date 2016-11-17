package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bluefoxcode/goplay-api/boot"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
)

var (
	db *sql.DB
)

// Hero is model
type Hero struct {
	Name        string `sql:"name" json:"name"`
	Description string `sql:"description" json:"description"`
}

func main() {
	info := boot.LoadConfig()

	var err error
	db, err = sql.Open("postgres", info.DatabaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	initDB()

	router := mux.NewRouter()
	// heroes := controllers.Heroes{}
	router.HandleFunc("/heroes", index).Methods("GET")
	router.HandleFunc("/heroes", create).Methods("POST")
	// router.HandleFunc("/heroes", heroes.Show).Methods("GET")
	// router.HandleFunc("/heroes", heroes.Update).Methods("GET")
	// router.HandleFunc("/heroes", heroes.Delete).Methods("GET")

	n := negroni.New(negroni.NewLogger())
	// n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router)

	n.Run(":" + info.Port)
}

func initDB() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS heroes (name text, description text)")
	checkErr(err)

	rows, err := db.Query("SELECT COUNT(*) as count FROM  heroes")
	checkErr(err)
	defer rows.Close()
	if checkCount(rows) < 1 {
		populateHeroes()
	}

}

func populateHeroes() {

	defaultHeroes := []Hero{
		Hero{
			Name:        "Superman",
			Description: "Man of Steel",
		},
		Hero{
			Name:        "Batman",
			Description: "Dark Knight",
		},
	}

	for _, hero := range defaultHeroes {
		_, err := db.Exec("INSERT INTO heroes(name, description) values($1, $2)", hero.Name, hero.Description)
		checkErr(err)
	}

}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Index for heroes db
func index(w http.ResponseWriter, req *http.Request) {

	r := render.New(render.Options{
		IndentJSON: true,
	})

	rows, err := db.Query("SELECT * FROM HEROES")
	if err != nil {
		log.Println(fmt.Sprintf("Error creating database: %q", err))
	}

	defer rows.Close()

	heroes := make([]*Hero, 0)

	for rows.Next() {
		hero := new(Hero)

		if err := rows.Scan(&hero.Name, &hero.Description); err != nil {
			errorCode := http.StatusInternalServerError
			r.JSON(w, errorCode, map[string]string{"code": strconv.Itoa(errorCode), "message": fmt.Sprintf("%q", err)})
			return
		}
		heroes = append(heroes, hero)
	}

	r.JSON(w, http.StatusOK, heroes)
}

func create(w http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{
		IndentJSON: true,
	})
	var hero Hero
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	checkErr(err)
	err = req.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &hero); err != nil {
		w.WriteHeader(422)
		err := json.NewEncoder(w).Encode(err)
		checkErr(err)
	}

	_, err = db.Exec("INSERT INTO heroes values($1, $2)", hero.Name, hero.Description)
	checkErr(err)

	r.JSON(w, http.StatusOK, hero)
}
