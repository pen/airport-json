# airport-json

macOSの`airport -I`をJSONで出力する

## 使用例

```bash
$ export SSID=`airport-json | jq -r '.SSID'`
```

```bash
$ airport-json -b -m | jq -C | lv
```

## オプション

- `-x` : `airport -Ix`の結果をJSONで出力
- `-b` : `airport -I`と`airport -Ix`の両方を実行
- `-m` : JSONのトップレベルのキーの有無を切替

## インストール

### homebrew

```shell
$ brew install pen/tap/airport-json
```

### go install

```shell
$ go install github.com/pen/airport-json/cmd/airport-json@latest
```

### ダウンロード

[Releases](https://github.com/pen/airport-json/releases)

## 注意

`airport -I`の出力で、JSONのキーとして不便そうな項目名は以下の変換をしている:

| airport       | airport-json  |
|:---           |:---           |
| 802.11 auth   | IEEE80211Auth |
| link auth     | linkAuth      |
| op mode       | opMode        |
