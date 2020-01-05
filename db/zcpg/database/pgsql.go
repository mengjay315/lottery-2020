package database

import (
	"database/sql"
	"github.com/mengjay315/lottery/db/zcpg"
	"log"
)

const (
	host    string = "localhost"
	port    string = "5432"
	user    string = "postgres"
	pass    string = "zchainbtc" // 密码？// 123456!(server)
	dbname  string = "lottery_db"
)

func InitDB() (db *sql.DB) {
	db, err := zcpg.Connect(host, port, user, pass, dbname)
	if err != nil {
		log.Fatalf("connect postgresql error %v", err)
	}

	db.Ping()
	return db
}



