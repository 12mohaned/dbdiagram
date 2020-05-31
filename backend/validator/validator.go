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
** Validate the Table Information
**/
func ValidateTableName(TableName string) bool {
	isValid, _ := regexp.MatchString("[a-zA-Z]{3,}", TableName)
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
