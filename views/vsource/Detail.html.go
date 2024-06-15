// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/Detail.html:1
package vsource

//line views/vsource/Detail.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/source"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vsource/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/Detail.html:12
type Detail struct {
	layout.Basic
	Source *source.Source
	Schema *schema.Schema
}

//line views/vsource/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/source/`)
//line views/vsource/Detail.html:20
	qw422016.E().S(p.Source.Key)
//line views/vsource/Detail.html:20
	qw422016.N().S(`/edit"><button type="button">Edit</button></a></div>
    <h3>`)
//line views/vsource/Detail.html:21
	components.StreamSVGRef(qw422016, p.Source.IconWithFallback(), 20, 20, `icon`, ps)
//line views/vsource/Detail.html:21
	qw422016.N().S(` `)
//line views/vsource/Detail.html:21
	qw422016.E().S(p.Source.Name())
//line views/vsource/Detail.html:21
	qw422016.N().S(`</h3>
    `)
//line views/vsource/Detail.html:22
	if p.Source.Description != "" {
//line views/vsource/Detail.html:22
		qw422016.N().S(`<em>`)
//line views/vsource/Detail.html:22
		qw422016.E().S(p.Source.Description)
//line views/vsource/Detail.html:22
		qw422016.N().S(`</em>`)
//line views/vsource/Detail.html:22
	}
//line views/vsource/Detail.html:22
	qw422016.N().S(`
    <p>
      <a href="/s/`)
//line views/vsource/Detail.html:24
	qw422016.E().S(p.Source.Key)
//line views/vsource/Detail.html:24
	qw422016.N().S(`"><button type="button">Workspace</button></a>
      <a href="/source/`)
//line views/vsource/Detail.html:25
	qw422016.E().S(p.Source.Key)
//line views/vsource/Detail.html:25
	qw422016.N().S(`/refresh"><button type="button">Refresh</button></a>
    </p>
  </div>

  <div class="card">
    <h3>Models</h3>
    <table>
      <thead>
        <tr>
          <th>Key</th>
          <th>Title</th>
          <th>Plural</th>
          <th>Path</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vsource/Detail.html:41
	if p.Schema != nil && len(p.Schema.Models) > 0 {
//line views/vsource/Detail.html:42
		for _, m := range p.Schema.Models {
//line views/vsource/Detail.html:42
			qw422016.N().S(`        `)
//line views/vsource/Detail.html:43
			streammodelRow(qw422016, p.Source.Key, m, 4, as, ps)
//line views/vsource/Detail.html:43
			qw422016.N().S(`
`)
//line views/vsource/Detail.html:44
		}
//line views/vsource/Detail.html:45
	} else {
//line views/vsource/Detail.html:45
		qw422016.N().S(`        <tr><td colspan="3"><em>no models</em></td></tr>
`)
//line views/vsource/Detail.html:47
	}
//line views/vsource/Detail.html:47
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vsource/Detail.html:51
}

//line views/vsource/Detail.html:51
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/Detail.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Detail.html:51
	p.StreamBody(qw422016, as, ps)
//line views/vsource/Detail.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Detail.html:51
}

//line views/vsource/Detail.html:51
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsource/Detail.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Detail.html:51
	p.WriteBody(qb422016, as, ps)
//line views/vsource/Detail.html:51
	qs422016 := string(qb422016.B)
//line views/vsource/Detail.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Detail.html:51
	return qs422016
//line views/vsource/Detail.html:51
}

//line views/vsource/Detail.html:53
func streammodelRow(qw422016 *qt422016.Writer, srcKey string, m *model.Model, indent int, as *app.State, ps *cutil.PageState) {
//line views/vsource/Detail.html:53
	qw422016.N().S(`<tr>`)
//line views/vsource/Detail.html:55
	components.StreamIndent(qw422016, true, indent+1)
//line views/vsource/Detail.html:55
	qw422016.N().S(`<td><a href="/source/`)
//line views/vsource/Detail.html:56
	qw422016.E().S(srcKey)
//line views/vsource/Detail.html:56
	qw422016.N().S(`/model/`)
//line views/vsource/Detail.html:56
	qw422016.E().S(m.String())
//line views/vsource/Detail.html:56
	qw422016.N().S(`">`)
//line views/vsource/Detail.html:56
	qw422016.E().S(m.Key)
//line views/vsource/Detail.html:56
	qw422016.N().S(`</a></td>`)
//line views/vsource/Detail.html:57
	components.StreamIndent(qw422016, true, indent+1)
//line views/vsource/Detail.html:59
	sing, plur := util.StringForms(util.StringToTitle(m.Key))

//line views/vsource/Detail.html:59
	qw422016.N().S(`<td>`)
//line views/vsource/Detail.html:61
	if m.Name() == sing {
//line views/vsource/Detail.html:61
		qw422016.N().S(`<em>`)
//line views/vsource/Detail.html:62
		qw422016.E().S(m.Name())
//line views/vsource/Detail.html:62
		qw422016.N().S(`</em>`)
//line views/vsource/Detail.html:63
	} else {
//line views/vsource/Detail.html:64
		qw422016.E().S(m.Name())
//line views/vsource/Detail.html:65
	}
//line views/vsource/Detail.html:65
	qw422016.N().S(`</td>`)
//line views/vsource/Detail.html:67
	components.StreamIndent(qw422016, true, indent+1)
//line views/vsource/Detail.html:67
	qw422016.N().S(`<td>`)
//line views/vsource/Detail.html:69
	if m.PluralName() == plur {
//line views/vsource/Detail.html:69
		qw422016.N().S(`<em>`)
//line views/vsource/Detail.html:70
		qw422016.E().S(m.PluralName())
//line views/vsource/Detail.html:70
		qw422016.N().S(`</em>`)
//line views/vsource/Detail.html:71
	} else {
//line views/vsource/Detail.html:72
		qw422016.E().S(m.PluralName())
//line views/vsource/Detail.html:73
	}
//line views/vsource/Detail.html:73
	qw422016.N().S(`</td>`)
//line views/vsource/Detail.html:75
	components.StreamIndent(qw422016, true, indent+1)
//line views/vsource/Detail.html:75
	qw422016.N().S(`<td>`)
//line views/vsource/Detail.html:76
	qw422016.E().S(m.Pkg.ToPath())
//line views/vsource/Detail.html:76
	qw422016.N().S(`</td>`)
//line views/vsource/Detail.html:77
	components.StreamIndent(qw422016, true, indent)
//line views/vsource/Detail.html:77
	qw422016.N().S(`</tr>`)
//line views/vsource/Detail.html:79
}

//line views/vsource/Detail.html:79
func writemodelRow(qq422016 qtio422016.Writer, srcKey string, m *model.Model, indent int, as *app.State, ps *cutil.PageState) {
//line views/vsource/Detail.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Detail.html:79
	streammodelRow(qw422016, srcKey, m, indent, as, ps)
//line views/vsource/Detail.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Detail.html:79
}

//line views/vsource/Detail.html:79
func modelRow(srcKey string, m *model.Model, indent int, as *app.State, ps *cutil.PageState) string {
//line views/vsource/Detail.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Detail.html:79
	writemodelRow(qb422016, srcKey, m, indent, as, ps)
//line views/vsource/Detail.html:79
	qs422016 := string(qb422016.B)
//line views/vsource/Detail.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Detail.html:79
	return qs422016
//line views/vsource/Detail.html:79
}
