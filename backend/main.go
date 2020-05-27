//A updated verision of dbdiagram

package main

import (
	"backend/database"
	"backend/validator"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	Password = "tarekandamr12/"
	dbname   = "dbdiagram"
)

/*
Table
*/
type Table struct {
	Tablename string
	Row       []row
}
type row struct {
	Columnname  string
	Columnvalue string
}

type Tables struct {
	Tables []Table
}

/**
Recieves Request from users to go to diagrams page
*/
func HomeHandler(Response http.ResponseWriter, Request *http.Request) {
	Request.ParseForm()
	CommanBody := Request.FormValue("command")
	var table Table
	var Colnames []string
	var Colvalues []string
	var tables []Table
	QueryTokenized := strings.Fields(CommanBody)
	if len(QueryTokenized) >= 4 {
		isCreate := validator.CheckCreate(QueryTokenized[0])
		if isCreate {
			isTable := validator.CheckTable(QueryTokenized[1])
			if isTable {
				table = CreateTable(QueryTokenized, QueryTokenized[2])
				fmt.Println(table)
				if len(table.Row) != 0 {
					for i := 0; i < len(table.Row); i++ {
						Colnames = append(Colnames, table.Row[i].Columnname)
						Colvalues = append(Colvalues, table.Row[i].Columnvalue)
					}
					database.CreateNewTable(table.Tablename, Colnames, Colvalues)
				}
			}
		}
	} else {
		if len(QueryTokenized) == 3 {
			CreateRef(QueryTokenized)
		}
	}
	databasetable := database.GetTables()
	for i := 0; i < len(databasetable); i++ {
		t := ConvertTables(databasetable[i].TableName, databasetable[i].Columnname, databasetable[i].Columnvalue)
		tables = append(tables, t)
	}
	T := Tables{Tables: tables}
	template, _ := template.ParseFiles("Home.html")
	template.Execute(Response, T)
}

/**
* * Function which convert Database Table to tables
 */
func ConvertTables(tablename string, columnname []string, columnvalue []string) Table {
	var rows []row
	for i := 0; i < len(columnname); i++ {
		Row := row{Columnname: columnname[i], Columnvalue: columnvalue[i]}
		rows = append(rows, Row)
	}
	return Table{Tablename: tablename, Row: rows}
}
func main() {
	http.HandleFunc("/Home", HomeHandler)
	http.ListenAndServe(":8000", nil)
}

/**
* * creates a new Table and Draw it
 */
func CreateTable(query []string, tableName string) Table {
	// var primaryKey string
	var table Table
	var names []string
	var values []string
	var j int32
	var rows []row
	if !validator.ValidateTableName(query[2]) {
		return Table{}
	}
	for i := 3; i < len(query); i++ {
		if i%2 != 0 {
			names = append(names, query[i])
		} else {
			if validator.ValidateColumnType(query[i]) {
				values = append(values, query[i])

			} else {
				return Table{}
			}
			row := row{Columnname: names[j], Columnvalue: values[j]}
			rows = append(rows, row)
			j++
		}
	}
	table = Table{Tablename: tableName, Row: rows}
	return table
}

/**
* * Creates Relation between Two Tables
 */
func CreateRef(query []string) bool {
	var tablename1 string
	var tablename2 string
	if !validator.ValidateReference(query[1]) {
		return false
	}
	if database.TableExists(query[0]) {
		if database.TableExists(query[2]) {
			tablename1 = query[0]
			tablename2 = query[2]
			if strings.EqualFold(tablename1, tablename2) {
				return false
			}
			database.CreateNewRelation(tablename1, tablename2)
		}
	}
	return true
}
