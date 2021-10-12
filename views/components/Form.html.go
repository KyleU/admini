// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Code generated by Project Forge, see https://projectforge.dev for details. -->

//line views/components/Form.html:2
package components

//line views/components/Form.html:2
import (
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/vutil"
)

//line views/components/Form.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Form.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Form.html:7
func StreamFormInput(qw422016 *qt422016.Writer, key string, id string, value string) {
//line views/components/Form.html:8
	if id == "" {
//line views/components/Form.html:8
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:9
		qw422016.E().S(key)
//line views/components/Form.html:9
		qw422016.N().S(`" value="`)
//line views/components/Form.html:9
		qw422016.E().S(value)
//line views/components/Form.html:9
		qw422016.N().S(`" />`)
//line views/components/Form.html:10
	} else {
//line views/components/Form.html:10
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:11
		qw422016.E().S(id)
//line views/components/Form.html:11
		qw422016.N().S(`" name="`)
//line views/components/Form.html:11
		qw422016.E().S(key)
//line views/components/Form.html:11
		qw422016.N().S(`" value="`)
//line views/components/Form.html:11
		qw422016.E().S(value)
//line views/components/Form.html:11
		qw422016.N().S(`" />`)
//line views/components/Form.html:12
	}
//line views/components/Form.html:13
}

//line views/components/Form.html:13
func WriteFormInput(qq422016 qtio422016.Writer, key string, id string, value string) {
//line views/components/Form.html:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:13
	StreamFormInput(qw422016, key, id, value)
//line views/components/Form.html:13
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:13
}

//line views/components/Form.html:13
func FormInput(key string, id string, value string) string {
//line views/components/Form.html:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:13
	WriteFormInput(qb422016, key, id, value)
//line views/components/Form.html:13
	qs422016 := string(qb422016.B)
//line views/components/Form.html:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:13
	return qs422016
//line views/components/Form.html:13
}

//line views/components/Form.html:15
func StreamFormInputNumber(qw422016 *qt422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:16
	if id == "" {
//line views/components/Form.html:16
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:17
		qw422016.E().S(key)
//line views/components/Form.html:17
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:17
		qw422016.E().V(value)
//line views/components/Form.html:17
		qw422016.N().S(`" />`)
//line views/components/Form.html:18
	} else {
//line views/components/Form.html:18
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:19
		qw422016.E().S(id)
//line views/components/Form.html:19
		qw422016.N().S(`" name="`)
//line views/components/Form.html:19
		qw422016.E().S(key)
//line views/components/Form.html:19
		qw422016.N().S(`" type="number" value="`)
//line views/components/Form.html:19
		qw422016.E().V(value)
//line views/components/Form.html:19
		qw422016.N().S(`" />`)
//line views/components/Form.html:20
	}
//line views/components/Form.html:21
}

//line views/components/Form.html:21
func WriteFormInputNumber(qq422016 qtio422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:21
	StreamFormInputNumber(qw422016, key, id, value)
//line views/components/Form.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:21
}

//line views/components/Form.html:21
func FormInputNumber(key string, id string, value interface{}) string {
//line views/components/Form.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:21
	WriteFormInputNumber(qb422016, key, id, value)
//line views/components/Form.html:21
	qs422016 := string(qb422016.B)
//line views/components/Form.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:21
	return qs422016
//line views/components/Form.html:21
}

//line views/components/Form.html:23
func StreamFormInputTimestamp(qw422016 *qt422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:24
	if id == "" {
//line views/components/Form.html:24
		qw422016.N().S(`<input name="`)
//line views/components/Form.html:25
		qw422016.E().S(key)
//line views/components/Form.html:25
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:25
		qw422016.E().V(value)
//line views/components/Form.html:25
		qw422016.N().S(`" />`)
//line views/components/Form.html:26
	} else {
//line views/components/Form.html:26
		qw422016.N().S(`<input id="`)
//line views/components/Form.html:27
		qw422016.E().S(id)
//line views/components/Form.html:27
		qw422016.N().S(`" name="`)
//line views/components/Form.html:27
		qw422016.E().S(key)
//line views/components/Form.html:27
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/Form.html:27
		qw422016.E().V(value)
//line views/components/Form.html:27
		qw422016.N().S(`" />`)
//line views/components/Form.html:28
	}
//line views/components/Form.html:29
}

//line views/components/Form.html:29
func WriteFormInputTimestamp(qq422016 qtio422016.Writer, key string, id string, value interface{}) {
//line views/components/Form.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:29
	StreamFormInputTimestamp(qw422016, key, id, value)
//line views/components/Form.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:29
}

