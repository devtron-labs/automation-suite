package testUtils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host                  = "localhost"
	port                  = 5432
	user                  = "postgres"
	password              = ""
	dbname                = "orchestrator"
	GETDATA               = "GetData"
	UPDATE_OR_DELETE_DATA = "UpdateOrDeleteData"
	INSERT_DATA           = "InsertData"
)

func ConnectToDB(queryType string, query string) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	CheckError(err)
	switch queryType {
	case GETDATA:
		GetData(db, query)
	case UPDATE_OR_DELETE_DATA:
		UpdateDeleteData(db, query)
	case INSERT_DATA:
		InsertData(db, query)
	}
	defer db.Close()
}

func UpdateDeleteData(db *sql.DB, deleteOrUpdateStmt string) {
	_, e := db.Exec(deleteOrUpdateStmt)
	CheckError(e)
}

func GetData(db *sql.DB, getQuery string) *sql.Rows {
	rows, err := db.Query(`getQuery`)
	CheckError(err)
	return rows
}

func InsertData(db *sql.DB, insertQuery string) {
	_, e := db.Exec(insertQuery)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
