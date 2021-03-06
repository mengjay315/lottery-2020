package internal

const (
	CreateProgramTable = `CREATE TABLE IF NOT EXISTS programs (
        id SERIAL PRIMARY KEY,
        depart  CHAR(32),
		number  INT4
    );`

	// insert
	insertProgramRow = `INSERT INTO programs (
		depart, number)
	VALUES ($1, $2) `

	InsertProgramRow = insertProgramRow + `RETURNING id;`


    // update
	updateProgramRow = `UPDATE programs SET number = $1 
        WHERE depart = $2`

	UpdateProgramRow = updateProgramRow + `RETURNING id;`

	SelectProgram = `SELECT depart, number
		FROM programs WHERE depart = $1;`

	SelectProgramRes = `SELECT depart, number
		FROM programs;`
)
