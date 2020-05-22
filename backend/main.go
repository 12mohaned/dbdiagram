//A updated verision of dbdiagram

package main

import (
	"backend/validator"
	"fmt"
	"html/template"
	"net/http"
	"strings"
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

/**
Recieves Request from users to go to diagrams page
*/
func HomeHandler(Response http.ResponseWriter, Request *http.Request) {
	Request.ParseForm()
	CommanBody := Request.FormValue("command")
	var table Table
	QueryTokenized := strings.Fields(CommanBody)
	if len(QueryTokenized) >= 4 {
		isCreate := strings.Contains(strings.ToLower(QueryTokenized[0]), "create")
		if isCreate {
			isTable := strings.Contains(strings.ToLower(QueryTokenized[1]), "table")
			if isTable {
				table = CreateTable(QueryTokenized, QueryTokenized[2])
			}
		}
	}
	fmt.Println(table.Row)
	template, _ := template.ParseFiles("Home.html")
	template.Execute(Response, table)
}

func main() {
	http.HandleFunc("/Home", HomeHandler)
	http.ListenAndServe(":8000", nil)
}

/**
* * Function which creates new Table and Draw it
 */
func CreateTable(query []string, tableName string) Table {
	// var primaryKey string
	var table Table
	var names []string
	var values []string
	var j int32
	var rows []row
	if !validator.ValidateTableName(query[2]) {
		return table
	}
	for i := 3; i < len(query); i++ {
		if i%2 != 0 {
			names = append(names, query[i])
		} else {
			if validator.ValidateColumnType(query[i]) {
				values = append(values, query[i])
			} else {
				return table
			}
			row := row{Columnname: names[j], Columnvalue: values[j]}
			rows = append(rows, row)
			j++
		}
	}
	table = Table{Tablename: tableName, Row: rows}
	return table
}
