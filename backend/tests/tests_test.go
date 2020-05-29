package tests

import (
	"backend/validator"
	"testing"
)

//Validator Tests

/**
Testing CheckCreate() Method in a Test Driven Data Way
**/
func TestCheckCreate(test *testing.T) {
	var tests = []struct {
		command string
		output  bool
	}{
		{"Create", true},
		{"CReAte", true},
		{"cReAte", true},
		{"crEaTE", true},
		{"Crea", false},
		{"CREATE", true},
		{"TanyaRomanov", false},
		{"MT", false},
		{"create", true},
		{"Reference", false},
		{"Table", false},
	}

	for _, tt := range tests {
		testname := tt.command
		test.Run(testname, func(test *testing.T) {
			answ := validator.CheckCreate(tt.command)
			if answ != tt.output {
				test.Error("Got", answ, "Wanted", tt.output)
			}
		})
	}
}

/**
Testing ValidateTableName() Method in a Test Driven Data Way
**/
func TestValidateTableName(test *testing.T) {
	var tests = []struct {
		TableName string
		output    bool
	}{
		{"Employee", true},
		{"Table", true},
		{"People", true},
		{"zo", false},
		{"zoo", true},
		{"a", false},
		{"Animals", true},
		{"2", false},
		{"1", false},
		{"Universities", true},
	}
	for _, tt := range tests {
		testname := tt.TableName
		test.Run(testname, func(test *testing.T) {
			answ := validator.ValidateTableName(tt.TableName)
			if answ != tt.output {
				test.Error("Got", answ, "Wanted", tt.output)
			}
		})
	}
}

/**
Testing ValidateTableName() Method in a Test Driven Data Way
**/
func TestCheckTable(test *testing.T) {
	var tests = []struct {
		TableName string
		output    bool
	}{
		{"Table", true},
		{"table", true},
		{"TaBle", true},
		{"zo", false},
		{"TaBLE", true},
		{"a", false},
		{"Animals", false},
		{"2", false},
		{"1", false},
		{"tAbLe", true},
	}
	for _, tt := range tests {
		testname := tt.TableName
		test.Run(testname, func(test *testing.T) {
			answ := validator.CheckTable(tt.TableName)
			if answ != tt.output {
				test.Error("Got", answ, "Wanted", tt.output)
			}
		})
	}
}

/**
Testing ValidateColumnType() Method in a Test Driven Data Way
**/
func TestValidateColumnType(test *testing.T) {
	var tests = []struct {
		ColumnType string
		output     bool
	}{
		{"varchar", true},
		{"int", true},
		{"float", true},
		{"TimeStamp", true},
		{"Text", true},
		{"double", false},
		{"VARCHAR", true},
		{"INT", true},
		{"FloAT", true},
		{"TimeSTAMP", true},
		{"Time", false},
		{"DOUBLE", false},
	}
	for _, tt := range tests {
		testname := tt.ColumnType
		test.Run(testname, func(test *testing.T) {
			answ := validator.ValidateColumnType(tt.ColumnType)
			if answ != tt.output {
				test.Error("Got", answ, "Wanted", tt.output)
			}
		})
	}
}

/**
Testing ValidateReference() Method in a Test Driven Data Way
**/
func TestValidateReference(test *testing.T) {
	var tests = []struct {
		RefName string
		output  bool
	}{
		{"Reference", true},
		{"ReFerence", true},
		{"RefErence", true},
		{"zo", false},
		{"reference", true},
		{"a", false},
		{"Animals", false},
		{"2", false},
		{"1", false},
		{"REFERENCE", true},
	}
	for _, tt := range tests {
		testname := tt.RefName
		test.Run(testname, func(test *testing.T) {
			answ := validator.ValidateReference(tt.RefName)
			if answ != tt.output {
				test.Error("Got", answ, "Wanted", tt.output)
			}
		})
	}
}

//Testing ValidatePrimaryKey
/**
Testing ValidatePrimaryKey() Method in a Test Driven Data Way
**/
func TestValidatePrimaryKey(test *testing.T) {
	var tests = []struct {
		PrimaryKey string
		output     bool
	}{
		{"PrimaryKey", true},
		{"primarykey", true},
		{"RefErence", false},
		{"primarykey", true},
		{"reference", false},
		{"a", false},
		{"Animals", false},
		{"2", false},
		{"1", false},
		{"productname", false},
	}
	for _, tt := range tests {
		testname := tt.PrimaryKey
		test.Run(testname, func(test *testing.T) {
			answ := validator.ValidatePrimaryKey(tt.PrimaryKey)
			if answ != tt.output {
				test.Error("Got", answ, "Wanted", tt.output)
			}
		})
	}
}