//line views/components/Form.html:29
func FormInputTimestamp(key string, id string, value interface{}) string {
//line views/components/Form.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:29
	WriteFormInputTimestamp(qb422016, key, id, value)
//line views/components/Form.html:29
	qs422016 := string(qb422016.B)
//line views/components/Form.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:29
	return qs422016
//line views/components/Form.html:29
}

//line views/components/Form.html:31
func StreamFormTextarea(qw422016 *qt422016.Writer, key string, id string, rows int, value string) {
//line views/components/Form.html:32
	if id == "" {
//line views/components/Form.html:32
		qw422016.N().S(`<textarea rows="`)
//line views/components/Form.html:33
		qw422016.N().D(rows)
//line views/components/Form.html:33
		qw422016.N().S(`" name="`)
//line views/components/Form.html:33
		qw422016.E().S(key)
//line views/components/Form.html:33
		qw422016.N().S(`">`)
//line views/components/Form.html:33
		qw422016.E().S(value)
//line views/components/Form.html:33
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:34
	} else {
//line views/components/Form.html:34
		qw422016.N().S(`<textarea rows="`)
//line views/components/Form.html:35
		qw422016.N().D(rows)
//line views/components/Form.html:35
		qw422016.N().S(`" id="`)
//line views/components/Form.html:35
		qw422016.E().S(id)
//line views/components/Form.html:35
		qw422016.N().S(`" name="`)
//line views/components/Form.html:35
		qw422016.E().S(key)
//line views/components/Form.html:35
		qw422016.N().S(`">`)
//line views/components/Form.html:35
		qw422016.E().S(value)
//line views/components/Form.html:35
		qw422016.N().S(`</textarea>`)
//line views/components/Form.html:36
	}
//line views/components/Form.html:37
}

//line views/components/Form.html:37
func WriteFormTextarea(qq422016 qtio422016.Writer, key string, id string, rows int, value string) {
//line views/components/Form.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:37
	StreamFormTextarea(qw422016, key, id, rows, value)
//line views/components/Form.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:37
}

//line views/components/Form.html:37
func FormTextarea(key string, id string, rows int, value string) string {
//line views/components/Form.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:37
	WriteFormTextarea(qb422016, key, id, rows, value)
//line views/components/Form.html:37
	qs422016 := string(qb422016.B)
//line views/components/Form.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:37
	return qs422016
//line views/components/Form.html:37
}

//line views/components/Form.html:39
func StreamFormSelect(qw422016 *qt422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:40
	if id == "" {
//line views/components/Form.html:40
		qw422016.N().S(`<select name="`)
//line views/components/Form.html:41
		qw422016.E().S(key)
//line views/components/Form.html:41
		qw422016.N().S(`">`)
//line views/components/Form.html:42
	} else {
//line views/components/Form.html:42
		qw422016.N().S(`<select name="`)
//line views/components/Form.html:43
		qw422016.E().S(key)
//line views/components/Form.html:43
		qw422016.N().S(`" id="`)
//line views/components/Form.html:43
		qw422016.E().S(id)
//line views/components/Form.html:43
		qw422016.N().S(`">`)
//line views/components/Form.html:44
	}
//line views/components/Form.html:45
	for idx, opt := range opts {
//line views/components/Form.html:47
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:52
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Form.html:53
		if opt == value {
//line views/components/Form.html:53
			qw422016.N().S(`<option selected="selected" value="`)
//line views/components/Form.html:54
			qw422016.E().S(opt)
//line views/components/Form.html:54
			qw422016.N().S(`">`)
//line views/components/Form.html:54
			qw422016.E().S(title)
//line views/components/Form.html:54
			qw422016.N().S(`</option>`)
//line views/components/Form.html:55
		} else {
//line views/components/Form.html:55
			qw422016.N().S(`<option value="`)
//line views/components/Form.html:56
			qw422016.E().S(opt)
//line views/components/Form.html:56
			qw422016.N().S(`">`)
//line views/components/Form.html:56
			qw422016.E().S(title)
//line views/components/Form.html:56
			qw422016.N().S(`</option>`)
//line views/components/Form.html:57
		}
//line views/components/Form.html:58
	}
//line views/components/Form.html:59
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:59
	qw422016.N().S(`</select>`)
//line views/components/Form.html:61
}

//line views/components/Form.html:61
func WriteFormSelect(qq422016 qtio422016.Writer, key string, id string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:61
	StreamFormSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:61
}

//line views/components/Form.html:61
func FormSelect(key string, id string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:61
	WriteFormSelect(qb422016, key, id, value, opts, titles, indent)
//line views/components/Form.html:61
	qs422016 := string(qb422016.B)
//line views/components/Form.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:61
	return qs422016
//line views/components/Form.html:61
}

