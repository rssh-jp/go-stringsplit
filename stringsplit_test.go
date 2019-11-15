package stringsplit

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	const str = `aaa,"bb,b"ccc{ddd,},eee`

	conf := NewConfiguration(",")
	conf.Append("{", "}")
	conf.Append("\"", "\"")
	res := Execute(str, conf)

	t.Logf("%+v\n", res)

	expect := []string{"aaa", `"bb,b"ccc{ddd,}`, "eee"}

	if len(res) != len(expect) {
		t.Error("Could not match result length")
	}

	for i := 0; i < len(res); i++ {
		e := expect[i]
		a := res[i]
		if a != e {
			t.Errorf("Could not match string\nexpect: %s\nactual: %s\n", e, a)
		}
	}
}
