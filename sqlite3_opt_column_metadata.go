// +build sqlite_column_metadata

package sqlite3

/*
#ifndef USE_LIBSQLITE3
#cgo CFLAGS: -DSQLITE_ENABLE_COLUMN_METADATA
#include <sqlite3-binding.h>
#else
#include <sqlite3.h>
#endif
*/
import "C"

// ColumnTableName returns the table that is the origin of a particular result
// column in a SELECT statement.
//
// See https://www.sqlite.org/c3ref/column_database_name.html
func (s *SQLiteStmt) ColumnTableName(n int) string {
	return C.GoString(C.sqlite3_column_table_name(s.s, C.int(n)))
}

// ColumnOriginName returns the column that is the origin of a particular result
// column in a SELECT statement.
//
// See https://www.sqlite.org/c3ref/column_database_name.html
func (s *SQLiteStmt) ColumnOriginName(n int) string {
	return C.GoString(C.sqlite3_column_origin_name(s.s, C.int(n)))
}

// ColumnDatabaseName returns the attached database name that is the origin of 
// a particular result column in a SELECT statement.
//
// See https://www.sqlite.org/c3ref/column_database_name.html
func (s *SQLiteStmt) ColumnDatabaseName(n int) string {
	return C.GoString(C.sqlite3_column_database_name(s.s, C.int(n)))
}

// ColumnTableName returns the table that is the origin of a particular result
// column in a SELECT statement.
//
// See https://www.sqlite.org/c3ref/column_database_name.html
func (r *SQLiteRows) ColumnTableName(n int) string {
	if r.s == nil {
		return ""
	}
	return r.s.ColumnTableName(n)
}

// ColumnOriginName returns the column that is the origin of a particular result
// column in a SELECT statement.
//
// See https://www.sqlite.org/c3ref/column_database_name.html
func (r *SQLiteRows) ColumnOriginName(n int) string {
	if r.s == nil {
		return ""
	}
	return r.s.ColumnOriginName(n)
}

// ColumnTableName returns the table that is the origin of a particular result
// column in a SELECT statement.
//
// See https://www.sqlite.org/c3ref/column_database_name.html
func (r *SQLiteRows) ColumnDatabaseName(n int) string {
	if r.s == nil {
		return ""
	}
	return r.s.ColumnDatabaseName(n)
}
