// Code generated by qtc from "Download.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsite/Download.html:2
package vsite

//line views/vsite/Download.html:2
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/site/download"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vsite/Download.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/Download.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/Download.html:11
type Download struct {
	layout.Basic
	Links download.Links
}

//line views/vsite/Download.html:16
func (p *Download) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Download.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsite/Download.html:18
	components.StreamSVGRefIcon(qw422016, "download", ps)
//line views/vsite/Download.html:18
	qw422016.N().S(` Download `)
//line views/vsite/Download.html:18
	qw422016.E().S(util.AppName)
//line views/vsite/Download.html:18
	qw422016.N().S(` `)
//line views/vsite/Download.html:18
	qw422016.E().S(as.BuildInfo.Version)
//line views/vsite/Download.html:18
	qw422016.N().S(`</h3>
  </div>
  `)
//line views/vsite/Download.html:20
	streamdownloadShowLinks(qw422016, "server", "database", "Server Version", "A command line interface that can launch a web server or run commands", p.Links, as.BuildInfo.Version, ps)
//line views/vsite/Download.html:20
	qw422016.N().S(`
  `)
//line views/vsite/Download.html:21
	streamdownloadShowLinks(qw422016, "desktop", "desktop", "Desktop Version", "Standalone application using your platform's native web viewer", p.Links, as.BuildInfo.Version, ps)
//line views/vsite/Download.html:21
	qw422016.N().S(`
  `)
//line views/vsite/Download.html:22
	streamdownloadShowLinks(qw422016, "mobile", "mobile", "Mobile Version", "Libraries and apps for Mobile Platforms", p.Links, as.BuildInfo.Version, ps)
//line views/vsite/Download.html:22
	qw422016.N().S(`
`)
//line views/vsite/Download.html:23
}

//line views/vsite/Download.html:23
func (p *Download) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Download.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/Download.html:23
	p.StreamBody(qw422016, as, ps)
//line views/vsite/Download.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/Download.html:23
}

//line views/vsite/Download.html:23
func (p *Download) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/Download.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/Download.html:23
	p.WriteBody(qb422016, as, ps)
//line views/vsite/Download.html:23
	qs422016 := string(qb422016.B)
//line views/vsite/Download.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/Download.html:23
	return qs422016
//line views/vsite/Download.html:23
}

//line views/vsite/Download.html:25
func streamdownloadShowLinks(qw422016 *qt422016.Writer, mode string, icon string, title string, desc string, links download.Links, v string, ps *cutil.PageState) {
//line views/vsite/Download.html:25
	qw422016.N().S(`
`)
//line views/vsite/Download.html:26
	modeLinks := links.GetByModes(mode)

//line views/vsite/Download.html:27
	if len(modeLinks) > 0 {
//line views/vsite/Download.html:27
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsite/Download.html:29
		components.StreamSVGRefIcon(qw422016, icon, ps)
//line views/vsite/Download.html:29
		qw422016.E().S(title)
//line views/vsite/Download.html:29
		qw422016.N().S(`</h3>
    <em>`)
//line views/vsite/Download.html:30
		qw422016.E().S(desc)
//line views/vsite/Download.html:30
		qw422016.N().S(`</em>
    <table class="mt">
      <tbody>
`)
//line views/vsite/Download.html:33
		var currentOS string

//line views/vsite/Download.html:34
		for _, link := range modeLinks {
//line views/vsite/Download.html:35
			if currentOS != link.OS {
//line views/vsite/Download.html:36
				if currentOS != "" {
//line views/vsite/Download.html:36
					qw422016.N().S(`          </td>
        </tr>
`)
//line views/vsite/Download.html:39
				}
//line views/vsite/Download.html:40
				currentOS = link.OS

//line views/vsite/Download.html:40
				qw422016.N().S(`        <tr>
          <td style="width: 25%;">`)
//line views/vsite/Download.html:42
				components.StreamSVGRef(qw422016, link.OSIcon(), 20, 20, "icon", ps)
//line views/vsite/Download.html:42
				qw422016.E().S(link.OSString())
//line views/vsite/Download.html:42
				qw422016.N().S(`</td>
          <td>
`)
//line views/vsite/Download.html:44
			}
//line views/vsite/Download.html:45
			if link.OS == download.OSLinux && (link.Arch == download.ArchPPC64 || link.Arch == download.ArchMIPS64Hard || link.Arch == download.ArchMIPSHard) {
//line views/vsite/Download.html:45
				qw422016.N().S(`            <div class="mt"></div>
`)
//line views/vsite/Download.html:47
			}
//line views/vsite/Download.html:47
			qw422016.N().S(`            <a href="https://github.com/kyleu/admini/releases/download/v`)
//line views/vsite/Download.html:48
			qw422016.E().S(v)
//line views/vsite/Download.html:48
			qw422016.N().S(`/`)
//line views/vsite/Download.html:48
			qw422016.E().S(link.URL)
//line views/vsite/Download.html:48
			qw422016.N().S(`"><button>`)
//line views/vsite/Download.html:48
			qw422016.E().S(link.ArchString())
//line views/vsite/Download.html:48
			qw422016.N().S(`</button></a>
`)
//line views/vsite/Download.html:49
		}
//line views/vsite/Download.html:49
		qw422016.N().S(`          </td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsite/Download.html:55
	}
//line views/vsite/Download.html:56
}

//line views/vsite/Download.html:56
func writedownloadShowLinks(qq422016 qtio422016.Writer, mode string, icon string, title string, desc string, links download.Links, v string, ps *cutil.PageState) {
//line views/vsite/Download.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/Download.html:56
	streamdownloadShowLinks(qw422016, mode, icon, title, desc, links, v, ps)
//line views/vsite/Download.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/Download.html:56
}

//line views/vsite/Download.html:56
func downloadShowLinks(mode string, icon string, title string, desc string, links download.Links, v string, ps *cutil.PageState) string {
//line views/vsite/Download.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/Download.html:56
	writedownloadShowLinks(qb422016, mode, icon, title, desc, links, v, ps)
//line views/vsite/Download.html:56
	qs422016 := string(qb422016.B)
//line views/vsite/Download.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/Download.html:56
	return qs422016
//line views/vsite/Download.html:56
}
