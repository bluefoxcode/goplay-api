package hero

import (
	"database/sql"
	"fmt"

	"github.com/bluefoxcode/goplay-api/lib/util"
)

var table = "hero"

// Item defines the model.
type Item struct {
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// List gets all items.
func List(db Connection) (Item, bool, error) {
	result := Item{}
	err := db.Get(&result, fmt.Sprintf(`
    SELECT name, description
    FROM %v`, table))
	return result, err == sql.ErrNoRows, err
}

func getCount(db Connection) (int, bool, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %v`, table))
	return result, err == sql.ErrNoRows, err
}

// Initialize sets up the database and prepopulates it.
func Initialize(db Connection) {
	var err error
	err = createTable(db)
	util.CheckErr(err)

	count, _, err := getCount(db)

	if count < 1 {
		populateDB(db)
	}

}

func createTable(db Connection) (err error) {
	_, err = db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %v 
			(
			name text, 
			description text
			)`, table))

	return err
}

func populateDB(db Connection) {

	defaultItems := []Item{
		Item{
			Name:        "Superman",
			Description: "Man of Steel",
		},
		Item{
			Name:        "Batman",
			Description: "Dark Knight",
		},
	}

	for _, item := range defaultItems {
		_, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v 
		(name, description) 
		VALUES
		($1,$2)
		`, table),
			item.Name, item.Description)
		util.CheckErr(err)
	}

}

// Index for heroes db
// func index(w http.ResponseWriter, req *http.Request) {

// 	r := render.New(render.Options{
// 		IndentJSON: true,
// 	})

// 	rows, err := db.Query("SELECT * FROM HEROES")
// 	if err != nil {
// 		log.Println(fmt.Sprintf("Error creating database: %q", err))
// 	}

// 	defer rows.Close()

// 	heroes := make([]*Hero, 0)

// 	for rows.Next() {
// 		hero := new(Hero)

// 		if err := rows.Scan(&hero.Name, &hero.Description); err != nil {
// 			errorCode := http.StatusInternalServerError
// 			r.JSON(w, errorCode, map[string]string{"code": strconv.Itoa(errorCode), "message": fmt.Sprintf("%q", err)})
// 			return
// 		}
// 		heroes = append(heroes, hero)
// 	}

// 	r.JSON(w, http.StatusOK, heroes)
// }

// func create(w http.ResponseWriter, req *http.Request) {
// 	r := render.New(render.Options{
// 		IndentJSON: true,
// 	})
// 	var hero Hero
// 	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
// 	checkErr(err)
// 	err = req.Body.Close()
// 	checkErr(err)

// 	if err := json.Unmarshal(body, &hero); err != nil {
// 		w.WriteHeader(422)
// 		err := json.NewEncoder(w).Encode(err)
// 		checkErr(err)
// 	}

// 	_, err = db.Exec("INSERT INTO heroes values($1, $2)", hero.Name, hero.Description)
// 	checkErr(err)

// 	r.JSON(w, http.StatusOK, hero)
// }
