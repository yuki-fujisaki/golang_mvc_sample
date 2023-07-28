# golang_mvc_sample

## ディレクトリ構造
```
.
├── cmd
│   └── cmd
│       ├── server.go # serverを起動するmain関数を書く
│       └── router
│           └── router.go # controllerをroutingに登録する。server.goから呼び出される。
├── db
│   └── migrations
│       └── 20210405064956_user.sql # dbのマイグレーションを書く
├── go.mod # go.mod.initで生成
└── pkg
    ├── controller
    │   └── user.go # リクエストを変換し、modelのロジックを呼び出し、viewに変換する
    ├── model
    │   └── user.go # gormに依存するstruct
    └── view
        └── user.go # jsonに変換できるstruct

7 directories, 7 files
```
## ライブラリ
- [gorm](https://github.com/go-gorm/gorm)
    - ORM
- [dbmate](https://github.com/amacneil/dbmate)
    - マイグレーションファイルの作成と、実際のマイグレーションに使う。
- [go-chi](https://github.com/go-chi/chi)
    - ルーティングとサーバーの起動に使う。
