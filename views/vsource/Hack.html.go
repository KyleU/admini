// Code generated by qtc from "Hack.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/Hack.html:1
package vsource

//line views/vsource/Hack.html:1
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/schema"
	"github.com/kyleu/admini/views/layout"
)

//line views/vsource/Hack.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/Hack.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/Hack.html:8
type Hack struct {
	layout.Basic
	Schema *schema.Schema
	Result string
}

//line views/vsource/Hack.html:14
func (p *Hack) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/Hack.html:14
	qw422016.N().S(`
  <div class="card">
    <h3>Source Hack!</h3>
    <pre>`)
//line views/vsource/Hack.html:17
	qw422016.E().S(p.Result)
//line views/vsource/Hack.html:17
	qw422016.N().S(`</pre>
  </div>
`)
//line views/vsource/Hack.html:19
}

//line views/vsource/Hack.html:19
func (p *Hack) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/Hack.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Hack.html:19
	p.StreamBody(qw422016, as, ps)
//line views/vsource/Hack.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Hack.html:19
}

//line views/vsource/Hack.html:19
func (p *Hack) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsource/Hack.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Hack.html:19
	p.WriteBody(qb422016, as, ps)
//line views/vsource/Hack.html:19
	qs422016 := string(qb422016.B)
//line views/vsource/Hack.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Hack.html:19
	return qs422016
//line views/vsource/Hack.html:19
}
