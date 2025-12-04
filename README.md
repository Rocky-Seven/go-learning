# Go言語 学習リポジトリ 🚀

Go言語の学習用リポジトリへようこそ！

このプロジェクトはGitHub Codespacesを使用して、ブラウザ上で簡単にGo言語を学習できる環境を提供します。

## 📚 ブログ記事との連動

このリポジトリは、ブログ記事と連動した実践的な学習環境を提供しています。

各記事で解説したコードを実際に動かしながら学習できます。

### 公開済みの記事

#### 第1回：Go言語入門「はじめの一歩から学ぶ基礎文法」
- **記事URL**: [https://my-studies.org/introduction-to-go-language-learn-basic-grammar-from-beginning/](https://my-studies.org/introduction-to-go-language-learn-basic-grammar-from-beginning/)
- **対応ディレクトリ**: `basics/02-variables-and-types/`
- **学べる内容**:
  - パッケージの概念
  - 変数宣言（`var`と`:=`の使い分け）
  - 基本的なデータ型（int, string, bool, float64）
  - `fmt.Println`での出力

#### 第2回：Go言語で条件分岐とループをマスターしよう！制御構文の基礎
- **対応ディレクトリ**: `basics/03-control-flow/`
- **学べる内容**:
  - if文による条件分岐
  - switch文による多分岐
  - for文によるループ
  - break/continueでのループ制御

## 🎯 このリポジトリについて

このリポジトリは、Go言語の基礎から実践的な内容まで段階的に学習できるように構成されています。

GitHub Codespacesを使用することで、ローカル環境の構築なしにすぐに学習を開始できます。

## 🚀 クイックスタート

### GitHub Codespacesで始める（推奨）

1. このリポジトリのページで「Code」ボタンをクリック
2. 「Codespaces」タブを選択
3. 「Create codespace on main」をクリック

数分で開発環境が起動します！

### ローカルで始める

```bash
# リポジトリをクローン
git clone https://github.com/Rocky-Seven/go-learning.git
cd go-learning

# Goのバージョン確認
go version
```

## 📁 ディレクトリ構成

```
go-learning/
├── .devcontainer/           # Codespaces設定
│   └── devcontainer.json
├── basics/                  # 基礎編
│   ├── 01-hello/           # Hello World
│   ├── 02-variables-and-types/  # 変数と型（第1回記事対応）
│   ├── 03-control-flow/    # 制御構文（第2回記事対応）
│   ├── 04-functions/       # 関数
│   ├── 05-slices-and-maps/ # スライスとマップ
│   └── BASICS.md           # 基礎編の詳細ガイド
├── intermediate/            # 中級編
│   ├── 01-structs/
│   ├── 02-interfaces/
│   └── 03-goroutines/
├── advanced/                # 応用編
│   ├── 01-web-server/
│   ├── 02-database/
│   └── 03-testing/
├── exercises/               # 演習問題
└── README.md
```

## 🎓 学習カリキュラム

### 基礎編（Basics）

- ✅ **Hello World**: 最初のGoプログラム
- ✅ **変数と型**: データ型、変数宣言、定数（[第1回記事](https://my-studies.org/introduction-to-go-language-learn-basic-grammar-from-beginning/)）
- ✅ **制御構文**: if, for, switch（第2回記事）
- **関数**: 関数定義、引数、戻り値
- **配列とスライス**: データ構造の基礎
- **マップ**: キー・バリュー型のデータ構造

### 中級編（Intermediate）

- **構造体**: カスタム型の定義
- **メソッド**: レシーバー付き関数
- **インターフェース**: 抽象化と多態性
- **エラーハンドリング**: エラー処理のベストプラクティス
- **ゴルーチン**: 並行処理の基礎
- **チャネル**: ゴルーチン間の通信

### 応用編（Advanced）

- **Webサーバー**: net/httpパッケージ
- **データベース接続**: database/sqlの使用
- **テスティング**: testing パッケージ
- **RESTful API**: 実践的なAPI開発
- **パッケージ管理**: Go Modules

## 💻 コードの実行方法

### プログラムを実行

```bash
# ディレクトリに移動
cd basics/02-variables-and-types

# プログラムを実行
go run main.go

# ビルドして実行
go build
./02-variables-and-types
```

### テストを実行

```bash
# テストを実行
go test

# カバレッジ付きで実行
go test -cover

# 詳細表示
go test -v
```

### コードの整形と検証

```bash
# ファイルをフォーマット
go fmt ./...

# コードの静的解析
go vet ./...
```

## 📖 学習の進め方

### 推奨する学習フロー

1. **ブログ記事を読む**: まず対応するブログ記事で概念を理解
2. **コードを確認**: GitHubの対応ディレクトリでコードを確認
3. **実際に動かす**: `go run main.go`で実行してみる
4. **コードを改変**: 自分で値を変えて動作を確認
5. **演習問題に挑戦**: `exercises/`や各ディレクトリのREADME.mdで理解度をチェック

### 詳細な学習ガイド

各編の詳細な学習ガイドを用意しています：

- **基礎編**: `basics/BASICS.md` - 全セクションの詳細なコード例と演習問題

### おすすめの学習順序

基礎編 → 中級編 → 応用編 → 自分でプロジェクト作成

## ❓ よくある質問

**Q: Codespacesが起動しない**  
A: ブラウザのキャッシュをクリアして再試行してください

**Q: `go: command not found`**  
A: Codespacesの再起動を試してください。`.devcontainer/devcontainer.json`の設定を確認してください

**Q: モジュールのエラーが出る**  
A: `go mod tidy`を実行してください

**Q: ブログ記事とコードが対応していない**  
A: 各記事の「対応ディレクトリ」セクションで正しいディレクトリを確認してください

## 🤝 コントリビューション

改善提案やバグ報告は、IssueまたはPull Requestでお願いします！

1. このリポジトリをフォーク
2. 新しいブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add some amazing feature'`)
4. ブランチにプッシュ (`git push origin feature/amazing-feature`)
5. Pull Requestを作成

## 📄 ライセンス

MIT License - 自由に学習・改変してください

## 🎉 最後に

Happy Coding! 🎉

質問があれば、Issueでお気軽にどうぞ！

一緒にGo言語を楽しく学んでいきましょう！