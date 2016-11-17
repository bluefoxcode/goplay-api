package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

type Heroes struct{}

func (c Heroes) Index(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})

	// find all heroes in the database
	heroes := []*models.hero{}
	if err := models.Heroes.NewQuery().Run(&heroes); err != nil {
		panic(err)
	}

	// render response
	r.JSON(res, 200, heroes)
}
