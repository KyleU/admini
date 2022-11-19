// Code generated by qtc from "FormVertical.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/FormVertical.html:2
package components

//line views/components/FormVertical.html:2
import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/vutil"
)

//line views/components/FormVertical.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/FormVertical.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/FormVertical.html:12
func StreamFormVerticalInput(qw422016 *qt422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:12
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:14
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:14
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:15
	qw422016.E().S(title)
//line views/components/FormVertical.html:15
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:16
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:16
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:17
	StreamFormInput(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:17
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:18
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:18
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:20
}

//line views/components/FormVertical.html:20
func WriteFormVerticalInput(qq422016 qtio422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:20
	StreamFormVerticalInput(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:20
}

//line views/components/FormVertical.html:20
func FormVerticalInput(key string, title string, value string, indent int, help ...string) string {
//line views/components/FormVertical.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:20
	WriteFormVerticalInput(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:20
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:20
	return qs422016
//line views/components/FormVertical.html:20
}

//line views/components/FormVertical.html:22
func StreamFormVerticalInputPassword(qw422016 *qt422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:22
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:24
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:24
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:25
	qw422016.E().S(title)
//line views/components/FormVertical.html:25
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:26
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:26
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:27
	StreamFormInputPassword(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:27
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:28
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:28
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:30
}

//line views/components/FormVertical.html:30
func WriteFormVerticalInputPassword(qq422016 qtio422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/FormVertical.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:30
	StreamFormVerticalInputPassword(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:30
}

//line views/components/FormVertical.html:30
func FormVerticalInputPassword(key string, title string, value string, indent int, help ...string) string {
//line views/components/FormVertical.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:30
	WriteFormVerticalInputPassword(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:30
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:30
	return qs422016
//line views/components/FormVertical.html:30
}

//line views/components/FormVertical.html:32
func StreamFormVerticalInputNumber(qw422016 *qt422016.Writer, key string, title string, value int, indent int, help ...string) {
//line views/components/FormVertical.html:32
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:34
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:34
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:35
	qw422016.E().S(title)
//line views/components/FormVertical.html:35
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:36
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:36
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:37
	StreamFormInputNumber(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:37
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:38
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:38
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:40
}

//line views/components/FormVertical.html:40
func WriteFormVerticalInputNumber(qq422016 qtio422016.Writer, key string, title string, value int, indent int, help ...string) {
//line views/components/FormVertical.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:40
	StreamFormVerticalInputNumber(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:40
}

//line views/components/FormVertical.html:40
func FormVerticalInputNumber(key string, title string, value int, indent int, help ...string) string {
//line views/components/FormVertical.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:40
	WriteFormVerticalInputNumber(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:40
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:40
	return qs422016
//line views/components/FormVertical.html:40
}

//line views/components/FormVertical.html:42
func StreamFormVerticalInputFloat(qw422016 *qt422016.Writer, key string, title string, value float64, indent int, help ...string) {
//line views/components/FormVertical.html:42
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:44
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:44
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:45
	qw422016.E().S(title)
//line views/components/FormVertical.html:45
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:46
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:46
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:47
	StreamFormInputFloat(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:47
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:48
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:48
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:50
}

//line views/components/FormVertical.html:50
func WriteFormVerticalInputFloat(qq422016 qtio422016.Writer, key string, title string, value float64, indent int, help ...string) {
//line views/components/FormVertical.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:50
	StreamFormVerticalInputFloat(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:50
}

//line views/components/FormVertical.html:50
func FormVerticalInputFloat(key string, title string, value float64, indent int, help ...string) string {
//line views/components/FormVertical.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:50
	WriteFormVerticalInputFloat(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:50
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:50
	return qs422016
//line views/components/FormVertical.html:50
}

//line views/components/FormVertical.html:52
func StreamFormVerticalInputTimestamp(qw422016 *qt422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:52
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:54
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:54
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:55
	qw422016.E().S(title)
//line views/components/FormVertical.html:55
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:56
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:56
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:57
	StreamFormInputTimestamp(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:57
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:58
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:58
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:60
}

//line views/components/FormVertical.html:60
func WriteFormVerticalInputTimestamp(qq422016 qtio422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:60
	StreamFormVerticalInputTimestamp(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:60
}

//line views/components/FormVertical.html:60
func FormVerticalInputTimestamp(key string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/FormVertical.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:60
	WriteFormVerticalInputTimestamp(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:60
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:60
	return qs422016
//line views/components/FormVertical.html:60
}

//line views/components/FormVertical.html:62
func StreamFormVerticalInputTimestampDay(qw422016 *qt422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:62
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:64
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:64
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:65
	qw422016.E().S(title)
//line views/components/FormVertical.html:65
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:66
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:66
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:67
	StreamFormInputTimestampDay(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:67
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:68
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:68
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:70
}

//line views/components/FormVertical.html:70
func WriteFormVerticalInputTimestampDay(qq422016 qtio422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/FormVertical.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:70
	StreamFormVerticalInputTimestampDay(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:70
}

//line views/components/FormVertical.html:70
func FormVerticalInputTimestampDay(key string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/FormVertical.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:70
	WriteFormVerticalInputTimestampDay(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:70
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:70
	return qs422016
//line views/components/FormVertical.html:70
}

//line views/components/FormVertical.html:72
func StreamFormVerticalInputUUID(qw422016 *qt422016.Writer, key string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/FormVertical.html:72
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:74
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:74
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:75
	qw422016.E().S(title)
//line views/components/FormVertical.html:75
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:76
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:76
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:77
	StreamFormInputUUID(qw422016, key, "input-"+key, value, help...)
//line views/components/FormVertical.html:77
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:78
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:78
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:80
}

//line views/components/FormVertical.html:80
func WriteFormVerticalInputUUID(qq422016 qtio422016.Writer, key string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/FormVertical.html:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:80
	StreamFormVerticalInputUUID(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:80
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:80
}

//line views/components/FormVertical.html:80
func FormVerticalInputUUID(key string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/FormVertical.html:80
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:80
	WriteFormVerticalInputUUID(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:80
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:80
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:80
	return qs422016
//line views/components/FormVertical.html:80
}

//line views/components/FormVertical.html:82
func StreamFormVerticalTextarea(qw422016 *qt422016.Writer, key string, title string, rows int, value string, indent int, help ...string) {
//line views/components/FormVertical.html:82
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:84
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:84
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:85
	qw422016.E().S(title)
//line views/components/FormVertical.html:85
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:86
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:86
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:87
	StreamFormTextarea(qw422016, key, "input-"+key, rows, value, help...)
//line views/components/FormVertical.html:87
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:88
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:88
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:90
}

//line views/components/FormVertical.html:90
func WriteFormVerticalTextarea(qq422016 qtio422016.Writer, key string, title string, rows int, value string, indent int, help ...string) {
//line views/components/FormVertical.html:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:90
	StreamFormVerticalTextarea(qw422016, key, title, rows, value, indent, help...)
//line views/components/FormVertical.html:90
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:90
}

//line views/components/FormVertical.html:90
func FormVerticalTextarea(key string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/FormVertical.html:90
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:90
	WriteFormVerticalTextarea(qb422016, key, title, rows, value, indent, help...)
//line views/components/FormVertical.html:90
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:90
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:90
	return qs422016
//line views/components/FormVertical.html:90
}

//line views/components/FormVertical.html:92
func StreamFormVerticalSelect(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:92
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:94
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:94
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:95
	qw422016.E().S(title)
//line views/components/FormVertical.html:95
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:96
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:96
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:97
	StreamFormSelect(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/FormVertical.html:97
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:98
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:98
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:100
}

//line views/components/FormVertical.html:100
func WriteFormVerticalSelect(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:100
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:100
	StreamFormVerticalSelect(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:100
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:100
}

//line views/components/FormVertical.html:100
func FormVerticalSelect(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:100
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:100
	WriteFormVerticalSelect(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:100
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:100
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:100
	return qs422016
//line views/components/FormVertical.html:100
}

//line views/components/FormVertical.html:102
func StreamFormVerticalDatalist(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:102
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:104
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:104
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:105
	qw422016.E().S(title)
//line views/components/FormVertical.html:105
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:106
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:106
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:107
	StreamFormDatalist(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/FormVertical.html:107
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:108
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:108
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:110
}

//line views/components/FormVertical.html:110
func WriteFormVerticalDatalist(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:110
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:110
	StreamFormVerticalDatalist(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:110
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:110
}

//line views/components/FormVertical.html:110
func FormVerticalDatalist(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:110
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:110
	WriteFormVerticalDatalist(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:110
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:110
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:110
	return qs422016
//line views/components/FormVertical.html:110
}

//line views/components/FormVertical.html:112
func StreamFormVerticalRadio(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:112
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:114
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:114
	qw422016.N().S(`<div><label`)
//line views/components/FormVertical.html:115
	streamtitleFor(qw422016, help)
//line views/components/FormVertical.html:115
	qw422016.N().S(`>`)
//line views/components/FormVertical.html:115
	qw422016.E().S(title)
//line views/components/FormVertical.html:115
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:116
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:116
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:118
	StreamFormRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/FormVertical.html:119
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:119
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:121
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:121
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:123
}

//line views/components/FormVertical.html:123
func WriteFormVerticalRadio(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/FormVertical.html:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:123
	StreamFormVerticalRadio(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:123
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:123
}

//line views/components/FormVertical.html:123
func FormVerticalRadio(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/FormVertical.html:123
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:123
	WriteFormVerticalRadio(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/FormVertical.html:123
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:123
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:123
	return qs422016
//line views/components/FormVertical.html:123
}

//line views/components/FormVertical.html:125
func StreamFormVerticalBoolean(qw422016 *qt422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/FormVertical.html:126
	StreamFormVerticalRadio(qw422016, key, title, fmt.Sprint(value), []string{"true", "false"}, []string{"True", "False"}, indent)
//line views/components/FormVertical.html:127
}

//line views/components/FormVertical.html:127
func WriteFormVerticalBoolean(qq422016 qtio422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/FormVertical.html:127
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:127
	StreamFormVerticalBoolean(qw422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:127
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:127
}

//line views/components/FormVertical.html:127
func FormVerticalBoolean(key string, title string, value bool, indent int, help ...string) string {
//line views/components/FormVertical.html:127
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:127
	WriteFormVerticalBoolean(qb422016, key, title, value, indent, help...)
//line views/components/FormVertical.html:127
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:127
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:127
	return qs422016
//line views/components/FormVertical.html:127
}

//line views/components/FormVertical.html:129
func StreamFormVerticalCheckbox(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) {
//line views/components/FormVertical.html:129
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:131
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:131
	qw422016.N().S(`<div><label>`)
//line views/components/FormVertical.html:132
	qw422016.E().S(title)
//line views/components/FormVertical.html:132
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:133
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:133
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:135
	StreamFormCheckbox(qw422016, key, values, opts, titles, linebreaks, indent+2)
//line views/components/FormVertical.html:136
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:136
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:138
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:138
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:140
}

//line views/components/FormVertical.html:140
func WriteFormVerticalCheckbox(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) {
//line views/components/FormVertical.html:140
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:140
	StreamFormVerticalCheckbox(qw422016, key, title, values, opts, titles, linebreaks, indent, help...)
//line views/components/FormVertical.html:140
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:140
}

//line views/components/FormVertical.html:140
func FormVerticalCheckbox(key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) string {
//line views/components/FormVertical.html:140
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:140
	WriteFormVerticalCheckbox(qb422016, key, title, values, opts, titles, linebreaks, indent, help...)
//line views/components/FormVertical.html:140
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:140
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:140
	return qs422016
//line views/components/FormVertical.html:140
}

//line views/components/FormVertical.html:142
func StreamFormVerticalInputTags(qw422016 *qt422016.Writer, key string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/FormVertical.html:142
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:144
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:144
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:145
	qw422016.E().S(title)
//line views/components/FormVertical.html:145
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:146
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:146
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:147
	StreamFormInputTags(qw422016, key, "input-"+key, values, ps, help...)
//line views/components/FormVertical.html:147
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:148
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:148
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:150
}

//line views/components/FormVertical.html:150
func WriteFormVerticalInputTags(qq422016 qtio422016.Writer, key string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/FormVertical.html:150
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:150
	StreamFormVerticalInputTags(qw422016, key, title, values, ps, indent, help...)
//line views/components/FormVertical.html:150
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:150
}

//line views/components/FormVertical.html:150
func FormVerticalInputTags(key string, title string, values []string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/FormVertical.html:150
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:150
	WriteFormVerticalInputTags(qb422016, key, title, values, ps, indent, help...)
//line views/components/FormVertical.html:150
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:150
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:150
	return qs422016
//line views/components/FormVertical.html:150
}

//line views/components/FormVertical.html:152
func StreamFormVerticalIconPicker(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/FormVertical.html:152
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/FormVertical.html:154
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:154
	qw422016.N().S(`<em class="title">`)
//line views/components/FormVertical.html:155
	qw422016.E().S(title)
//line views/components/FormVertical.html:155
	qw422016.N().S(`</em>`)
//line views/components/FormVertical.html:156
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/FormVertical.html:156
	qw422016.N().S(`<div>`)
//line views/components/FormVertical.html:157
	StreamIconPicker(qw422016, key, value, ps, indent)
//line views/components/FormVertical.html:157
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:158
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/FormVertical.html:158
	qw422016.N().S(`</div>`)
//line views/components/FormVertical.html:160
}

//line views/components/FormVertical.html:160
func WriteFormVerticalIconPicker(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/FormVertical.html:160
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/FormVertical.html:160
	StreamFormVerticalIconPicker(qw422016, key, title, value, ps, indent)
//line views/components/FormVertical.html:160
	qt422016.ReleaseWriter(qw422016)
//line views/components/FormVertical.html:160
}

//line views/components/FormVertical.html:160
func FormVerticalIconPicker(key string, title string, value string, ps *cutil.PageState, indent int) string {
//line views/components/FormVertical.html:160
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/FormVertical.html:160
	WriteFormVerticalIconPicker(qb422016, key, title, value, ps, indent)
//line views/components/FormVertical.html:160
	qs422016 := string(qb422016.B)
//line views/components/FormVertical.html:160
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/FormVertical.html:160
	return qs422016
//line views/components/FormVertical.html:160
}
