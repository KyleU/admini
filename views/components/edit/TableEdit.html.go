// Code generated by qtc from "TableEdit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/TableEdit.html:2
package edit

//line views/components/edit/TableEdit.html:2
import (
	"admini.dev/admini/app/util"
)

//line views/components/edit/TableEdit.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/TableEdit.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/TableEdit.html:6
func StreamTableEditor(qw422016 *qt422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap, action string, method string, title string) {
//line views/components/edit/TableEdit.html:6
	qw422016.N().S(`
  <form action="`)
//line views/components/edit/TableEdit.html:7
	qw422016.E().S(action)
//line views/components/edit/TableEdit.html:7
	qw422016.N().S(`" method="`)
//line views/components/edit/TableEdit.html:7
	qw422016.E().S(method)
//line views/components/edit/TableEdit.html:7
	qw422016.N().S(`">
    <table class="mt expanded">
      <tbody>
`)
//line views/components/edit/TableEdit.html:10
	for _, arg := range columns {
//line views/components/edit/TableEdit.html:11
		switch arg.Type {
//line views/components/edit/TableEdit.html:12
		case "bool":
//line views/components/edit/TableEdit.html:12
			qw422016.N().S(`        `)
//line views/components/edit/TableEdit.html:13
			StreamBoolTable(qw422016, arg.Key, arg.Title, values.GetBoolOpt(arg.Key), 3, arg.Description)
//line views/components/edit/TableEdit.html:13
			qw422016.N().S(`
`)
//line views/components/edit/TableEdit.html:14
		default:
//line views/components/edit/TableEdit.html:14
			qw422016.N().S(`        `)
//line views/components/edit/TableEdit.html:15
			StreamStringTable(qw422016, arg.Key, key+"-"+arg.Key, arg.Title, values.GetStringOpt(arg.Key), 3, arg.Description)
//line views/components/edit/TableEdit.html:15
			qw422016.N().S(`
`)
//line views/components/edit/TableEdit.html:16
		}
//line views/components/edit/TableEdit.html:17
	}
//line views/components/edit/TableEdit.html:17
	qw422016.N().S(`        <tr>
          <td colspan="2"><button type="submit">`)
//line views/components/edit/TableEdit.html:19
	qw422016.E().S(title)
//line views/components/edit/TableEdit.html:19
	qw422016.N().S(`</button></td>
        </tr>
      </tbody>
    </table>
  </form>
`)
//line views/components/edit/TableEdit.html:24
}

//line views/components/edit/TableEdit.html:24
func WriteTableEditor(qq422016 qtio422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap, action string, method string, title string) {
//line views/components/edit/TableEdit.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/TableEdit.html:24
	StreamTableEditor(qw422016, key, columns, values, action, method, title)
//line views/components/edit/TableEdit.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/TableEdit.html:24
}

//line views/components/edit/TableEdit.html:24
func TableEditor(key string, columns []*util.FieldDesc, values util.ValueMap, action string, method string, title string) string {
//line views/components/edit/TableEdit.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/TableEdit.html:24
	WriteTableEditor(qb422016, key, columns, values, action, method, title)
//line views/components/edit/TableEdit.html:24
	qs422016 := string(qb422016.B)
//line views/components/edit/TableEdit.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/TableEdit.html:24
	return qs422016
//line views/components/edit/TableEdit.html:24
}
