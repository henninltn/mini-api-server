mini-api-server
===============

勉強用のAPIサーバー.


## Enviroment 
環境構築の手順.
1. Go本体のインストール
2. Goの補助ツールのインストール
3. Spacemacsの設定
4. データベース(MongoDB)のインストール

### Golang
goはインストール済み.

インストールされているかの確認にバージョンの確認がよく使われる.

デスクトップ左上の「アプリケーション」-> 「システムツール」にあるUXTermを開いて以下のコマンドを打つ.
```
$ go version
go version go1.6.3 darwin/amd64
```
($は$を除いたコマンドをシェルに打つという意味)


### Go Tools
Go言語を書くための支援ツールのインストール.

すでにインストールされていたgoを使う.

goはコンパイラとパッケージマネージャ(パケージ管理ツール)の役割も果たしているので、以下のように簡単にパッケージをインストールできる.
他に何ができるのか興味があるならシェルに```go help```とタイプしてみるといい.

```
$ go get       golang.org/x/tools/cmd/goimports
$ go get -u -v github.com/nsf/gocode
$ go get    -v github.com/rogpeppe/godef
$ go get -u -v golang.org/x/tools/cmd/oracle
$ go get -u    github.com/golang/lint/golint
$ go get -u -v golang.org/x/tools/cmd/gorename
```

「permission denied」みたいなの出たらコマンドの前に sudo 付けてみる.


### Spacemacs

[spacemacs](https://github.com/syl20bnr/spacemacs)

上記の layers -> +lang -> go に詳しい説明がある.

取り敢えず```~/.spacemacs```の上の方にある```dotspacemacs-configuration-layers '( ... )```の中に```go```って書く.

ついでにauto-completionとsyntax-checking入れとけばって書いてるので入れとく.

それぞれ layers -> auto-completion、layers -> syntax-checking に詳しい説明がある.

同じく```~/.spacemacs```の```dotspacemacs-configuration-layers '( ... )```の中に以下を記述.
```
dotspacemacs-configuration-layers
'(
  (auto-completion :variables
                   auto-completion-return-key-behavior 'complete
                   auto-completion-tab-key-behavior 'cycle
                   auto-completion-complete-with-key-sequence nil
                   auto-completion-complete-with-key-sequence-delay 1.0
                   auto-completion-private-snippets-directory nil)
  go
  syntax-checking
)
```

### MongoDB
データベースを使うのでインストールする.

今回はMongoDBを使う.

まずパケージの状態を最新にしてから、mongodbをインストールする.
```
$ pacman -Syu
$ pacman -S mongodb
```


## Installation
1. GitLabからソースコードをclone
2. フレームワークGinのパッケージのインストール

https://gitlab.com/hennin/mini-api-server にソースコードを上げてある.

インストール済みのgitを使って上記のコードを```$GOPATH/src/gitlab.com/ユーザ名```にコピーしてくる.

ユーザ名のところは自分で[GitLab](https://gitlab.com/users/sign_in)に行ってADOCUSのメールアドレスで登録したものにする.

分からなければhenninにしといてもいい.

```
$ echo $GOPATH
(or $ print $GOPATH)
/home/adocus/Desktop/Development/Golang

$ mkdir -p $GOPATH/src/gitlab.com/ユーザ名
$ cd $GOPATH/src/gitlab.com/ユーザ名

// print working directory 作業ディレクトリを表示 の略
$ pwd
/home/adocus/Desktop/Development/Golang/src/gitlab.com/ユーザ名

$ git clone https://gitlab.com/hennin/mini-api-server.git

$ ls
mini-api-server
$ cd mini-api-server
$ ls
README.md db handler main.go
```

あとフレームワークとしてGin使ってるのでパッケージマネージャでそれもインストールしておく.
```
$ go get github.com/gin-gonic/gin
```


## Usage

### Summary

1. MongoDBの起動
2. Goで実行
3. ウェブブラウザでアクセス
```
$ systemctl start mongodb
$ cd $GOPATH/src/gitlab.com/ユーザ名/mini-api-server
$ go run main.go
```
http://localhost:3000/ にアクセス
取り敢えず以上のコマンドで動きます.


### Detail

1. MongoDBの起動

まずは先ほどインストールしたMongoDBを起動する必要がある.

Arch LinuxではMongoDBはサービスとして扱われるので、systemctlを使って起動、停止を行う.
```
// 起動
$ systemctl start mongodb

// 停止
$ systemctl stop mongodb
```

2. Goでの実行

MongoDBの起動後に、main.goのあるディレクトリで```go run main.go```と打てばいい.

```
$ cd $GOPATH/src/gitlab.com/ユーザ名/mini-api-server
$ go run main.go
```

ここまできて「あれ、Goってコンパイル言語じゃね？」って思った人.
```go run```を使うとインタプリタ言語っぽくすぐに実行できるけど、ちゃんとコンパイルしてます.

もちろん```go build main.go```でコンパイルしてから生成された実行ファイルmainを実行してもいい.
```
$ cd $GOPATH/src/gitlab.com/ユーザ名/mini-api-server
$ go build main.go
$ ls
README.md db handler main main.go
$ ./main.go
```

なお、実行ファイルを実行する時は、そのまま```main```と打つとシェルが```main```っていうコマンドと勘違いするので、カレントディレクトリ(現在のディレクトリ)を指す```.```を使って実行ファイルを指定する.

3. ウェブブラウザでアクセス
http://localhost:3000/

実はこれ、ちょっと気持ち悪いけど、自分のコンピュータの一部(mini-api-serverディレクトリ以下)をサーバーとして、自分のコンピュータのみにアクセスを許可して公開している.


## Author
[hennin](https://gitlab.com/hennin)