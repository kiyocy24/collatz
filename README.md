# collatz
コラッツ操作をするCLIツール

## 使い方
https://github.com/kiyocy24/collatz/releases の最新からOSとCPUに対応したバイナリをダウンロードしてCLIを叩くだけ

```
# Mac
./collatz -n 10

# Linux
./collatz -n 10

# Windows
.\collatz.exe -n 10
```

### コマンドオプション
```
   --num value, -n value     num (default: 0)            // コラッツ操作の入力値
   --output value, -o value  output file path            // 出力ファイル名、未入力ならコンソール出力
   --end value, -e value     start num (default: 0)      // nからendまで数値に対してコラッツ操作を実行
   --help, -h                show help (default: false)  // ヘルプ
```

## ソースから実行
```
go run main.go
```

## ビルド
```
go build
```