// Code generated by qtc from "String.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/fieldview/String.html:1
package fieldview

//line views/components/fieldview/String.html:1
import (
	"strings"
)

//line views/components/fieldview/String.html:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/String.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/String.html:5
func StreamString(qw422016 *qt422016.Writer, v string) {
//line views/components/fieldview/String.html:6
	if len(v) == 0 {
//line views/components/fieldview/String.html:6
		qw422016.N().S(`<em>empty string</em>`)
//line views/components/fieldview/String.html:8
	} else if strings.Contains(v, "\n") {
//line views/components/fieldview/String.html:8
		qw422016.N().S(`<div class="pre">`)
//line views/components/fieldview/String.html:9
		qw422016.E().S(strings.TrimSpace(v))
//line views/components/fieldview/String.html:9
		qw422016.N().S(`</div>`)
//line views/components/fieldview/String.html:10
	} else {
//line views/components/fieldview/String.html:11
		qw422016.E().S(v)
//line views/components/fieldview/String.html:12
	}
//line views/components/fieldview/String.html:13
}

//line views/components/fieldview/String.html:13
func WriteString(qq422016 qtio422016.Writer, v string) {
//line views/components/fieldview/String.html:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/String.html:13
	StreamString(qw422016, v)
//line views/components/fieldview/String.html:13
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/String.html:13
}

//line views/components/fieldview/String.html:13
func String(v string) string {
//line views/components/fieldview/String.html:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/String.html:13
	WriteString(qb422016, v)
//line views/components/fieldview/String.html:13
	qs422016 := string(qb422016.B)
//line views/components/fieldview/String.html:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/String.html:13
	return qs422016
//line views/components/fieldview/String.html:13
}
