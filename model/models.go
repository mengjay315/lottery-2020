package model

import (
	"github.com/mengjay315/lottery/db/dbtypes"
	"github.com/mengjay315/lottery/db/zcpg"
	"github.com/mengjay315/lottery/db/zcpg/database"
	"log"
)

func CreateTable() error {
	db := database.InitDB()
	err := zcpg.CreateTables(db)
	if err != nil {
		log.Fatalf("create postgresql table error %v", err)
	}

	// 给表 vid赋值 0


	return err
}

func GetPerson(name string) (person dbtypes.PersonBasic, err error) {
	db := database.InitDB()
	person, err = zcpg.QueryPerson(db, name)
	if err != nil {
		log.Fatal(err)
	}

	return person, nil
}

func GetProgram(depart string) (program dbtypes.ProgramBasic, err error) {
	db := database.InitDB()
	program, err = zcpg.QueryProgram(db, depart)
	if err != nil {
		log.Fatal(err)
	}

	return program, nil
}

func GetVoteNums() (res []*dbtypes.ProgramBasic, err error) {
	db := database.InitDB()
	res, err = zcpg.QueryVoteRes(db)
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}

func GetAllPerson() (persons []*dbtypes.PersonBasic, err error) {
	db := database.InitDB()
	persons, err = zcpg.QueryAllPersons(db)
	if err != nil {
		log.Fatal(err)
	}

	return persons, nil
}

func GetSignRes() (res []*dbtypes.PersonBasic, err error) {
	db := database.InitDB()
	res, err = zcpg.QuerySignInRes(db)
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}

func GetVid() (res dbtypes.VoteId, err error) {
	db := database.InitDB()
	res, err = zcpg.QueryVid(db)
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}
