package zcpg

import (
	"database/sql"
	"github.com/mengjay315/lottery/db/dbtypes"
	"github.com/mengjay315/lottery/db/zcpg/internal"
)

func InsertPerson(db *sql.DB, dbPerson *dbtypes.PersonBasic) (uint64, error) {
	insertStatement := internal.InsertPersonRow
	var id uint64

	err := db.QueryRow(insertStatement,
		dbPerson.Name, dbPerson.Memo, dbPerson.Time, dbPerson.Able).Scan(&id)
	return id, err
}

func UpdatePerson(db *sql.DB, dbPerson *dbtypes.PersonBasic) (uint64, error) {
	updateStatement := internal.UpdatePersonRow
	var id uint64

	err := db.QueryRow(updateStatement,
		dbPerson.Able, dbPerson.Name).Scan(&id)
	return id, err
}


func InsertProgram(db *sql.DB, dbProgram *dbtypes.ProgramBasic) (uint64, error) {
	insertStatement := internal.InsertProgramRow
	var id uint64

	err := db.QueryRow(insertStatement,
		dbProgram.Depart, dbProgram.Number).Scan(&id)
	return id, err
}

func UpdateProgram(db *sql.DB, dbProgram *dbtypes.ProgramBasic) (uint64, error) {
	updateStatement := internal.UpdateProgramRow
	var id uint64

	err := db.QueryRow(updateStatement,
		dbProgram.Number, dbProgram.Depart).Scan(&id)
	return id, err
}

func UpdateVoteID(db *sql.DB, dbVoteID *dbtypes.VoteId) (uint64, error) {
	updateStatement := internal.UpdateVidRow
	var id uint64

	err := db.QueryRow(updateStatement,
		dbVoteID.Vid).Scan(&id)
	return id, err
}

func QueryPerson(db *sql.DB, name string) (person dbtypes.PersonBasic, err error) {
	personQueryStmt := internal.SelectPersonal
	rows, err := db.Query(personQueryStmt, name)
	checkErr(err)
	defer rows.Close()
	for rows.Next(){

		res := dbtypes.PersonBasic{}
		err := rows.Scan(&res.Name, &res.Memo, &res.Time, &res.Able)
		checkErr(err)
		person = res
	}

	return person, err

}

func QueryProgram(db *sql.DB, depart string) (program dbtypes.ProgramBasic, err error) {
	programQueryStmt := internal.SelectProgram
	rows, err := db.Query(programQueryStmt, depart)
	checkErr(err)
	defer rows.Close()
	for rows.Next(){

		res := dbtypes.ProgramBasic{}
		err := rows.Scan(&res.Depart, &res.Number)
		checkErr(err)
		program = res
	}

	return program, err
}

func QueryVoteRes(db *sql.DB) (res []*dbtypes.ProgramBasic, err error) {
	voteQueryStmt := internal.SelectProgramRes

	rows, err := db.Query(voteQueryStmt)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		program := dbtypes.ProgramBasic{}
		err := rows.Scan(&program.Depart, &program.Number)
		checkErr(err)
		res = append(res, &program)
	}
	return res, err
}

func QueryAllPersons(db *sql.DB) (res []*dbtypes.PersonBasic, err error) {
	personsQueryStmt := internal.SelectPersonRes

	rows, err := db.Query(personsQueryStmt)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		person := dbtypes.PersonBasic{}
		err := rows.Scan(&person.Name, &person.Memo, &person.Time, &person.Able)
		checkErr(err)
		res = append(res, &person)
	}
	return res, err
}

func QuerySignInRes(db *sql.DB) (res []*dbtypes.PersonBasic, err error) {
	personLimitQueryStmt := internal.SelectPersonResLimit

	rows, err := db.Query(personLimitQueryStmt)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		person := dbtypes.PersonBasic{}
		err := rows.Scan(&person.Name, &person.Memo, &person.Time, &person.Able)
		checkErr(err)
		res = append(res, &person)
	}
	return res, err
}

func QueryVid(db *sql.DB) (res dbtypes.VoteId, err error) {
	vidQueryStmt := internal.SelectVid
	rows, err := db.Query(vidQueryStmt)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		vid := dbtypes.VoteId{}
		err := rows.Scan(&vid.Vid)
		checkErr(err)
		res = vid
	}
	return res, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
