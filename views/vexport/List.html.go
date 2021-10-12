// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vexport/List.html:1
package vexport

//line views/vexport/List.html:1
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/export"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/layout"
)

//line views/vexport/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vexport/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vexport/List.html:10
type List struct {
	layout.Basic
	Act   *action.Action
	Req   *cutil.WorkspaceRequest
	Model *model.Model
}

//line views/vexport/List.html:17
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexport/List.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vexport/List.html:19
	qw422016.E().S(ps.Title)
//line views/vexport/List.html:19
	qw422016.N().S(`</h3>
    <ul>
`)
//line views/vexport/List.html:21
	for _, f := range export.AllFormats {
//line views/vexport/List.html:21
		qw422016.N().S(`      <li><a href="`)
//line views/vexport/List.html:22
		qw422016.E().S(p.Req.RouteAct(p.Act, 0, f.Language, f.Flavor))
//line views/vexport/List.html:22
		qw422016.N().S(`">`)
//line views/vexport/List.html:22
		qw422016.E().S(f.Language)
//line views/vexport/List.html:22
		qw422016.N().S(` (`)
//line views/vexport/List.html:22
		qw422016.E().S(f.Flavor)
//line views/vexport/List.html:22
		qw422016.N().S(`)</a></li>
`)
//line views/vexport/List.html:23
	}
//line views/vexport/List.html:23
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vexport/List.html:26
}

//line views/vexport/List.html:26
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexport/List.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexport/List.html:26
	p.StreamBody(qw422016, as, ps)
//line views/vexport/List.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/vexport/List.html:26
}

//line views/vexport/List.html:26
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vexport/List.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexport/List.html:26
	p.WriteBody(qb422016, as, ps)
//line views/vexport/List.html:26
	qs422016 := string(qb422016.B)
//line views/vexport/List.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexport/List.html:26
	return qs422016
//line views/vexport/List.html:26
}
