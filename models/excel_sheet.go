package models

type ExcelSheet struct {
	Name        string
	InitialCell string
	Headers     []string
	Values      [][]string
}