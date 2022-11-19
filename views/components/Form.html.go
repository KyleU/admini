// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Form.html:2
package components

//line views/components/Form.html:2
import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vutil"
)

//line views/components/Form.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Form.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Form.html:14
func StreamFormInput(qw422016 *qt422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/Form.html:15
	if id == "" {
//line views/components/Form.html:15
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:16
		qw422016.E().S(key)
//line views/components/Form.html:16
		qw422016.N().S(`" value="`)
//line views/components/Form.html:16
		qw422016.E().S(value)
//line views/components/Form.html:16
		qw422016.N().S(`"`)
//line views/components/Form.html:16
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:16
		qw422016.N().S(`/>`)
//line views/components/Form.html:17
	} else {
//line views/components/Form.html:17
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:18
		qw422016.E().S(id)
//line views/components/Form.html:18
		qw422016.N().S(`" name="`)
//line views/components/Form.html:18
		qw422016.E().S(key)
//line views/components/Form.html:18
		qw422016.N().S(`" value="`)
//line views/components/Form.html:18
		qw422016.E().S(value)
//line views/components/Form.html:18
		qw422016.N().S(`"`)
//line views/components/Form.html:18
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:18
		qw422016.N().S(`/>`)
//line views/components/Form.html:19
	}
//line views/components/Form.html:20
}

