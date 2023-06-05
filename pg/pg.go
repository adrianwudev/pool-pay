package mypg

import (
	"encoding/json"
	"strconv"
)

type PgErr struct {
	Severity      string `json:"severity"`
	Code          string `json:"code"`
	Message       string `json:"message"`
	Detail        string `json:"detail"`
	Hint          string `json:"hint"`
	Position      int    `json:"position"`
	InternalPos   int    `json:"internal_position"`
	InternalQuery string `json:"internal_query"`
	Where         string `json:"where"`
	SchemaName    string `json:"schema_name"`
	TableName     string `json:"table_name"`
	ColumnName    string `json:"column_name"`
	DataTypeName  string `json:"data_type_name"`
	Constraint    string `json:"constraint_name"`
	File          string `json:"file"`
	Line          int    `json:"line"`
	Routine       string `json:"routine"`
}

func GetPgError(err error) PgErr {
	byteErr, _ := json.Marshal(err)
	var pgError PgErr
	json.Unmarshal((byteErr), &pgError)
	return pgError
}

func IsDuplicateKeyError(pgErr PgErr) bool {
	code, _ := strconv.Atoi(pgErr.Code)
	return code == 23505
}
