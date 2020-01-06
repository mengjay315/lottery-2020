package main

import (
	"github.com/mengjay315/lottery/model"
	"github.com/mengjay315/lottery/routers"
	"log"
)

func main() {

	err := model.CreateTable()
	if err != nil {
		log.Fatalf("create postgresql table error %v", err)
	}

	router := routers.InitRouter()
	router.Run("192.168.60.251:8087")
}
