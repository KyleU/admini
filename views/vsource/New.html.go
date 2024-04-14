// Code generated by qtc from "New.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/New.html:1
package vsource

//line views/vsource/New.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/views/components/edit"
	"admini.dev/admini/views/layout"
)

//line views/vsource/New.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/New.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/New.html:9
type New struct {
	layout.Basic
	Origin schema.Origin
}

//line views/vsource/New.html:14
func (p *New) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/New.html:14
	qw422016.N().S(`
  <div class="card">
    <h3>New Source</h3>
    <form class="mt" action="/source" method="post" enctype="application/x-www-form-urlencoded">
      <table class="expanded">
        <tbody>
          `)
//line views/vsource/New.html:20
	edit.StreamStringTable(qw422016, "key", "", "Key", "", 5)
//line views/vsource/New.html:20
	qw422016.N().S(`
          `)
//line views/vsource/New.html:21
	edit.StreamStringTable(qw422016, "title", "", "Title", "", 5)
//line views/vsource/New.html:21
	qw422016.N().S(`
          `)
//line views/vsource/New.html:22
	edit.StreamIconsTable(qw422016, "icon", "Icon", "", ps, 5)
//line views/vsource/New.html:22
	qw422016.N().S(`
          `)
//line views/vsource/New.html:23
	edit.StreamTextareaTable(qw422016, "description", "", "Description", 8, "", 5)
//line views/vsource/New.html:23
	qw422016.N().S(`
          `)
//line views/vsource/New.html:24
	edit.StreamRadioTable(qw422016, "type", "Type", p.Origin.Key, schema.AllOrigins.Keys(), schema.AllOrigins.Titles(), 5)
//line views/vsource/New.html:24
	qw422016.N().S(`
        </tbody>
      </table>
      <button class="mt" type="submit">Save Changes</button>
      <a href="/source"><button type="button">Cancel</button></a>
    </form>
  </div>
`)
//line views/vsource/New.html:31
}

//line views/vsource/New.html:31
func (p *New) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/New.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/New.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vsource/New.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/New.html:31
}

//line views/vsource/New.html:31
func (p *New) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsource/New.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/New.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vsource/New.html:31
	qs422016 := string(qb422016.B)
//line views/vsource/New.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/New.html:31
	return qs422016
//line views/vsource/New.html:31
}
