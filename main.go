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
	router.Run("172.168.0.191:8087")
}


