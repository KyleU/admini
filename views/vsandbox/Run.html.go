// Code generated by qtc from "Run.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsandbox/Run.html:2
package vsandbox

//line views/vsandbox/Run.html:2
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vsandbox/Run.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsandbox/Run.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsandbox/Run.html:9
type Run struct {
	layout.Basic
	Key    string
	Title  string
	Icon   string
	Result any
}

//line views/vsandbox/Run.html:17
func (p *Run) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Run.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsandbox/Run.html:19
	components.StreamSVGIcon(qw422016, p.Icon, ps)
//line views/vsandbox/Run.html:19
	qw422016.E().S(p.Title)
//line views/vsandbox/Run.html:19
	qw422016.N().S(`</h3>
    <div class="mt">`)
//line views/vsandbox/Run.html:20
	components.StreamJSON(qw422016, p.Result)
//line views/vsandbox/Run.html:20
	qw422016.N().S(`</div>
  </div>
`)
//line views/vsandbox/Run.html:22
}

//line views/vsandbox/Run.html:22
func (p *Run) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Run.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/Run.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vsandbox/Run.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/Run.html:22
}

//line views/vsandbox/Run.html:22
func (p *Run) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/Run.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/Run.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vsandbox/Run.html:22
	qs422016 := string(qb422016.B)
//line views/vsandbox/Run.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/Run.html:22
	return qs422016
//line views/vsandbox/Run.html:22
}
