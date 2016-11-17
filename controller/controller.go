package controller

import (
	"github.com/bluefoxcode/goplay-api/controller/hero"
)

// LoadRoutes loads the routes for each of the controllers
func LoadRoutes() {
	hero.Load()
}
