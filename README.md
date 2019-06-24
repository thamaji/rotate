rotate
====

雑にログローテーションするだけのツール。単機能の小さなシングルバイナリがほしかった。

標準入力のテキストを１行ずつ読んで、`-p` 引数で指定したパターンのファイルに書き込む。

`-p` の指定方法は golang の time パッケージのもの。

## Install

```
$ sudo curl -o /usr/local/bin/rotate -fSsL https://github.com/thamaji/rotate/releases/download/v1.0.1/rotate
$ sudo chmod +x /usr/local/bin/rotate
```

## Example

```
$ yes | ./rotate -p 2006-01-02-15-04-05.log -o log
$ ls log
2019-06-24-19-42-05.log  2019-06-24-19-42-07.log  2019-06-24-19-42-09.log
2019-06-24-19-42-06.log  2019-06-24-19-42-08.log
```
