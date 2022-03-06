// Code generated by qtc from "Menu.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Menu.html:2
package layout

//line views/layout/Menu.html:2
import (
	"strings"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/menu"
	"github.com/kyleu/admini/views/components"
	"github.com/kyleu/admini/views/vutil"
)

//line views/layout/Menu.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Menu.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Menu.html:11
func StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:11
	qw422016.N().S(`<div class="menu-container">`)
//line views/layout/Menu.html:13
	vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:13
	qw422016.N().S(`<div class="menu">`)
//line views/layout/Menu.html:15
	vutil.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:15
	qw422016.N().S(`<ul class="level-0">`)
//line views/layout/Menu.html:17
	for _, i := range ps.Menu {
//line views/layout/Menu.html:18
		StreamMenuItem(qw422016, i, []string{}, ps.Breadcrumbs, 3, ps)
//line views/layout/Menu.html:19
	}
//line views/layout/Menu.html:20
	vutil.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:20
	qw422016.N().S(`</ul>`)
//line views/layout/Menu.html:22
	vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:22
	qw422016.N().S(`</div>`)
//line views/layout/Menu.html:24
	vutil.StreamIndent(qw422016, true, 1)
//line views/layout/Menu.html:24
	qw422016.N().S(`</div>`)
//line views/layout/Menu.html:26
}

//line views/layout/Menu.html:26
func WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:26
	StreamMenu(qw422016, ps)
//line views/layout/Menu.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:26
}

//line views/layout/Menu.html:26
func Menu(ps *cutil.PageState) string {
//line views/layout/Menu.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:26
	WriteMenu(qb422016, ps)
//line views/layout/Menu.html:26
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:26
	return qs422016
//line views/layout/Menu.html:26
}

//line views/layout/Menu.html:28
func StreamMenuItem(qw422016 *qt422016.Writer, i *menu.Item, path []string, breadcrumbs cutil.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:30
	path = append(path, i.Key)
	active, final := breadcrumbs.Active(i, path)

//line views/layout/Menu.html:33
	if i.Key == "" {
//line views/layout/Menu.html:34
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:34
		qw422016.N().S(`<li class="separator"></li>`)
//line views/layout/Menu.html:36
	} else if len(i.Children) > 0 {
//line views/layout/Menu.html:37
		itemID := strings.Join(path, "--")

//line views/layout/Menu.html:38
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:39
		if active {
//line views/layout/Menu.html:39
			qw422016.N().S(`<li class="active">`)
//line views/layout/Menu.html:39
		} else {
//line views/layout/Menu.html:39
			qw422016.N().S(`<li>`)
//line views/layout/Menu.html:39
		}
//line views/layout/Menu.html:40
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:40
		qw422016.N().S(`<input id="`)
//line views/layout/Menu.html:41
		qw422016.E().S(itemID)
//line views/layout/Menu.html:41
		qw422016.N().S(`-input" type="checkbox"`)
//line views/layout/Menu.html:41
		if active {
//line views/layout/Menu.html:41
			qw422016.N().S(` `)
//line views/layout/Menu.html:41
			qw422016.N().S(`checked="checked"`)
//line views/layout/Menu.html:41
		}
//line views/layout/Menu.html:41
		qw422016.N().S(` `)
//line views/layout/Menu.html:41
		qw422016.N().S(`hidden />`)
//line views/layout/Menu.html:42
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:43
		if final {
//line views/layout/Menu.html:43
			qw422016.N().S(`<label class="final" for="`)
//line views/layout/Menu.html:43
			qw422016.E().S(itemID)
//line views/layout/Menu.html:43
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:43
			qw422016.E().S(i.Description)
//line views/layout/Menu.html:43
			qw422016.N().S(`">`)
//line views/layout/Menu.html:43
		} else {
//line views/layout/Menu.html:43
			qw422016.N().S(`<label for="`)
//line views/layout/Menu.html:43
			qw422016.E().S(itemID)
//line views/layout/Menu.html:43
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:43
			qw422016.E().S(i.Description)
//line views/layout/Menu.html:43
			qw422016.N().S(`">`)
//line views/layout/Menu.html:43
		}
//line views/layout/Menu.html:44
		if i.Route != "" {
//line views/layout/Menu.html:45
			vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:45
			qw422016.N().S(`<a class="label-link" href="`)
//line views/layout/Menu.html:46
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:46
			qw422016.N().S(`">`)
//line views/layout/Menu.html:46
			components.StreamSVGRef(qw422016, `link`, 15, 15, ``, ps)
//line views/layout/Menu.html:46
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:47
		}
//line views/layout/Menu.html:48
		components.StreamExpandCollapse(qw422016, indent+3, ps)
//line views/layout/Menu.html:49
		vutil.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:50
		if i.Icon != "" {
//line views/layout/Menu.html:51
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:52
		}
//line views/layout/Menu.html:53
		qw422016.E().S(i.Title)
//line views/layout/Menu.html:54
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:54
		qw422016.N().S(`</label>`)
//line views/layout/Menu.html:56
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:56
		qw422016.N().S(`<ul class="level-`)
//line views/layout/Menu.html:57
		qw422016.N().D(len(path))
//line views/layout/Menu.html:57
		qw422016.N().S(`">`)
//line views/layout/Menu.html:58
		for _, i := range i.Children {
//line views/layout/Menu.html:59
			StreamMenuItem(qw422016, i, path, breadcrumbs, indent+2, ps)
//line views/layout/Menu.html:60
		}
//line views/layout/Menu.html:61
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:61
		qw422016.N().S(`</ul>`)
//line views/layout/Menu.html:63
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:63
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:65
	} else {
//line views/layout/Menu.html:67
		finalClass := "item"
		if active {
			finalClass += " active"
		}
		if final {
			finalClass += " final"
		}

//line views/layout/Menu.html:75
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:75
		qw422016.N().S(`<li>`)
//line views/layout/Menu.html:77
		if i.Description == "" {
//line views/layout/Menu.html:77
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:78
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:78
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:78
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:78
			qw422016.N().S(`">`)
//line views/layout/Menu.html:79
			if i.Icon != "" {
//line views/layout/Menu.html:80
				components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:81
			}
//line views/layout/Menu.html:82
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:82
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:84
		} else {
//line views/layout/Menu.html:84
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:85
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:85
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:85
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:85
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:85
			qw422016.E().S(i.Description)
//line views/layout/Menu.html:85
			qw422016.N().S(`">`)
//line views/layout/Menu.html:86
			if i.Icon != "" {
//line views/layout/Menu.html:87
				components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:88
			}
//line views/layout/Menu.html:89
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:89
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:91
		}
//line views/layout/Menu.html:91
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:93
	}
//line views/layout/Menu.html:94
}

//line views/layout/Menu.html:94
func WriteMenuItem(qq422016 qtio422016.Writer, i *menu.Item, path []string, breadcrumbs cutil.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:94
	StreamMenuItem(qw422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:94
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:94
}

//line views/layout/Menu.html:94
func MenuItem(i *menu.Item, path []string, breadcrumbs cutil.Breadcrumbs, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:94
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:94
	WriteMenuItem(qb422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:94
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:94
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:94
	return qs422016
//line views/layout/Menu.html:94
}
