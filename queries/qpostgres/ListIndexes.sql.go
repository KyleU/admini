// Code generated by qtc from "ListIndexes.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/qpostgres/ListIndexes.sql:1
package qpostgres

//line queries/qpostgres/ListIndexes.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/qpostgres/ListIndexes.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/qpostgres/ListIndexes.sql:1
func StreamListIndexes(qw422016 *qt422016.Writer, schema string) {
//line queries/qpostgres/ListIndexes.sql:1
	qw422016.N().S(`
select
  n.nspname as schema_name,
  t.relname as table_name,
  i.relname as index_name,
  case when idx.indisprimary then 1 else 0 end as pk,
  case when idx.indisunique then 1 else 0 end as u,
  array_to_string(array_agg(a.attname), ',') as column_names
from
  pg_class t,
  pg_class i,
  pg_index idx,
  pg_attribute a,
  pg_namespace n
where
  t.oid = idx.indrelid
  and i.oid = idx.indexrelid
  and a.attrelid = t.oid
  and n.oid = t.relnamespace
  and a.attnum = any(idx.indkey)
  and t.relkind = 'r'
  and n.nspname not in ('information_schema', 'pg_catalog')
  `)
//line queries/qpostgres/ListIndexes.sql:23
	if schema != "" {
//line queries/qpostgres/ListIndexes.sql:23
		qw422016.N().S(`
  and n.nspname = '`)
//line queries/qpostgres/ListIndexes.sql:24
		qw422016.E().S(schema)
//line queries/qpostgres/ListIndexes.sql:24
		qw422016.N().S(`'
  `)
//line queries/qpostgres/ListIndexes.sql:25
	}
//line queries/qpostgres/ListIndexes.sql:25
	qw422016.N().S(`
group by
  n.nspname,
  t.relname,
  i.relname,
  idx.indisprimary,
  idx.indisunique
order by
  t.relname,
  i.relname;

-- `)
//line queries/qpostgres/ListIndexes.sql:36
}

//line queries/qpostgres/ListIndexes.sql:36
func WriteListIndexes(qq422016 qtio422016.Writer, schema string) {
//line queries/qpostgres/ListIndexes.sql:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/qpostgres/ListIndexes.sql:36
	StreamListIndexes(qw422016, schema)
//line queries/qpostgres/ListIndexes.sql:36
	qt422016.ReleaseWriter(qw422016)
//line queries/qpostgres/ListIndexes.sql:36
}

//line queries/qpostgres/ListIndexes.sql:36
func ListIndexes(schema string) string {
//line queries/qpostgres/ListIndexes.sql:36
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/qpostgres/ListIndexes.sql:36
	WriteListIndexes(qb422016, schema)
//line queries/qpostgres/ListIndexes.sql:36
	qs422016 := string(qb422016.B)
//line queries/qpostgres/ListIndexes.sql:36
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/qpostgres/ListIndexes.sql:36
	return qs422016
//line queries/qpostgres/ListIndexes.sql:36
}
