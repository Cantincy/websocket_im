package main

import (
	"log"
	"newim/dao"
	"newim/entity"
	"newim/router"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		entity.Svr.DispatchMsg()
	}()

	engine := router.NewRouter()

	engine.Run(":9090")
}
