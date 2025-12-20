# 04. 関数（Functions）

このディレクトリでは、Go言語の関数について学習します。

## 📖 対応するブログ記事: [2025.12.20公開](https://my-studies.org/go-language-learn-basic-grammar-from-beginning03/)

**「Go言語入門（3）：関数を完全理解！引数・戻り値・エラーハンドリング入門」**

関数はプログラムの基本的な構成要素であり、コードを整理し再利用可能にするために不可欠です。

## 🎯 学習内容

### 1. 基本的な関数
- 関数の定義方法
- 引数の受け取り方
- 戻り値の返し方

### 2. 複数の戻り値
- Goの特徴的な機能
- 複数の値を同時に返す
- `_`で不要な戻り値を無視

### 3. エラーハンドリング
- `error`型の基本
- `if err != nil`パターン
- エラーメッセージの作成

### 4. 可変長引数
- `...`を使った柔軟な引数
- スライスとしての扱い

### 5. 名前付き戻り値
- 戻り値に名前をつける
- 裸のreturn

## 💻 実行方法

```bash
# このディレクトリに移動
cd basics/04-functions

# プログラムを実行
go run main.go
```

## 📝 コードのポイント

### 関数の基本構造

```go
func 関数名(引数名 型) 戻り値の型 {
    // 処理
    return 戻り値
}
```

### 複数の戻り値（Goの特徴）

```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

// 使用例
quotient, remainder := divmod(17, 5)
```

Goでは関数から複数の値を返すことができます。これは特にエラーハンドリングで重要です。

### エラーハンドリングのパターン

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("0で割ることはできません")
    }
    return a / b, nil
}

// 使用例
result, err := divide(10, 0)
if err != nil {
    fmt.Println("エラー:", err)
    return
}
fmt.Println("結果:", result)
```

**重要なポイント：**
- 関数は`(結果, error)`の形で値を返す
- 呼び出し側で必ず`if err != nil`でチェック
- エラーがない場合は`nil`を返す

これはGoの最も重要な慣習の1つです。

### 可変長引数

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 使用例
sum(1, 2, 3)        // 6
sum(1, 2, 3, 4, 5)  // 15
```

`...`を使うと任意の数の引数を受け取れます。

### 名前付き戻り値

```go
func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // 裸のreturn
}
```

戻り値に名前をつけると、`return`だけで値を返せます。

## ✏️ 演習問題

理解を深めるために、以下の課題に挑戦してみましょう。

### 演習1: 温度変換関数
摂氏から華氏への変換関数と、華氏から摂氏への変換関数を作成してください。

計算式：
- 華氏 = 摂氏 × 9/5 + 32
- 摂氏 = (華氏 - 32) × 5/9

<details>
<summary>ヒント</summary>

```go
func celsiusToFahrenheit(celsius float64) float64 {
    return celsius*9/5 + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5 / 9
}
```
</details>

### 演習2: 素数判定関数
引数として与えられた整数が素数かどうかを判定する関数を作成してください。

<details>
<summary>ヒント</summary>

```go
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
```
</details>

### 演習3: 文字列反転関数
文字列を反転させる関数を作成してください。

例: "hello" → "olleh"

<details>
<summary>ヒント</summary>

```go
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```
</details>

### 演習4: 最大公約数（GCD）
ユークリッドの互除法を使って、2つの整数の最大公約数を求める関数を作成してください。

<details>
<summary>ヒント</summary>

```go
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}
```
</details>

## 🔍 よくあるエラーと解決方法

### エラー1: 戻り値の数が合わない

```go
// ❌ 誤り
func divide(a, b int) (int, error) {
    return a / b  // エラー：戻り値が1つしかない
}

// ✅ 正しい
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### エラー2: エラーチェックを忘れる

```go
// ❌ 悪い例（エラーチェックなし）
result, _ := divide(10, 0)
fmt.Println(result)  // エラーが発生しているのに気づかない

// ✅ 正しい（エラーチェックあり）
result, err := divide(10, 0)
if err != nil {
    fmt.Println("エラー:", err)
    return
}
fmt.Println(result)
```

### エラー3: 関数内で変数を宣言したつもりが再代入

```go
// ❌ 誤り
var result int
result := add(10, 20)  // エラー：resultは既に宣言されている

// ✅ 正しい（再代入）
var result int
result = add(10, 20)

// または新規宣言
result := add(10, 20)
```

## 💡 重要なポイント

### 1. エラーは無視しない

Goでは、エラーは値として扱われます。必ずチェックしましょう。

```go
// 悪い例
result, _ := someFunction()

// 良い例
result, err := someFunction()
if err != nil {
    // エラー処理
}
```

### 2. 関数は小さく、単一責任で

1つの関数は1つのことをするべきです。

```go
// 悪い例：複数の責任
func processUserDataAndSendEmail(user User) error {
    // ユーザーデータの処理
    // メールの送信
    // データベースへの保存
}

// 良い例：責任を分離
func validateUser(user User) error { ... }
func saveUser(user User) error { ... }
func sendWelcomeEmail(user User) error { ... }
```

### 3. 早期リターンを活用

エラーチェックは早期リターンで読みやすくなります。

```go
func process(data string) error {
    if data == "" {
        return errors.New("data is empty")
    }
    
    if len(data) > 100 {
        return errors.New("data too long")
    }
    
    // メインの処理
    return nil
}
```

## 🚀 次のステップ

関数をマスターしたら、次はスライスとマップについて学びましょう。

👉 次のディレクトリ: `basics/05-slices-and-maps/`

## 📚 参考リソース

- [A Tour of Go - Functions](https://go.dev/tour/moretypes/1)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Multiple Return Values](https://gobyexample.com/multiple-return-values)
- [Go by Example - Errors](https://gobyexample.com/errors)

---

🎉 関数を理解すれば、コードの再利用性と保守性が大幅に向上します！

質問があれば、リポジトリのIssueでお気軽にどうぞ。