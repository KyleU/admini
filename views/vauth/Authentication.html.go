// Code generated by qtc from "Authentication.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vauth/Authentication.html:2
package vauth

//line views/vauth/Authentication.html:2
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/auth"
	"admini.dev/admini/views/components"
)

//line views/vauth/Authentication.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vauth/Authentication.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vauth/Authentication.html:9
func StreamAuthentication(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vauth/Authentication.html:9
	qw422016.N().S(`
`)
//line views/vauth/Authentication.html:11
	prvs, err := as.Auth.Providers(ps.Logger)
	if err != nil {
		panic(err)
	}

//line views/vauth/Authentication.html:15
	qw422016.N().S(`  <div class="card">
    <div class="right"><a href="#modal-available"><button type="button">Available</button></a></div>
    <h3>`)
//line views/vauth/Authentication.html:18
	components.StreamSVGRefIcon(qw422016, `user`, ps)
//line views/vauth/Authentication.html:18
	qw422016.N().S(`Authentication</h3>
`)
//line views/vauth/Authentication.html:19
	if len(prvs) == 0 {
//line views/vauth/Authentication.html:19
		qw422016.N().S(`    <em class="mt">no authentication providers configured, why not <a href="#modal-available">add one</a>?</em>
`)
//line views/vauth/Authentication.html:21
	} else {
//line views/vauth/Authentication.html:21
		qw422016.N().S(`    <ul class="mt">
`)
//line views/vauth/Authentication.html:23
		for _, prv := range prvs {
//line views/vauth/Authentication.html:23
			qw422016.N().S(`      <li><a href="/auth/`)
//line views/vauth/Authentication.html:24
			qw422016.N().U(prv.ID)
//line views/vauth/Authentication.html:24
			qw422016.N().S(`?refer=`)
//line views/vauth/Authentication.html:24
			qw422016.N().U(`/admin`)
//line views/vauth/Authentication.html:24
			qw422016.N().S(`">`)
//line views/vauth/Authentication.html:24
			qw422016.E().S(auth.AvailableProviderNames[prv.ID])
//line views/vauth/Authentication.html:24
			qw422016.N().S(`</a></li>
`)
//line views/vauth/Authentication.html:25
		}
//line views/vauth/Authentication.html:25
		qw422016.N().S(`    </ul>
`)
//line views/vauth/Authentication.html:27
	}
//line views/vauth/Authentication.html:27
	qw422016.N().S(`  </div>

  <div id="modal-available" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Available Authentication Providers</h2>
      </div>
      <div class="modal-body">
        <ul>
`)
//line views/vauth/Authentication.html:39
	for _, x := range auth.AvailableProviderKeys {
//line views/vauth/Authentication.html:39
		qw422016.N().S(`          <li title="`)
//line views/vauth/Authentication.html:40
		qw422016.E().S(auth.ProviderUsage(x, prvs.Contains(x)))
//line views/vauth/Authentication.html:40
		qw422016.N().S(`">
`)
//line views/vauth/Authentication.html:41
		if prvs.Contains(x) {
//line views/vauth/Authentication.html:41
			qw422016.N().S(`            `)
//line views/vauth/Authentication.html:42
			qw422016.E().S(auth.AvailableProviderNames[x])
//line views/vauth/Authentication.html:42
			qw422016.N().S(`
`)
//line views/vauth/Authentication.html:43
		} else {
//line views/vauth/Authentication.html:43
			qw422016.N().S(`            <em>`)
//line views/vauth/Authentication.html:44
			qw422016.E().S(auth.AvailableProviderNames[x])
//line views/vauth/Authentication.html:44
			qw422016.N().S(`</em>
`)
//line views/vauth/Authentication.html:45
		}
//line views/vauth/Authentication.html:45
		qw422016.N().S(`          </li>
`)
//line views/vauth/Authentication.html:47
	}
//line views/vauth/Authentication.html:47
	qw422016.N().S(`        </ul>
      </div>
    </div>
  </div>
`)
//line views/vauth/Authentication.html:52
}

//line views/vauth/Authentication.html:52
func WriteAuthentication(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vauth/Authentication.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vauth/Authentication.html:52
	StreamAuthentication(qw422016, as, ps)
//line views/vauth/Authentication.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/vauth/Authentication.html:52
}

//line views/vauth/Authentication.html:52
func Authentication(as *app.State, ps *cutil.PageState) string {
//line views/vauth/Authentication.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vauth/Authentication.html:52
	WriteAuthentication(qb422016, as, ps)
//line views/vauth/Authentication.html:52
	qs422016 := string(qb422016.B)
//line views/vauth/Authentication.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vauth/Authentication.html:52
	return qs422016
//line views/vauth/Authentication.html:52
}
