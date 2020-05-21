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
