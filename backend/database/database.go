package database

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	Password = "tarekandamr12/"
	dbname   = "dbdiagram"
)

type DatabaseTable struct {
	TableName   string
	Columnname  []string
	Columnvalue []string
}

func initConnection() string {
	postgresconnection := "user=" + user + " " + "password=" + Password + " " + "dbname=" + dbname + " " + "sslmode=disable"
	return postgresconnection
}

/**
*  CreateNewGroup
* * Create a New Table
 */
func CreateNewTable(tableName string, columnnames []string, columnvalues []string) {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `INSERT INTO tables (tablename, columnname, columnvalues)
                    Values ($1,$2,$3)`
	row, err := db.Query(sqlStatement, tableName, pq.Array(columnnames), pq.Array(columnvalues))
	if err != nil {
		fmt.Println(row)
		panic(err)
	}

}

/**
*  GetTables
* * Return all Tables Created
 */
func GetTables() []DatabaseTable {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `select tablename, columnname, columnvalues from tables`
	row, err := db.Query(sqlStatement)
	var Tables []DatabaseTable
	defer row.Close()
	for row.Next() {
		var tablename string
		var columnnames []string
		var columnvalues []string
		row.Scan(&tablename, pq.Array(&columnnames), pq.Array(&columnvalues))
		table := DatabaseTable{TableName: tablename, Columnname: columnnames, Columnvalue: columnvalues}
		Tables = append(Tables, table)
	}
	return Tables
}

/**
*  DeleteTables
* * Delete Table with corresponding name
 */
func DeleteTables(tablename string) bool {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `delete from tables where tablename=$1`
	row, err := db.Query(sqlStatement, tablename)
	if err != nil {
		fmt.Println(row)
		panic(err)
	}
	return true
}

/**
*  TableExists
* * Check if a Table Does Exists or not
 */
func TableExists(tableName string) bool {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `select tablename from tables where tablename = $1`
	row, err := db.Query(sqlStatement, tableName)
	if err != nil {
		fmt.Println(row)
		panic(err)
	}
	return row.Next()
}

/**
*  CreateRelation
* * Create a new Relation
 */
func CreateNewRelation(tablename1 string, tablename2 string) {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `INSERT INTO relations (tablename1, tablename2)
    Values ($1,$2)`
	row, err := db.Query(sqlStatement, tablename1, tablename2)
	if err != nil {
		fmt.Println(row)
		panic(err)
	}
}
