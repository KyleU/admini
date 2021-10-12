// Code generated by qtc from "Folder.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Folder.html:1
package vaction

//line views/vaction/Folder.html:1
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views/components"
	"github.com/kyleu/admini/views/layout"
)

//line views/vaction/Folder.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Folder.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Folder.html:9
type Folder struct {
	layout.Basic
	Req *cutil.WorkspaceRequest
	Act *action.Action
}

//line views/vaction/Folder.html:15
func (p *Folder) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Folder.html:15
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-action"><button type="button">Action</button></a></div>
    <h3>`)
//line views/vaction/Folder.html:18
	components.StreamSVGRef(qw422016, p.Act.IconWithFallback(), 20, 20, "icon", ps)
//line views/vaction/Folder.html:18
	qw422016.E().S(p.Act.Name())
//line views/vaction/Folder.html:18
	qw422016.N().S(`</h3>
    <em>Folder</em>
  </div>
`)
//line views/vaction/Folder.html:21
	StreamResultChildren(qw422016, p.Req, p.Act, 1)
//line views/vaction/Folder.html:22
	StreamResultDebug(qw422016, p.Req, p.Act)
//line views/vaction/Folder.html:23
}

//line views/vaction/Folder.html:23
func (p *Folder) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Folder.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Folder.html:23
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Folder.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Folder.html:23
}

//line views/vaction/Folder.html:23
func (p *Folder) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Folder.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Folder.html:23
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Folder.html:23
	qs422016 := string(qb422016.B)
//line views/vaction/Folder.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Folder.html:23
	return qs422016
//line views/vaction/Folder.html:23
}
