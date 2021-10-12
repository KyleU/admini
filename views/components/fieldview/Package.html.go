// Code generated by qtc from "Package.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/fieldview/Package.html:1
package fieldview

//line views/components/fieldview/Package.html:1
import (
	"github.com/kyleu/admini/app/util"
)

//line views/components/fieldview/Package.html:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/Package.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/Package.html:5
func StreamPackage(qw422016 *qt422016.Writer, v util.Pkg) {
//line views/components/fieldview/Package.html:6
	for idx, x := range v {
//line views/components/fieldview/Package.html:7
		qw422016.E().S(x)
//line views/components/fieldview/Package.html:7
		if len(v) < idx {
//line views/components/fieldview/Package.html:7
			qw422016.N().S(`/`)
//line views/components/fieldview/Package.html:7
		}
//line views/components/fieldview/Package.html:8
	}
//line views/components/fieldview/Package.html:9
}

//line views/components/fieldview/Package.html:9
func WritePackage(qq422016 qtio422016.Writer, v util.Pkg) {
//line views/components/fieldview/Package.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/Package.html:9
	StreamPackage(qw422016, v)
//line views/components/fieldview/Package.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/Package.html:9
}

//line views/components/fieldview/Package.html:9
func Package(v util.Pkg) string {
//line views/components/fieldview/Package.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/Package.html:9
	WritePackage(qb422016, v)
//line views/components/fieldview/Package.html:9
	qs422016 := string(qb422016.B)
//line views/components/fieldview/Package.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/Package.html:9
	return qs422016
//line views/components/fieldview/Package.html:9
}
