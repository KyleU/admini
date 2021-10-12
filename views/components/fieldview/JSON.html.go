// Code generated by qtc from "JSON.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/fieldview/JSON.html:1
package fieldview

//line views/components/fieldview/JSON.html:1
import (
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
)

//line views/components/fieldview/JSON.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/JSON.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/JSON.html:6
func StreamJSON(qw422016 *qt422016.Writer, v interface{}) {
//line views/components/fieldview/JSON.html:8
	b, ok := v.([]byte)
	if ok {
		_ = util.FromJSON(b, &v)
	}

//line views/components/fieldview/JSON.html:13
	out, err := cutil.FormatJSON(v)

//line views/components/fieldview/JSON.html:14
	if err == nil {
//line views/components/fieldview/JSON.html:15
		qw422016.N().S(out)
//line views/components/fieldview/JSON.html:16
	} else {
//line views/components/fieldview/JSON.html:17
		qw422016.E().S(err.Error())
//line views/components/fieldview/JSON.html:18
	}
//line views/components/fieldview/JSON.html:19
}

//line views/components/fieldview/JSON.html:19
func WriteJSON(qq422016 qtio422016.Writer, v interface{}) {
//line views/components/fieldview/JSON.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/JSON.html:19
	StreamJSON(qw422016, v)
//line views/components/fieldview/JSON.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/JSON.html:19
}

//line views/components/fieldview/JSON.html:19
func JSON(v interface{}) string {
//line views/components/fieldview/JSON.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/JSON.html:19
	WriteJSON(qb422016, v)
//line views/components/fieldview/JSON.html:19
	qs422016 := string(qb422016.B)
//line views/components/fieldview/JSON.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/JSON.html:19
	return qs422016
//line views/components/fieldview/JSON.html:19
}

//line views/components/fieldview/JSON.html:21
func StreamJSONInline(qw422016 *qt422016.Writer, v interface{}) {
//line views/components/fieldview/JSON.html:23
	b, ok := v.([]byte)
	if ok {
		_ = util.FromJSON(b, &v)
	}

//line views/components/fieldview/JSON.html:28
	qw422016.E().S(util.ToJSON(v))
//line views/components/fieldview/JSON.html:29
}

//line views/components/fieldview/JSON.html:29
func WriteJSONInline(qq422016 qtio422016.Writer, v interface{}) {
//line views/components/fieldview/JSON.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/JSON.html:29
	StreamJSONInline(qw422016, v)
//line views/components/fieldview/JSON.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/JSON.html:29
}

//line views/components/fieldview/JSON.html:29
func JSONInline(v interface{}) string {
//line views/components/fieldview/JSON.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/JSON.html:29
	WriteJSONInline(qb422016, v)
//line views/components/fieldview/JSON.html:29
	qs422016 := string(qb422016.B)
//line views/components/fieldview/JSON.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/JSON.html:29
	return qs422016
//line views/components/fieldview/JSON.html:29
}