//line views/components/Form.html:63
func StreamFormRadio(qw422016 *qt422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:64
	for idx, opt := range opts {
//line views/components/Form.html:66
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:71
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:72
		if opt == value {
//line views/components/Form.html:72
			qw422016.N().S(`<label><input type="radio" name="`)
//line views/components/Form.html:73
			qw422016.E().S(key)
//line views/components/Form.html:73
			qw422016.N().S(`" value="`)
//line views/components/Form.html:73
			qw422016.E().S(opt)
//line views/components/Form.html:73
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:73
			qw422016.N().S(` `)
//line views/components/Form.html:73
			qw422016.E().S(title)
//line views/components/Form.html:73
			qw422016.N().S(`</label>`)
//line views/components/Form.html:74
		} else {
//line views/components/Form.html:74
			qw422016.N().S(`<label><input type="radio" name="`)
//line views/components/Form.html:75
			qw422016.E().S(key)
//line views/components/Form.html:75
			qw422016.N().S(`" value="`)
//line views/components/Form.html:75
			qw422016.E().S(opt)
//line views/components/Form.html:75
			qw422016.N().S(`" />`)
//line views/components/Form.html:75
			qw422016.N().S(` `)
//line views/components/Form.html:75
			qw422016.E().S(title)
//line views/components/Form.html:75
			qw422016.N().S(`</label>`)
//line views/components/Form.html:76
		}
//line views/components/Form.html:77
	}
//line views/components/Form.html:78
}

//line views/components/Form.html:78
func WriteFormRadio(qq422016 qtio422016.Writer, key string, value string, opts []string, titles []string, indent int) {
//line views/components/Form.html:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:78
	StreamFormRadio(qw422016, key, value, opts, titles, indent)
//line views/components/Form.html:78
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:78
}

//line views/components/Form.html:78
func FormRadio(key string, value string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:78
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:78
	WriteFormRadio(qb422016, key, value, opts, titles, indent)
//line views/components/Form.html:78
	qs422016 := string(qb422016.B)
//line views/components/Form.html:78
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:78
	return qs422016
//line views/components/Form.html:78
}

//line views/components/Form.html:80
func StreamFormCheckbox(qw422016 *qt422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/Form.html:81
	for idx, opt := range opts {
//line views/components/Form.html:83
		title := opt
		if idx < len(titles) {
			title = titles[idx]
		}

//line views/components/Form.html:88
		vutil.StreamIndent(qw422016, true, indent)
//line views/components/Form.html:89
		if util.StringArrayContains(values, opt) {
//line views/components/Form.html:89
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:90
			qw422016.E().S(key)
//line views/components/Form.html:90
			qw422016.N().S(`" value="`)
//line views/components/Form.html:90
			qw422016.E().S(opt)
//line views/components/Form.html:90
			qw422016.N().S(`" checked="checked" />`)
//line views/components/Form.html:90
			qw422016.N().S(` `)
//line views/components/Form.html:90
			qw422016.E().S(title)
//line views/components/Form.html:90
			qw422016.N().S(`</label>`)
//line views/components/Form.html:91
		} else {
//line views/components/Form.html:91
			qw422016.N().S(`<label><input type="checkbox" name="`)
//line views/components/Form.html:92
			qw422016.E().S(key)
//line views/components/Form.html:92
			qw422016.N().S(`" value="`)
//line views/components/Form.html:92
			qw422016.E().S(opt)
//line views/components/Form.html:92
			qw422016.N().S(`" />`)
//line views/components/Form.html:92
			qw422016.N().S(` `)
//line views/components/Form.html:92
			qw422016.E().S(title)
//line views/components/Form.html:92
			qw422016.N().S(`</label>`)
//line views/components/Form.html:93
		}
//line views/components/Form.html:94
	}
//line views/components/Form.html:95
}

//line views/components/Form.html:95
func WriteFormCheckbox(qq422016 qtio422016.Writer, key string, values []string, opts []string, titles []string, indent int) {
//line views/components/Form.html:95
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Form.html:95
	StreamFormCheckbox(qw422016, key, values, opts, titles, indent)
//line views/components/Form.html:95
	qt422016.ReleaseWriter(qw422016)
//line views/components/Form.html:95
}

//line views/components/Form.html:95
func FormCheckbox(key string, values []string, opts []string, titles []string, indent int) string {
//line views/components/Form.html:95
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Form.html:95
	WriteFormCheckbox(qb422016, key, values, opts, titles, indent)
//line views/components/Form.html:95
	qs422016 := string(qb422016.B)
//line views/components/Form.html:95
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Form.html:95
	return qs422016
//line views/components/Form.html:95
}
