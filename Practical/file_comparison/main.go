//ファイル比較プログラム
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type DiffResult struct {
	LineNum int
	Type    string // "added", "removed", "modified"
	File1   string
	File2   string
}

func main() {
	// コマンドライン引数の定義
	file1 := flag.String("file1", "", "比較元ファイル")
	file2 := flag.String("file2", "", "比較先ファイル")
	outputFile := flag.String("output", "", "結果出力ファイル (省略時は標準出力)")
	ignoreCase := flag.Bool("ignore-case", false, "大文字小文字を区別しない")
	ignoreSpace := flag.Bool("ignore-space", false, "空白を無視")
	
	flag.Parse()

	if *file1 == "" || *file2 == "" {
		fmt.Println("使用方法: go run main.go -file1 <ファイル1> -file2 <ファイル2> [オプション]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// ファイル拡張子を確認
	ext1 := strings.ToLower(filepath.Ext(*file1))
	ext2 := strings.ToLower(filepath.Ext(*file2))

	var diffs []DiffResult
	var err error

	// ファイルタイプに応じて比較方法を選択
	if ext1 == ".csv" && ext2 == ".csv" {
		diffs, err = compareCSV(*file1, *file2, *ignoreCase, *ignoreSpace)
	} else {
		diffs, err = compareTXT(*file1, *file2, *ignoreCase, *ignoreSpace)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	// 結果の出力
	output := os.Stdout
	if *outputFile != "" {
		output, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "出力ファイル作成エラー: %v\n", err)
			os.Exit(1)
		}
		defer output.Close()
	}

	printDiffs(output, diffs, *file1, *file2)
}

// テキストファイルの比較
func compareTXT(file1, file2 string, ignoreCase, ignoreSpace bool) ([]DiffResult, error) {
	lines1, err := readLines(file1)
	if err != nil {
		return nil, fmt.Errorf("ファイル1読み込みエラー: %w", err)
	}

	lines2, err := readLines(file2)
	if err != nil {
		return nil, fmt.Errorf("ファイル2読み込みエラー: %w", err)
	}

	return diffLines(lines1, lines2, ignoreCase, ignoreSpace), nil
}

// CSVファイルの比較
func compareCSV(file1, file2 string, ignoreCase, ignoreSpace bool) ([]DiffResult, error) {
	records1, err := readCSV(file1)
	if err != nil {
		return nil, fmt.Errorf("CSV1読み込みエラー: %w", err)
	}

	records2, err := readCSV(file2)
	if err != nil {
		return nil, fmt.Errorf("CSV2読み込みエラー: %w", err)
	}

	// CSVレコードを文字列に変換
	lines1 := make([]string, len(records1))
	lines2 := make([]string, len(records2))

	for i, record := range records1 {
		lines1[i] = strings.Join(record, ",")
	}
	for i, record := range records2 {
		lines2[i] = strings.Join(record, ",")
	}

	return diffLines(lines1, lines2, ignoreCase, ignoreSpace), nil
}

// 行単位でファイルを読み込む
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// CSVファイルを読み込む
func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil && err != io.EOF {
		return nil, err
	}

	return records, nil
}

// 行の差分を検出
func diffLines(lines1, lines2 []string, ignoreCase, ignoreSpace bool) []DiffResult {
	var diffs []DiffResult
	maxLen := len(lines1)
	if len(lines2) > maxLen {
		maxLen = len(lines2)
	}

	for i := 0; i < maxLen; i++ {
		var line1, line2 string

		if i < len(lines1) {
			line1 = lines1[i]
		}
		if i < len(lines2) {
			line2 = lines2[i]
		}

		// 正規化
		cmp1 := normalize(line1, ignoreCase, ignoreSpace)
		cmp2 := normalize(line2, ignoreCase, ignoreSpace)

		if cmp1 != cmp2 {
			var diffType string
			if line1 == "" {
				diffType = "added"
			} else if line2 == "" {
				diffType = "removed"
			} else {
				diffType = "modified"
			}

			diffs = append(diffs, DiffResult{
				LineNum: i + 1,
				Type:    diffType,
				File1:   line1,
				File2:   line2,
			})
		}
	}

	return diffs
}

// 文字列の正規化
func normalize(s string, ignoreCase, ignoreSpace bool) string {
	if ignoreSpace {
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "\t", "")
	}
	if ignoreCase {
		s = strings.ToLower(s)
	}
	return s
}

// 差分結果の出力
func printDiffs(w io.Writer, diffs []DiffResult, file1, file2 string) {
	if len(diffs) == 0 {
		fmt.Fprintln(w, "差分はありません。ファイルは同一です。")
		return
	}

	fmt.Fprintf(w, "=== ファイル比較結果 ===\n")
	fmt.Fprintf(w, "ファイル1: %s\n", file1)
	fmt.Fprintf(w, "ファイル2: %s\n", file2)
	fmt.Fprintf(w, "差分数: %d\n\n", len(diffs))

	for _, diff := range diffs {
		fmt.Fprintf(w, "--- 行 %d ---\n", diff.LineNum)
		
		switch diff.Type {
		case "added":
			fmt.Fprintf(w, "  [追加] %s\n", diff.File2)
		case "removed":
			fmt.Fprintf(w, "  [削除] %s\n", diff.File1)
		case "modified":
			fmt.Fprintf(w, "  [変更前] %s\n", diff.File1)
			fmt.Fprintf(w, "  [変更後] %s\n", diff.File2)
		}
		fmt.Fprintln(w)
	}
}