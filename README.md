# go-stringsplit

[![Go](https://github.com/rssh-jp/go-stringsplit/actions/workflows/go.yml/badge.svg)](https://github.com/rssh-jp/go-stringsplit/actions/workflows/go.yml)

区切り文字を指定して文字列を分割するライブラリ。  
特定のセクション（括弧内・クォート内など）では区切らないよう制御できる。

## インストール

```bash
go get github.com/rssh-jp/go-stringsplit
```

## 使い方

### 基本的な使用例

```go
package main

import (
	"log"

	"github.com/rssh-jp/go-stringsplit"
)

func main() {
	const str = `aaa,"bb,b"ccc{ddd,},eee`
	const delimiter = ","

	conf := stringsplit.NewConfiguration(delimiter)
	conf.Append("{", "}")   // {} 内は区切らない
	conf.Append(`"`, `"`)  // "" 内は区切らない

	res, err := stringsplit.Execute(str, conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res) // [aaa "bb,b"ccc{ddd,} eee]
}
```

### シンプルな使用例（セクション1つ）

```go
res, err := stringsplit.ExecuteSimple(str, ",", "{", "}")
```

## API リファレンス

### 関数

| 関数 | シグネチャ | 説明 |
|------|-----------|------|
| `NewConfiguration` | `func NewConfiguration(delimiter string) Configuration` | 区切り文字を指定して設定を生成する |
| `Execute` | `func Execute(str string, config Configuration) ([]string, error)` | 設定に従って文字列を分割する |
| `ExecuteSimple` | `func ExecuteSimple(str, delimiter, begin, end string) ([]string, error)` | セクション1つを指定して文字列を分割する |

### `Configuration` メソッド

| メソッド | シグネチャ | 説明 |
|---------|-----------|------|
| `Append` | `func (c *Configuration) Append(begin, end string)` | スキップするセクションの開始・終了文字列を追加する |

## 動作の詳細

- `config.Delimiter` で指定した文字列が区切り文字となる。
- `config.Append(begin, end)` で登録したセクション内（`begin` から `end` の間）では、区切り文字が出現しても分割しない。
- セクションは複数登録可能。
- 同一文字列を `begin` と `end` に指定することで、ダブルクォートのような対称セクションにも対応できる。

## 開発

```bash
# ビルド確認
go build ./...

# テスト実行
go test ./...

# 静的解析
go vet ./...
```

