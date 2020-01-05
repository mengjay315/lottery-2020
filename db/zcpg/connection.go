package zcpg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Start the PostgreSQL driver
	"strings"
)

// Connect opens a connection to a PostgreSQL database. The caller is
// responsible for calling Close() on the returned db when finished using it.
// The input host may be an IP address for TCP connection, or an absolute path
// to a UNIX domain socket. An empty string should be provided for UNIX sockets.

func Connect(host, port, user, pass, dbname string) (*sql.DB, error) {
	var psqlInfo string
	if pass == "" {
		psqlInfo = fmt.Sprintf("host=%s user=%s "+
			"dbname=%s sslmode=disable",
			host, user, dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, user, pass, dbname)
	}

	// Only add port arg for TCP connection since UNIX domain sockets (specified
	// by a "/" prefix) do not have a port.
	if !strings.HasPrefix(host, "/") {
		psqlInfo += fmt.Sprintf(" port=%s", port)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// 最大打开连接数和最大空闲连接数
	// 默认最大连接数600
	//db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(500)

	return db, err
}

