/*package main

import (
	"errors"

    "fmt"
)

// すべての関数を定義
func divmod(a, b int) (int, int) {
    return a / b, a % b
}


func main() {
    // すべての処理を記述
    _, remainder := divmod(20, 3)
    fmt.Println(remainder)
}



func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("0で割ることはできません")
    }
    return a / b, nil  // エラーがない場合はnilを返す
}

func main() {
    // 正常なケース
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("エラー:", err)
    } else {
        fmt.Println("結果:", result)
    }
    // 出力: 結果: 5
    
    // エラーケース
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("エラー:", err)
    } else {
        fmt.Println("結果:", result)
    }
    // 出力: エラー: 0で割ることはできません
}
	
	func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println("合計:", sum(1, 2, 3))           // 出力: 合計: 6
    fmt.Println("合計:", sum(1, 2, 3, 4, 5))     // 出力: 合計: 15
    fmt.Println("合計:", sum(10, 20))            // 出力: 合計: 30
}
	
	func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // 名前付き戻り値は return だけでOK
}

func main() {
    area, perimeter := rectangle(5, 10)
    fmt.Printf("面積: %d, 周囲: %d\n", area, perimeter)
    // 出力: 面積: 50, 周囲: 30
}
	*/
package main

import (
    "errors"
    "fmt"
)

// 四則演算を行う関数（エラーハンドリング付き）
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

func main() {
    fmt.Println("=== 簡易計算機 ===")
    
    // テストケース
    testCases := []struct {
        a, b     float64
        operator string
    }{
        {10, 5, "+"},
        {10, 5, "-"},
        {10, 5, "*"},
        {10, 5, "/"},
        {10, 0, "/"},  // エラーケース
        {10, 5, "%"},  // エラーケース
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
}