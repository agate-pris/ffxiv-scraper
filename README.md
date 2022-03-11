# FFXIV Scraper

## 概要

[FINAL FANTASY XIV, The Lodestone](https://jp.finalfantasyxiv.com/lodestone/) から情報を取得し、クエストのコンプリート状況を `quest.csv` に出力します。

**悪用・多用厳禁！！**

多用するとサーバーに負荷が掛かります。

## 使い方

ブラウザで [The Lodestone]((https://jp.finalfantasyxiv.com/lodestone/) にログインし、ウェブ開発ツールを開き、クッキーの `ldst_sess` の値をコピーし、引数 `-s` に、キャラクターIDを引数 `-c` に渡してください。

```bash
./ffxiv-scraper.exe -s <ldst_sess> -c <character-id>
```

ヘルプは以下の方法で表示します。

```bash
./ffxiv-scraper.exe -h
```

## Build & Run

```bash
go build -tags forceposix
go run -tags forceposix .
```

## 既知の問題

- シーズナルクエストには対応していません。
