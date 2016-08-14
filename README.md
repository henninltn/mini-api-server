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

goはコンパイラとパッケージマネージャ(パケージ管理ツール)の役割も果たしているので, 以下のように簡単にパッケージをインストールできる.
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

それぞれ layers -> auto-completion, layers -> syntax-checking に詳しい説明がある.

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

まずパケージの状態を最新にしてから, mongodbをインストールする.
```
$ pacman -Syu
$ pacman -S mongodb
```

ついでにこのプログラムが使用するデータベースとコレクションを作っておく。
```
$ mongo
> show dbs
test
> use mini-api-server
> show collections

> db.createCollection('thread')
> db.createCollection('respose')
> show collections
thread response
> show dbs
test mini-api-server
> db.thread.find()

> db.thread.insert({"title":"test"})
> db.thread.find()
{"id" : ObjectId("------------------------"), "title" : "test"}
> exit
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

[GIN-debug] GET    /thread ......
[GIN-debug] GET    /thread/:id ......
......
```
http://localhost:8080/thread にブラウザからアクセス


### Detail

#### MongoDBの起動

まずは先ほどインストールしたMongoDBを起動する必要がある.

Arch LinuxではMongoDBはサービスとして扱われるので, systemctlを使って起動, 停止を行う.

```
// 起動
$ systemctl start mongodb

// 停止
$ systemctl stop mongodb
```

#### Goでの実行

MongoDBの起動後に, main.goのあるディレクトリで```go run main.go```と打てばいい.

```
$ cd $GOPATH/src/gitlab.com/ユーザ名/mini-api-server
$ go run main.go
```

ここまできて「あれ, Goってコンパイル言語じゃね？」って思った人.
```go run```を使うとインタプリタ言語っぽくすぐに実行できるけど, ちゃんとコンパイルしてます.

もちろん```go build main.go```でコンパイルしてから生成された実行ファイルmainを実行してもいい.

```
$ cd $GOPATH/src/gitlab.com/ユーザ名/mini-api-server
$ go build main.go
$ ls
README.md db handler main main.go
$ ./main.go
```

なお, 実行ファイルを実行する時は, そのまま```main```と打つとシェルが```main```っていうコマンドと勘違いするので, カレントディレクトリ(現在のディレクトリ)を指す```.```を使って実行ファイルを指定する.

#### ウェブブラウザでアクセス
普通にブラウザにURL打てばGETになるので http://localhost:8080/thread にアクセス.

さっきMongoDBで作成した```{"id":"------------------------",title":"test"}```が見える.

実はこれ, ちょっと気持ち悪いけど, 自分のコンピュータの一部(mini-api-serverディレクトリ以下)をサーバーとして, 自分のコンピュータのみにアクセスを許可して公開している.

ちなみにGETは検索, POSTは新規作成, PATCHは編集, DELETEは削除.

POST, PATCH, DELETEのテストについては、シェルスクリプトを書いてtest_sh以下に置いておいた.

アクセスするURL, 実行する回数, リクエストデータのJSONを指定する.

アクセスするURLはgo run main.goしたとこの下に書いてある.

```
[GIN-debug] GET  /thread    ......
[GIN-debug] GET  /thread:id ......
[GIN-debug] POST /thead     ......
......
```



##### POST(新規作成)

```
$ cd test_sh
$ ./post_test http://localhost:8080/thread 1 '{"title":"post test"}'
```
このあとブラウザで http://localhost:8080/thread アクセス

ちなみにidは自動生成される.


##### PATCH(編集)
idは http://localhost:8080/thread からどれかコピーしてきて.

```
$ ./patch_test http://localhost:8008/thread 1 '{"id":"------------------------","title":"update!"}'
```
またブラウザにアクセスして確認すると、idをコピーしたやつのtitleがupdate!になってるはず.


##### DELETE(削除)

```
$ ./delete_test http://localhost:8080/thread 1 '{"id":"------------------------"}'
```
例によってidは適当にコピってくる.ブラウザで確認するとidコピったやつが消えてるはず.


##### GET(検索)
先にPOSTでいっぱい作っときます.

そしてブラウザに http://localhost:8080/thread/------------------------ でアクセス.

もちろん適当にidをコピる. idコピったやつだけ表示されるはず.

つまり、/thread でアクセスすると全件検索で、/thead/:id でアクセスすると1件検索になる.


## Author
[hennin](https://gitlab.com/hennin)
