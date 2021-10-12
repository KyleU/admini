// Code generated by qtc from "Boolean.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/fieldedit/Boolean.html:1
package fieldedit

//line views/components/fieldedit/Boolean.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldedit/Boolean.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldedit/Boolean.html:1
func StreamBoolean(qw422016 *qt422016.Writer, x interface{}, nullable bool, k string) {
//line views/components/fieldedit/Boolean.html:2
	b, isBool := x.(bool)

//line views/components/fieldedit/Boolean.html:3
	if isBool && b {
//line views/components/fieldedit/Boolean.html:3
		qw422016.N().S(`<label class="radiolabel"><input value="true" name="`)
//line views/components/fieldedit/Boolean.html:4
		qw422016.E().S(k)
//line views/components/fieldedit/Boolean.html:4
		qw422016.N().S(`" type="radio" checked="checked" /> true</label>`)
//line views/components/fieldedit/Boolean.html:5
	} else {
//line views/components/fieldedit/Boolean.html:5
		qw422016.N().S(`<label class="radiolabel"><input value="true" name="`)
//line views/components/fieldedit/Boolean.html:6
		qw422016.E().S(k)
//line views/components/fieldedit/Boolean.html:6
		qw422016.N().S(`" type="radio" /> true</label>`)
//line views/components/fieldedit/Boolean.html:7
	}
//line views/components/fieldedit/Boolean.html:9
	if isBool && !b {
//line views/components/fieldedit/Boolean.html:9
		qw422016.N().S(`<label class="radiolabel"><input value="false" name="`)
//line views/components/fieldedit/Boolean.html:10
		qw422016.E().S(k)
//line views/components/fieldedit/Boolean.html:10
		qw422016.N().S(`" type="radio" checked="checked" /> false</label>`)
//line views/components/fieldedit/Boolean.html:11
	} else {
//line views/components/fieldedit/Boolean.html:11
		qw422016.N().S(`<label class="radiolabel"><input value="false" name="`)
//line views/components/fieldedit/Boolean.html:12
		qw422016.E().S(k)
//line views/components/fieldedit/Boolean.html:12
		qw422016.N().S(`" type="radio" /> false</label>`)
//line views/components/fieldedit/Boolean.html:13
	}
//line views/components/fieldedit/Boolean.html:15
	if nullable {
//line views/components/fieldedit/Boolean.html:16
		if x == nil {
//line views/components/fieldedit/Boolean.html:16
			qw422016.N().S(`<label class="radiolabel"><input value="∅" name="`)
//line views/components/fieldedit/Boolean.html:17
			qw422016.E().S(k)
//line views/components/fieldedit/Boolean.html:17
			qw422016.N().S(`" type="radio" checked="checked" /> nil</label>`)
//line views/components/fieldedit/Boolean.html:18
		} else {
//line views/components/fieldedit/Boolean.html:18
			qw422016.N().S(`<label class="radiolabel"><input value="∅" name="`)
//line views/components/fieldedit/Boolean.html:19
			qw422016.E().S(k)
//line views/components/fieldedit/Boolean.html:19
			qw422016.N().S(`" type="radio" /> nil</label>`)
//line views/components/fieldedit/Boolean.html:20
		}
//line views/components/fieldedit/Boolean.html:21
	}
//line views/components/fieldedit/Boolean.html:22
}

//line views/components/fieldedit/Boolean.html:22
func WriteBoolean(qq422016 qtio422016.Writer, x interface{}, nullable bool, k string) {
//line views/components/fieldedit/Boolean.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldedit/Boolean.html:22
	StreamBoolean(qw422016, x, nullable, k)
//line views/components/fieldedit/Boolean.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldedit/Boolean.html:22
}

//line views/components/fieldedit/Boolean.html:22
func Boolean(x interface{}, nullable bool, k string) string {
//line views/components/fieldedit/Boolean.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldedit/Boolean.html:22
	WriteBoolean(qb422016, x, nullable, k)
//line views/components/fieldedit/Boolean.html:22
	qs422016 := string(qb422016.B)
//line views/components/fieldedit/Boolean.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldedit/Boolean.html:22
	return qs422016
//line views/components/fieldedit/Boolean.html:22
}
