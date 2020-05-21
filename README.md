# go-stringsplit
文字列を分割する
# usage
```
package main

import (
	"log"

	"github.com/rssh-jp/go-stringsplit"
)

func main() {
	const str = `aaa,"bb,b"ccc{ddd,},eee`
	const delimiter = ","
	const begin1 = "{"
	const end1 = "}"
	const begin2 = "\""
	const end2 = "\""

	conf := stringsplit.NewConfiguration(delimiter)
	conf.Append(begin1, end1)
	conf.Append(begin2, end2)
	res, err := stringsplit.Execute(str, conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)  // [aaa "bb,b"ccc{ddd,} eee]
}
```

# detail
元となる文字列と、それをどの文字列で分割しないのかを指定する。  
それと、`"`内や`/* */`内など特定の文字列内では区切りたくない場合はそれを指定できる。  
`conf := stringsplit.NewConfiguration`の引数に区切りたい文字列を入れ  
`conf.Append`の引数に区切りたくない特定の文字列を入れる
