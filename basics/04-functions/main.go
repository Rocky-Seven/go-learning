package main

import (
	"errors"
	"fmt"
)

// ===== 1. 基本的な関数 =====

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

// ===== 2. 複数の戻り値 =====

// 割り算の商と余りを返す
func divmod(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// ===== 3. エラーハンドリング =====

// エラーを返す関数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return a / b, nil
}

// ===== 4. 可変長引数 =====

// 任意の数の整数の合計を計算
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 可変長引数の平均を計算
func average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("引数が空です")
	}

	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers)), nil
}

// ===== 5. 名前付き戻り値 =====

// 長方形の面積と周囲を計算
func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 裸のreturn
}

// ===== 6. 実践的な関数：計算機 =====

// 四則演算を行う関数
func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("0で割ることはできません")
		}
		return a / b, nil
	default:
		return 0, errors.New("未対応の演算子です: " + operator)
	}
}

// ===== メイン関数 =====

func main() {
	fmt.Println("=== Go言語 関数の学習 ===\n")

	// 1. 基本的な関数
	fmt.Println("【1. 基本的な関数】")
	sayHello()
	greet("太郎")
	result := add(10, 20)
	fmt.Println("10 + 20 =", result)
	product := multiply(5, 6)
	fmt.Println("5 × 6 =", product)
	fmt.Println()

	// 2. 複数の戻り値
	fmt.Println("【2. 複数の戻り値】")
	q, r := divmod(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余り %d\n", q, r)

	// 戻り値の一部を無視
	_, remainder := divmod(20, 3)
	fmt.Println("余りのみ:", remainder)
	fmt.Println()

	// 3. エラーハンドリング
	fmt.Println("【3. エラーハンドリング】")

	// 正常なケース
	divResult, err := divide(10, 2)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("10 ÷ 2 =", divResult)
	}

	// エラーケース
	divResult, err = divide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("結果:", divResult)
	}
	fmt.Println()

	// 4. 可変長引数
	fmt.Println("【4. 可変長引数】")
	fmt.Println("合計:", sum(1, 2, 3))
	fmt.Println("合計:", sum(1, 2, 3, 4, 5))
	fmt.Println("合計:", sum(10, 20))

	avg, err := average(10, 20, 30, 40, 50)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("平均: %.1f\n", avg)
	}
	fmt.Println()

	// 5. 名前付き戻り値
	fmt.Println("【5. 名前付き戻り値】")
	area, perimeter := rectangle(5, 10)
	fmt.Printf("面積: %d, 周囲: %d\n", area, perimeter)
	fmt.Println()

	// 6. 実践：簡易計算機
	fmt.Println("【6. 実践：簡易計算機】")

	testCases := []struct {
		a, b     float64
		operator string
	}{
		{10, 5, "+"},
		{10, 5, "-"},
		{10, 5, "*"},
		{10, 5, "/"},
		{10, 0, "/"},
		{10, 5, "%"},
	}

	for _, tc := range testCases {
		result, err := calculate(tc.a, tc.b, tc.operator)
		if err != nil {
			fmt.Printf("%.1f %s %.1f = エラー: %v\n",
				tc.a, tc.operator, tc.b, err)
		} else {
			fmt.Printf("%.1f %s %.1f = %.1f\n",
				tc.a, tc.operator, tc.b, result)
		}
	}

	fmt.Println("\n=== 学習完了！ ===")
	fmt.Println("次はスライスとマップを学びましょう")
}