//line views/components/Form.html:20
func WriteFormInput(qq422016 qtio422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/Form.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:20
	StreamFormInput(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:20
}

//line views/components/Form.html:20
func FormInput(key string, id string, value string, placeholder ...string) string {
//line views/components/Form.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:20
	WriteFormInput(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:20
	qs422016 := string(qb422016.B)
//line views/components/Form.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:20
	return qs422016
//line views/components/Form.html:20
}

//line views/components/Form.html:22
func StreamFormInputPassword(qw422016 *qt422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/Form.html:23
	if id == "" {
//line views/components/Form.html:23
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:24
		qw422016.E().S(key)
//line views/components/Form.html:24
		qw422016.N().S(`" type="password" value="`)
//line views/components/Form.html:24
		qw422016.E().S(value)
//line views/components/Form.html:24
		qw422016.N().S(`"`)
//line views/components/Form.html:24
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:24
		qw422016.N().S(`/>`)
//line views/components/Form.html:25
	} else {
//line views/components/Form.html:25
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:26
		qw422016.E().S(id)
//line views/components/Form.html:26
		qw422016.N().S(`" name="`)
//line views/components/Form.html:26
		qw422016.E().S(key)
//line views/components/Form.html:26
		qw422016.N().S(`" type="password" value="`)
//line views/components/Form.html:26
		qw422016.E().S(value)
//line views/components/Form.html:26
		qw422016.N().S(`"`)
//line views/components/Form.html:26
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:26
		qw422016.N().S(`/>`)
//line views/components/Form.html:27
	}
//line views/components/Form.html:28
}

//line views/components/Form.html:28
func WriteFormInputPassword(qq422016 qtio422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/Form.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:28
	StreamFormInputPassword(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:28
}

//line views/components/Form.html:28
func FormInputPassword(key string, id string, value string, placeholder ...string) string {
//line views/components/Form.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:28
	WriteFormInputPassword(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:28
	qs422016 := string(qb422016.B)
//line views/components/Form.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:28
	return qs422016
//line views/components/Form.html:28
}

//line views/components/Form.html:30
func StreamFormInputNumber(qw422016 *qt422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/Form.html:31
	if id == "" {
//line views/components/Form.html:31
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:32
		qw422016.E().S(key)
//line views/components/Form.html:32
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:32
		qw422016.E().V(value)
//line views/components/Form.html:32
		qw422016.N().S(`"`)
//line views/components/Form.html:32
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:32
		qw422016.N().S(`/>`)
//line views/components/Form.html:33
	} else {
//line views/components/Form.html:33
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:34
		qw422016.E().S(id)
//line views/components/Form.html:34
		qw422016.N().S(`" name="`)
//line views/components/Form.html:34
		qw422016.E().S(key)
//line views/components/Form.html:34
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:34
		qw422016.E().V(value)
//line views/components/Form.html:34
		qw422016.N().S(`"`)
//line views/components/Form.html:34
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:34
		qw422016.N().S(`/>`)
//line views/components/Form.html:35
	}
//line views/components/Form.html:36
}

//line views/components/Form.html:36
func WriteFormInputNumber(qq422016 qtio422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/Form.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:36
	StreamFormInputNumber(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:36
}

//line views/components/Form.html:36
func FormInputNumber(key string, id string, value any, placeholder ...string) string {
//line views/components/Form.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:36
	WriteFormInputNumber(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:36
	qs422016 := string(qb422016.B)
//line views/components/Form.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:36
	return qs422016
//line views/components/Form.html:36
}

//line views/components/Form.html:38
func StreamFormInputFloat(qw422016 *qt422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/Form.html:39
	if id == "" {
//line views/components/Form.html:39
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:40
		qw422016.E().S(key)
//line views/components/Form.html:40
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:40
		qw422016.E().V(value)
//line views/components/Form.html:40
		qw422016.N().S(`" step="0.001"`)
//line views/components/Form.html:40
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:40
		qw422016.N().S(`/>`)
//line views/components/Form.html:41
	} else {
//line views/components/Form.html:41
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:42
		qw422016.E().S(id)
//line views/components/Form.html:42
		qw422016.N().S(`" name="`)
//line views/components/Form.html:42
		qw422016.E().S(key)
//line views/components/Form.html:42
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:42
		qw422016.E().V(value)
//line views/components/Form.html:42
		qw422016.N().S(`" step="0.001"`)
//line views/components/Form.html:42
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:42
		qw422016.N().S(`/>`)
//line views/components/Form.html:43
	}
//line views/components/Form.html:44
}

//line views/components/Form.html:44
func WriteFormInputFloat(qq422016 qtio422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/Form.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:44
	StreamFormInputFloat(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:44
}

//line views/components/Form.html:44
func FormInputFloat(key string, id string, value any, placeholder ...string) string {
//line views/components/Form.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:44
	WriteFormInputFloat(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:44
	qs422016 := string(qb422016.B)
//line views/components/Form.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:44
	return qs422016
//line views/components/Form.html:44
}

//line views/components/Form.html:46
func StreamFormInputTimestamp(qw422016 *qt422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/Form.html:47
	if id == "" {
//line views/components/Form.html:47
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:48
		qw422016.E().S(key)
//line views/components/Form.html:48
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:48
		qw422016.E().S(util.TimeToFull(value))
//line views/components/Form.html:48
		qw422016.N().S(`"`)
//line views/components/Form.html:48
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:48
		qw422016.N().S(`/>`)
//line views/components/Form.html:49
	} else {
//line views/components/Form.html:49
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:50
		qw422016.E().S(id)
//line views/components/Form.html:50
		qw422016.N().S(`" name="`)
//line views/components/Form.html:50
		qw422016.E().S(key)
//line views/components/Form.html:50
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:50
		qw422016.E().S(util.TimeToFull(value))
//line views/components/Form.html:50
		qw422016.N().S(`"`)
//line views/components/Form.html:50
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:50
		qw422016.N().S(`/>`)
//line views/components/Form.html:51
	}
//line views/components/Form.html:52
}

//line views/components/Form.html:52
func WriteFormInputTimestamp(qq422016 qtio422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/Form.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:52
	StreamFormInputTimestamp(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:52
}

//line views/components/Form.html:52
func FormInputTimestamp(key string, id string, value *time.Time, placeholder ...string) string {
//line views/components/Form.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:52
	WriteFormInputTimestamp(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:52
	qs422016 := string(qb422016.B)
//line views/components/Form.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:52
	return qs422016
//line views/components/Form.html:52
}

//line views/components/Form.html:54
func StreamFormInputTimestampDay(qw422016 *qt422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/Form.html:55
	if id == "" {
//line views/components/Form.html:55
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:56
		qw422016.E().S(key)
//line views/components/Form.html:56
		qw422016.N().S(`" type="date" value="`)
//line views/components/Form.html:56
		qw422016.E().S(util.TimeToFull(value))
//line views/components/Form.html:56
		qw422016.N().S(`"`)
//line views/components/Form.html:56
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:56
		qw422016.N().S(`/>`)
//line views/components/Form.html:57
	} else {
//line views/components/Form.html:57
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:58
		qw422016.E().S(id)
//line views/components/Form.html:58
		qw422016.N().S(`" name="`)
//line views/components/Form.html:58
		qw422016.E().S(key)
//line views/components/Form.html:58
		qw422016.N().S(`" type="date" value="`)
//line views/components/Form.html:58
		qw422016.E().S(util.TimeToYMD(value))
//line views/components/Form.html:58
		qw422016.N().S(`"`)
//line views/components/Form.html:58
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:58
		qw422016.N().S(`/>`)
//line views/components/Form.html:59
	}
//line views/components/Form.html:60
}

//line views/components/Form.html:60
func WriteFormInputTimestampDay(qq422016 qtio422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/Form.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:60
	StreamFormInputTimestampDay(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:60
}

//line views/components/Form.html:60
func FormInputTimestampDay(key string, id string, value *time.Time, placeholder ...string) string {
//line views/components/Form.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:60
	WriteFormInputTimestampDay(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:60
	qs422016 := string(qb422016.B)
//line views/components/Form.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:60
	return qs422016
//line views/components/Form.html:60
}

//line views/components/Form.html:62
func StreamFormInputUUID(qw422016 *qt422016.Writer, key string, id string, value *uuid.UUID, placeholder ...string) {
//line views/components/Form.html:64
	v := ""
	if value != nil {
		v = value.String()
	}

//line views/components/Form.html:69
	StreamFormInput(qw422016, key, id, v, placeholder...)
//line views/components/Form.html:70
}

//line views/components/Form.html:70
func WriteFormInputUUID(qq422016 qtio422016.Writer, key string, id string, value *uuid.UUID, placeholder ...string) {
//line views/components/Form.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:70
	StreamFormInputUUID(qw422016, key, id, value, placeholder...)
//line views/components/Form.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:70
}

//line views/components/Form.html:70
func FormInputUUID(key string, id string, value *uuid.UUID, placeholder ...string) string {
//line views/components/Form.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:70
	WriteFormInputUUID(qb422016, key, id, value, placeholder...)
//line views/components/Form.html:70
	qs422016 := string(qb422016.B)
//line views/components/Form.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:70
	return qs422016
//line views/components/Form.html:70
}

//line views/components/Form.html:72
func StreamFormTextarea(qw422016 *qt422016.Writer, key string, id string, rows int, value string, placeholder ...string) {
//line views/components/Form.html:73
	if id == "" {
//line views/components/Form.html:73
		qw422016.N().S(`<textarea rows="`)
//line views/components/Form.html:74
		qw422016.N().D(rows)
//line views/components/Form.html:74
		qw422016.N().S(`" name="`)
//line views/components/Form.html:74
		qw422016.E().S(key)
//line views/components/Form.html:74
		qw422016.N().S(`"`)
//line views/components/Form.html:74
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:74
		qw422016.N().S(`>`)
//line views/components/Form.html:74
		qw422016.E().S(value)
//line views/components/Form.html:74
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:75
	} else {
//line views/components/Form.html:75
		qw422016.N().S(`<textarea rows="`)
//line views/components/Form.html:76
		qw422016.N().D(rows)
//line views/components/Form.html:76
		qw422016.N().S(`" id="`)
//line views/components/Form.html:76
		qw422016.E().S(id)
//line views/components/Form.html:76
		qw422016.N().S(`" name="`)
//line views/components/Form.html:76
		qw422016.E().S(key)
//line views/components/Form.html:76
		qw422016.N().S(`"`)
//line views/components/Form.html:76
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:76
		qw422016.N().S(`>`)
//line views/components/Form.html:76
		qw422016.E().S(value)
//line views/components/Form.html:76
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:77
	}
//line views/components/Form.html:78
}

//line views/components/Form.html:78
func WriteFormTextarea(qq422016 qtio422016.Writer, key string, id string, rows int, value string, placeholder ...string) {
//line views/components/Form.html:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:78
	StreamFormTextarea(qw422016, key, id, rows, value, placeholder...)
//line views/components/Form.html:78
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:78
}

//line views/components/Form.html:78
func FormTextarea(key string, id string, rows int, value string, placeholder ...string) string {
//line views/components/Form.html:78
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:78
	WriteFormTextarea(qb422016, key, id, rows, value, placeholder...)
//line views/components/Form.html:78
	qs422016 := string(qb422016.B)
//line views/components/Form.html:78
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:78
	return qs422016
//line views/components/Form.html:78
}

//line views/components/Form.html:80
func StreamFormSelect(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:80
	qw422016.N().S(`<select name="`)
//line views/components/Form.html:81
	qw422016.E().S(key)
//line views/components/Form.html:81
	qw422016.N().S(`"`)
//line views/components/Form.html:81
	if id == `` {
//line views/components/Form.html:81
		qw422016.N().S(` `)
//line views/components/Form.html:81
		qw422016.N().S(`id="`)
//line views/components/Form.html:81
		qw422016.E().S(id)
//line views/components/Form.html:81
		qw422016.N().S(`"`)
//line views/components/Form.html:81
	}
//line views/components/Form.html:81
	qw422016.N().S(`>`)
//line views/components/Form.html:82
	for idx, opt := range opts {
//line views/components/Form.html:84
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:89
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Form.html:90
		if opt == value {
//line views/components/Form.html:90
			qw422016.N().S(`<option selected="selected" value="`)
//line views/components/Form.html:91
			qw422016.E().S(opt)
//line views/components/Form.html:91
			qw422016.N().S(`">`)
//line views/components/Form.html:91
			qw422016.E().S(title)
//line views/components/Form.html:91
			qw422016.N().S(`</option>`)
//line views/components/Form.html:92
		} else {
//line views/components/Form.html:92
			qw422016.N().S(`<option value="`)
//line views/components/Form.html:93
			qw422016.E().S(opt)
//line views/components/Form.html:93
			qw422016.N().S(`">`)
//line views/components/Form.html:93
			qw422016.E().S(title)
//line views/components/Form.html:93
			qw422016.N().S(`</option>`)
//line views/components/Form.html:94
		}
//line views/components/Form.html:95
	}
//line views/components/Form.html:96
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:96
	qw422016.N().S(`</select>`)
//line views/components/Form.html:98
}

//line views/components/Form.html:98
func WriteFormSelect(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:98
	StreamFormSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:98
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:98
}

//line views/components/Form.html:98
func FormSelect(key string, id string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:98
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:98
	WriteFormSelect(qb422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:98
	qs422016 := string(qb422016.B)
//line views/components/Form.html:98
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:98
	return qs422016
//line views/components/Form.html:98
}

//line views/components/Form.html:100
func StreamFormDatalist(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int, placeholder ...string) {
//line views/components/Form.html:101
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:101
	qw422016.N().S(`<input id="`)
//line views/components/Form.html:102
	qw422016.E().S(id)
//line views/components/Form.html:102
	qw422016.N().S(`" list="`)
//line views/components/Form.html:102
	qw422016.E().S(id)
//line views/components/Form.html:102
	qw422016.N().S(`-list" name="`)
//line views/components/Form.html:102
	qw422016.E().S(key)
//line views/components/Form.html:102
	qw422016.N().S(`" value="`)
//line views/components/Form.html:102
	qw422016.E().S(value)
//line views/components/Form.html:102
	qw422016.N().S(`"`)
//line views/components/Form.html:102
	streamphFor(qw422016, placeholder)
//line views/components/Form.html:102
	qw422016.N().S(`/>`)
//line views/components/Form.html:103
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:103
	qw422016.N().S(`<datalist id="`)
//line views/components/Form.html:104
	qw422016.E().S(id)
//line views/components/Form.html:104
	qw422016.N().S(`-list">`)
//line views/components/Form.html:105
	for idx, opt := range opts {
//line views/components/Form.html:107
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:112
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Form.html:112
		qw422016.N().S(`<option value="`)
//line views/components/Form.html:113
		qw422016.E().S(opt)
//line views/components/Form.html:113
		qw422016.N().S(`">`)
//line views/components/Form.html:113
		qw422016.E().S(title)
//line views/components/Form.html:113
		qw422016.N().S(`</option>`)
//line views/components/Form.html:114
	}
//line views/components/Form.html:115
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:115
	qw422016.N().S(`</datalist>`)
//line views/components/Form.html:117
}

//line views/components/Form.html:117
func WriteFormDatalist(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int, placeholder ...string) {
//line views/components/Form.html:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:117
	StreamFormDatalist(qw422016, key, id, value, opts, titles, indent, placeholder...)
//line views/components/Form.html:117
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:117
}

//line views/components/Form.html:117
func FormDatalist(key string, id string, value string, opts []string, titles []string, indent int, placeholder ...string) string {
//line views/components/Form.html:117
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:117
	WriteFormDatalist(qb422016, key, id, value, opts, titles, indent, placeholder...)
//line views/components/Form.html:117
	qs422016 := string(qb422016.B)
//line views/components/Form.html:117
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:117
	return qs422016
//line views/components/Form.html:117
}

//line views/components/Form.html:119
func StreamFormRadio(qw422016 *qt422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:120
	for idx, opt := range opts {
//line views/components/Form.html:122
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:127
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:128
		if opt == value {
//line views/components/Form.html:128
			qw422016.N().S(`<label class="radio-label"><input type="radio" name="`)
//line views/components/Form.html:129
			qw422016.E().S(key)
//line views/components/Form.html:129
			qw422016.N().S(`" value="`)
//line views/components/Form.html:129
			qw422016.E().S(opt)
//line views/components/Form.html:129
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:129
			qw422016.N().S(` `)
//line views/components/Form.html:129
			qw422016.E().S(title)
//line views/components/Form.html:129
			qw422016.N().S(`</label>`)
//line views/components/Form.html:130
		} else {
//line views/components/Form.html:130
			qw422016.N().S(`<label class="radio-label"><input type="radio" name="`)
//line views/components/Form.html:131
			qw422016.E().S(key)
//line views/components/Form.html:131
			qw422016.N().S(`" value="`)
//line views/components/Form.html:131
			qw422016.E().S(opt)
//line views/components/Form.html:131
			qw422016.N().S(`" />`)
//line views/components/Form.html:131
			qw422016.N().S(` `)
//line views/components/Form.html:131
			qw422016.E().S(title)
//line views/components/Form.html:131
			qw422016.N().S(`</label>`)
//line views/components/Form.html:132
		}
//line views/components/Form.html:133
	}
//line views/components/Form.html:134
}

//line views/components/Form.html:134
func WriteFormRadio(qq422016 qtio422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:134
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:134
	StreamFormRadio(qw422016, key, value, opts, titles, indent)
//line views/components/Form.html:134
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:134
}

//line views/components/Form.html:134
func FormRadio(key string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:134
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:134
	WriteFormRadio(qb422016, key, value, opts, titles, indent)
//line views/components/Form.html:134
	qs422016 := string(qb422016.B)
//line views/components/Form.html:134
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:134
	return qs422016
//line views/components/Form.html:134
}

//line views/components/Form.html:136
func StreamFormCheckbox(qw422016 *qt422016.Writer, key string, values []string, opts []string, titles []string, linebreaks bool, indent int) {
//line views/components/Form.html:137
	for idx, opt := range opts {
//line views/components/Form.html:139
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:144
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:145
		if slices.Contains(values, opt) {
//line views/components/Form.html:145
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:146
			qw422016.E().S(key)
//line views/components/Form.html:146
			qw422016.N().S(`" value="`)
//line views/components/Form.html:146
			qw422016.E().S(opt)
//line views/components/Form.html:146
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:146
			qw422016.N().S(` `)
//line views/components/Form.html:146
			qw422016.E().S(title)
//line views/components/Form.html:146
			qw422016.N().S(`</label>`)
//line views/components/Form.html:147
		} else {
//line views/components/Form.html:147
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:148
			qw422016.E().S(key)
//line views/components/Form.html:148
			qw422016.N().S(`" value="`)
//line views/components/Form.html:148
			qw422016.E().S(opt)
//line views/components/Form.html:148
			qw422016.N().S(`" />`)
//line views/components/Form.html:148
			qw422016.N().S(` `)
//line views/components/Form.html:148
			qw422016.E().S(title)
//line views/components/Form.html:148
			qw422016.N().S(`</label>`)
//line views/components/Form.html:149
		}
//line views/components/Form.html:150
		if slices.Contains(values, opt) {
//line views/components/Form.html:150
			qw422016.N().S(`<br />`)
//line views/components/Form.html:152
		}
//line views/components/Form.html:153
	}
//line views/components/Form.html:154
}

//line views/components/Form.html:154
func WriteFormCheckbox(qq422016 qtio422016.Writer, key string, values []string, opts []string, titles []string, linebreaks bool, indent int) {
//line views/components/Form.html:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:154
	StreamFormCheckbox(qw422016, key, values, opts, titles, linebreaks, indent)
//line views/components/Form.html:154
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:154
}

//line views/components/Form.html:154
func FormCheckbox(key string, values []string, opts []string, titles []string, linebreaks bool, indent int) string {
//line views/components/Form.html:154
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:154
	WriteFormCheckbox(qb422016, key, values, opts, titles, linebreaks, indent)
//line views/components/Form.html:154
	qs422016 := string(qb422016.B)
//line views/components/Form.html:154
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:154
	return qs422016
//line views/components/Form.html:154
}

//line views/components/Form.html:156
func StreamFormInputTags(qw422016 *qt422016.Writer, key string, id string, values []string, ps *cutil.PageState, placeholder ...string) {
//line views/components/Form.html:158
	ps.AddIcon("times")
	ps.AddIcon("plus")

//line views/components/Form.html:160
	qw422016.N().S(`<div class="tag-editor">`)
//line views/components/Form.html:162
	if id == "" {
//line views/components/Form.html:162
		qw422016.N().S(`<input class="result" name="`)
//line views/components/Form.html:163
		qw422016.E().S(key)
//line views/components/Form.html:163
		qw422016.N().S(`" value="`)
//line views/components/Form.html:163
		qw422016.E().S(strings.Join(values, `, `))
//line views/components/Form.html:163
		qw422016.N().S(`"`)
//line views/components/Form.html:163
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:163
		qw422016.N().S(`/>`)
//line views/components/Form.html:164
	} else {
//line views/components/Form.html:164
		qw422016.N().S(`<input class="result" id="`)
//line views/components/Form.html:165
		qw422016.E().S(id)
//line views/components/Form.html:165
		qw422016.N().S(`" name="`)
//line views/components/Form.html:165
		qw422016.E().S(key)
//line views/components/Form.html:165
		qw422016.N().S(`" value="`)
//line views/components/Form.html:165
		qw422016.E().S(strings.Join(values, `, `))
//line views/components/Form.html:165
		qw422016.N().S(`"`)
//line views/components/Form.html:165
		streamphFor(qw422016, placeholder)
//line views/components/Form.html:165
		qw422016.N().S(`/>`)
//line views/components/Form.html:166
	}
//line views/components/Form.html:166
	qw422016.N().S(`<div class="tags"></div><div class="clear"></div></div>`)
//line views/components/Form.html:170
}

//line views/components/Form.html:170
func WriteFormInputTags(qq422016 qtio422016.Writer, key string, id string, values []string, ps *cutil.PageState, placeholder ...string) {
//line views/components/Form.html:170
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:170
	StreamFormInputTags(qw422016, key, id, values, ps, placeholder...)
//line views/components/Form.html:170
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:170
}

//line views/components/Form.html:170
func FormInputTags(key string, id string, values []string, ps *cutil.PageState, placeholder ...string) string {
//line views/components/Form.html:170
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:170
	WriteFormInputTags(qb422016, key, id, values, ps, placeholder...)
//line views/components/Form.html:170
	qs422016 := string(qb422016.B)
//line views/components/Form.html:170
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:170
	return qs422016
//line views/components/Form.html:170
}

//line views/components/Form.html:172
func streamphFor(qw422016 *qt422016.Writer, phs []string) {
//line views/components/Form.html:173
	if len(phs) > 0 {
//line views/components/Form.html:173
		qw422016.N().S(` `)
//line views/components/Form.html:173
		qw422016.N().S(`placeholder="`)
//line views/components/Form.html:173
		qw422016.E().S(strings.Join(phs, "; "))
//line views/components/Form.html:173
		qw422016.N().S(`"`)
//line views/components/Form.html:173
	}
//line views/components/Form.html:174
}

//line views/components/Form.html:174
func writephFor(qq422016 qtio422016.Writer, phs []string) {
//line views/components/Form.html:174
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:174
	streamphFor(qw422016, phs)
//line views/components/Form.html:174
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:174
}

//line views/components/Form.html:174
func phFor(phs []string) string {
//line views/components/Form.html:174
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:174
	writephFor(qb422016, phs)
//line views/components/Form.html:174
	qs422016 := string(qb422016.B)
//line views/components/Form.html:174
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:174
	return qs422016
//line views/components/Form.html:174
}
