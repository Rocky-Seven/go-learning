package main

import "fmt"

func main() {
	fmt.Println("=== Go言語 変数と型の学習 ===\n")

	// 1. varによる変数宣言
	fmt.Println("【1. varによる変数宣言】")
	var name string = "太郎"
	var age int = 25
	var isStudent bool = true

	fmt.Println("名前:", name)
	fmt.Println("年齢:", age)
	fmt.Println("学生:", isStudent)
	fmt.Println()

	// 2. 型推論を使った宣言
	fmt.Println("【2. 型推論】")
	var city = "三重" // stringと推論される
	var count = 100   // intと推論される

	fmt.Println("都市:", city)
	fmt.Println("カウント:", count)
	fmt.Println()

	// 3. 短縮宣言（:=）
	fmt.Println("【3. 短縮宣言（:=）】")
	country := "日本"
	temperature := 25.5
	hasLicense := false

	fmt.Println("国:", country)
	fmt.Println("気温:", temperature, "度")
	fmt.Println("免許保持:", hasLicense)
	fmt.Println()

	// 4. 基本的なデータ型
	fmt.Println("【4. 基本的なデータ型】")

	// 整数（int）
	population := 1000000
	fmt.Printf("人口: %d (型: %T)\n", population, population)

	// 文字列（string）
	message := "こんにちは、Go言語！"
	fmt.Printf("メッセージ: %s (型: %T)\n", message, message)

	// 真偽値（bool）
	isActive := true
	fmt.Printf("アクティブ: %t (型: %T)\n", isActive, isActive)

	// 浮動小数点数（float64）
	price := 1980.50
	fmt.Printf("価格: %.2f円 (型: %T)\n", price, price)
	fmt.Println()

	// 5. 複数変数の同時宣言
	fmt.Println("【5. 複数変数の同時宣言】")
	var x, y, z int = 10, 20, 30
	fmt.Printf("座標: (%d, %d, %d)\n", x, y, z)

	a, b, c := "A", "B", "C"
	fmt.Printf("文字: %s, %s, %s\n", a, b, c)
	fmt.Println()

	// 6. ゼロ値
	fmt.Println("【6. ゼロ値（初期値を指定しない場合）】")
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	fmt.Printf("int: %d\n", zeroInt)
	fmt.Printf("float64: %f\n", zeroFloat)
	fmt.Printf("bool: %t\n", zeroBool)
	fmt.Printf("string: '%s'\n", zeroString)
	fmt.Println()

	// 7. 定数
	fmt.Println("【7. 定数】")
	const Pi = 3.14159
	const AppName = "Go学習アプリ"

	fmt.Println("円周率:", Pi)
	fmt.Println("アプリ名:", AppName)

	// 定数を使った計算
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("半径%.1fの円の面積: %.2f\n", radius, area)
	fmt.Println()

	// 8. 実践例：プロフィール表示
	fmt.Println("【8. 実践例：プロフィール】")
	fmt.Println("=== プロフィール ===")
	
	var profileName string = "田中太郎"
	profileAge := 25
	height := 175.5
	isEmployed := true

	fmt.Println("名前:", profileName)
	fmt.Println("年齢:", profileAge, "歳")
	fmt.Println("身長:", height, "cm")
	fmt.Println("就業状況:", isEmployed)

	fmt.Println("\n=== 学習完了！ ===")
	fmt.Println("次は制御構文（if, for, switch）を学びましょう")
}
