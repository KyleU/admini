package sandbox

type Sandbox struct {
	Key string
	Run func() (interface{}, error)
}

var testbed = &Sandbox{Key: "testbed", Run: func() (interface{}, error) {
	return "OK", nil
}}

type Sandboxes []*Sandbox

func (s Sandboxes) Get(key string) *Sandbox {
	for _, v := range s {
		if v.Key == key {
			return v
		}
	}
	return nil
}

var AllSandboxes = Sandboxes{testbed}
