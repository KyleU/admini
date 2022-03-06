// Code generated by qtc from "WSTable.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vresult/WSTable.html:1
package vresult

//line views/vresult/WSTable.html:1
import (
	"fmt"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/filter"
	"github.com/kyleu/admini/app/lib/schema/field"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/views/components"
	"github.com/kyleu/admini/views/vutil"
)

//line views/vresult/WSTable.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vresult/WSTable.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vresult/WSTable.html:13
func StreamWSTable(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, locs map[string]cutil.Locations, indent int, showNum bool, opts *filter.Options) {
//line views/vresult/WSTable.html:14
	vutil.StreamIndent(qw422016, true, indent)
//line views/vresult/WSTable.html:14
	qw422016.N().S(`<table class="result-table">`)
//line views/vresult/WSTable.html:16
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSTable.html:16
	qw422016.N().S(`<thead>`)
//line views/vresult/WSTable.html:18
	vutil.StreamIndent(qw422016, true, indent+2)
//line views/vresult/WSTable.html:18
	qw422016.N().S(`<tr>`)
//line views/vresult/WSTable.html:20
	if showNum {
//line views/vresult/WSTable.html:21
		vutil.StreamIndent(qw422016, true, indent+3)
//line views/vresult/WSTable.html:21
		qw422016.N().S(`<th class="no-padding"><div class="resize"></div></th>`)
//line views/vresult/WSTable.html:23
	}
//line views/vresult/WSTable.html:24
	for fIdx, f := range r.Fields {
//line views/vresult/WSTable.html:25
		if m != nil && m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSTable.html:26
			streamtableHeader(qw422016, ws, r, m, fIdx, f, locs[f.Key], true, indent+3, opts.Params)
//line views/vresult/WSTable.html:27
		}
//line views/vresult/WSTable.html:28
	}
//line views/vresult/WSTable.html:29
	for fIdx, f := range r.Fields {
//line views/vresult/WSTable.html:30
		if m == nil || !m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSTable.html:31
			streamtableHeader(qw422016, ws, r, m, fIdx, f, locs[f.Key], false, indent+3, opts.Params)
//line views/vresult/WSTable.html:32
		}
//line views/vresult/WSTable.html:33
	}
//line views/vresult/WSTable.html:34
	vutil.StreamIndent(qw422016, true, indent+3)
//line views/vresult/WSTable.html:34
	qw422016.N().S(`<th class="tfill"></th>`)
//line views/vresult/WSTable.html:36
	vutil.StreamIndent(qw422016, true, indent+2)
//line views/vresult/WSTable.html:36
	qw422016.N().S(`</tr>`)
//line views/vresult/WSTable.html:38
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSTable.html:38
	qw422016.N().S(`</thead>`)
//line views/vresult/WSTable.html:40
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSTable.html:40
	qw422016.N().S(`<tbody>`)
//line views/vresult/WSTable.html:42
	for rIdx, row := range r.Data {
//line views/vresult/WSTable.html:43
		StreamWSRow(qw422016, ws, act, rIdx, row, r.Fields, m, indent+2, showNum)
//line views/vresult/WSTable.html:44
	}
//line views/vresult/WSTable.html:46
	if opts.Params.HasNextPage(r.Count) || opts.Params.HasPreviousPage() {
//line views/vresult/WSTable.html:47
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/vresult/WSTable.html:47
		qw422016.N().S(`<tr><td colspan="`)
//line views/vresult/WSTable.html:48
		qw422016.N().D(len(r.Fields) + 1)
//line views/vresult/WSTable.html:48
		qw422016.N().S(`">`)
//line views/vresult/WSTable.html:48
		components.StreamPagination(qw422016, r.Count, opts.Params, ws.PS.URI)
//line views/vresult/WSTable.html:48
		qw422016.N().S(`</td></tr>`)
//line views/vresult/WSTable.html:49
	}
//line views/vresult/WSTable.html:50
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSTable.html:50
	qw422016.N().S(`</tbody>`)
//line views/vresult/WSTable.html:52
	vutil.StreamIndent(qw422016, true, indent)
//line views/vresult/WSTable.html:52
	qw422016.N().S(`</table>`)
//line views/vresult/WSTable.html:54
}

//line views/vresult/WSTable.html:54
func WriteWSTable(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, locs map[string]cutil.Locations, indent int, showNum bool, opts *filter.Options) {
//line views/vresult/WSTable.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSTable.html:54
	StreamWSTable(qw422016, ws, act, r, m, locs, indent, showNum, opts)
//line views/vresult/WSTable.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSTable.html:54
}

//line views/vresult/WSTable.html:54
func WSTable(ws *cutil.WorkspaceRequest, act *action.Action, r *result.Result, m *model.Model, locs map[string]cutil.Locations, indent int, showNum bool, opts *filter.Options) string {
//line views/vresult/WSTable.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSTable.html:54
	WriteWSTable(qb422016, ws, act, r, m, locs, indent, showNum, opts)
//line views/vresult/WSTable.html:54
	qs422016 := string(qb422016.B)
//line views/vresult/WSTable.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSTable.html:54
	return qs422016
//line views/vresult/WSTable.html:54
}

//line views/vresult/WSTable.html:56
func streamtableHeader(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, r *result.Result, m *model.Model, fIdx int, f *field.Field, locs cutil.Locations, pk bool, indent int, params *filter.Params) {
//line views/vresult/WSTable.html:57
	vutil.StreamIndent(qw422016, true, indent)
//line views/vresult/WSTable.html:59
	tooltip := fmt.Sprintf(`%s: ordinal %d (%s)`, f.Key, fIdx, f.Type)
	srt := m != nil && f.Type.Sortable()
	cls := ""
	icon := ""
	if pk {
		cls = "pkcol"
		icon = "lock"
	}

//line views/vresult/WSTable.html:68
	components.StreamTableHeader(qw422016, "x", f.Key, f.Name(), params, icon, ws.PS.URI, tooltip, srt, cls, true, ws.PS)
//line views/vresult/WSTable.html:69
}

//line views/vresult/WSTable.html:69
func writetableHeader(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, r *result.Result, m *model.Model, fIdx int, f *field.Field, locs cutil.Locations, pk bool, indent int, params *filter.Params) {
//line views/vresult/WSTable.html:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSTable.html:69
	streamtableHeader(qw422016, ws, r, m, fIdx, f, locs, pk, indent, params)
//line views/vresult/WSTable.html:69
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSTable.html:69
}

//line views/vresult/WSTable.html:69
func tableHeader(ws *cutil.WorkspaceRequest, r *result.Result, m *model.Model, fIdx int, f *field.Field, locs cutil.Locations, pk bool, indent int, params *filter.Params) string {
//line views/vresult/WSTable.html:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSTable.html:69
	writetableHeader(qb422016, ws, r, m, fIdx, f, locs, pk, indent, params)
//line views/vresult/WSTable.html:69
	qs422016 := string(qb422016.B)
//line views/vresult/WSTable.html:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSTable.html:69
	return qs422016
//line views/vresult/WSTable.html:69
}
