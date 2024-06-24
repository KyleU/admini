// Code generated by qtc from "Modules.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Modules.html:1
package vadmin

//line views/vadmin/Modules.html:1
import (
	"runtime/debug"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/components/view"
	"admini.dev/admini/views/layout"
)

//line views/vadmin/Modules.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Modules.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Modules.html:12
type Modules struct {
	layout.Basic
	Modules []*debug.Module
}

//line views/vadmin/Modules.html:17
func (p *Modules) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vadmin/Modules.html:19
	qw422016.E().S(util.AppName)
//line views/vadmin/Modules.html:19
	qw422016.N().S(` v`)
//line views/vadmin/Modules.html:19
	qw422016.E().S(as.BuildInfo.Version)
//line views/vadmin/Modules.html:19
	qw422016.N().S(`</div>
    <h3>`)
//line views/vadmin/Modules.html:20
	components.StreamSVGIcon(qw422016, `archive`, ps)
//line views/vadmin/Modules.html:20
	qw422016.N().S(` Go Modules</h3>
    `)
//line views/vadmin/Modules.html:21
	streammoduleTable(qw422016, p.Modules, ps)
//line views/vadmin/Modules.html:21
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/Modules.html:23
}

//line views/vadmin/Modules.html:23
func (p *Modules) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:23
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Modules.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:23
}

//line views/vadmin/Modules.html:23
func (p *Modules) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Modules.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:23
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Modules.html:23
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:23
	return qs422016
//line views/vadmin/Modules.html:23
}

//line views/vadmin/Modules.html:25
func streammoduleTable(qw422016 *qt422016.Writer, mods []*debug.Module, ps *cutil.PageState) {
//line views/vadmin/Modules.html:25
	qw422016.N().S(`
    <div class="overflow full-width">
      <table class="mt">
        <thead>
          <tr>
            <th>Name</th>
            <th>Version</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Modules.html:35
	for _, m := range mods {
//line views/vadmin/Modules.html:35
		qw422016.N().S(`          <tr>
            <td>`)
//line views/vadmin/Modules.html:37
		view.StreamURL(qw422016, "https://"+m.Path, m.Path, true, ps)
//line views/vadmin/Modules.html:37
		qw422016.N().S(`</td>
            <td title="`)
//line views/vadmin/Modules.html:38
		qw422016.E().S(m.Sum)
//line views/vadmin/Modules.html:38
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:38
		qw422016.E().S(m.Version)
//line views/vadmin/Modules.html:38
		qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Modules.html:40
	}
//line views/vadmin/Modules.html:40
	qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vadmin/Modules.html:44
}

//line views/vadmin/Modules.html:44
func writemoduleTable(qq422016 qtio422016.Writer, mods []*debug.Module, ps *cutil.PageState) {
//line views/vadmin/Modules.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:44
	streammoduleTable(qw422016, mods, ps)
//line views/vadmin/Modules.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:44
}

//line views/vadmin/Modules.html:44
func moduleTable(mods []*debug.Module, ps *cutil.PageState) string {
//line views/vadmin/Modules.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:44
	writemoduleTable(qb422016, mods, ps)
//line views/vadmin/Modules.html:44
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:44
	return qs422016
//line views/vadmin/Modules.html:44
}
