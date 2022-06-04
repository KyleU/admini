// Code generated by qtc from "ModelFieldList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/ModelFieldList.html:1
package vsource

//line views/vsource/ModelFieldList.html:1
import (
	"fmt"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/components/fieldview"
)

//line views/vsource/ModelFieldList.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/ModelFieldList.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/ModelFieldList.html:12
func StreamModelFieldList(qw422016 *qt422016.Writer, m *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vsource/ModelFieldList.html:12
	if len(m.Fields) > 0 {
//line views/vsource/ModelFieldList.html:12
		qw422016.N().S(`    <div class="card">
      <h3>`)
//line views/vsource/ModelFieldList.html:14
		qw422016.E().S(util.StringPlural(len(m.Fields), `field`))
//line views/vsource/ModelFieldList.html:14
		qw422016.N().S(`</h3>
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Title</th>
            <th>Plural</th>
            <th>Type</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vsource/ModelFieldList.html:25
		for _, f := range m.Fields {
//line views/vsource/ModelFieldList.html:26
			sing, plur := util.StringForms(util.StringToTitle(f.Key))

//line views/vsource/ModelFieldList.html:26
			qw422016.N().S(`          <tr>
            <td>`)
//line views/vsource/ModelFieldList.html:28
			fieldview.StreamString(qw422016, f.Key)
//line views/vsource/ModelFieldList.html:28
			qw422016.N().S(`</td>
            <td>
              <div class="toggle">
                <input id="toggle-`)
//line views/vsource/ModelFieldList.html:31
			qw422016.E().S(f.Key)
//line views/vsource/ModelFieldList.html:31
			qw422016.N().S(`-title" type="checkbox" hidden />
                <label for="toggle-`)
//line views/vsource/ModelFieldList.html:32
			qw422016.E().S(f.Key)
//line views/vsource/ModelFieldList.html:32
			qw422016.N().S(`-title">
`)
//line views/vsource/ModelFieldList.html:33
			if f.Name() == sing {
//line views/vsource/ModelFieldList.html:33
				qw422016.N().S(`                  <em>`)
//line views/vsource/ModelFieldList.html:34
				qw422016.E().S(f.Name())
//line views/vsource/ModelFieldList.html:34
				qw422016.N().S(`</em>
`)
//line views/vsource/ModelFieldList.html:35
			} else {
//line views/vsource/ModelFieldList.html:35
				qw422016.N().S(`                  `)
//line views/vsource/ModelFieldList.html:36
				qw422016.E().S(f.Name())
//line views/vsource/ModelFieldList.html:36
				qw422016.N().S(`
`)
//line views/vsource/ModelFieldList.html:37
			}
//line views/vsource/ModelFieldList.html:37
			qw422016.N().S(`                </label>
                <div class="x">
                  `)
//line views/vsource/ModelFieldList.html:40
			components.StreamFormInput(qw422016, fmt.Sprintf("f.%s.title", f.Key), f.Key+"-title", f.Name())
//line views/vsource/ModelFieldList.html:40
			qw422016.N().S(`
                </div>
              </div>
            </td>
            <td>
              <div class="toggle">
                <input id="toggle-`)
//line views/vsource/ModelFieldList.html:46
			qw422016.E().S(f.Key)
//line views/vsource/ModelFieldList.html:46
			qw422016.N().S(`-plural" type="checkbox" hidden />
                <label for="toggle-`)
//line views/vsource/ModelFieldList.html:47
			qw422016.E().S(f.Key)
//line views/vsource/ModelFieldList.html:47
			qw422016.N().S(`-plural">
`)
//line views/vsource/ModelFieldList.html:48
			if f.PluralName() == plur {
//line views/vsource/ModelFieldList.html:48
				qw422016.N().S(`                  <em>`)
//line views/vsource/ModelFieldList.html:49
				qw422016.E().S(f.PluralName())
//line views/vsource/ModelFieldList.html:49
				qw422016.N().S(`</em>
`)
//line views/vsource/ModelFieldList.html:50
			} else {
//line views/vsource/ModelFieldList.html:50
				qw422016.N().S(`                  `)
//line views/vsource/ModelFieldList.html:51
				qw422016.E().S(f.PluralName())
//line views/vsource/ModelFieldList.html:51
				qw422016.N().S(`
`)
//line views/vsource/ModelFieldList.html:52
			}
//line views/vsource/ModelFieldList.html:52
			qw422016.N().S(`                </label>
                <div class="x">
                  `)
//line views/vsource/ModelFieldList.html:55
			components.StreamFormInput(qw422016, fmt.Sprintf("f.%s.plural", f.Key), f.Key+"-plural", f.PluralName())
//line views/vsource/ModelFieldList.html:55
			qw422016.N().S(`
                </div>
              </div>
            </td>
            <td>`)
//line views/vsource/ModelFieldList.html:59
			fieldview.StreamType(qw422016, f.Type)
//line views/vsource/ModelFieldList.html:59
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vsource/ModelFieldList.html:61
		}
//line views/vsource/ModelFieldList.html:61
		qw422016.N().S(`        </tbody>
      </table>
      <div class="mt">
        <button type="submit">Save All Changes</button>
        <button type="reset">Reset</button>
      </div>
    </div>
`)
//line views/vsource/ModelFieldList.html:69
	}
//line views/vsource/ModelFieldList.html:69
}

//line views/vsource/ModelFieldList.html:69
func WriteModelFieldList(qq422016 qtio422016.Writer, m *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vsource/ModelFieldList.html:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/ModelFieldList.html:69
	StreamModelFieldList(qw422016, m, as, ps)
//line views/vsource/ModelFieldList.html:69
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/ModelFieldList.html:69
}

//line views/vsource/ModelFieldList.html:69
func ModelFieldList(m *model.Model, as *app.State, ps *cutil.PageState) string {
//line views/vsource/ModelFieldList.html:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/ModelFieldList.html:69
	WriteModelFieldList(qb422016, m, as, ps)
//line views/vsource/ModelFieldList.html:69
	qs422016 := string(qb422016.B)
//line views/vsource/ModelFieldList.html:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/ModelFieldList.html:69
	return qs422016
//line views/vsource/ModelFieldList.html:69
}
