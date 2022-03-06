// Code generated by qtc from "Page.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->
// This is the base page interface. All pages implement this interface.
//

//line views/layout/Page.html:4
package layout

//line views/layout/Page.html:4
import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
)

//line views/layout/Page.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Page.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Page.html:10
type Page interface {
//line views/layout/Page.html:10
	Head(as *app.State, ps *cutil.PageState) string
//line views/layout/Page.html:10
	StreamHead(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:10
	WriteHead(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:10
	Nav(as *app.State, ps *cutil.PageState) string
//line views/layout/Page.html:10
	StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:10
	WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:10
	Menu(ps *cutil.PageState) string
//line views/layout/Page.html:10
	StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState)
//line views/layout/Page.html:10
	WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState)
//line views/layout/Page.html:10
	Body(as *app.State, ps *cutil.PageState) string
//line views/layout/Page.html:10
	StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:10
	WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:10
}
