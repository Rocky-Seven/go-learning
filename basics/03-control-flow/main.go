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

	// if-else
	age = 15
	if age >= 18 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}

	// if-else if-else
	score := 75
	if score >= 90 {
		fmt.Println("評価: A")
	} else if score >= 70 {
		fmt.Println("評価: B")
	} else if score >= 50 {
		fmt.Println("評価: C")
	} else {
		fmt.Println("評価: D")
	}

	// 初期化付きif文
	if num := 10; num%2 == 0 {
		fmt.Println(num, "は偶数です")
	}

	fmt.Println()

	// 2. switch文
	fmt.Println("【2. switch文】")
	day := "月曜日"
	switch day {
	case "月曜日":
		fmt.Println("週の始まりです")
	case "金曜日":
		fmt.Println("週末が近いです")
	case "土曜日", "日曜日":
		fmt.Println("休日です")
	default:
		fmt.Println("平日です")
	}

	// 条件式を使うswitch
	score = 85
	switch {
	case score >= 90:
		fmt.Println("優秀です")
	case score >= 70:
		fmt.Println("良好です")
	case score >= 50:
		fmt.Println("合格です")
	default:
		fmt.Println("不合格です")
	}

	fmt.Println()

	// 3. for文
	fmt.Println("【3. for文】")

	// 伝統的なfor文
	fmt.Println("カウントアップ:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// whileスタイルのfor文
	fmt.Println("whileスタイル:")
	count := 0
	for count < 3 {
		fmt.Printf("カウント: %d ", count)
		count++
	}
	fmt.Println()

	// rangeを使った反復処理
	fmt.Println("range を使った反復:")
	fruits := []string{"りんご", "バナナ", "みかん"}
	for index, fruit := range fruits {
		fmt.Printf("%d番目: %s\n", index, fruit)
	}

	// インデックスを無視
	fmt.Println("インデックスなし:")
	for _, fruit := range fruits {
		fmt.Printf("- %s\n", fruit)
	}

	fmt.Println()

	// 4. break と continue
	fmt.Println("【4. break と continue】")

	// break
	fmt.Println("break の例:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// continue
	fmt.Println("continue の例:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println()

	// 5. 実践例：簡単な数当てゲーム（自動プレイ版）
	fmt.Println("【5. 実践例：数当てゲーム】")
	targetNumber := 7
	guesses := []int{5, 8, 7} // シミュレーション用の予想

	fmt.Println("1から10までの数字を当ててください")

	for attempt := 0; attempt < len(guesses); attempt++ {
		guess := guesses[attempt]
		fmt.Printf("%d回目の挑戦: %d\n", attempt+1, guess)

		if guess == targetNumber {
			fmt.Println("正解です！おめでとうございます！")
			break
		} else if guess < targetNumber {
			fmt.Println("もっと大きい数字です")
		} else {
			fmt.Println("もっと小さい数字です")
		}

		if attempt == len(guesses)-1 {
			fmt.Println("残念！正解は", targetNumber, "でした")
		}
	}

	fmt.Println("\n=== 学習完了！ ===")
}
