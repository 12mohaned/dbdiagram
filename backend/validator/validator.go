package validator

import (
	"regexp"
	"strings"
)

/**
** Validate the Table Information
**/
func ValidateTableName(TableName string) bool {
	isValid, _ := regexp.MatchString("[a-zA-Z]{3,}", TableName)
	return isValid
}

/**
** Check if the ColumnType is Right
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
