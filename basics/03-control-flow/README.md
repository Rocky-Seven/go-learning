# 03. 制御構文（Control Flow）

このディレクトリでは、Go言語の制御構文（条件分岐とループ）を学習します。

## 📖 対応するブログ記事

**「Go言語で条件分岐とループをマスターしよう！制御構文の基礎」**

制御構文を理解することで、プログラムに「判断」と「繰り返し」の機能を持たせることができます。

## 🎯 学習内容

### 1. if文による条件分岐
- 基本的なif文の構文
- if-else文とif-else if-else文
- 初期化付きif文（Goの特徴的な機能）

### 2. switch文による多分岐
- 基本的なswitch文
- 複数条件のcase
- 条件式を使わないswitch
- fallthroughキーワード

### 3. for文によるループ
- 伝統的なfor文
- whileスタイルのfor文
- 無限ループとbreak
- rangeを使った反復処理

### 4. ループの制御
- breakでループを抜ける
- continueで次の反復に進む

## 💻 実行方法

```bash
# このディレクトリに移動
cd basics/03-control-flow

# プログラムを実行
go run main.go
```

## 🔍 コードのポイント

### if文の特徴

```go
// Goではカッコ()が不要
if age >= 18 {
    fmt.Println("成人です")
}

// 初期化付きif文（変数のスコープがif文内に限定される）
if num := 10; num%2 == 0 {
    fmt.Println("偶数です")
}
// numはここでは使えない
```

### switchの特徴

```go
// breakが不要（自動的に終了）
switch day {
case "月曜日":
    fmt.Println("週の始まり")
case "土曜日", "日曜日":  // 複数条件をカンマで区切る
    fmt.Println("休日")
}

// 条件式を使わないswitch
switch {
case score >= 90:
    fmt.Println("優秀")
case score >= 70:
    fmt.Println("良好")
}
```

### forループの柔軟性

```go
// 伝統的なfor文
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// whileスタイル（条件のみ）
for count < 10 {
    count++
}

// 無限ループ
for {
    if condition {
        break
    }
}

// rangeで反復処理
fruits := []string{"りんご", "バナナ", "みかん"}
for index, fruit := range fruits {
    fmt.Printf("%d: %s\n", index, fruit)
}

// インデックスを無視
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

## ✏️ 演習問題

理解を深めるために、以下の課題に挑戦してみましょう。

### 演習1: FizzBuzz
1から30までの数字を出力するプログラムを作成してください。
- 3の倍数のときは数字の代わりに「Fizz」
- 5の倍数のときは「Buzz」
- 15の倍数のときは「FizzBuzz」

<details>
<summary>ヒント</summary>

```go
for i := 1; i <= 30; i++ {
    if i%15 == 0 {
        // 15の倍数の処理
    } else if i%3 == 0 {
        // 3の倍数の処理
    } else if i%5 == 0 {
        // 5の倍数の処理
    } else {
        // それ以外の処理
    }
}
```
</details>

### 演習2: 成績判定プログラム
5人の学生の点数を配列で用意し、以下を出力してください：
- 各学生の評価（A, B, C, D）
- クラス全体の平均点
- 合格者数（60点以上）

<details>
<summary>ヒント</summary>

```go
scores := []int{85, 92, 58, 73, 45}
total := 0
passed := 0

for _, score := range scores {
    // 評価の判定
    // 合計の計算
    // 合格判定
}
```
</details>

### 演習3: 九九の表
九九の表を出力するプログラムを作成してください。

期待される出力例：
```
1x1=1  1x2=2  1x3=3  ...
2x1=2  2x2=4  2x3=6  ...
...
```

<details>
<summary>ヒント</summary>

```go
for i := 1; i <= 9; i++ {
    for j := 1; j <= 9; j++ {
        fmt.Printf("%dx%d=%d\t", i, j, i*j)
    }
    fmt.Println()
}
```
</details>

## 🚀 次のステップ

制御構文をマスターしたら、次は関数について学びましょう。

👉 次のディレクトリ: `basics/04-functions/`

## 💡 Tips

- **if文のカッコは省略できない**: 1行でも`{}`が必要です
- **switchにbreakは不要**: 自動的に終了します
- **forだけでループは十分**: whileやdo-whileは必要ありません
- **rangeは便利**: スライスや配列の反復処理に最適です

## 📚 参考資料

- [A Tour of Go - Flow control statements](https://go.dev/tour/flowcontrol/1)
- [Effective Go - Control structures](https://go.dev/doc/effective_go#control-structures)

---

🎉 制御構文を理解すれば、プログラムに複雑な振る舞いを持たせることができます！

質問があれば、リポジトリのIssueでお気軽にどうぞ。