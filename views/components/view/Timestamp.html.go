// Code generated by qtc from "Timestamp.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Timestamp.html:2
package view

//line views/components/view/Timestamp.html:2
import (
	"time"

	"admini.dev/admini/app/util"
)

//line views/components/view/Timestamp.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Timestamp.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Timestamp.html:8
func StreamTimestamp(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/view/Timestamp.html:8
	qw422016.N().S(`<span class="timestamp nowrap" data-timestamp="`)
//line views/components/view/Timestamp.html:9
	qw422016.E().S(util.TimeToJS(value))
//line views/components/view/Timestamp.html:9
	qw422016.N().S(`">`)
//line views/components/view/Timestamp.html:9
	qw422016.E().S(util.TimeToFull(value))
//line views/components/view/Timestamp.html:9
	qw422016.N().S(`</span>`)
//line views/components/view/Timestamp.html:10
}

//line views/components/view/Timestamp.html:10
func WriteTimestamp(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/view/Timestamp.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Timestamp.html:10
	StreamTimestamp(qw422016, value)
//line views/components/view/Timestamp.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Timestamp.html:10
}

//line views/components/view/Timestamp.html:10
func Timestamp(value *time.Time) string {
//line views/components/view/Timestamp.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Timestamp.html:10
	WriteTimestamp(qb422016, value)
//line views/components/view/Timestamp.html:10
	qs422016 := string(qb422016.B)
//line views/components/view/Timestamp.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Timestamp.html:10
	return qs422016
//line views/components/view/Timestamp.html:10
}

//line views/components/view/Timestamp.html:12
func StreamTimestampMillis(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/view/Timestamp.html:12
	qw422016.N().S(`<span class="timestamp nowrap" data-timestamp="`)
//line views/components/view/Timestamp.html:13
	qw422016.E().S(util.TimeToRFC3339(value))
//line views/components/view/Timestamp.html:13
	qw422016.N().S(`">`)
//line views/components/view/Timestamp.html:13
	qw422016.E().S(util.TimeToFullMS(value))
//line views/components/view/Timestamp.html:13
	qw422016.N().S(`</span>`)
//line views/components/view/Timestamp.html:14
}

//line views/components/view/Timestamp.html:14
func WriteTimestampMillis(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/view/Timestamp.html:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Timestamp.html:14
	StreamTimestampMillis(qw422016, value)
//line views/components/view/Timestamp.html:14
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Timestamp.html:14
}

//line views/components/view/Timestamp.html:14
func TimestampMillis(value *time.Time) string {
//line views/components/view/Timestamp.html:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Timestamp.html:14
	WriteTimestampMillis(qb422016, value)
//line views/components/view/Timestamp.html:14
	qs422016 := string(qb422016.B)
//line views/components/view/Timestamp.html:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Timestamp.html:14
	return qs422016
//line views/components/view/Timestamp.html:14
}

//line views/components/view/Timestamp.html:16
func StreamTimestampRelative(qw422016 *qt422016.Writer, value *time.Time, static bool) {
//line views/components/view/Timestamp.html:16
	qw422016.N().S(`<span class="`)
//line views/components/view/Timestamp.html:17
	if !static {
//line views/components/view/Timestamp.html:17
		qw422016.N().S(`reltime`)
//line views/components/view/Timestamp.html:17
		qw422016.N().S(` `)
//line views/components/view/Timestamp.html:17
	}
//line views/components/view/Timestamp.html:17
	qw422016.N().S(`nowrap" data-timestamp="`)
//line views/components/view/Timestamp.html:17
	qw422016.E().S(util.TimeToFull(value))
//line views/components/view/Timestamp.html:17
	qw422016.N().S(`">`)
//line views/components/view/Timestamp.html:17
	qw422016.E().S(util.TimeRelative(value))
//line views/components/view/Timestamp.html:17
	qw422016.N().S(`</span>`)
//line views/components/view/Timestamp.html:18
}

//line views/components/view/Timestamp.html:18
func WriteTimestampRelative(qq422016 qtio422016.Writer, value *time.Time, static bool) {
//line views/components/view/Timestamp.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Timestamp.html:18
	StreamTimestampRelative(qw422016, value, static)
//line views/components/view/Timestamp.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Timestamp.html:18
}

//line views/components/view/Timestamp.html:18
func TimestampRelative(value *time.Time, static bool) string {
//line views/components/view/Timestamp.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Timestamp.html:18
	WriteTimestampRelative(qb422016, value, static)
//line views/components/view/Timestamp.html:18
	qs422016 := string(qb422016.B)
//line views/components/view/Timestamp.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Timestamp.html:18
	return qs422016
//line views/components/view/Timestamp.html:18
}

//line views/components/view/Timestamp.html:20
func StreamTimestampDay(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/view/Timestamp.html:21
	qw422016.E().S(util.TimeToYMD(value))
//line views/components/view/Timestamp.html:22
}

//line views/components/view/Timestamp.html:22
func WriteTimestampDay(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/view/Timestamp.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Timestamp.html:22
	StreamTimestampDay(qw422016, value)
//line views/components/view/Timestamp.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Timestamp.html:22
}

//line views/components/view/Timestamp.html:22
func TimestampDay(value *time.Time) string {
//line views/components/view/Timestamp.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Timestamp.html:22
	WriteTimestampDay(qb422016, value)
//line views/components/view/Timestamp.html:22
	qs422016 := string(qb422016.B)
//line views/components/view/Timestamp.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Timestamp.html:22
	return qs422016
//line views/components/view/Timestamp.html:22
}
