package stringsplit

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	const str = `aaa,"bb,b"ccc{ddd,},eee`
	const delimiter = ","
	const begin1 = "{"
	const end1 = "}"
	const begin2 = "\""
	const end2 = "\""

	t.Logf("test configure\n%-10s: %s\n%-10s: %s,\n%-10s: %s\n%-10s: %s\n%-10s: %s\n%-10s: %s\n", "src", str, "delimiter", delimiter, "begin1", begin1, "end1", end1, "begin2", begin2, "end2", end2)

	conf := NewConfiguration(delimiter)
	conf.Append(begin1, end1)
	conf.Append(begin2, end2)
	res, err := Execute(str, conf)
	if err != nil {
		t.Fatal(err)
	}

	for index, item := range res {
		t.Logf("result %d: %s\n", index+1, item)
	}

	expect := []string{"aaa", `"bb,b"ccc{ddd,}`, "eee"}

	if len(res) != len(expect) {
		t.Fatal("Could not match result length")
	}

	for i := 0; i < len(res); i++ {
		e := expect[i]
		a := res[i]
		if a != e {
			t.Errorf("Could not match string\nexpect: %s\nactual: %s\n", e, a)
		}
	}
}

func TestExecuteSimple(t *testing.T) {
	const str = `aaa,bb,bccc{ddd,},eee`
	const delimiter = ","
	const begin = "{"
	const end = "}"

	t.Logf("test configure\n%-10s: %s\n%-10s: %s\n%-10s: %s\n%-10s: %s\n", "src", str, "delimiter", delimiter, "begin", begin, "end", end)

	res, err := ExecuteSimple(str, delimiter, begin, end)
	if err != nil {
		t.Fatal(err)
	}

	for index, item := range res {
		t.Logf("result %d: %s\n", index+1, item)
	}

	expect := []string{"aaa", "bb", "bccc{ddd,}", "eee"}

	if len(res) != len(expect) {
		t.Fatal("Could not match result length")
	}

	for i := 0; i < len(res); i++ {
		e := expect[i]
		a := res[i]
		if a != e {
			t.Errorf("Could not match string\nexpect: %s\nactual: %s\n", e, a)
		}
	}
}
