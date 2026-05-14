package stringsplit_test

import (
	"testing"

	"github.com/rssh-jp/go-stringsplit"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		delimiter string
		sections  [][2]string
		expect    []string
	}{
		{
			name:      "空文字列",
			str:       "",
			delimiter: ",",
			expect:    []string{""},
		},
		{
			name:      "区切り文字なし（全体を1要素として返す）",
			str:       "aaa",
			delimiter: ",",
			expect:    []string{"aaa"},
		},
		{
			name:      "区切り文字のみ",
			str:       ",",
			delimiter: ",",
			expect:    []string{"", ""},
		},
		{
			name:      "通常の分割",
			str:       "aaa,bbb,ccc",
			delimiter: ",",
			expect:    []string{"aaa", "bbb", "ccc"},
		},
		{
			name:      "末尾に区切り文字（末尾の空要素を含む）",
			str:       "aaa,bbb,",
			delimiter: ",",
			expect:    []string{"aaa", "bbb", ""},
		},
		{
			name:      "先頭に区切り文字（先頭の空要素を含む）",
			str:       ",aaa,bbb",
			delimiter: ",",
			expect:    []string{"", "aaa", "bbb"},
		},
		{
			name:      "連続した区切り文字（空要素を含む）",
			str:       "aaa,,bbb",
			delimiter: ",",
			expect:    []string{"aaa", "", "bbb"},
		},
		{
			name:      "複数文字の区切り文字",
			str:       "aaa, bbb, ccc",
			delimiter: ", ",
			expect:    []string{"aaa", "bbb", "ccc"},
		},
		{
			name:      "1文字のセクション内は区切らない",
			str:       `aaa,"bb,b"ccc{ddd,},eee`,
			delimiter: ",",
			sections:  [][2]string{{"{", "}"}, {`"`, `"`}},
			expect:    []string{"aaa", `"bb,b"ccc{ddd,}`, "eee"},
		},
		{
			name:      "セクション末尾直後の区切り文字",
			str:       "{aaa,bbb},ccc",
			delimiter: ",",
			sections:  [][2]string{{"{", "}"}},
			expect:    []string{"{aaa,bbb}", "ccc"},
		},
		{
			name:      "複数のセクション",
			str:       "{aaa,bbb},{ccc,ddd},eee",
			delimiter: ",",
			sections:  [][2]string{{"{", "}"}},
			expect:    []string{"{aaa,bbb}", "{ccc,ddd}", "eee"},
		},
		{
			name:      "複数文字の begin/end（セクション内は区切らない）",
			str:       "aaa<<bb,b>>ccc,ddd",
			delimiter: ",",
			sections:  [][2]string{{"<<", ">>"}},
			expect:    []string{"aaa<<bb,b>>ccc", "ddd"},
		},
		{
			name:      "複数文字の begin/end（末尾区切り文字）",
			str:       "aaa<<bb,b>>ccc,",
			delimiter: ",",
			sections:  [][2]string{{"<<", ">>"}},
			expect:    []string{"aaa<<bb,b>>ccc", ""},
		},
		{
			name:      "セクションが文字列末尾まで続く（end なし）は分割する",
			str:       "aaa,{bbb,ccc",
			delimiter: ",",
			sections:  [][2]string{{"{", "}"}},
			expect:    []string{"aaa", "{bbb", "ccc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := stringsplit.NewConfiguration(tt.delimiter)
			for _, sec := range tt.sections {
				conf.Append(sec[0], sec[1])
			}
			got, err := stringsplit.Execute(tt.str, conf)
			if err != nil {
				t.Fatalf("stringsplit.Execute returned error: %v", err)
			}
			if len(got) != len(tt.expect) {
				t.Fatalf("length mismatch: expect %d, got %d\n  expect: %v\n  got:    %v",
					len(tt.expect), len(got), tt.expect, got)
			}
			for i := range got {
				if got[i] != tt.expect[i] {
					t.Errorf("index %d: expect %q, got %q", i, tt.expect[i], got[i])
				}
			}
		})
	}
}

func TestExecuteSimple(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		delimiter string
		begin     string
		end       string
		expect    []string
	}{
		{
			name:      "セクション内は区切らない",
			str:       "aaa,bb,bccc{ddd,},eee",
			delimiter: ",",
			begin:     "{",
			end:       "}",
			expect:    []string{"aaa", "bb", "bccc{ddd,}", "eee"},
		},
		{
			name:      "マッチするセクションなし",
			str:       "aaa,bbb,ccc",
			delimiter: ",",
			begin:     "[",
			end:       "]",
			expect:    []string{"aaa", "bbb", "ccc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringsplit.ExecuteSimple(tt.str, tt.delimiter, tt.begin, tt.end)
			if err != nil {
				t.Fatalf("stringsplit.ExecuteSimple returned error: %v", err)
			}
			if len(got) != len(tt.expect) {
				t.Fatalf("length mismatch: expect %d, got %d\n  expect: %v\n  got:    %v",
					len(tt.expect), len(got), tt.expect, got)
			}
			for i := range got {
				if got[i] != tt.expect[i] {
					t.Errorf("index %d: expect %q, got %q", i, tt.expect[i], got[i])
				}
			}
		})
	}
}

