// +build sqlite_column_metadata

package sqlite3

import "testing"

func TestColumnMetadata(t *testing.T) {
	d := SQLiteDriver{}
	conn, err := d.Open(":memory:")
	if err != nil {
		t.Fatal("failed to get database connection:", err)
	}
	defer conn.Close()
	sqlite3conn := conn.(*SQLiteConn)

	_, err = sqlite3conn.Exec(`CREATE TABLE foo (namea string)`, nil)
	if err != nil {
		t.Fatal("Failed to create table:", err)
	}
	_, err = sqlite3conn.Exec(`CREATE TABLE bar (nameb string)`, nil)
	if err != nil {
		t.Fatal("Failed to create table:", err)
	}

	stmt, err := sqlite3conn.Prepare(`SELECT *, 1+2+3 FROM foo JOIN bar ON foo.name = bar.name`)
	if err != nil {
		t.Fatal(err)
	}

	if exp, got := "foo", stmt.(*SQLiteStmt).ColumnTableName(0); exp != got {
		t.Fatalf("Incorrect table name returned; expected: %s, got: %s", exp, got)
	}
	if exp, got := "bar", stmt.(*SQLiteStmt).ColumnTableName(1); exp != got {
		t.Fatalf("Incorrect table name returned; expected: %s, got: %s", exp, got)
	}
	if exp, got := "", stmt.(*SQLiteStmt).ColumnTableName(2); exp != got {
		t.Fatalf("Incorrect table name returned for expression; expected: %s, got: %s", exp, got)
	}
	if exp, got := "", stmt.(*SQLiteStmt).ColumnTableName(3); exp != got {
		t.Fatalf("Incorrect table name returned for out-of-range; expected: %s, got: %s", exp, got)
	}

	
	if exp, got := "namea", stmt.(*SQLiteStmt).ColumnOriginName(0); exp != got {
		t.Fatalf("Incorrect origin name returned; expected: %s, got: %s", exp, got)
	}
	if exp, got := "nameb", stmt.(*SQLiteStmt).ColumnOriginName(1); exp != got {
		t.Fatalf("Incorrect origin name returned; expected: %s, got: %s", exp, got)
	}
	if exp, got := "", stmt.(*SQLiteStmt).ColumnOriginName(2); exp != got {
		t.Fatalf("Incorrect origin name returned for expression; expected: %s, got: %s", exp, got)
	}
	if exp, got := "", stmt.(*SQLiteStmt).ColumnOriginName(3); exp != got {
		t.Fatalf("Incorrect origin name returned for out-of-range; expected: %s, got: %s", exp, got)
	}

	if exp, got := ":memory:", stmt.(*SQLiteStmt).ColumnDatabaseName(0); exp != got {
		t.Fatalf("Incorrect database name returned; expected: %s, got: %s", exp, got)
	}
	if exp, got := ":memory:", stmt.(*SQLiteStmt).ColumnDatabaseName(1); exp != got {
		t.Fatalf("Incorrect database name returned; expected: %s, got: %s", exp, got)
	}
	if exp, got := "", stmt.(*SQLiteStmt).ColumnDatabaseName(2); exp != got {
		t.Fatalf("Incorrect database name returned for expression; expected: %s, got: %s", exp, got)
	}
	if exp, got := "", stmt.(*SQLiteStmt).ColumnDatabaseName(3); exp != got {
		t.Fatalf("Incorrect database name returned for out-of-range; expected: %s, got: %s", exp, got)
	}
}
