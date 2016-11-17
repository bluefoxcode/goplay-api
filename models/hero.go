package models

type Hero struct {
	Name        string `sql:"name"`
	Description string `sql:"description"`
}
