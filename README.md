## 参考
[(zenn)自作のスキーマを使ってGraphQLサーバーを作ろう](https://zenn.dev/hsaki/books/golang-graphql/viewer/originalserver)

---

## TODO
下記ファイルにスキーマを定義
*.graphql

定義後下記を実行
```shell
go run github.com/99designs/gqlgen generate
```

---

### 初期設定
```shell
$ cp .env.sample .env

$ cp ./src/.env.sample ./src/.env

$ docker compose build

$ docker compose up -d
```

起動後下記にアクセス   
[http://localhost/](http://localhost/)


