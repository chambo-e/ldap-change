//go:generate statik -src=./dist -f

package main

import (
	"log"

	_ "github.com/chambo-e/ldap-change/statik"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

func main() {
	r := gin.Default()

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	r.StaticFS("/", statikFS)

	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}
