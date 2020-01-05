package internal

const (
	CreateVidTable = `CREATE TABLE IF NOT EXISTS vids (
        id SERIAL PRIMARY KEY,
		vid  INT4
    );`

	// insert
	insertVidRow = `INSERT INTO vids (
		vid)
	VALUES ($1) `

	InsertVidRow = insertPersonRow + `RETURNING id;`

	SelectVid = `SELECT vid
		FROM vids;`

	// update
	updateVidRow = `UPDATE vids SET vid = $1`

	UpdateVidRow = updateVidRow + `RETURNING id;`

)

