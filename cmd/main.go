package main

import (
	"log"

	"github.com/akmuhammetakmyradov/test/internal/app"
	"github.com/akmuhammetakmyradov/test/pkg/config"
)

func main() {
	config, err := config.LoadConfiguration()
	if err != nil {
		log.Println(err)
	}

	if err = app.InitApp(config); err != nil {
		log.Fatalln("err in InitApp: ", err)
	}
}
