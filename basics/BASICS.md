# Go言語 基礎編ガイド 📚

このガイドでは、Go言語の基礎を段階的に学習します。各セクションには具体的なコード例と演習問題が含まれています。

## 📖 ブログ記事との連動について

このリポジトリは、ブログ記事と連動した実践的な学習環境を提供しています。

**学習の進め方：**
1. まず対応するブログ記事を読んで概念を理解
2. このガイドのコード例を実際に動かして確認
3. 演習問題で理解度をチェック

### 公開済みのブログ記事

#### 📝 第1回：Go言語入門「はじめの一歩から学ぶ基礎文法」
- **記事URL**: [https://my-studies.org/introduction-to-go-language-learn-basic-grammar-from-beginning/](https://my-studies.org/introduction-to-go-language-learn-basic-grammar-from-beginning/)
- **対応セクション**: [2. 変数と型](#2-変数と型)
- **学べる内容**: パッケージの概念、変数宣言（`var`と`:=`）、基本データ型、`fmt.Println`

#### 📝 第2回：Go言語で条件分岐とループをマスターしよう！制御構文の基礎
- **記事URL**:[[2025年12月6日公開](https://my-studies.org/go-language-learn-basic-grammar-from-beginning02/)]
- **対応セクション**: [3. 制御構文](#3-制御構文)
- **学べる内容**: if文、switch文、for文、break/continue

---

## 目次

1. [Hello World](#1-hello-world)
2. [変数と型](#2-変数と型) ✅ 第1回記事対応
3. [制御構文](#3-制御構文) ✅ 第2回記事対応
4. [関数](#4-関数)
5. [配列とスライス](#5-配列とスライス)
6. [マップ](#6-マップ)

---

## 1. Hello World

### ファイル構成

```
basics/01-hello/
├── main.go
└── README.md
```

### `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to Go Programming!")
}
```

### 実行方法

```bash
cd basics/01-hello
go run main.go
```

### 学習ポイント

- `package main`: 実行可能なプログラムはmainパッケージに属する
- `import "fmt"`: 標準ライブラリのインポート
- `func main()`: プログラムのエントリーポイント
- `fmt.Println()`: 標準出力に文字列を表示

### 演習問題

1. 自分の名前を表示するプログラムを作成
2. 複数行にわたるメッセージを表示
3. `fmt.Printf()` を使って「Hello, %s!」の形式で表示

---

## 2. 変数と型

> 📖 **関連記事**: [第1回 - Go言語入門「はじめの一歩から学ぶ基礎文法」](https://my-studies.org/introduction-to-go-language-learn-basic-grammar-from-beginning/)

### ファイル構成

```
basics/02-variables-and-types/
├── main.go
├── types.go
└── README.md
```

### `main.go`

```go
package main

import "fmt"

func main() {
    // 変数宣言の方法
    
    // 方法1: var キーワード
    var name string = "太郎"
    var age int = 25
    
    // 方法2: 型推論
    var city = "東京"
    
    // 方法3: 短縮宣言（関数内のみ）
    country := "日本"
    
    // 複数変数の同時宣言
    var x, y, z int = 1, 2, 3
    
    // 定数の宣言
    const Pi = 3.14159
    const AppName = "GoLearning"
    
    // 出力
    fmt.Println("名前:", name)
    fmt.Println("年齢:", age)
    fmt.Println("都市:", city)
    fmt.Println("国:", country)
    fmt.Printf("座標: (%d, %d, %d)\n", x, y, z)
    fmt.Println("円周率:", Pi)
}
```

### `types.go`

```go
package main

import "fmt"

func demonstrateTypes() {
    // 基本型
    
    // 整数型
    var intVar int = 42
    var int8Var int8 = 127
    var int16Var int16 = 32767
    var int32Var int32 = 2147483647
    var int64Var int64 = 9223372036854775807
    
    // 符号なし整数型
    var uintVar uint = 42
    var uint8Var uint8 = 255
    
    // 浮動小数点型
    var float32Var float32 = 3.14
    var float64Var float64 = 3.141592653589793
    
    // 真偽値型
    var boolVar bool = true
    
    // 文字列型
    var stringVar string = "こんにちは"
    
    // 文字型（rune = int32のエイリアス）
    var runeVar rune = 'あ'
    
    // バイト型（byte = uint8のエイリアス）
    var byteVar byte = 65 // 'A'
    
    // 型の表示
    fmt.Printf("int: %d, 型: %T\n", intVar, intVar)
    fmt.Printf("float64: %f, 型: %T\n", float64Var, float64Var)
    fmt.Printf("bool: %t, 型: %T\n", boolVar, boolVar)
    fmt.Printf("string: %s, 型: %T\n", stringVar, stringVar)
    fmt.Printf("rune: %c, 型: %T\n", runeVar, runeVar)
    fmt.Printf("byte: %c, 型: %T\n", byteVar, byteVar)
    
    // ゼロ値（初期化されない変数のデフォルト値）
    var zeroInt int       // 0
    var zeroFloat float64 // 0.0
    var zeroBool bool     // false
    var zeroString string // ""
    
    fmt.Println("\nゼロ値:")
    fmt.Printf("int: %d, float64: %f, bool: %t, string: '%s'\n", 
        zeroInt, zeroFloat, zeroBool, zeroString)
}
```

### 実行方法

```bash
cd basics/02-variables-and-types
go run main.go
go run types.go
```

### 演習問題

1. 自分の情報（名前、年齢、住所）を変数に格納して表示
2. 異なる型の変数を使った計算プログラムを作成
3. 定数を使って消費税率を定義し、価格計算を実装

---

## 3. 制御構文

> 📖 **関連記事**: 第2回 - Go言語で条件分岐とループをマスターしよう！制御構文の基礎

### ファイル構成

```
basics/03-control-flow/
├── main.go
└── README.md
```

### `main.go`

記事のすべてのサンプルコードを統合した実行可能なプログラムです。

```go
package main

import "fmt"

func main() {
    fmt.Println("=== Go言語 制御構文の学習 ===\n")

    // 1. if文の基本
    fmt.Println("【1. if文の基本】")
    age := 20
    if age >= 18 {
        fmt.Println("成人です")
    }
    
    // ... (他のコードは basics/03-control-flow/main.go を参照)
}
```

完全なコードは `basics/03-control-flow/` ディレクトリを参照してください。

### 個別のサンプルファイル（オプション）

学習を細かく分けたい場合は、以下のように分割することもできます：

```
basics/03-control-flow/
├── if.go          # if文の例
├── for.go         # forループの例
├── switch.go      # switch文の例
├── main.go        # 統合版
└── README.md
```

### `if.go`

```go
package main

import "fmt"

func main() {
    // 基本的なif文
    age := 20
    if age >= 20 {
        fmt.Println("成人です")
    }
    
    // if-else文
    score := 75
    if score >= 80 {
        fmt.Println("優秀です")
    } else {
        fmt.Println("もう少し頑張りましょう")
    }
    
    // if-else if-else文
    temperature := 25
    if temperature >= 30 {
        fmt.Println("暑いです")
    } else if temperature >= 20 {
        fmt.Println("快適です")
    } else if temperature >= 10 {
        fmt.Println("涼しいです")
    } else {
        fmt.Println("寒いです")
    }
    
    // 初期化付きif文（スコープ内でのみ有効）
    if num := 10; num%2 == 0 {
        fmt.Println(num, "は偶数です")
    }
    // fmt.Println(num) // エラー: numはスコープ外
}
```

### `for.go`

```go
package main

import "fmt"

func main() {
    // 基本的なforループ
    fmt.Println("1. 基本的なfor:")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // 条件のみのfor（whileのような使い方）
    fmt.Println("\n2. 条件のみのfor:")
    count := 0
    for count < 3 {
        fmt.Println("カウント:", count)
        count++
    }
    
    // 無限ループ（breakで抜ける）
    fmt.Println("\n3. 無限ループとbreak:")
    n := 0
    for {
        if n >= 3 {
            break
        }
        fmt.Println("n =", n)
        n++
    }
    
    // continueの使用
    fmt.Println("\n4. continueの使用:")
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue // i=2の時はスキップ
        }
        fmt.Println(i)
    }
    
    // 範囲ループ（range）
    fmt.Println("\n5. rangeを使ったループ:")
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Printf("インデックス: %d, 値: %d\n", index, value)
    }
    
    // インデックスのみ
    fmt.Println("\n6. インデックスのみ:")
    for index := range numbers {
        fmt.Println("インデックス:", index)
    }
    
    // 値のみ（_でインデックスを無視）
    fmt.Println("\n7. 値のみ:")
    for _, value := range numbers {
        fmt.Println("値:", value)
    }
}
```

### `switch.go`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 基本的なswitch
    day := 3
    switch day {
    case 1:
        fmt.Println("月曜日")
    case 2:
        fmt.Println("火曜日")
    case 3:
        fmt.Println("水曜日")
    case 4:
        fmt.Println("木曜日")
    case 5:
        fmt.Println("金曜日")
    case 6, 7:
        fmt.Println("週末")
    default:
        fmt.Println("不正な日")
    }
    
    // 条件式付きswitch
    score := 85
    switch {
    case score >= 90:
        fmt.Println("A評価")
    case score >= 80:
        fmt.Println("B評価")
    case score >= 70:
        fmt.Println("C評価")
    case score >= 60:
        fmt.Println("D評価")
    default:
        fmt.Println("F評価")
    }
    
    // 型スイッチ
    var x interface{} = "hello"
    switch v := x.(type) {
    case int:
        fmt.Println("整数:", v)
    case string:
        fmt.Println("文字列:", v)
    case bool:
        fmt.Println("真偽値:", v)
    default:
        fmt.Println("不明な型")
    }
    
    // 時刻によるswitch
    hour := time.Now().Hour()
    switch {
    case hour < 12:
        fmt.Println("おはようございます")
    case hour < 18:
        fmt.Println("こんにちは")
    default:
        fmt.Println("こんばんは")
    }
}
```

### 演習問題

1. 1から100までの数字のうち、3の倍数と5の倍数を表示するプログラム（FizzBuzz）
2. じゃんけんゲーム（ユーザー入力と判定）を実装
3. 九九の表を表示するプログラム（ネストしたforループ）

詳しい演習問題は `basics/03-control-flow/README.md` を参照してください。

---

## 4. 関数

> 📝 今後のブログ記事で解説予定

### ファイル構成

```
basics/04-functions/
├── basic.go
├── advanced.go
└── README.md
```

### `basic.go`

```go
package main

import "fmt"

// 引数なし、戻り値なし
func sayHello() {
    fmt.Println("こんにちは！")
}

// 引数あり、戻り値なし
func greet(name string) {
    fmt.Println("こんにちは、", name, "さん！")
}

// 引数あり、戻り値あり
func add(a int, b int) int {
    return a + b
}

// 同じ型の引数は省略可能
func multiply(a, b int) int {
    return a * b
}

// 複数の戻り値
func divmod(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// 名前付き戻り値
func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return // 名前付き戻り値は return だけでOK
}

// 可変長引数
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    // 関数の呼び出し
    sayHello()
    greet("太郎")
    
    // 戻り値の受け取り
    result := add(10, 20)
    fmt.Println("10 + 20 =", result)
    
    product := multiply(5, 6)
    fmt.Println("5 × 6 =", product)
    
    // 複数の戻り値
    q, r := divmod(17, 5)
    fmt.Printf("17 ÷ 5 = %d 余り %d\n", q, r)
    
    // 名前付き戻り値
    area, perimeter := rectangle(5, 10)
    fmt.Printf("面積: %d, 周囲: %d\n", area, perimeter)
    
    // 戻り値の一部を無視（_を使用）
    _, remainder := divmod(20, 3)
    fmt.Println("余りのみ:", remainder)
    
    // 可変長引数
    fmt.Println("合計:", sum(1, 2, 3, 4, 5))
    fmt.Println("合計:", sum(10, 20))
}
```

### `advanced.go`

```go
package main

import "fmt"

// 高階関数: 関数を引数に取る
func apply(f func(int, int) int, a, b int) int {
    return f(a, b)
}

// 無名関数とクロージャ
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 再帰関数
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// フィボナッチ数列（再帰）
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// defer文: 関数の終了時に実行
func deferExample() {
    defer fmt.Println("3. 最後に実行")
    fmt.Println("1. 最初に実行")
    fmt.Println("2. 次に実行")
}

func main() {
    // 高階関数の使用
    addFunc := func(a, b int) int { return a + b }
    multiplyFunc := func(a, b int) int { return a * b }
    
    fmt.Println("Apply add:", apply(addFunc, 10, 5))
    fmt.Println("Apply multiply:", apply(multiplyFunc, 10, 5))
    
    // クロージャの使用
    fmt.Println("\nクロージャ:")
    increment := counter()
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
    fmt.Println(increment()) // 3
    
    // 新しいカウンター
    newIncrement := counter()
    fmt.Println(newIncrement()) // 1
    
    // 再帰関数
    fmt.Println("\n階乗:")
    fmt.Println("5! =", factorial(5)) // 120
    
    fmt.Println("\nフィボナッチ数列:")
    for i := 0; i < 10; i++ {
        fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
    }
    
    // defer文
    fmt.Println("\ndefer文の実行順序:")
    deferExample()
}
```

### 演習問題

1. 温度変換関数（摂氏→華氏、華氏→摂氏）を作成
2. 素数判定関数を実装
3. 文字列を反転させる関数を作成
4. 最大公約数を求める関数（ユークリッドの互除法）を実装

---

## 5. 配列とスライス

> 📝 今後のブログ記事で解説予定

### ファイル構成

```
basics/05-arrays-slices/
├── arrays.go
├── slices.go
└── README.md
```

*(以下、既存のコードと同じため省略)*

---

## 6. マップ

> 📝 今後のブログ記事で解説予定

### ファイル構成

```
basics/06-maps/
├── main.go
└── README.md
```

*(以下、既存のコードと同じため省略)*

---

## 学習の進め方

### ステップバイステップガイド

#### 📖 ブログ記事がある場合（推奨）

1. **ブログ記事を読む**
   - まず対応するブログ記事で概念を理解
   - わかりやすい解説とサンプルコードを確認

2. **このガイドのコードを実行**
   - ブログで学んだ内容を実際に動かして確認
   - `go run` コマンドで実行して動作を確認

3. **コードを改変して実験**
   - 値を変えて実行
   - 新しい機能を追加
   - エラーを意図的に起こして学習

4. **演習問題に挑戦**
   - 各セクションの演習問題を解く
   - 理解度をチェック

#### 📚 ブログ記事がまだない場合

1. **このガイドのコードを読む**
   - コード例とコメントから学習
   
2. **実際に動かして確認**
   - `go run` コマンドで実行
   
3. **演習問題で理解を深める**

### 各ディレクトリの作成

```bash
mkdir -p basics/01-hello
mkdir -p basics/02-variables-and-types
mkdir -p basics/03-control-flow
mkdir -p basics/04-functions
mkdir -p basics/05-arrays-slices
mkdir -p basics/06-maps
```

### 実行コマンド一覧

```bash
# Hello World
cd basics/01-hello && go run main.go

# 変数と型（第1回記事対応）
cd basics/02-variables-and-types && go run main.go
cd basics/02-variables-and-types && go run types.go

# 制御構文（第2回記事対応）
cd basics/03-control-flow && go run main.go
# または個別に
cd basics/03-control-flow && go run if.go
cd basics/03-control-flow && go run for.go
cd basics/03-control-flow && go run switch.go

# 関数
cd basics/04-functions && go run basic.go
cd basics/04-functions && go run advanced.go

# 配列とスライス
cd basics/05-arrays-slices && go run arrays.go
cd basics/05-arrays-slices && go run slices.go

# マップ
cd basics/06-maps && go run main.go
```

---

*(総合演習問題、よくあるエラーと解決方法、次のステップ、参考リソース、チェックリスト、まとめは既存のものと同じため省略)*

---

## まとめ

お疲れ様でした！この基礎編では以下のことを学びました：

1. **Hello World**: Goプログラムの基本構造
2. **変数と型**: データの扱い方と型システム（✅ 第1回記事で解説済み）
3. **制御構文**: プログラムのフロー制御（✅ 第2回記事で解説済み）
4. **関数**: コードの再利用と構造化
5. **配列とスライス**: データのコレクション管理
6. **マップ**: キー・バリュー型データの扱い方

これらの知識は、Go言語プログラミングの基盤となります。しっかりと理解してから次のステップに進みましょう！

**ブログ記事も合わせて読むことで、より深い理解が得られます！** 📖

質問や不明点があれば、[Issues](https://github.com/Rocky-Seven/go-learning/issues)で気軽に聞いてください。

Happy Coding! 🎉
