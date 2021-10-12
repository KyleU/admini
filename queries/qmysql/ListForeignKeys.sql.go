// Code generated by qtc from "ListForeignKeys.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/qmysql/ListForeignKeys.sql:1
package qmysql

//line queries/qmysql/ListForeignKeys.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/qmysql/ListForeignKeys.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/qmysql/ListForeignKeys.sql:1
func StreamListForeignKeys(qw422016 *qt422016.Writer, schema string) {
//line queries/qmysql/ListForeignKeys.sql:1
	qw422016.N().S(`
select
  x.constraint_name as constraint_name,
  x.ordinal_position as ordinal,
  x.table_schema as schema_name,
  x.table_name as table_name,
  x.column_name as column_name,
  x.table_schema as foreign_schema_name,
  x.referenced_table_name as foreign_table_name,
  x.referenced_column_name as foreign_column_name
from
  information_schema.key_column_usage x
where
  x.referenced_table_name is not null`)
//line queries/qmysql/ListForeignKeys.sql:14
	if schema != "" {
//line queries/qmysql/ListForeignKeys.sql:14
		qw422016.N().S(`
  and x.table_schema = '`)
//line queries/qmysql/ListForeignKeys.sql:15
		qw422016.E().S(schema)
//line queries/qmysql/ListForeignKeys.sql:15
		qw422016.N().S(`'`)
//line queries/qmysql/ListForeignKeys.sql:15
	}
//line queries/qmysql/ListForeignKeys.sql:15
	qw422016.N().S(`
order by
  x.constraint_name,
  x.ordinal_position;
-- `)
//line queries/qmysql/ListForeignKeys.sql:19
}

//line queries/qmysql/ListForeignKeys.sql:19
func WriteListForeignKeys(qq422016 qtio422016.Writer, schema string) {
//line queries/qmysql/ListForeignKeys.sql:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/qmysql/ListForeignKeys.sql:19
	StreamListForeignKeys(qw422016, schema)
//line queries/qmysql/ListForeignKeys.sql:19
	qt422016.ReleaseWriter(qw422016)
//line queries/qmysql/ListForeignKeys.sql:19
}

//line queries/qmysql/ListForeignKeys.sql:19
func ListForeignKeys(schema string) string {
//line queries/qmysql/ListForeignKeys.sql:19
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/qmysql/ListForeignKeys.sql:19
	WriteListForeignKeys(qb422016, schema)
//line queries/qmysql/ListForeignKeys.sql:19
	qs422016 := string(qb422016.B)
//line queries/qmysql/ListForeignKeys.sql:19
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/qmysql/ListForeignKeys.sql:19
	return qs422016
//line queries/qmysql/ListForeignKeys.sql:19
}
