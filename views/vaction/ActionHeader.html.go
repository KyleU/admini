// Code generated by qtc from "ActionHeader.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/ActionHeader.html:1
package vaction

//line views/vaction/ActionHeader.html:1
import (
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/components"
)

//line views/vaction/ActionHeader.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/ActionHeader.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/ActionHeader.html:7
func StreamActionHeader(qw422016 *qt422016.Writer, act *action.Action, subtitle string, ps *cutil.PageState) {
//line views/vaction/ActionHeader.html:7
	qw422016.N().S(`
  <div class="right"><a href="#modal-action"><button type="button">Action</button></a></div>
  <h3>`)
//line views/vaction/ActionHeader.html:9
	components.StreamSVGIcon(qw422016, act.IconWithFallback(), ps)
//line views/vaction/ActionHeader.html:9
	qw422016.N().S(` `)
//line views/vaction/ActionHeader.html:9
	qw422016.E().S(act.Name())
//line views/vaction/ActionHeader.html:9
	qw422016.N().S(`</h3>`)
//line views/vaction/ActionHeader.html:9
	if subtitle != "" {
//line views/vaction/ActionHeader.html:9
		qw422016.N().S(`
  <em>`)
//line views/vaction/ActionHeader.html:10
		qw422016.E().S(subtitle)
//line views/vaction/ActionHeader.html:10
		qw422016.N().S(`</em>`)
//line views/vaction/ActionHeader.html:10
	}
//line views/vaction/ActionHeader.html:10
	qw422016.N().S(`
`)
//line views/vaction/ActionHeader.html:11
}

//line views/vaction/ActionHeader.html:11
func WriteActionHeader(qq422016 qtio422016.Writer, act *action.Action, subtitle string, ps *cutil.PageState) {
//line views/vaction/ActionHeader.html:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/ActionHeader.html:11
	StreamActionHeader(qw422016, act, subtitle, ps)
//line views/vaction/ActionHeader.html:11
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/ActionHeader.html:11
}

//line views/vaction/ActionHeader.html:11
func ActionHeader(act *action.Action, subtitle string, ps *cutil.PageState) string {
//line views/vaction/ActionHeader.html:11
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/ActionHeader.html:11
	WriteActionHeader(qb422016, act, subtitle, ps)
//line views/vaction/ActionHeader.html:11
	qs422016 := string(qb422016.B)
//line views/vaction/ActionHeader.html:11
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/ActionHeader.html:11
	return qs422016
//line views/vaction/ActionHeader.html:11
}
