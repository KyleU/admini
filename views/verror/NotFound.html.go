// Code generated by qtc from "NotFound.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Code generated by Project Forge, see https://projectforge.dev for details. -->

//line views/verror/NotFound.html:2
package verror

//line views/verror/NotFound.html:2
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views/layout"
)

//line views/verror/NotFound.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/verror/NotFound.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/verror/NotFound.html:8
type NotFound struct {
	layout.Basic
	Path string
}

//line views/verror/NotFound.html:13
func (p *NotFound) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/NotFound.html:13
	qw422016.N().S(`
  <div class="card">
    <h3>404</h3>
    <em>Page not found</em>
    <p>We can't find anything at <code>`)
//line views/verror/NotFound.html:17
	qw422016.E().S(p.Path)
//line views/verror/NotFound.html:17
	qw422016.N().S(`</code></p>
    <p>Sorry about that, maybe try the <a href="/">home page</a>?</p>
  </div>
`)
//line views/verror/NotFound.html:20
}

//line views/verror/NotFound.html:20
func (p *NotFound) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/NotFound.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/NotFound.html:20
	p.StreamBody(qw422016, as, ps)
//line views/verror/NotFound.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/verror/NotFound.html:20
}

//line views/verror/NotFound.html:20
func (p *NotFound) Body(as *app.State, ps *cutil.PageState) string {
//line views/verror/NotFound.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/NotFound.html:20
	p.WriteBody(qb422016, as, ps)
//line views/verror/NotFound.html:20
	qs422016 := string(qb422016.B)
//line views/verror/NotFound.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/NotFound.html:20
	return qs422016
//line views/verror/NotFound.html:20
}
