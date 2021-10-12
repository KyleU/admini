// Code generated by qtc from "New.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/New.html:1
package vproject

//line views/vproject/New.html:1
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/views/components"
	"github.com/kyleu/admini/views/layout"
)

//line views/vproject/New.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/New.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/New.html:10
type New struct {
	layout.Basic
	Project          *project.Project
	AvailableSources source.Sources
}

//line views/vproject/New.html:16
func (p *New) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/New.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>New Project</h3>
    <form class="mt" action="/project" method="post" enctype="application/x-www-form-urlencoded">
      <table>
        <tbody>
          `)
//line views/vproject/New.html:22
	components.StreamTableInput(qw422016, "key", "Key", "", 5)
//line views/vproject/New.html:22
	qw422016.N().S(`
          `)
//line views/vproject/New.html:23
	components.StreamTableInput(qw422016, "title", "Title", "", 5)
//line views/vproject/New.html:23
	qw422016.N().S(`
          `)
//line views/vproject/New.html:24
	components.StreamTableIcons(qw422016, "icon", "Icon", "", ps, 5)
//line views/vproject/New.html:24
	qw422016.N().S(`
          `)
//line views/vproject/New.html:25
	components.StreamTableInput(qw422016, "description", "Description", "", 5)
//line views/vproject/New.html:25
	qw422016.N().S(`
          `)
//line views/vproject/New.html:26
	components.StreamTableCheckbox(qw422016, "sources", "Sources", nil, p.AvailableSources.Keys(), p.AvailableSources.Names(), 5)
//line views/vproject/New.html:26
	qw422016.N().S(`
        </tbody>
      </table>
      <div class="mt">
        <button type="submit">Save Changes</button>
        <a href="/project"><button type="button">Cancel</button></a>
      </div>
    </form>
  </div>
`)
//line views/vproject/New.html:35
}

//line views/vproject/New.html:35
func (p *New) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/New.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/New.html:35
	p.StreamBody(qw422016, as, ps)
//line views/vproject/New.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/New.html:35
}

//line views/vproject/New.html:35
func (p *New) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/New.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/New.html:35
	p.WriteBody(qb422016, as, ps)
//line views/vproject/New.html:35
	qs422016 := string(qb422016.B)
//line views/vproject/New.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/New.html:35
	return qs422016
//line views/vproject/New.html:35
}
