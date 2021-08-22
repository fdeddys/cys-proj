package main

import (
	_ "github.com/fdeddys/tes/database"
	"github.com/fdeddys/tes/router"
)

func main() {
	router.InitRouter().Run(":8888")
}
