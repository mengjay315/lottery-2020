package dbtypes


type PersonBasic struct {
	Name       string  `json:"name"`
	Memo       string  `json:"memo"`
	Time       int64   `json:"time"`
	Able       int     `json:"able"`
}

type ProgramBasic struct {
	Depart       string  `json:"depart"`
	Number       int  	`json:"number"`
}

type VoteId struct {
	Vid       int  	`json:"vid"`
}


