# gocode

使用 go/parser 分析 go 代碼

go/parser 可用於分析 golang 源碼，這對某些時候非常有用，例如識別出包的導出內容然後自動將它們導入到腳本以供腳本調用。

這段代碼是本喵使用 go/parser 進行源碼分析的一些處理代碼

# example

example 顯示了如何調用本庫提供的代碼

# gocode/gocode

gocode 檔案夾下是一個二進制程序，它可以接收檔案或檔案夾，並分析裏面的golang代碼最終將 golang 中定義的聲明進行打印

```
./gocode -h
```

```
./gocode file main.go
```

```
./gocode dir /opt/google/go/src/ -r
```
