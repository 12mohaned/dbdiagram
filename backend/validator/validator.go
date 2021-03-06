package validator

import (
	"regexp"
	"strings"
)

/**
** Check the create command
**/
func CheckCreate(command string) bool {
	if strings.Contains(strings.ToLower(command), "create") {
		return true
	}
	return false
}

/**
** Check the table keyword
**/
func CheckTable(table string) bool {
	if strings.Contains(strings.ToLower(table), "table") {
		return true
	}
	return false
}

/**
** Check the Database keyword
**/
func CheckDatabase(databasename string) bool {
	if strings.Contains(strings.ToLower(databasename), "database") {
		return true
	}
	return false
}

/**
** Validate the Database Name
**/
func ValidateTableName(databasename string) bool {
	isValid, _ := regexp.MatchString("[a-zA-Z]{3,40}", databasename)
	return isValid
}

/**
** Validate the Table Name
**/
func ValidateDatabaseName(databasename string) bool {
	isValid, _ := regexp.MatchString("[a-zA-Z]{3,40}", databasename)
	return isValid
}

/**
** Check if the Column Type is Right
**/
func ValidateColumnType(ColumnType string) bool {

	if strings.EqualFold(ColumnType, "varchar") {
		return true
	}

	if strings.EqualFold(ColumnType, "int") {
		return true
	}

	if strings.EqualFold(ColumnType, "float") {
		return true
	}

	if strings.EqualFold(ColumnType, "TimeStamp") {
		return true
	}

	if strings.EqualFold(ColumnType, "Text") {
		return true
	}
	return false
}

/**
** Check if there's is a Primary Key
**/
func ValidatePrimaryKey(PrimaryKey string) bool {
	if strings.EqualFold(PrimaryKey, "PrimaryKey") {
		return true
	}
	return false
}

/**
** Check the Reference Keyword
**/
func ValidateReference(Reference string) bool {
	if strings.EqualFold(Reference, "Reference") {
		return true
	}
	return false
}

/**
** Check the Drop Keyword
**/
func ValidateDrop(Name string, table string) bool {
	if strings.EqualFold(Name, "drop") {
		if CheckTable(table) {
			return true
		}
	}
	return false
}
