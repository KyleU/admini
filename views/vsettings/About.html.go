// Code generated by qtc from "About.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Code generated by Project Forge, see https://projectforge.dev for details. -->

//line views/vsettings/About.html:2
package vsettings

//line views/vsettings/About.html:2
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/util"
)

//line views/vsettings/About.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsettings/About.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsettings/About.html:7
func StreamAbout(qw422016 *qt422016.Writer, as *app.State) {
//line views/vsettings/About.html:7
	qw422016.N().S(`
  <div class="card">
`)
//line views/vsettings/About.html:9
	if util.AppSource != "" {
//line views/vsettings/About.html:9
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vsettings/About.html:10
		qw422016.E().S(util.AppSource)
//line views/vsettings/About.html:10
		qw422016.N().S(`"><button>Github</button></a></div>
`)
//line views/vsettings/About.html:11
	}
//line views/vsettings/About.html:11
	qw422016.N().S(`    <h3 title="github:`)
//line views/vsettings/About.html:12
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vsettings/About.html:12
	qw422016.N().S(`">`)
//line views/vsettings/About.html:12
	qw422016.E().S(util.AppName)
//line views/vsettings/About.html:12
	qw422016.N().S(` `)
//line views/vsettings/About.html:12
	qw422016.E().S(as.BuildInfo.String())
//line views/vsettings/About.html:12
	qw422016.N().S(`</h3>
`)
//line views/vsettings/About.html:13
	if util.AppLegal != "" {
//line views/vsettings/About.html:13
		qw422016.N().S(`    <div>`)
//line views/vsettings/About.html:14
		qw422016.N().S(util.AppLegal)
//line views/vsettings/About.html:14
		qw422016.N().S(`</div>
`)
//line views/vsettings/About.html:15
	}
//line views/vsettings/About.html:16
	if util.AppURL != "" {
//line views/vsettings/About.html:16
		qw422016.N().S(`    <p><a href="`)
//line views/vsettings/About.html:17
		qw422016.N().S(util.AppURL)
//line views/vsettings/About.html:17
		qw422016.N().S(`">`)
//line views/vsettings/About.html:17
		qw422016.N().S(util.AppURL)
//line views/vsettings/About.html:17
		qw422016.N().S(`</a></p>
`)
//line views/vsettings/About.html:18
	}
//line views/vsettings/About.html:18
	qw422016.N().S(`  </div>
`)
//line views/vsettings/About.html:20
}

//line views/vsettings/About.html:20
func WriteAbout(qq422016 qtio422016.Writer, as *app.State) {
//line views/vsettings/About.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsettings/About.html:20
	StreamAbout(qw422016, as)
//line views/vsettings/About.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vsettings/About.html:20
}

//line views/vsettings/About.html:20
func About(as *app.State) string {
//line views/vsettings/About.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsettings/About.html:20
	WriteAbout(qb422016, as)
//line views/vsettings/About.html:20
	qs422016 := string(qb422016.B)
//line views/vsettings/About.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsettings/About.html:20
	return qs422016
//line views/vsettings/About.html:20
}
