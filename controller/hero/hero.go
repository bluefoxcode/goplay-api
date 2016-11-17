package hero

import (
	"log"
	"net/http"

	"github.com/bluefoxcode/goplay-api/lib/util"
	"github.com/bluefoxcode/goplay-api/model/hero"
	"github.com/gorilla/mux"
)

var (
	url = "/heroes"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
}

// Load the routes.
func Load() {
	router.HandleFunc("/heroes", Index).Methods("GET")
}

// Index displays list of heroes
func Index(w http.ResponseWriter, r *http.Request) {
	c := util.Context(w, r)
	items, _, err := hero.List(c.DB)
	if err != nil {
		panic(err)
	}
	log.Println(items)
	// r := render.New(render.Options{
	// 	IndentJSON: true,
	// })

	// rows, err := db.Query("SELECT * FROM HEROES")
	// if err != nil {
	// 	log.Println(fmt.Sprintf("Error creating database: %q", err))
	// }

	// defer rows.Close()

	// heroes := make([]*Hero, 0)

	// for rows.Next() {
	// 	hero := new(Hero)

	// 	if err := rows.Scan(&hero.Name, &hero.Description); err != nil {
	// 		errorCode := http.StatusInternalServerError
	// 		r.JSON(w, errorCode, map[string]string{"code": strconv.Itoa(errorCode), "message": fmt.Sprintf("%q", err)})
	// 		return
	// 	}
	// 	heroes = append(heroes, hero)
	// }

	// r.JSON(w, http.StatusOK, heroes)
}
