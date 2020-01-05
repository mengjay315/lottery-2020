package zcpg

import (
	"database/sql"
	"fmt"
	"github.com/lexkong/log"
	"github.com/mengjay315/lottery/db/zcpg/internal"
)

var createTableStatements = map[string]string{
	"personnels":         internal.CreatePersonTable, // 员工+祝福语+是否投过票（两个字段）
	"programs":           internal.CreateProgramTable, // 部门+节目票数 (两个字段)
	"vids":               internal.CreateVidTable, // 投票信号
}

// CreateTables creates all tables required by ZChain if they do not already exist.
func CreateTables(db *sql.DB) error {
	// Create all of the data tables.
	for tableName, createCommand := range createTableStatements {
		err := createTable(db, tableName, createCommand)
		if err != nil {
			return err
		}
	}

	return nil
}

// createTable creates a table with the given name using the provided SQL
// statement, if it does not already exist.
func createTable(db *sql.DB, tableName, stmt string) error {
	exists, err := TableExists(db, tableName)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Printf("Creating the %s table.\n", tableName)
		_, err = db.Exec(stmt)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Table the %s exists.\n", tableName)
	}

	return err
}

// TableExists checks if the specified table exists.
func TableExists(db *sql.DB, tableName string) (bool, error) {
	rows, err := db.Query(`select relname from pg_class where relname = $1`,
		tableName)
	if err != nil {
		return false, err
	}

	defer func() {
		if e := rows.Close(); e != nil {
			log.Errorf(e, "Close of Query failed")
		}
	}()
	return rows.Next(), nil
}
