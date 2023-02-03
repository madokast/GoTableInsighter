/*
数据源
*/

package datasource

import (
	_ "ti/internal/logger"
)

type tableName = string
type columnName = string
type columnType = string
type intCells = []int
type stringCells = []string
type doubleCells = []float64

type TableMeta struct {
	TaskId       string
	TableName    tableName
	TableDataSrc ITableDataSrc
}

type ITableDataSrc interface{}

type ContentTableDataSrc struct {
	ColumnNames []columnName
	ColumnType  []columnType
	Lines       []string
}

type LocalCSVTableDataSrc struct {
	Path string
}

type DatebaseTableDataSrc struct {
	driverName string
	// root:PoD2020@sics@tcp(192.168.13.93:3306)/testdata
	dataSourceName string
}
