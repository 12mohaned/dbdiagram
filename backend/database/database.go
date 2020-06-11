package database

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

/* Cookie Set-up and Information*/
var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	Password = ""
	dbname   = "dbdiagram"
)

type DatabaseTable struct {
	TableName   string
	Columnname  []string
	Columnvalue []string
	PrimaryKey  string
}

func initConnection() string {
	postgresconnection := "user=" + user + " " + "password=" + Password + " " + "dbname=" + dbname + " " + "sslmode=disable"
	return postgresconnection
}

// /**
// *  CreateNewDatabase
// * * Create a New Database
//  */
// func CreateNewDatabase(databasename string) {
// 	postgresconnection := initConnection()
// 	db, err := sql.Open("postgres", postgresconnection)
// 	if err != nil {
// 		panic(err)
// 	}
// 	sqlStatement := `INSERT INTO database(databasename)
//                     Values ($1)`
// 	row, err := db.Query(sqlStatement, databasename)
// 	if err != nil {
// 		fmt.Println(row)
// 		panic(err)
// 	}

// }

/**
*  CreateNewTable
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
* * Return all Tables
 */
func GetTables() []DatabaseTable {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `select tablename, columnname, columnvalues, primarykey from tables`
	row, err := db.Query(sqlStatement)
	var Tables []DatabaseTable
	defer row.Close()
	for row.Next() {
		var tablename string
		var columnnames []string
		var columnvalues []string
		var primarykey string
		row.Scan(&tablename, pq.Array(&columnnames), pq.Array(&columnvalues), &primarykey)
		table := DatabaseTable{TableName: tablename, Columnname: columnnames, Columnvalue: columnvalues, PrimaryKey: primarykey}
		Tables = append(Tables, table)
	}
	return Tables
}

/**
*  DeleteTables
* * Delete Table with corresponding name
 */
func DeleteTables(tablename string, databasename string) bool {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `delete from tables where tablename = $1`
	row, err := db.Query(sqlStatement, tablename, databasename)
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
	return true
}

/**
Delete a Table
*/
func Droptable(tablename string) {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `delete from tables where tablename = $1`
	row, err := db.Query(sqlStatement, tablename)
	if err != nil {
		fmt.Println(row)
		panic(err)
	}
}

/**
*  ColumnExists
* * Check if a column Exists
 */
func ColumnExists(columnname string) bool {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `select columnname from tables`
	row, err := db.Query(sqlStatement)
	defer row.Close()
	for row.Next() {
		var columnnames []string
		row.Scan(pq.Array(&columnnames))
		if checkColumnname(columnnames, columnname) {
			return true
		}
	}
	return false
}

func checkColumnname(columns []string, columnname string) bool {
	for i := 0; i < len(columns); i++ {
		if columnname == columns[i] {
			return true
		}
	}
	return false
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

/**
*  AddprimaryKey
* * Add a Primarykey to the table
 */
func AddprimaryKey(columnname string, tablename string) {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	sqlStatement := `Update tables set primarykey =$1 where tablename = $2`
	row, err := db.Query(sqlStatement, columnname, tablename)
	if err != nil {
		fmt.Println(row)
		panic(err)
	}
}

/**
Get PrimaryKey
*/
func GetPrimaryKey(tablename string) string {
	postgresconnection := initConnection()
	db, err := sql.Open("postgres", postgresconnection)
	if err != nil {
		panic(err)
	}
	var PrimaryKey string
	sqlStatement := `select primarykey from tables where tablename=$1`
	row, err := db.Query(sqlStatement, tablename)
	if err != nil {
		fmt.Println(row)
		panic(err)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&PrimaryKey)
	}
	return PrimaryKey
}

// /**
// Get Database with corresponding name
// */
// func GetDatabase(databasename string) bool {
// 	postgresconnection := initConnection()
// 	db, err := sql.Open("postgres", postgresconnection)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var PrimaryKey string
// 	sqlStatement := `select databasename from database where databasename =$1`
// 	row, err := db.Query(sqlStatement, databasename)
// 	if err != nil {
// 		fmt.Println(row)
// 		panic(err)
// 	}
// 	return row.Next()
// }
