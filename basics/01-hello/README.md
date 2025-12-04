# 01. Hello World

最初のGoプログラムへようこそ！

このディレクトリでは、Goプログラミングの基本となる「Hello World」プログラムを学習します。

## 🎯 学習内容

- Goプログラムの基本構造
- `package main`の役割
- `import`文によるパッケージのインポート
- `func main()`のエントリーポイント
- `fmt.Println()`による出力

## 💻 実行方法

```bash
# このディレクトリに移動
cd basics/01-hello

# プログラムを実行
go run main.go
```

## 📝 コードの解説

```go
package main
```
- すべてのGoプログラムはパッケージに属します
- 実行可能なプログラムは`main`パッケージである必要があります

```go
import "fmt"
```
- `fmt`パッケージをインポートします
- `fmt`は「format」の略で、入出力に関する機能を提供します

```go
func main() {
    // ...
}
```
- プログラムのエントリーポイント（開始地点）です
- Goプログラムは必ず`main`関数から実行されます

```go
fmt.Println("Hello, World!")
```
- 標準出力に文字列を表示します
- `Println`は「Print line」の略で、末尾に改行を追加します

## ✏️ 演習問題

理解を深めるために、以下の課題に挑戦してみましょう。

### 演習1: 自己紹介プログラム
自分の名前、年齢、出身地を表示するプログラムを作成してください。

<details>
<summary>ヒント</summary>

```go
fmt.Println("名前: 田中太郎")
fmt.Println("年齢: 25歳")
fmt.Println("出身: 三重県")
```
</details>

### 演習2: 複数行のメッセージ
3行以上のメッセージを表示するプログラムを作成してください。

### 演習3: fmt.Printf()を使う
`fmt.Printf()`を使って、「Hello, %s!」の形式で名前を表示してください。

<details>
<summary>ヒント</summary>

```go
name := "太郎"
fmt.Printf("Hello, %s!\n", name)
```
</details>

## 🔍 重要なポイント

### Goプログラムの3つの基本要素

1. **パッケージ宣言**: `package main`
2. **インポート**: `import "fmt"`
3. **main関数**: `func main() { ... }`

この3つがあれば、Goプログラムは実行できます！

### よくあるエラー

```go
// ❌ 誤り: パッケージ宣言がない
import "fmt"
func main() {
    fmt.Println("Hello")
}

// ✅ 正しい
package main
import "fmt"
func main() {
    fmt.Println("Hello")
}
```

```go
// ❌ 誤り: main関数がない
package main
import "fmt"
// エラー: プログラムが実行できない

// ✅ 正しい
package main
import "fmt"
func main() {
    fmt.Println("Hello")
}
```

## 🚀 次のステップ

Hello Worldプログラムを理解したら、次は変数とデータ型について学びましょう。

👉 次のディレクトリ: `basics/02-variables-and-types/`

## 💡 豆知識

**なぜ「Hello, World!」なのか？**

「Hello, World!」は、1978年に出版されたBrian KernighanとDennis Ritchieの著書「The C Programming Language」で使用されて以来、プログラミング学習の伝統的な最初のプログラムとなっています。

---

おめでとうございます！あなたは最初のGoプログラムを実行しました 🎉

次のステップに進んで、さらにGoの世界を探検しましょう！