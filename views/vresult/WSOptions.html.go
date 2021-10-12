// Code generated by qtc from "WSOptions.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vresult/WSOptions.html:1
package vresult

//line views/vresult/WSOptions.html:1
import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/filter"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/views/components"
	"github.com/kyleu/admini/views/vutil"
)

//line views/vresult/WSOptions.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vresult/WSOptions.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vresult/WSOptions.html:11
func StreamWSOptions(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, indent int, opts *filter.Options) {
//line views/vresult/WSOptions.html:11
	qw422016.N().S(`<div class="result-options">`)
//line views/vresult/WSOptions.html:13
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSOptions.html:13
	qw422016.N().S(`<form action="???" class="search" title="search"><input type="search" name="q" placeholder=" " /><div class="search-image"><svg><use xlink:href="#svg-searchbox" /></svg></div></form>`)
//line views/vresult/WSOptions.html:18
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSOptions.html:18
	qw422016.N().S(`<a href="#modal-sort" class="opt">`)
//line views/vresult/WSOptions.html:19
	components.StreamSVGRef(qw422016, `sort`, 16, 16, `icon`, ws.PS)
//line views/vresult/WSOptions.html:19
	qw422016.N().S(`Sort</a>`)
//line views/vresult/WSOptions.html:20
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSOptions.html:20
	qw422016.N().S(`<a href="#modal-filter" class="opt">`)
//line views/vresult/WSOptions.html:21
	components.StreamSVGRef(qw422016, `settings`, 16, 16, `icon`, ws.PS)
//line views/vresult/WSOptions.html:21
	qw422016.N().S(`Filter</a>`)
//line views/vresult/WSOptions.html:22
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSOptions.html:22
	qw422016.N().S(`<a href="#modal-group" class="opt">`)
//line views/vresult/WSOptions.html:23
	components.StreamSVGRef(qw422016, `folder`, 16, 16, `icon`, ws.PS)
//line views/vresult/WSOptions.html:23
	qw422016.N().S(`Group</a>`)
//line views/vresult/WSOptions.html:24
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSOptions.html:24
	qw422016.N().S(`<div><a href="#modal-download" class="opt">`)
//line views/vresult/WSOptions.html:25
	components.StreamSVGRef(qw422016, `download`, 16, 16, `icon`, ws.PS)
//line views/vresult/WSOptions.html:25
	qw422016.N().S(`Download</a></div>`)
//line views/vresult/WSOptions.html:26
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSOptions.html:26
	qw422016.N().S(`<div class="clear"></div>`)
//line views/vresult/WSOptions.html:28
	vutil.StreamIndent(qw422016, true, indent)
//line views/vresult/WSOptions.html:28
	qw422016.N().S(`</div>`)
//line views/vresult/WSOptions.html:30
}

//line views/vresult/WSOptions.html:30
func WriteWSOptions(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, indent int, opts *filter.Options) {
//line views/vresult/WSOptions.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSOptions.html:30
	StreamWSOptions(qw422016, ws, act, r, m, indent, opts)
//line views/vresult/WSOptions.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSOptions.html:30
}

//line views/vresult/WSOptions.html:30
func WSOptions(ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, indent int, opts *filter.Options) string {
//line views/vresult/WSOptions.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSOptions.html:30
	WriteWSOptions(qb422016, ws, act, r, m, indent, opts)
//line views/vresult/WSOptions.html:30
	qs422016 := string(qb422016.B)
//line views/vresult/WSOptions.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSOptions.html:30
	return qs422016
//line views/vresult/WSOptions.html:30
}

//line views/vresult/WSOptions.html:32
func StreamWSModals(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, indent int, opts *filter.Options) {
//line views/vresult/WSOptions.html:32
	qw422016.N().S(`<div id="modal-sort" class="modal" style="display: none;"><a class="backdrop" href="#"></a><div class="modal-content"><div class="modal-header"><a href="#" class="modal-close">×</a><h2>Sort Options</h2></div><div class="modal-body">TODO: Sorting...</div></div></div><div id="modal-filter" class="modal" style="display: none;"><a class="backdrop" href="#"></a><div class="modal-content"><div class="modal-header"><a href="#" class="modal-close">×</a><h2>Filters</h2></div><div class="modal-body">TODO: Filters...</div></div></div><div id="modal-group" class="modal" style="display: none;"><a class="backdrop" href="#"></a><div class="modal-content"><div class="modal-header"><a href="#" class="modal-close">×</a><h2>Grouping Options</h2></div><div class="modal-body">TODO: Grouping...</div></div></div><div id="modal-download" class="modal" style="display: none;"><a class="backdrop" href="#"></a><div class="modal-content"><div class="modal-header"><a href="#" class="modal-close">×</a><h2>Download</h2></div><div class="modal-body">TODO: Downloads...</div></div></div>`)
//line views/vresult/WSOptions.html:81
}

//line views/vresult/WSOptions.html:81
func WriteWSModals(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, indent int, opts *filter.Options) {
//line views/vresult/WSOptions.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSOptions.html:81
	StreamWSModals(qw422016, ws, act, r, m, indent, opts)
//line views/vresult/WSOptions.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSOptions.html:81
}

//line views/vresult/WSOptions.html:81
func WSModals(ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, indent int, opts *filter.Options) string {
//line views/vresult/WSOptions.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSOptions.html:81
	WriteWSModals(qb422016, ws, act, r, m, indent, opts)
//line views/vresult/WSOptions.html:81
	qs422016 := string(qb422016.B)
//line views/vresult/WSOptions.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSOptions.html:81
	return qs422016
//line views/vresult/WSOptions.html:81
}
