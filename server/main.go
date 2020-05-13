package main

import (
	"github.com/go_nuxt_2/server/environment/router"
)

func main() {
	e := router.NewRouter()

	e.Start(":8080")
}
