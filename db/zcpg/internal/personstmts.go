package internal

const (
	CreatePersonTable = `CREATE TABLE IF NOT EXISTS personnels (
        id SERIAL PRIMARY KEY,
        name VARCHAR(12),
        memo VARCHAR(64),
		time  INT8,
		able  INT4
    );`

	//name VARCHAR(12),
	//able  INT2
	//memo VARCHAR(16),

	// insert
	insertPersonRow = `INSERT INTO personnels (
		name, memo, time, able)
	VALUES ($1, $2, $3, $4) `

	InsertPersonRow = insertPersonRow + `RETURNING id;`

	SelectPersonal = `SELECT name, memo, time, able
		FROM personnels WHERE name = $1;`

	// update
	updatePersonRow = `UPDATE personnels SET able = $1 
        WHERE name = $2`

	UpdatePersonRow = updatePersonRow + `RETURNING id;`

	SelectPersonRes = `SELECT name, memo, time, able
		FROM personnels;`

	SelectPersonResLimit = `SELECT name, memo, time, able
		FROM personnels ORDER BY time DESC LIMIT 3;`
)
