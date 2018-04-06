# go-wiki

## Settings

* install

```markdown
go get github/mmm888/go-wiki
```

* create wiki directory (init)

```markdown
cd $GOPATH/src/mmm888/go-wiki
mkdir wiki
git -C wiki init
```

* create wiki directory (exist)

```markdown
cd $GOPATH/src/mmm888/go-wiki
git clone https://github.com/hoge/fuga.git wiki
```

* run (go run)

```markdown
cd $GOPATH/src/mmm888/go-wiki
make run
```

* run (docker run)

```markdown
cd $GOPATH/src/mmm888/go-wiki
docker-compose up -d
```

## TODO

* test を書く
* errorhandler 作成
  * エラー処理を multierror などでラップする
  * 存在しないファイルの場合 errorhandler にリダイレクト

## Memo

* vue.js
  * 分割編集
* diff: 差分表示に色付け
* diff: 見やすくする
* blackfriday のオプション確認
* 認証機能追加
