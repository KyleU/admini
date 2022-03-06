// Code generated by qtc from "About.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/About.html:2
package views

//line views/About.html:2
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/components"
	"github.com/kyleu/admini/views/layout"
)

//line views/About.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/About.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/About.html:10
type About struct{ layout.Basic }

//line views/About.html:12
func (p *About) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:12
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/About.html:14
	components.StreamSVGRefIcon(qw422016, `app`, ps)
//line views/About.html:14
	qw422016.E().S(util.AppName)
//line views/About.html:14
	qw422016.N().S(`</h3>
    <em>v`)
//line views/About.html:15
	qw422016.E().S(as.BuildInfo.Version)
//line views/About.html:15
	qw422016.N().S(`, started `)
//line views/About.html:15
	qw422016.E().S(util.TimeRelative(&as.Started))
//line views/About.html:15
	qw422016.N().S(`</em>
  </div>
  <div class="card">
    <h3>Help</h3>
`)
//line views/About.html:19
	qw422016.N().S(`    <p>Coming soon...</p>
`)
//line views/About.html:21
	qw422016.N().S(`  </div>
  `)
//line views/About.html:23
	StreamFeedback(qw422016)
//line views/About.html:23
	qw422016.N().S(`
`)
//line views/About.html:24
}

//line views/About.html:24
func (p *About) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:24
	p.StreamBody(qw422016, as, ps)
//line views/About.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:24
}

//line views/About.html:24
func (p *About) Body(as *app.State, ps *cutil.PageState) string {
//line views/About.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:24
	p.WriteBody(qb422016, as, ps)
//line views/About.html:24
	qs422016 := string(qb422016.B)
//line views/About.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:24
	return qs422016
//line views/About.html:24
}

//line views/About.html:26
func StreamFeedback(qw422016 *qt422016.Writer) {
//line views/About.html:26
	qw422016.N().S(`
<div class="card">
  <h3>Feedback</h3>
  <p>For now, email <a href="mailto:kyle@kyleu.com">Kyle U</a></p>
</div>
`)
//line views/About.html:31
}

//line views/About.html:31
func WriteFeedback(qq422016 qtio422016.Writer) {
//line views/About.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:31
	StreamFeedback(qw422016)
//line views/About.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:31
}

//line views/About.html:31
func Feedback() string {
//line views/About.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:31
	WriteFeedback(qb422016)
//line views/About.html:31
	qs422016 := string(qb422016.B)
//line views/About.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:31
	return qs422016
//line views/About.html:31
}
