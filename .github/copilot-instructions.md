# go-stringsplit — Copilot Instructions

## プロジェクト概要

区切り文字を指定して文字列を分割するライブラリ。  
特定のセクション（括弧内・クォート内など）では区切り文字を無視する機能を提供する。

## パッケージ構成

```
go-stringsplit/
├── stringsplit.go      # Execute() / ExecuteSimple() — 分割ロジック本体
├── config.go           # Configuration 構造体・NewConfiguration() / Append()
├── section.go          # Section / Sections 構造体・インデックス管理
├── stringsplit_test.go # テスト
└── cmd/stringsplit/
    └── main.go         # 動作確認用 CLI エントリーポイント
```

## コーディングルール

- Go 1.26 以上の構文・機能を使用する
- 外部依存ライブラリは持たない（標準ライブラリのみ）
- エラーは呼び出し元に返す。内部でのログ出力は禁止
- `time.Sleep` など処理とは無関係な待機処理を混在させない

## 設計上の注意点

### 分割ロジック（stringsplit.go）

- `Execute()` は2ステップで動作する
  1. `findSection()` で文字列内のスキップセクション（`begin`〜`end`）の位置を全て走査し `Sections` を構築する
  2. `config.Delimiter` を順に探し、その位置が `Sections.IsInIndex()` に含まれる場合はスキップして分割しない
- `ExecuteSimple()` は `Configuration` に1つのセクションを追加して `Execute()` を呼ぶ薄いラッパー

### セクション管理（section.go）

- `Section` は文字列ペア（`Begin`/`End`）とインデックスペア（`BeginIndex`/`EndIndex`）の両方を兼用する構造体
- `Sections.IsInIndex(index)` でデリミタが保護区間内かを判定する

### Configuration（config.go）

- `Configuration.Delimiter` が分割文字列
- `Configuration.Sections` がスキップ区間の定義リスト
- `FindSectionByBeginString()` は `begin` 文字列から対応するセクション定義を返す

## 新しい機能を追加する場合

- `stringsplit.go` に公開関数を追加する
- `stringsplit_test.go` に対応するテストを追加する
- README.md の API リファレンステーブルも更新する

## テスト・検証コマンド

```bash
# ビルド確認
go build ./...

# テスト実行
go test ./...

# 静的解析
go vet ./...
```

## 依存更新

外部依存なし。`go mod tidy` のみで十分。

```bash
go mod tidy
```
