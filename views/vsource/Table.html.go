// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/Table.html:1
package vsource

//line views/vsource/Table.html:1
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/views/components"
)

//line views/vsource/Table.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/Table.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/Table.html:8
func StreamTable(qw422016 *qt422016.Writer, sources source.Sources, as *app.State, ps *cutil.PageState) {
//line views/vsource/Table.html:8
	qw422016.N().S(`    <table>
      <thead>
        <tr>
          <th style="width: 30%;">Name</th>
          <th>Description</th>
          <th>Type</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vsource/Table.html:19
	for _, s := range sources {
//line views/vsource/Table.html:19
		qw422016.N().S(`        <tr>
          <td>
            <a href="/source/`)
//line views/vsource/Table.html:22
		qw422016.E().S(s.Key)
//line views/vsource/Table.html:22
		qw422016.N().S(`">`)
//line views/vsource/Table.html:22
		components.StreamSVGRef(qw422016, s.IconWithFallback(), 16, 16, "icon", ps)
//line views/vsource/Table.html:22
		qw422016.E().S(s.Name())
//line views/vsource/Table.html:22
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vsource/Table.html:24
		qw422016.E().S(s.Description)
//line views/vsource/Table.html:24
		qw422016.N().S(`</td>
          <td>`)
//line views/vsource/Table.html:25
		qw422016.E().S(s.Type.String())
//line views/vsource/Table.html:25
		qw422016.N().S(`</td>
          <td class="shrink">
            <a href="/s/`)
//line views/vsource/Table.html:27
		qw422016.E().S(s.Key)
//line views/vsource/Table.html:27
		qw422016.N().S(`"><button type="button">workspace</button></a>
          </td>
        </tr>
`)
//line views/vsource/Table.html:30
	}
//line views/vsource/Table.html:30
	qw422016.N().S(`      </tbody>
    </table>
`)
//line views/vsource/Table.html:33
}

//line views/vsource/Table.html:33
func WriteTable(qq422016 qtio422016.Writer, sources source.Sources, as *app.State, ps *cutil.PageState) {
//line views/vsource/Table.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Table.html:33
	StreamTable(qw422016, sources, as, ps)
//line views/vsource/Table.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Table.html:33
}

//line views/vsource/Table.html:33
func Table(sources source.Sources, as *app.State, ps *cutil.PageState) string {
//line views/vsource/Table.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Table.html:33
	WriteTable(qb422016, sources, as, ps)
//line views/vsource/Table.html:33
	qs422016 := string(qb422016.B)
//line views/vsource/Table.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Table.html:33
	return qs422016
//line views/vsource/Table.html:33
}
